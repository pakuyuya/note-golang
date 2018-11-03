package main

import (
	"encoding/xml"
	"fmt"
)

type Root struct {
	TagA, TagB string
}

func main() {
	xmldoc := []byte(`
		<Roots>
		  <TagA>xxx</TagA>
		  <TagB>mmm</TagB>
		  <TagC>nnn</TagC>
		  <TagA>aaa</TagA>
		</Roots>`)

	r := Root{}

	e := xml.Unmarshal(xmldoc, &r)
	fmt.Println(r)
	fmt.Println(e)
}
