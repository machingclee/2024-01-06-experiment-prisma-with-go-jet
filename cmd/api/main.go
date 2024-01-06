package main

import (
	"authentication/.gen/udemy/public/model"
	. "authentication/.gen/udemy/public/table"
	"authentication/pkg/pgsql"
	"encoding/json"
	"fmt"
	"log"

	. "github.com/go-jet/jet/v2/postgres"
)

func main() {
	db := pgsql.NewDB()
	statement := SELECT(Student.AllColumns).FROM(Student).WHERE(Student.ID.EQ(Int(1)))
	dest := model.Student{}
	err := statement.Query(db, &dest)
	if err != nil {
		log.Fatal(err)
	}
	jsonText, _ := json.MarshalIndent(dest, "", "\t")
	fmt.Println(string(jsonText))
}
