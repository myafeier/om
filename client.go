package om

import (
	"fmt"
	//"io"
	"log"

	"bufio"
	"bytes"
	"io"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Client struct {
	debug  *log.Logger
	error  *log.Logger
	Ip     string
	Port   string
	Prefix string
	c      net.Conn
	sync.RWMutex
}

func NewOmClient(ip, port, uri string) (httpClient *Client, err error) {

	httpClient = &Client{
		Ip:     ip,
		Port:   port,
		Prefix: uri,
	}
	httpClient.debug = log.New(os.Stdout, "[OM Debug] ", log.Lshortfile|log.Ltime)
	httpClient.error = log.New(os.Stdout, "[OM Error] ", log.Lshortfile|log.Ltime)
	err = httpClient.Conn()
	return
}
func (self *Client) Conn() (err error) {
	self.Lock()
	defer self.Unlock()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", OMIp+":"+OMPort)
	if err != nil {
		self.error.Println("address:", err)
		return
	}
	self.c, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		self.error.Println("dialTcp:", err)
		return
	}
	return

}

func (self *Client) PostXML(postString string) (code int, Response []byte, err error) {
	if self.c == nil {
		err = self.Conn()
		if err != nil {
			return
		}
	}

	self.RLock()
	defer self.RUnlock()

	request := fmt.Sprintf("POST "+self.Prefix+" HTTP/1.0\nContent-Type:text/html\nContent-Length: %d\n\n%s", len(postString), postString)
	_, err = self.c.Write([]byte(request))
	if err != nil {
		self.error.Println("Write:", err)
	}

	ResponseBuffer := new(bytes.Buffer)
	var start bool

	buf := bufio.NewReader(self.c)

	for {
		var line []byte
		line, _, err = buf.ReadLine()
		if err != nil && err != io.EOF {
			self.error.Println("error:", err)
			return
		}
		if err == io.EOF {
			err = nil
			break
		}

		if strings.Contains(string(line), "HTTP/1.") {

			reg := regexp.MustCompile(`\d{3}`)
			bf := bytes.NewBuffer(reg.Find(line)).String()
			code, err = strconv.Atoi(bf)
			if err != nil {
				self.error.Println("Error found when convert to number!", err)
				return
			}
		}
		if strings.Contains(string(line), "<?xml") {
			start = true
		}
		if start {
			_, err = ResponseBuffer.Write(line)
			if err != nil {
				self.error.Println("error:", err)
				return
			}
		}

	}
	log.Println(ResponseBuffer.String())
	Response = ResponseBuffer.Bytes()
	return
}

//注意关闭资源
func (self *Client) Close() {
	self.c.Close()
}
