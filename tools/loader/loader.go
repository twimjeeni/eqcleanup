package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"reflect"
)

var zoneid int = 48
var zonename string

var lootids []int = []int{}
var gridids []int = []int{}

func main() {
	idPtr := flag.Int("id", 48, "Zoneid to parse")
	flag.StringVar(&zonename, "name", "", "Zonename to parse")
	flag.Parse()
	if zonename == "" {
		fmt.Println("Not a valid zone name:", zonename)
		os.Exit(1)
	}
	fmt.Println("Zone:", zonename)
	zoneid = *idPtr
	db, err := sqlx.Open("mysql", fmt.Sprintf("root@tcp(127.0.0.1:3306)/eqmac?charset=utf8&parseTime=true"))
	if err != nil {
		fmt.Println("Error connecting to DB:", err.Error())
		os.Exit(1)
	}
	defer db.Close()
	generateNPCTypes(db)
	generateSpawngroups(db)
	generateLoot(db)
	generateGrids(db)
	return
}

type M map[string]interface{} // just an alias

func mapFields(x interface{}) M {
	o := make(M)
	v := reflect.ValueOf(x).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		// skip unexported fields
		if f.PkgPath != "" {
			continue
		}
		o[f.Name] = v.FieldByIndex([]int{i}).Interface()
	}
	return o
}
