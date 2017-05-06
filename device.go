package om

import (
	"encoding/xml"
	"log"
)

//查询设备信息
type DeviceInfo struct {
	XMLName      xml.Name                  `xml:"DeviceInfo"`
	Manufacturer string                    `xml:"manufacturer,omitempty"` //生产商
	Model        string                    `xml:"model,omitempty"`        //硬件版本
	Version      string                    `xml:"version,omitempty"`      //软件版本
	Mac          string                    `xml:"mac,omitempty"`          //物理地址
	Devices      `xml:"devices,omitempty"` //设备详情
}

//设备详情
type Devices struct {
	XMLName xml.Name `xml:"devices"`
	Exts    []*Ext   `xml:"exts,omitempty"`  //分机
	Trunks  []*Trunk `xml:"line,omitempty"`  //中继
	Groups  []*Group `xml:"group,omitempty"` //分机组
}

//查询设备信息
func (self *DeviceInfo) Query() {
	queryString := `<?xml version="1.0" encoding="utf-8" ?>
			<Control attribute="Query">
			<DeviceInfo/>
			</Control>`
	log.Println(queryString)
}
