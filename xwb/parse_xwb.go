package xwb

import (
	"encoding/json"
	"fmt"

	"github.com/eturella/go-parser-test/util"
	"github.com/xwb1989/sqlparser"
)

// Exec ....
func Exec(qs []string, debug bool) []util.ExecTime {
	var r []util.ExecTime
	for i, q := range qs {
		fmt.Printf("==> %s\n", q)

		in := util.MakeTimestamp()
		stmt, err := sqlparser.Parse(q)
		out := util.MakeTimestamp()
		if err != nil {
			fmt.Printf("KO ==> %v\n", err)
		} else {
			j, _ := json.MarshalIndent(stmt, "", "    ")
			fmt.Printf("OK => %s\n", j)
		}

		et := util.ExecTime{
			QNum:     i,
			Lib:      "xwb1989/sqlparser",
			Ok:       (err == nil),
			Millisec: out - in,
			Msg:      fmt.Sprintf("%v", err),
		}
		r = append(r, et)
	}
	return r

}
