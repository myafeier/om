package om

import "encoding/xml"

//来电
type Visitor struct {
	XMLName xml.Name `xml:"visitor"`
	Id      int      `xml:"id,attr"`          //来电的短期唯一标识符，用于作为对某个来电进行所有操作的唯一依据。当来电呼入时，OM为该来电分配一个编号（visitor id），当该来电释放后，该编号会随之失效，之后此编号可能会再分配给其他来电。
	From    string   `xml:"from,omitempty"`   //原始主叫号码
	To      string   `xml:"to,omitempty"`     //原始被叫号码（对于visitor而言，原始被叫为该visitor呼入OM的中继）
	CallId  int      `xml:"callid,omitempty"` //通话的相对唯一标识符
	State   string   `xml:"state,omitempty"`  //通话状态 Talk: 通话进行中 Progress: 呼叫处理过程中 Wait: 呼叫等待中
	Ext     *Ext     `xml:"ext,omitempty"`
	Menu    *Menu    `xml:"menu,omitempty"`
	Outer   *Outer   `xml:"outer,omitempty"`
}

//来电
type VisitorBasic struct {
	XMLName xml.Name `xml:"visitor"`
	Id      int      `xml:"id,attr"`               //来电的短期唯一标识符，用于作为对某个来电进行所有操作的唯一依据。当来电呼入时，OM为该来电分配一个编号（visitor id），当该来电释放后，该编号会随之失效，之后此编号可能会再分配给其他来电。
	From    string   `xml:"from,attr,omitempty"`   //原始主叫号码
	To      string   `xml:"to,attr,omitempty"`     //原始被叫号码（对于visitor而言，原始被叫为该visitor呼入OM的中继）
	CallId  int      `xml:"callid,attr,omitempty"` //通话的相对唯一标识符
	State   string   `xml:"state,attr,omitempty"`  //通话状态 Talk: 通话进行中 Progress: 呼叫处理过程中 Wait: 呼叫等待中
	Info    string   `xml:"info,omitempty"`        //按键信息事件(DTMF)
	Menu    *Menu    `xml:"menu,omitempty"`
}

//来电状态
type VisitorStatus struct {
	XMLName xml.Name `xml:"Status"`
	Visitor *Visitor `xml:"visitor"`
}

//查询来电
func (self *Visitor) Query() {

}

//查询来电的状态和属性参数
func (self *Visitor) QueryStatAndParam() {

}

//来电转接分机
func (self *Visitor) TransferExtension() {

}

//来电转接语音菜单
func (self *Visitor) TransferMenu() {

}

//来电转接外部电话
func (self *Visitor) TransferOuter() {

}

//呼叫状态变化的事件报告
func (self *Visitor) StatEvent() {

}

//来电呼入控制流程的事件报告
func (self *Visitor) IncomingEvent() {

}
