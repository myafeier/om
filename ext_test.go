package om

import (
	"encoding/xml"
	"testing"
)

func TestExt_Query(t *testing.T) {
	et := new(ExtStatus)
	et.Ext = new(Ext)
	et.Ext.Id = 1
	et.Ext.LineId = "line-1"
	et.Ext.StaffId = "6001"
	g1 := new(Group)
	g1.Id = 1
	g2 := new(Group)
	g2.Id = 2
	et.Ext.Groups = append(et.Ext.Groups, g1, g2)

	o := new(Outer)
	o.Id = 2
	o.CallId = 2344
	et.Ext.Outer = o
	x, err := xml.MarshalIndent(et, "", "	")
	if err != nil {
		t.Error(err)
	}
	t.Log(xml.Header + string(x))
	originString := xml.Header + string(x)

	s := new(ExtStatus)
	err = xml.Unmarshal([]byte(originString), s)
	if err != nil {
		t.Error(err)
	}
	t.Logf("groups length:=%d\n", len(s.Ext.Groups))
	t.Logf("et:%+v\n", *s.Ext)
	t.Logf("outer:%+v\n", *s.Ext.Outer)

	etr := new(ExtBasic)
	etr.Id = 2
	etr.Api = "yes"
	etr.Email = "hiall@aliyun.com"
	str, err := xml.Marshal(etr)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ExtResponse XML:%s\n", string(str))

}
