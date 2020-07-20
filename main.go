package main

import (
	"golang.org/x/net/html"
	"net"
	"net/http"
)

func scrape(url string) ([]string, error) {
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

func main() {
	host := "https://gedewahyu.com"
	url := host + "/"
	urls, err := scrape(url)
	if err != nil {
		println(err.Error())
	}
	for _, u := range urls {
		println(u)
	}
}
