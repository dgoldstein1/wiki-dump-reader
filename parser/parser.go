package parser

import (
	"encoding/xml"
	wiki "github.com/dgoldstein1/wiki-dump-reader/wikipedia"
	log "github.com/sirupsen/logrus"
	"os"
)

var logMsg = log.Infof
var logErr = log.Errorf
var logWarn = log.Warnf
var logFatal = log.Fatalf

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
	p.Title = wiki.CanonicalizeTitle(p.Title)
	logMsg("Parsing %s", p.Title)
	// find links on page
	e, links := ParseOutLinks(p.Text)
	logMsg("links found: %v", links)
	UpdateMetrics(1)
	return e
}

// finds links within string, which look like:
// '[[legal document]]'
func ParseOutLinks(text string) (e error, links []string) {

	return e, links
}
