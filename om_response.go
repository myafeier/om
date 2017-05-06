//OM响应结构体
package om

import "encoding/xml"

type Response struct {
	XMLName   xml.Name `xml:"Event"`
	Attribute string   `xml:"attribute,attr"`
	Outer     []*Outer `xml:"outer,omitempey"`
	Ext       []*Ext   `xml:"ext,omitempty"`
}
