package om

import (
	"context"
	"encoding/xml"
	"gopkg.in/macaron.v1"
	"io/ioutil"
	oLog "log"
	"net/http"
	"os"
)

var logError *oLog.Logger = oLog.New(os.Stdout, "[OM Handler Error] ", oLog.Ldate|oLog.Lshortfile)
var logDebug *oLog.Logger = oLog.New(os.Stdout, "[OM Handler Debug] ", oLog.Ldate|oLog.Lshortfile)

//对OM的请求进行处理
func OmServiceHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		request, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logError.Println(err)
			w.WriteHeader(400)
			return
		}
		logDebug.Println("Request:", string(request))
		reqBody := new(Event)
		err = xml.Unmarshal(request, reqBody)
		if err != nil {
			logError.Println(err)
			w.WriteHeader(400)
			return
		}
		logDebug.Println("xml:", *reqBody)

		ctx := context.WithValue(req.Context(), "omEvent", reqBody)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func OmMacaronHandler(ctx *macaron.Context) {
	logDebug.Println("in handler")
	request, err := ctx.Req.Body().Bytes()
	if err != nil {
		logError.Println(err)
		ctx.Status(400)
		return
	}
	logDebug.Println("Request:", string(request))
	reqBody := new(Event)
	err = xml.Unmarshal(request, reqBody)
	if err != nil {
		logError.Println(err)
		ctx.Status(400)
		return
	}
	logDebug.Println("xml:", *reqBody)
	ctx.Data["omEvent"] = reqBody

}
