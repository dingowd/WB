package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var res = regexp.MustCompile(`src=\"[\w[:punct:]]+\"`)

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
		if n.Type == html.ElementNode && n.Data == "a" {
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
func FilterLinks(links *map[string]struct{}, link string) {
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
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img" || n.Data == "font" || n.Data == "link") {
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

func download(link, to string, links *map[string]struct{}) {
	for k, _ := range *links {
		u, err := url.Parse(k)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		if !u.IsAbs() {
			dir, file := filepath.Split(k)
			if file == "" {
				continue
			}
			f, err := url.Parse(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
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
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}
			fileToCreate := filepath.Join(dirToCreate, file)
			req, err := http.Get(fileToDownload)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}
			local, err := os.Create(fileToCreate)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}
			io.Copy(local, req.Body)
			req.Body.Close()
			local.Close()
		} else {
			dir, file := filepath.Split(k)
			f, err := url.Parse(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
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
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}
			fileToCreate := filepath.Join(dirToCreate, file)
			local, err := os.Create(fileToCreate)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}
			req, err := http.Get(k)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}
			io.Copy(local, req.Body)
			req.Body.Close()
			local.Close()
		}
	}
}

func main() {
	depth := 2
	link := "https://www.win-rar.com/"
	to := "e:/download"
	u, err := url.Parse(link)
	if u.Host == "" {
		fmt.Fprintln(os.Stderr, "Link is not valid")
		return
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	to = strings.ReplaceAll(to, `\`, "/")
	if !strings.HasSuffix(to, "/") {
		to = to + "/"
	}
	to = to + u.Host
	refs := make(map[string]struct{})
	GetLinks(link, depth, 1, &refs)
	for k, _ := range refs {
		fmt.Println(k)
	}
	fmt.Println("*************************")
	FilterLinks(&refs, link)
	for k, _ := range refs {
		fmt.Println(k)
	}
	src := make(map[string]struct{})
	GetSrc(&refs, &src)
	fmt.Println("***********SRC***********")
	for k, _ := range src {
		fmt.Println(k)
	}
	download(link, to, &src)
	download(link, to, &refs)

	/*	check := make(map[string]struct{})
		check["https://www.win-rar.com"] = struct{}{}
		download(link, to, &check)*/
}
