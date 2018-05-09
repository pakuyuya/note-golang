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
		<Root>
		  <TagA>xxx</TagA>
		  <TabB>mmm</TagB>
		  <TabC>nnn</TagC>
		  <TagA>aaa</TagA>
		</Root>`)

	r := Root{}

	xml.Unmarshal(xmldoc, &r)
	fmt.Println(r)
}
