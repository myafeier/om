package om

import "encoding/xml"

//中继
type Trunk struct {
	XMLName xml.Name `xml:"trunk"`
	Id      int      `xml:"id,attr"` //中继的线路编号，由OM分配的固定值，作为对中继进行配置的唯一依据。
	//TrunkId string `xml:"trunkid"`//中继号，即，中继的电话号码，用户可自定义，作为对中继进行查询的唯一依据。
	LineId  string   `xml:"lineid"`  //中继的线路编号，是中继的唯一固定标识
	State   string   `xml: "state"`  //中继的线路状态 ready: 可用 active: 摘机、振铃或通话中 unwired:未接线 offline:离线 注：对于IP中继多路并发时，只要有一路空闲可用，其状态就是ready。
	Outer   *Outer   `xml:"outer"`   //去电
	Visitor *Visitor `xml:"visitor"` //来电
}

//获取中继列表（查询设备信息，即可获取所有中继的线路编号和电话号码）
func (self *Trunk) GetList() {

}

//查询中继
//该API用于查询指定中继（又称为外线）的相关信息，如：配置参数、线路状态、呼叫状态等。
func (self *Trunk) Query() {

}

//查询某条中继的状态
func (self *Trunk) GetStat() {

}
