package om

const (
	MenuFringWaitTransfer = 1 //自定义彩铃,正在接通顾客电话，请勿挂机
)

var (
	OutConnectPrefix = "9"    //外呼前缀
	OM_URI           = "/xml" //OM接收API请求的URI
	OMIp             = "192.168.1.21"
	OMPort           = "9080"
	//OMUrl  = "https://crm"
	//OMPort = "8443"
)

//呼叫权限
var CallRestrictionMap = map[int]string{0: "内线", 1: "市话", 2: "国内", 3: "国际"}

//呼叫转移方式
var FwdTypeMap = map[int]string{0: "关闭", 1: "全转", 2: "遇忙或无应答转"}

//API的功能开关
var ApiStatMap = map[int]string{0: "关闭API状态监控", 7: "开启API状态监控"}

//线路状态
var LineStatMap = map[string]string{"Ready": "空闲可用", "Active": "振铃、回铃或通话中", "Progress": "模拟分机摘机后等待拨号以及拨号过程中", "Offline": "IP分机离线", "Offhook": "模拟分机听催挂音时的状态"}

//中继状态
var TrunkStatMap = map[string]string{"Ready": "空闲可用", "Active": "摘机、振铃或通话中", "unwired": "未接线", "Offline": "离线"}

//分机组来电分配规则
var DistributionMap = map[string]string{"sequential": "顺选", "circular": "轮选", "group": "群振"}

//通话状态
var CallStatMap = map[string]string{"Talk": "通话进行中", "Progress": "呼叫处理过程中", "Wait": "呼叫等待中"}

//呼叫类型
var CdrStatMap = map[string]string{"IN": "打入", "OU": "打出", "FI": "呼叫转移入", "FW": "呼叫转移出", "LO": "内部通话", "CB": "双向外呼"}

//呼叫路由类型
var CdrRouteMap = map[string]string{"IP": "IP中继", "XO": "模拟中继", "IC": "内部", "OP": "总机"}

var OnOffMap = map[string]string{"on": "开启", "off": "关闭"}
var YesNoMap = map[string]string{"yes": "允许", "no": "不允许"}
