package om

import (
	"encoding/xml"
	"log"
	"strconv"
)

//分机
type Ext struct {
	XMLName         xml.Name `xml:"ext"`
	Id              int      `xml:"id,attr"`                    //分机号，即分机的电话号码，用户可自定义，作为对分机进行除配置以外的操作的唯一依据。
	LineId          string   `xml:"lineid,omitempty"`           //分机的线路编号，由OM分配的固定值，作为对分机进行配置的唯一依据。
	Groups          []*Group `xml:"group,omitempty"`            ///该分机属于分机组,可隶属多组 1~50
	StaffId         string   `xml:"staffid,omitempty"`          //工号，分机接通前会向来电方播放该工号
	CallPickup      string   `xml:"Call_Pickup,omitempty"`      //该分机属于被其他分机来电代接 yes/no
	FwdNumber       string   `xml:"Fwd_Number,omitempty"`       //呼叫转移号码,值为空时关闭
	CallRestriction int      `xml:"Call_Restriction,omitempty"` //呼叫权限,0: 内线，1: 市话，2: 国内，3: 国际
	OffLineNum      string   `xml:"Off_Line_Num,omitempty"`     //当该分机离线时来电将转移到号码
	Fork            string   `xml:"fork,omitempty"`             //同振号码 值为空时关闭
	Email           string   `xml:"email,omitempty"`            //电子邮箱
	Record          string   `xml:"record,omitempty"`           //分机录音状态 on: 开启，off: 关闭
	Api             string   `xml:"api,omitempty"`              ////API开关为7（0为关闭，7为开启）
	VoiceFile       string   `xml:"voicefile,omitempty"`        //语音文件，这里为分机队列中排队等待时播放的语音文件,只支持.dat和.pcm格式
	NoDisturb       string   `xml:"no_disturb,omitempty"`       //免打扰功能开关，开启免打扰后分机将屏蔽所有来电，但能主动发起呼叫,on/off
	FwdType         int      `xml:"Fwd_Type,omitempty"`         //呼叫转移方式,0: 关闭，1: 全转，2: 遇忙或无应答转
	Mobile          string   `xml:"mobile,omitempty"`           //分机绑定的手机号 该手机号可作为呼叫转移、离线转移的缺省配置 值为空时关闭
	Stat            string   `xml:"stat,omitempty"`             //线路状态,Ready: 空闲可用 Active: 振铃、回铃或通话中 Progress：模拟分机摘机后等待拨号以及拨号过程中 Offline: IP分机离线 Offhook：模拟分机听催挂音时的状态
	Outer           *Outer   `xml:"outer,omitempty"`            //去电
}

//分机响应
type ExtBasic struct {
	XMLName         xml.Name `xml:"ext"`
	Id              int      `xml:"id,attr"`                         //分机号，即分机的电话号码，用户可自定义，作为对分机进行除配置以外的操作的唯一依据。
	LineId          string   `xml:"lineid,attr,omitempty"`           //分机的线路编号，由OM分配的固定值，作为对分机进行配置的唯一依据。
	StaffId         string   `xml:"staffid,attr,omitempty"`          //工号，分机接通前会向来电方播放该工号
	CallPickup      string   `xml:"Call_Pickup,attr,omitempty"`      //该分机属于被其他分机来电代接 yes/no
	FwdNumber       string   `xml:"Fwd_Number,attr,omitempty"`       //呼叫转移号码,值为空时关闭
	CallRestriction int      `xml:"Call_Restriction,attr,omitempty"` //呼叫权限,0: 内线，1: 市话，2: 国内，3: 国际
	OffLineNum      string   `xml:"Off_Line_Num,attr,omitempty"`     //当该分机离线时来电将转移到号码
	Fork            string   `xml:"fork,attr,omitempty"`             //同振号码 值为空时关闭
	Email           string   `xml:"email,attr,omitempty"`            //电子邮箱
	Record          string   `xml:"record,attr,omitempty"`           //分机录音状态 on: 开启，off: 关闭
	Api             string   `xml:"api,attr,omitempty"`              ////API开关为7（0为关闭，7为开启）
	VoiceFile       string   `xml:"voicefile,attr,omitempty"`        //语音文件，这里为分机队列中排队等待时播放的语音文件,只支持.dat和.pcm格式
	NoDisturb       string   `xml:"no_disturb,attr,omitempty"`       //免打扰功能开关，开启免打扰后分机将屏蔽所有来电，但能主动发起呼叫,on/off
	FwdType         int      `xml:"Fwd_Type,attr,omitempty"`         //呼叫转移方式,0: 关闭，1: 全转，2: 遇忙或无应答转
	Mobile          string   `xml:"mobile,attr,omitempty"`           //分机绑定的手机号 该手机号可作为呼叫转移、离线转移的缺省配置 值为空时关闭
	Stat            string   `xml:"stat,attr,omitempty"`             //线路状态,Ready: 空闲可用 Active: 振铃、回铃或通话中 Progress：模拟分机摘机后等待拨号以及拨号过程中 Offline: IP分机离线 Offhook：模拟分机听催挂音时的状态
}

type ExtStatus struct {
	XMLName xml.Name `xml:"Status"`
	Ext     *Ext     `xml:"ext"`
}

// 查询分机
// 该API用于查询指定分机的相关信息，如：配置参数（分机号、录音开关等）、分机状态、分机的通话方等。
// OP:id 分机号 必须为OM上的有效分机，值不能为空
func (self *Ext) Query(id int) (result *ExtStatus) {
	queryString := `<?xml version="1.0" encoding="utf-8" ?>
			<Control attribute="Query">
			<ext id="` + strconv.Itoa(id) + `"/>
			</Control>`
	log.Println(queryString)

	return
}

//获取分机列表（查询设备信息，即可获取所有分机的线路编号和电话号码）
func (self *Ext) GetList() {

}

//查询某个分机的配置参数和状态
func (self *Ext) GetParamAndStat() {

}

//设置某个分机的配置参数
func (self *Ext) SetParam() {

}

//分机呼叫分机
func (self *Ext) CallExt() {

}

//分机呼叫外部电话
func (self *Ext) CallOuter() {

}

//让分机的当前呼叫保持
func (self *Ext) Hold() {

}

//让分机的当前呼叫解除保持
func (self *Ext) UnHold() {

}

//让分机的话筒静音/解除静音
