package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cavaliergopher/grab/v3"
)

func main() {
	// parsing flags
	output := pflag.StringP("output-file", "o", "", "logfile")
	input := pflag.StringP("input-file", "i", "", "Input file")
	prefix := pflag.StringP("directory-prefix", "P", ".", "Prefix to download")
	recursive := pflag.BoolP("recursive", "r", false, "Enable recursive loading")
	level := pflag.IntP("level", "l", 5, "Maximum recursive loading depth")
	convertLinks := pflag.BoolP("convert-links", "k", false, "Convert links to a document for offline viewing")
	pflag.Parse()
	// Init logger
	logger := NewLogger()
	logger.SetLevel("DEBUG")
	if *output != "" {
		file, err := os.Create(*output)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		logger.SetOutput(file)
	} else {
		logger.SetOutput(os.Stdout)
	}
	// get links to download
	links := make([]string, 0)
	if *input == "" {
		link := pflag.Args()[0] //link := "https://www.win-rar.com/"
		links = append(links, link)
	} else {
		file, err := os.Open(*input)
		if err != nil {
			logger.Log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			links = append(links, fileScanner.Text())
		}
	}
	if len(links) == 0 {
		logger.Log.Fatal("Check command. URL to download is invalid.")
	}
	// set depth
	depth := *level
	msg := "Depth to download: " + strconv.Itoa(depth)
	logger.Info(msg)
	// set directory to download
	to := *prefix
	if *recursive {
		var msg string
		URL := links[0]
		msg = "URL to download: " + URL
		u, err := url.Parse(URL)
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
		msg = "Download to " + to
		logger.Info(msg)
		to = to + u.Host
		// get refs
		msg = "Getting refs with depth " + strconv.Itoa(*level)
		logger.Info(msg)
		refs := make(map[string]struct{})
		GetLinks(URL, depth, 1, &refs)
		// filter refs
		msg = "Filtering refs"
		logger.Info(msg)
		FilterLinks(&refs, URL, logger)
		// get src
		msg = "Getting src with depth " + strconv.Itoa(*level)
		logger.Info(msg)
		src := make(map[string]struct{})
		GetSrc(&refs, &src)
		// download src
		msg = "Downloading src..."
		logger.Info(msg)
		download(URL, to, &src, logger)
		// download refs
		msg = "Downloading refs..."
		logger.Info(msg)
		download(URL, to, &refs, logger)
		// convert links
		if *convertLinks {
			msg = "Converting links..."
			logger.Info(msg)
			convLinks(*prefix, URL, logger)
		}
	} else {
		for _, v := range links {
			err := os.MkdirAll(*prefix, 0777)
			if err != nil {
				log.Fatal(err)
			}
			_, err = grab.Get(*prefix, v)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
			msg := "Downloaded " + v + " to " + *prefix
			logger.Info(msg)
		}
	}
}
