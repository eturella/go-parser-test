package srcd

import (
	"encoding/json"
	"fmt"
	"text/template/parse"

	"github.com/src-d/go-mysql-server/sql"
	"github.com/src-d/go-mysql-server/sql/parse"
)

// Exec ....
func Exec(qs []string, debug bool) {
	for _, q := range qs {
		fmt.Printf("==> %s\n", q)
		ctx := sql.NewEmptyContext()
		p, err := parse.Parse(ctx, q)
		if err != nil {
			fmt.Printf("KO ==> %v\n", err)
		} else if debug {
			j, _ := json.MarshalIndent(tree, "", "    ")
			fmt.Printf("OK => %s\n", j)
		}
	}

}
