package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// https://github.com/go-pg/pg/blob/master/example_model_test.go
var pgdb *pg.DB

type Test struct {
	Id        int
	Name      string
	CreatedAt time.Time `sql:"default:now()"`
}

func init() {
	pgdb = pg.Connect(&pg.Options{
		User:     "gotest",
		Password: "12345678",
		Database: "gotestdb",
	})

	var n int
	_, err := pgdb.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(n)

	// test için tablo varsa siliyor sonra tekrar oluşturuyor
	err = pgdb.DropTable((*Test)(nil), &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
	panicIf(err)

	err = pgdb.CreateTable((*Test)(nil), &orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
		Varchar:     64,
	})
	panicIf(err)

}

func main() {

	tests := []Test{
		{

			Name: "Mert",
		}, {

			Name: "Erdem",
		}, {

			Name: "Kübra",
		}}

	err := pgdb.Insert(&tests)

	panicIf(err)

	test := Test{
		Id: 1,
	}

	err = pgdb.Select(&test)
	if err != nil {
		panic(err)
	}
	fmt.Println(test)

	// update

	testu := Test{Id: 1}
	err = pgdb.Select(&testu)
	if err != nil {
		panic(err)
	}

	testu.Name = "Mert Acel"
	err = pgdb.Update(&testu)
	if err != nil {
		panic(err)
	}

	err = pgdb.Select(&testu)
	if err != nil {
		panic(err)
	}

	fmt.Println(testu)

}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
