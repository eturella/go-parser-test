package stack

import (
	"encoding/json"
	"fmt"

	"github.com/eturella/go-parser-test/util"
	"github.com/xelabs/go-mysqlstack/sqlparser"
)

// Exec ....
func Exec(qs []string, debug bool) []util.ExecTime {
	var r []util.ExecTime
	for i, q := range qs {
		fmt.Printf("==> %s\n", q)
		in := util.MakeTimestamp()

		sel, err := sqlparser.Parse(q)

		out := util.MakeTimestamp()
		if err != nil {
			fmt.Printf("KO ==> %v\n", err)
		} else if debug {
			tree := sel.(*sqlparser.Select)
			j, _ := json.MarshalIndent(tree, "", "    ")
			fmt.Printf("OK => %s\n", j)
		}
		et := util.ExecTime{
			QNum:     i,
			Lib:      "xelabs/go-mysqlstack",
			Ok:       (err == nil),
			Millisec: out - in,
			Msg:      fmt.Sprintf("%v", err),
		}
		r = append(r, et)
	}
	return r

}
