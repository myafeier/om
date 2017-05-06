package om

import "encoding/xml"

//分机组
type Group struct {
	XMLName      xml.Name   `xml:"group"`
	Id           int        `xml:"id,attr"`                //分机组的编号，有效值为1~50，是对分机组进行所有操作的唯一依据。用户使用时分机组时无需添加或删除，直接选取某个group id进行配置后即可使用，不用时清空其配置即可。
	VoiceFile    string     `xml:"voicefile,omitempty"`    //语音文件，支持dat和pcm两种格式，这里为呼叫等待时播放的音乐
	Distribution string     `xml:"distribution,omitempty"` //分机组来电分配规则，有效值：sequential（顺选）、circular（轮选）、group（群振） 默认值：circular
	Exts         []*Ext     `xml:"ext,omitempty"`          //分机，这里为该分机组中的分机成员
	Visitors     []*Visitor `xml:"visitor,omitempty"`      //来电，这里为正处于该分机组等待队列中的来电
}

//分机组
type GroupBasic struct {
	XMLName      xml.Name `xml:"group"`
	Id           int      `xml:"id,attr"`                     //分机组的编号，有效值为1~50，是对分机组进行所有操作的唯一依据。用户使用时分机组时无需添加或删除，直接选取某个group id进行配置后即可使用，不用时清空其配置即可。
	VoiceFile    string   `xml:"voicefile,attr,omitempty"`    //语音文件，支持dat和pcm两种格式，这里为呼叫等待时播放的音乐
	Distribution string   `xml:"distribution,attr,omitempty"` //分机组来电分配规则，有效值：sequential（顺选）、circular（轮选）、group（群振） 默认值：circular
}

type GroupStatus struct {
	XMLName xml.Name `xml:"Status"`
	Group   *Group   `xml:"group"`
}

//查询分机组
//该API用于查询分机组的相关信息，如：配置参数（分机成员、呼叫排队时播放的背景音乐、呼叫分配规则）、正在该分机组队列中等待的来电。
//分机组队列中的呼叫排队顺序为从上到下。
func (self *Group) Query() {

}

//设置分机组的配置参数
func (self *Group) SetParam() {

}

//查询分机组的配置参数和等待队列
func (self *Group) QueryParamAndQueue() {

}

//来电转接分机组
func (self *Group) TransferGroup() {

}
