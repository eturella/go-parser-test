package redshift

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/araddon/qlbridge/lex"
)

func TestCopyFrom(t *testing.T) {

	p := "../../copyfrom.sql"
	p, err := filepath.Abs(p)
	if err != nil {
		t.Errorf("il file dei test non esiste")
		t.FailNow()
	}
	b, err := ioutil.ReadFile(p) // just pass the file name
	if err != nil {
		t.Errorf("il file dei test non è leggibile: %s", p)
		t.FailNow()
	}
	tmp := string(b)
	examples := strings.Split(tmp, ";")
	fmt.Printf("%d esempi", len(examples))

	if len(examples) <= 0 {
		t.Errorf("i test non sono definiti")
		t.FailNow()
	}

	for i, sql := range examples {
		sql = strings.TrimSpace(sql)
		sql = strings.Replace(sql, "json", "myjson", -1)
		// sql = strings.Replace(sql, "format as", "format", 1)
		// sql = strings.Replace(sql, "delimiter as", "delimiter", 1)
		if len(sql) > 0 {
			l := NewRedshiftLexer(sql)
			fmt.Printf("\n\nn. esempio: %d\n", i+1)
			fmt.Printf("sql: %v\n\n", sql)
			// for _, goodToken := range tokens {
			// 	tok := l.NextToken()
			// 	fmt.Printf("got:%v  want:%v \n", tok, goodToken)
			// }
			tok := l.NextToken()
			for tok.T != lex.TokenEOF {
				fmt.Printf("got:%v  \n", tok)
				if tok.T == lex.TokenError {
					//fmt.Println("  --> errore")
					t.Errorf("Token ERROR: %s", sql)
					t.Fail()
				}
				tok = l.NextToken()
			}
		}
	}
}
