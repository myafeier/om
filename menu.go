/*语音菜单（menu）是OM的一种虚拟通话终端。它具有以下两个特点：
可以指定内容和指定次数地向通话方播放语音文件；
可以检查和汇报通话方输入的按键信息；*/
package om

import "encoding/xml"

//语音菜单
type Menu struct {
	XMLName    xml.Name   `xml:"menu"`
	Id         int        `xml:"id,attr"`    //语音菜单的编号，有效值为1~50，是对语音菜单进行所有操作的唯一依据。用户使用语音菜单时无需添加或删除，直接选取某个menu id进行配置即可使用，不用时清空其配置即可。
	VoiceFile  string     `xml:"voicefile"`  //语音文件，当呼叫转接到该菜单并接通后，OM会向通话方播放该文件。（只支持dat和pcm格式。）
	Repeat     int        `xml:"repeat"`     //语音文件的播放次数，取值范围 0-50，值为0时循环播放。
	InfoLength int        `xml:"infolength"` //拨号检测长度，当该菜单的通话方一旦拨号达到该长度后，OM就会将已统计到的按键信息（DTMF事件）推送给应用服务器，并重新开始统计。
	Exit       byte       `xml:"exit"`       //按键检查结束符，当该菜单的通话方一旦拨了该字符后，OM会立刻将已统计到的按键信息（DTMF事件）推送给应用服务器，并重新开始统计。
	Visitors   []*Visitor `xml:"visitor"`
	Outers     []*Outer   `xml:"outer"`
}

type MenuStatus struct {
	XMLName xml.Name `xml:"Status"`
	Menu    *Menu    `xml:"menu"`
}

//查询语音菜单
//该API用于查询语音菜单的相关信息，如：配置参数（语音文件、拨号检测长度、按键检查结束符）、转接到该菜单的呼叫信息等。
func (self *Menu) Query() {

}

//查询语音菜单的配置参数和呼叫等待队列
func (self *Menu) QueryParamAndQueue() {

}

//设置语音菜单的配置参数
func (self *Menu) SetParam() {

}

//来电转接语音菜单
func (self *Menu) VisitorTransferVoiceMenu() {

}

//去电转语音菜单
func (self *Menu) OuterVoiceMenu() {

}

//语音菜单呼外部电话
func (self *Menu) TransferOuter() {

}

//语音菜单呼分机
func (self *Menu) TransferExtension() {

}

//语音文件播放完毕的事件报告
func (self *Menu) VoiceOverEvent() {

}

//用户按键信息的事件报告
func (self *Menu) KeyPressEvent() {

}
