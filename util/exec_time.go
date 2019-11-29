package util

import (
	"fmt"
	"time"
)

// ExecTime ...
type ExecTime struct {
	// Lib ...
	Lib string
	// QNum ...
	QNum int
	// Ok ...
	Ok bool
	// Millisec ...
	Millisec int64
	// Msg
	Msg string
}

// MakeTimestamp in millisecondi
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / 1000
}

// Sintesi ....
func Sintesi(t *[]ExecTime) ExecTime {
	var ok int
	var tot int64
	ok, tot = 0, 0
	n := ""
	for _, v := range *t {
		n = v.Lib
		if v.Ok {
			ok++
		}
		tot = tot + v.Millisec
	}
	return ExecTime{
		Lib:      n,
		QNum:     -1,
		Millisec: tot,
		Msg:      fmt.Sprintf("%d / %d , %f", ok, len(*t), float64(ok/len(*t)*100)),
	}
}
