package om

import "encoding/xml"

//去电
type Outer struct {
	XMLName xml.Name `xml:"outer"`
	Id      int      `xml:"id,attr"`          //去电的短期唯一标识符，用于作为对某个去电进行所有操作的唯一依据。当去电呼出时，OM为该去电分配一个编号（outer id），到去电释放后，该编号失效，之后编号可能会再分配给其他去电。
	From    string   `xml:"from,omitempty"`   //原始主叫号码
	To      string   `xml:"to,omitempty"`     //原始被叫号码（对于visitor而言，原始被叫为该visitor呼入OM的中继）
	Trunk   string   `xml:"trunk,omitempty"`  //中继号，即该去电从该中继呼出
	CallId  int      `xml:"callid,omitempty"` //通话的相对唯一标识符
	State   string   `xml:"state,omitempty"`  //通话状态 Talk: 通话进行中 Progress: 呼叫处理过程中 Wait: 呼叫等待中
}

type OuterResponse struct {
	XMLName xml.Name `xml:"outer"`
	Id      int      `xml:"id,attr"`               //去电的短期唯一标识符，用于作为对某个去电进行所有操作的唯一依据。当去电呼出时，OM为该去电分配一个编号（outer id），到去电释放后，该编号失效，之后编号可能会再分配给其他去电。
	From    string   `xml:"from,attr,omitempty"`   //原始主叫号码
	To      string   `xml:"to,attr,omitempty"`     //原始被叫号码（对于visitor而言，原始被叫为该visitor呼入OM的中继）
	Trunk   string   `xml:"trunk,attr,omitempty"`  //中继号，即该去电从该中继呼出
	CallId  int      `xml:"callid,attr,omitempty"` //通话的相对唯一标识符
	State   string   `xml:"state,attr,omitempty"`  //通话状态 Talk: 通话进行中 Progress: 呼叫处理过程中 Wait: 呼叫等待中
}

//去电状态
type OuterStatus struct {
	XMLName xml.Name `xml:"Status"`
	Outer   *Outer   `xml:"outer"`
}

//查询去电
//该API用于查询指定去电的相关信息，如：去电的编号、主叫方、被叫方、通过的中继号码、通话状态、以及通话的相对唯一标识符。
func (self *Outer) Query() {

}

//查询去电的属性参数
func (self *Outer) QueryParam() {

}

//去电转接分机
func (self *Outer) TransferExtension() {

}

//去电转接外部电话
func (self *Outer) TransferOuter() {

}

//双向外呼
func (self *Outer) BidirectionalCall() {

}

//呼叫状态变化的事件报告
func (self *Outer) StatEvent() {

}
