package myfirst

import(
	"github.com/1backend/go-client"
)

var Token string

type Test struct {
	Name	string
	Foods	[]string
}


func GetImportedHi() error {
	var ret 
	return ret, goclient.New(Token).Call("dobi", "myfirst", "GET", "/imported-hi", map[string]interface{}{  }, &ret)
}

func GetHi() error {
	var ret 
	return ret, goclient.New(Token).Call("dobi", "myfirst", "GET", "/hi", map[string]interface{}{  }, &ret)
}

func GetSqlExample() error {
	var ret 
	return ret, goclient.New(Token).Call("dobi", "myfirst", "GET", "/sql-example", map[string]interface{}{  }, &ret)
}

