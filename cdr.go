package om

import "encoding/xml"

type Cdr struct {
	XMLName     xml.Name `xml:"Cdr"`
	Id          string   `xml:"id,attr"`               //话单id，用于关联多个cdr,无关联时为Null
	CallId      int      `xml:"callid,attr,omitempty"` //通话的相对唯一标识符
	Visitor     *Visitor `xml:"visitor,omitempty"`     //来电的编号 或者去电的编号，内部分机通话记录无此字段。
	Outer       *Outer   `xml:"outer,omitempty"`       //来电的编号 或者去电的编号，内部分机通话记录无此字段。
	Type        string   `xml:"type,omitempty"`        //IN（打入）/OU（打出）/FI（呼叫转移入）/FW（呼叫转移出）/LO（内部通话）/CB（双向外呼）
	Route       string   `xml:"Route,omitempty"`       //IP（IP中继）/XO（模拟中继）/IC（内部）/OP（总机）
	TimeStart   string   `xml:"TimeStart,omitempty"`   //呼叫起始时间，即发送或收到呼叫请求的时间
	TimeEnd     string   `xml:"TimeEnd,omitempty"`     //呼叫结束时间，即通话的一方挂断的时间
	Cpn         string   `xml:"CPN,omitempty"`         //主叫号码
	Cdpn        string   `xml:"CDPN,omitempty"`        //被叫号码
	Duration    int      `xml:"duration,omitempty"`    //通话时长，值为0说明未接通。
	TrunkNumber string   `xml:"TrunkNumber,omitempty"` //该路通话所经过的中继号码
	Recoding    string   `xml:"recoding,omitempty"`    //录音文件的相对保存路径，格式：年月日/录音文件 录音文件的命名格式：1.主叫号码_被叫号码_年月日-时分秒_通话的相对唯一标识符.文件格式2.主叫号码_被叫号码_年月日-时分秒_通话的相对唯一标识符_cd.文件格式3.主叫号码_被叫号码_年月日-时分秒_通话的相对唯一标识符_cg.文件格式 cg 为主叫录音，cd 为被叫录音。

}
