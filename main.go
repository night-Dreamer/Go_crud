package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/solenovex/web-tuborial/common"
	"github.com/solenovex/web-tuborial/controller"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "company"
)

func init() {
	//连接postgreSql数据库
	connstr := fmt.Sprintf("host=%s port=%d user=%s password = %s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	common.Db, err = sql.Open("postgres", connstr)
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	err = common.Db.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Connected!")
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	controller.RegisterRoutes()
	server.ListenAndServe()

}
