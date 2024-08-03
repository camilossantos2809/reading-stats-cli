package main

import (
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/html"
)

type ReadingProgress struct {
	Date     string
	Progress string
}

func ExtractProgress(doc *html.Node) []ReadingProgress {
	var progresses []ReadingProgress
	var currentDate string

	var parseHTMLTree func(*html.Node)
	parseHTMLTree = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "div" {
			for _, attr := range node.Attr {
				// Encontrar as datas
				if attr.Key == "style" && strings.Contains(attr.Val, "float:left; font-size:11px; color:#666666") {
					if node.FirstChild != nil && node.FirstChild.Type == html.TextNode {
						currentDate = node.FirstChild.Data
					}
				}
				// Encontrar os progressos
				if attr.Key == "style" && strings.Contains(attr.Val, "font-size:11px; float:left;margin-top:-5px;") {
					if node.FirstChild != nil && node.FirstChild.Type == html.TextNode {
						progress := ReadingProgress{
							Date:     currentDate,
							Progress: node.FirstChild.Data,
						}
						progresses = append(progresses, progress)
					}
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			parseHTMLTree(child)
		}
	}

	parseHTMLTree(doc)
	return progresses
}
