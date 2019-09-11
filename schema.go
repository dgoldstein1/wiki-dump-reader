package main

// checks if is valid wiki link
type IsValidWikiLink func(string) bool

// number of nodesVisited and nodesAdded
type asyncInt int32

type Redirect struct {
	Title string `xml:"title,attr"`
}

type Page struct {
	Title string `xml:"title"`
	Text  string `xml:"revision>text"`
}
