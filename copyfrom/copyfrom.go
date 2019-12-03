package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/eturella/go-parser-test/copyfrom/redshift"
)

func main() {
	// sql := `UPDATE users SET name = "bob", email = "email@email.com" WHERE id = 12 AND user_type >= 2 LIMIT 10;`
	// tokens := []lex.Token{
	// 	tv(lex.TokenUpdate, "UPDATE"),
	// 	tv(lex.TokenTable, "users"),
	// }
	// l := lex.NewSqlLexer(sql)
	// fmt.Printf("sql: %v\n", sql)
	// for _, goodToken := range tokens {
	// 	tok := l.NextToken()
	// 	fmt.Printf("got:%v  want:%v \n", tok, goodToken)
	// }

	// sql := `COPY users from 's3://mybucket/data/listing/' iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' delimiter '\t' ;`
	// // tokens := []lex.Token{
	// // 	tv(lex.TokenCopy, "COPY"),
	// // 	tv(lex.TokenTable, "users"),
	// // 	tv(lex.TokenFrom, "FROM"),
	// // 	tv(lex.TokenValue, "'s3://mybucket/data/listing/'"),
	// // }
	// l := redshift.NewRedshiftLexer(sql)
	// fmt.Printf("sql: %v\n", sql)
	// // for _, goodToken := range tokens {
	// // 	tok := l.NextToken()
	// // 	fmt.Printf("got:%v  want:%v \n", tok, goodToken)
	// // }
	// tok := l.NextToken()
	// for tok.T != lex.TokenEOF {
	// 	fmt.Printf("got:%v  \n", tok)
	// 	if tok.T == lex.TokenError {
	// 		fmt.Printf("%v", tok.Err(l))
	// 	}
	// 	tok = l.NextToken()
	// }

	tutti := true
	singolo := -1
	if len(os.Args) >= 2 {
		tutti = false
		singolo, _ = strconv.Atoi(os.Args[1])
	}

	p := "copyfrom.sql"
	p, _ = filepath.Abs(p)
	b, err := ioutil.ReadFile(p) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	tmp := string(b)
	examples := strings.Split(tmp, ";")
	fmt.Printf("%d esempi", len(examples))

	for i, sql := range examples {
		sql = strings.TrimSpace(sql)
		if len(sql) > 0 && (tutti || singolo == i+1) {
			fmt.Printf("\n\nn. esempio: %d\n", i+1)
			fmt.Printf("sql: %v\n\n", sql)
			ok := redshift.ParseCopy(sql)
			if ok {
				fmt.Println("OK")
			} else {
				fmt.Println("ERRORE")
			}
		}
	}

}

// func tv(t lex.TokenType, v string) lex.Token {
// 	return lex.Token{T: t, V: v}
// }
