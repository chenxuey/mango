package mango_logger

import (
	"mango/config"
	"sync"
	"gitlab.gaodun.com/golib/graylog"
)

func AddLog(p []byte, mlog *config.Logger) {
	AddGrayLog(string(p), map[string]interface{}{"x-response-id": mlog.Info})
}


var p sync.Mutex

//AddGrayLog 添加日志
func AddGrayLog(logInfo ...interface{}) {
	var logStr string
	m := make(map[string]interface{})
	if len(logInfo) == 1 {
		logStr = logInfo[0].(string)
	} else {
		logStr = logInfo[0].(string)
		m = logInfo[1].(map[string]interface{})
	}
	p.Lock()
	m["item"] = "test-mango"
	p.Unlock()
	graylog.GdLog(logStr, m)
}