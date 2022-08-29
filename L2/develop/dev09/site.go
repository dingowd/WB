package main

import (
	"bufio"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// GetLinks Рекурсивная функция для сбора ссылок
func GetLinks(URL string, depth, level int, links *map[string]struct{}) {
	if level > depth {
		return
	}
	//links := make([]string, 0)
	r, err := http.Get(URL)
	if err != nil {
		return
	}
	doc, err := html.Parse(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link") {
			for _, a := range n.Attr {
				if a.Key == "href" {
					(*links)[a.Val] = struct{}{}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	for k, _ := range *links {
		GetLinks(k, depth, level+1, links)
	}
	return
}

// FilterLinks Фильтрация ссылок не относящихся к сайту
func FilterLinks(links *map[string]struct{}, link string, logger *Lrus) {
	u, _ := url.Parse(link)
	for k, _ := range *links {
		l, _ := url.Parse(k)
		if l.IsAbs() {
			if u.Host != l.Host {
				delete(*links, k)
			}
		}
	}
}

// GetSrc Получение src
func GetSrc(hrefs, src *map[string]struct{}) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img" || n.Data == "font") {
			for _, a := range n.Attr {
				if a.Key == "src" {
					(*src)[a.Val] = struct{}{}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	for k, _ := range *hrefs {
		r, err := http.Get(k)
		if err != nil {
			continue
		}
		doc, err := html.Parse(r.Body)
		if err != nil {
			continue
		}
		f(doc)
	}
}

// download Скачивание файлов по локальным и абсолютным ссылкам
func download(link, to string, links *map[string]struct{}, logger *Lrus) {
	for k, _ := range *links {
		u, err := url.Parse(k)
		if err != nil {
			logger.Info(err.Error())
			continue
		}
		if !u.IsAbs() {
			dir, file := filepath.Split(k)
			if file == "" {
				continue
			}
			f, err := url.Parse(file)
			if err != nil {
				logger.Info(err.Error())
				continue
			}
			file = filepath.Base(f.Path)
			fileToDownload := link + dir + file
			if len(filepath.Ext(file)) < 2 {
				file = file + ".html"
			}
			dirToCreate := filepath.Join(to, dir)
			err = os.MkdirAll(dirToCreate, 0777)
			if err != nil {
				logger.Info(err.Error())
				continue
			}
			fileToCreate := filepath.Join(dirToCreate, file)
			req, err := http.Get(fileToDownload)
			if err != nil {
				logger.Info(err.Error())
				continue
			}
			local, err := os.Create(fileToCreate)
			if err != nil {
				logger.Info(err.Error())
				continue
			}
			io.Copy(local, req.Body)
			req.Body.Close()
			local.Close()
		} else {
			dir, file := filepath.Split(k)
			f, err := url.Parse(file)
			if err != nil {
				logger.Info(err.Error())
				continue
			}
			file = filepath.Base(f.Path)
			if filepath.Ext(file) == "" {
				file = file + ".html"
			}
			dir = strings.ReplaceAll(dir, link, "")
			if dir == "https://" || dir == "http://" {
				dir = ""
				file = "index.html"
			}
			dirToCreate := filepath.Join(to, dir)
			err = os.MkdirAll(dirToCreate, 0777)
			if err != nil {
				logger.Info(err.Error())
				continue
			}
			fileToCreate := filepath.Join(dirToCreate, file)
			local, err := os.Create(fileToCreate)
			if err != nil {
				logger.Info(err.Error())
				continue
			}
			req, err := http.Get(k)
			if err != nil {
				logger.Info(err.Error())
				continue
			}
			io.Copy(local, req.Body)
			req.Body.Close()
			local.Close()
		}
	}
}

// convLinks Функция для преобразования ссылок в локальные
func convLinks(dir, link string, logger *Lrus) {
	filePathWalkDir := func(root, ext string) ([]string, error) {
		var files []string
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && filepath.Ext(info.Name()) == ext {
				files = append(files, path)
			}
			return nil
		})
		return files, err
	}
	files, err := filePathWalkDir(dir, ".html")
	if err != nil {
		logger.Error(err.Error())
		return
	}
	u, _ := url.Parse(link)
	toChange := u.Scheme + "://" + u.Host + "/"
	for _, v := range files {
		file, err := os.Open(v)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		new := v + "1"
		newFile, err := os.Create(new)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			text := fileScanner.Text()
			text = strings.ReplaceAll(text, toChange, "./")
			text = strings.ReplaceAll(text, `href="/`, `href="./`)
			text = strings.ReplaceAll(text, `src="/`, `src="./`)
			text = text + "\n"
			newFile.WriteString(text)
		}
		file.Close()
		newFile.Close()
		os.Remove(v)
		os.Rename(new, v)
	}
}
