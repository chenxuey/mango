package config

import (
	"io"
	"log"
	"os"
)

type actionLog func(p []byte, mlog *Logger)

type Logger struct {
	Debug  bool
	out    io.Writer
	addLog actionLog
	*log.Logger
	Info interface{}
	Remark interface{}
}

func (ml *Logger) Write(p []byte) (int, error) {
	if ml.Debug {
		return ml.out.Write(p)
	}

	ml.addLog(p, ml)
	return 0, nil
}

func (ml *Logger) SetMangoOutput(f actionLog) {
	ml.addLog = f
}


func (ml *Logger) SetMangoDebug(f bool) {
	ml.Debug = f
}

func NewLogger(debug bool, f actionLog) *Logger {
	l := log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
	ml := Logger{out: os.Stderr, Debug: debug}
	ml.SetMangoOutput(f)
	l.SetOutput(&ml)
	ml.Logger = l
	return &ml
}

func NewSQL(debug bool, f actionLog) Logger {
	nl := Logger{out: os.Stderr, Debug: debug}
	nl.SetMangoOutput(f)
	return nl
}


