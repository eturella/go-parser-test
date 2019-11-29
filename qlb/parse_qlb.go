package qlb

import (
	"encoding/json"
	"fmt"

	td "github.com/araddon/qlbridge/datasource/mockcsvtestdata"
	"github.com/araddon/qlbridge/rel"
	"github.com/eturella/go-parser-test/util"
)

// Exec ....
func Exec(qs []string, debug bool) []util.ExecTime {
	// exec.RegisterSqlDriver()
	// exec.DisableRecover()
	td.SetContextToMockCsv()

	var r []util.ExecTime
	for i, q := range qs {
		fmt.Printf("==> %s\n", q)

		in := util.MakeTimestamp()

		//ctx := td.TestContext(q)
		//ctx := plan.NewContext(q)
		//job, err := exec.BuildSqlJob(ctx)

		//ctx := td.TestContext(q)
		stmt, err := rel.ParseSql(q)

		out := util.MakeTimestamp()
		if err != nil {
			fmt.Printf("KO ==> %v\n", err)
		} else {
			j, _ := json.MarshalIndent(stmt, "", "    ")
			fmt.Printf("OK => %s\n", j)
		}

		et := util.ExecTime{
			QNum:     i,
			Lib:      "araddon/qlbridge",
			Ok:       (err == nil),
			Millisec: out - in,
			Msg:      fmt.Sprintf("%v", err),
		}
		r = append(r, et)
	}
	return r

}
