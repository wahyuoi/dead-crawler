package main

import (
	"golang.org/x/net/html"
	"net"
	"net/http"
	"net/url"
)

func scrape(url string) ([]string, error) {
	println("? " + url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, &net.AddrError{
			Err:  resp.Status,
			Addr: url,
		}
	}
	tkz := html.NewTokenizer(resp.Body)

	var list []string
	for {
		tt := tkz.Next()
		switch {
		case tt == html.ErrorToken:
			return list, nil
		default:
			t := tkz.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				for _, att := range t.Attr {
					if att.Key == "href" {
						list = append(list, att.Val)
					}
				}
			}
		}
	}
}

func check(parent, current string, visited map[string]bool) map[string][]string {
	m := make(map[string][]string)

	if visited[current] {
		return m
	}
	visited[current] = true

	links, err := scrape(current)
	if err != nil {
		m[parent] = []string{current}
		return m
	}

	if !isSameDomain(parent, current) {
		return m
	}

	for _, link := range links {
		link, _ = completeLink(current, link)
		dead := check(current, link, visited)
		for k, v := range dead {
			for _, s := range v {
				m[k] = append(m[k], s)
				println(k + " --> " + s)
			}
		}
	}
	return m
}

func isSameDomain(parent string, current string) bool {
	p1, err1 := url.Parse(parent)
	p2, err2 := url.Parse(current)

	if err1 != nil || err2 != nil {
		return false
	}

	return p1.Host == p2.Host
}

func completeLink(current string, link string) (string, error) {
	parse, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	curParse, err := url.Parse(current)
	if err != nil {
		return "", err
	}

	if parse.Scheme == "" {
		parse.Scheme = curParse.Scheme
	}
	if parse.Host == "" {
		parse.Host = curParse.Host
	}

	// TODO if current link is a file (e.g. HTML), remove file from path
	if len(parse.Path) > 0 && len(curParse.Path) > 0 && parse.Path[0] != '/' {
		if curParse.Path[len(curParse.Path)-1] == '/' {
			parse.Path = curParse.Path + parse.Path
		} else {
			parse.Path = curParse.Path + "/" + parse.Path
		}
	}

	return parse.String(), nil
}
func main() {
	url := "http://kubernetes.io/"

	visited := make(map[string]bool)
	dead := check(url, url, visited)

	for k, v := range dead {
		if len(v) == 0 {
			continue
		}
		println(k)
		for _, link := range v {
			println("-- " + link)
		}
	}
}
