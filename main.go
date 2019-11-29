package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/eturella/go-parser-test/pgquery"
	"github.com/eturella/go-parser-test/qlb"
	"github.com/eturella/go-parser-test/stack"
	"github.com/eturella/go-parser-test/xwb"

	"github.com/eturella/go-parser-test/util"
	"github.com/jedib0t/go-pretty/table"
)

func main() {
	p := "queries.sql"
	p, _ = filepath.Abs(p)
	qs := util.LeggiQuery(p)

	var r, t []util.ExecTime

	fmt.Println("--------------------- lfittl/pg_query_go ---------------------")
	t = pgquery.Exec(qs, false)
	r = append(r, t...)
	r = append(r, util.Sintesi(&t))

	//fmt.Println("--------------------- SRC-D ---------------------")
	// t = srcd.Exec(qs, true)
	// r = append(r, t...)
	// r = append(r, util.Sintesi(&t))

	fmt.Println("--------------------- xwb1989/sqlparser ---------------------")
	t = xwb.Exec(qs, false)
	r = append(r, t...)
	r = append(r, util.Sintesi(&t))

	fmt.Println("--------------------- araddon/qlbridge ---------------------")
	t = qlb.Exec(qs, false)
	r = append(r, t...)
	r = append(r, util.Sintesi(&t))

	fmt.Println("--------------------- xelabs/go-mysqlstack ---------------------")
	t = stack.Exec(qs, false)
	r = append(r, t...)
	r = append(r, util.Sintesi(&t))

	tab := table.NewWriter()
	tab.SetOutputMirror(os.Stdout)
	tab.AppendHeader(table.Row{"Libreria", "Test n.", "Esito", "T(ms)", "Errore"})
	for _, v := range r {
		tab.AppendRow([]interface{}{v.Lib, v.QNum, v.Ok, float64(v.Millisec) / 1000, v.Msg})
	}
	tab.Render()

}
