package om

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	oLog "log"
	"net/http"
	"os"
)

var logError *oLog.Logger = oLog.New(os.Stdout, "[OM Handler Error] ", oLog.Ldate|oLog.Lshortfile)

//对OM的请求进行处理
func OmServiceHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		request, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logError.Println(err)
			w.WriteHeader(400)
			return
		}
		reqBody := new(Event)
		err = xml.Unmarshal(request, reqBody)
		if err != nil {
			logError.Println(err)
			w.WriteHeader(400)
			return
		}
		ctx := context.WithValue(req.Context(), "omEvent", reqBody)
		next.ServeHTTP(w, req.WithContext(ctx))
	})

}
