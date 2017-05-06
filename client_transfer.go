//呼叫转接命令（包含点击拨号功能）(Transfer)
/*简介
此类API用于把OM上的两个通话对象（如来电与分机）进行呼叫连接，从而使两者最终能够建立通话。
可操作对象及规则说明
OM包含的通话对象一共有6种，分别为：来电（visitor）、去电(outer)、分机(ext)、分机组(group)、语音菜单(menu)。大多情况下，这6种对象之间可以进行相互连接。*/
package om

import (
	"encoding/xml"
	"log"
)

//分机呼分机
//该API用于将一个分机呼叫另一个分机，从而使两者能够建立通话。
//分机呼叫分机成功只返回statusCode,否则会返回错误
// 错误演示：200 &{{ Event} FAILED [] []} <nil>
func (self *Client) TransferFromExtToExt(FromExtId, ToExtId string) (code int, result *Response, err error) {

	queryString := `<?xml version="1.0" encoding="utf-8" ?>
	<Transfer attribute="Connect">
  	<ext id="` + FromExtId + `"/>
	<ext id="` + ToExtId + `"/>
	</Transfer>`

	code, returnInterface, err := self.PostXML(queryString)
	if err != nil {
		self.error.Println(err)
		return

	}
	if len(returnInterface) > 1 {
		result = new(Response)
		err = xml.Unmarshal(returnInterface, result)
		if err != nil {
			self.error.Println(err)
			return
		}
	}
	self.debug.Printf("%d,%+v\n", code, result)

	return

}

//分机呼外部电话
//该API用于让分机向外部电话发起呼叫，从而使两者能够建立通话。
//注：分机呼分机请勿使用该API。
func (self *Client) TransferFromExtToOuter(fromExtId string, toOuterPhone string) (code int, result *Response, err error) {

	queryString := `<?xml version="1.0" encoding="utf-8" ?>
		<Transfer attribute="Connect">
  		<ext id="` + fromExtId + `"/>
		<outer to="` + OutConnectPrefix + "," + toOuterPhone + `"/>
		</Transfer>`
	code, returnInterface, err := self.PostXML(queryString)
	if err != nil {
		self.error.Println(err)
	}
	result = new(Response)
	err = xml.Unmarshal(returnInterface, result)
	if err != nil {
		self.error.Println(err)
		return
	}

	self.debug.Printf("%d,%+v\n", code, result)
	return

}

//给分机播放指定menu
func (self *Client) MenuTransferExt(toExt, MenuId string) {
	queryString := `<?xml version="1.0" encoding="utf-8" ?> <Transfer attribute="Connect"> <menu id="` + MenuId + `"/> <ext id="` + toExt + `"/> </Transfer>`
	code, returnInterface, err := self.PostXML(queryString)
	if err != nil {
		self.error.Println("Error_1:")
		self.error.Println(err)
	}
	log.Printf("%d,%v\n", code, returnInterface)
}
