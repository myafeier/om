package om

import "encoding/xml"

type Event struct {
	XMLName   xml.Name      `xml:"Event"`
	Attribute string        `xml:"attribute,attr"`
	Exts      []*ExtBasic   `xml:"ext,omitempty"`
	Visitor   *VisitorBasic `xml:"visitor,omitempty"`
	Menu      *Menu         `xml:"menu,omitempty"`
	Outer     *Outer        `xml:"outer,omitempty"`
}
