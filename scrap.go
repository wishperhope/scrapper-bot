package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// todo make it work for another website
var endContentMark = "2018 BOXNOVEL. All rights reserved"

var (
	minimumWord = 3
	pagination  = 10
)

func scrap(url string) ([]string, error) {

	var result []string
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return result, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Print("status code error: ", res.StatusCode, res.Status)
		return result, errors.New("non200")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print(err)
	}

	sel := doc.Find("p")
	var content strings.Builder

	part, count := 1, 1
	for i := range sel.Nodes {
		text := sel.Eq(i).Text()

		// todo make ignore word
		if len(text) >= minimumWord && !strings.Contains(text, endContentMark) {
			content.WriteString(sel.Eq(i).Text())
			content.WriteString("\n")
			count = count + 1
		}

		if count == pagination || i == len(sel.Nodes) || strings.Contains(text, endContentMark) {
			if content.String() != "" {
				result = append(result, content.String())
			}
			part = part + 1
			count = 1
			content.Reset()
		}

		if strings.Contains(text, endContentMark) {
			break
		}
	}

	return result, nil
}
