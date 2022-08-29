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
			fmt.Println("97:", err.Error())
		}
		if !u.IsAbs() {
			l, err := url.Parse(link)
			if err != nil {
				fmt.Println("102:", err.Error())
			}
			/*			if !strings.HasSuffix(to, "/") || !strings.HasSuffix(to, "\\") {
						to = to + "/"
					}*/
			toCreateDir := to + "/" + l.Host + filepath.Dir(k)
			os.MkdirAll(toCreateDir, 0777)
			fileToDownload := l.Scheme + "://" + l.Host + filepath.Dir(k) + "/" + filepath.Base(k)
			fileToDownload = strings.ReplaceAll(fileToDownload, `\`, "/")
			get, err := http.Get(fileToDownload)
			if err != nil {
				fmt.Println("110:", err.Error())
			}
			fileToCreate := toCreateDir + "/" + filepath.Base(k)
			f, err := os.Create(fileToCreate)
			if err != nil {
				fmt.Println("117:", err.Error())
			}
			io.Copy(f, get.Body)
			f.Close()
			get.Body.Close()
		}
	}
}

func main() {
	depth := 2
	link := "https://www.win-rar.com/"
	to := "e:/download"
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

}
