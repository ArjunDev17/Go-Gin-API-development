package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type logFormateLocal struct {
	TimeStamp    time.Time
	StatusCode   int
	Latency      time.Duration
	ClientIP     string
	Method       string
	Path         string
	ErrorMessage string
	RequestProto string
}

func FormatLogsJSON(param gin.LogFormatterParams) string {
	params := &logFormateLocal{
		ClientIP:     param.ClientIP,
		TimeStamp:    param.TimeStamp,
		Method:       param.Method,
		Path:         param.Path,
		RequestProto: param.Request.Proto,
		StatusCode:   param.StatusCode,
		Latency:      param.Latency,
		ErrorMessage: param.ErrorMessage,
	}
	j, _ := json.Marshal(params)
	fmt.Println(string(j))
	return string(j)
}
func FormatLogs(param gin.LogFormatterParams) string {
	return fmt.Sprintf("{%s-[%s]\"%s %s %s %d %s\"%s\"%s\"}\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
