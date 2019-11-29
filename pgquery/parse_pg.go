package pgquery

import (
	"encoding/json"
	"fmt"

	"github.com/eturella/go-parser-test/util"
	pg_query "github.com/lfittl/pg_query_go"
)

// Exec ....
func Exec(qs []string, debug bool) []util.ExecTime {
	var r []util.ExecTime
	for i, q := range qs {
		fmt.Printf("==> %s\n", q)
		in := util.MakeTimestamp()
		tree, err := pg_query.Parse(q)
		out := util.MakeTimestamp()
		if err != nil {
			fmt.Printf("KO ==> %v\n", err)
		} else if debug {
			j, _ := json.MarshalIndent(tree, "", "    ")
			fmt.Printf("OK => %s\n", j)
		}
		et := util.ExecTime{
			QNum:     i,
			Lib:      "lfittl/pg_query_go",
			Ok:       (err == nil),
			Millisec: out - in,
			Msg:      fmt.Sprintf("%v", err),
		}
		r = append(r, et)
	}
	return r

}
