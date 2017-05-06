package om

import (
	"encoding/xml"
	"testing"
)

func TestVisitor_QueryStatAndParam(t *testing.T) {
	v := new(Visitor)
	v.Id = 123
	v.CallId = 1232342
	v.From = "18987092886"
	v.To = "18687185198"
	v.State = "Talk"
	x, err := xml.MarshalIndent(v, "", "	")
	if err != nil {
		t.Error(err)
	}
	t.Log(xml.Header + string(x))
}
