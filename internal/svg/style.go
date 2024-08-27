package svg

import (
	"encoding/xml"
)

type Style struct {
	XMLName xml.Name `xml:"style"`
	Content string   `xml:",innerxml"`
}

func (style *Style) AppendContent(s string) {
	style.Content += s
}
