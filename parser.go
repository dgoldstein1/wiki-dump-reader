package main

import (
	"encoding/xml"
	wiki "github.com/dgoldstein1/crawler/wikipedia"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func Run(file string) {
	logMsg("Reading in file: %s", file)
	xmlFile, err := os.Open(file)
	if err != nil {
		logErr("Error opening file %v:", err)
		return
	}
	Parse(xmlFile)
}

// main look of parser. Adopted from https://github.com/dps/go-xml-parse/blob/master/go-xml-parse.go
func Parse(file *os.File) {
	defer file.Close()
	decoder := xml.NewDecoder(file)
	var inElement string
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			inElement = se.Name.Local
			// ...and its name is "page"
			if inElement == "page" {
				var p Page
				decoder.DecodeElement(&p, &se)
				HandlePage(p)
			}
		default:
		}
	}
}

// handles what happens when a page tag is discovered
func HandlePage(p Page) error {
	p.Title = CanonicalizeTitle(p.Title)
	logMsg("Parsing %s", p.Title)
	// find links on page
	err, links := ParseOutLinks(p.Text)
	if err != nil {
		logErr("Could not parse out links from node %s: %v", p.Title, err)
		return err
	}
	// update DBs
	neighborsAdded, err := wiki.AddEdgesIfDoNotExist(
		p.Title,
		links,
	)
	if err != nil {
		logErr("Could not add edges to graph for node %s: %v", p.Title, err)
		return err
	}
	// succesfully processed
	UpdateMetrics(len(neighborsAdded))
	return err
}

// finds links within string, which look like:
// '[[legal document]]'
var r, _ = regexp.Compile(`\[\[([^\[\]:]+)\]\]`)

func ParseOutLinks(text string) (e error, links []string) {
	links = r.FindAllString(text, -1)
	return e, links
}

// taken from https://github.com/dps/go-xml-parse/blob/master/go-xml-parse.go
func CanonicalizeTitle(title string) string {
	can := strings.ToLower(title)
	can = strings.Replace(can, " ", "_", -1)
	can = url.QueryEscape(can)
	return can
}
