package model

import (
	"fmt"
	"log"

	"github.com/solenovex/web-tuborial/common"
)

func Get_all() (data []common.Record, err error) {
	Sql := "SELECT * FROM staff"
	rows, err := common.Db.Query(Sql)
	if err != nil {
		log.Fatalln(err)
	}
	L := common.Record{}
	for rows.Next() {
		err = rows.Scan(&L.Id, &L.Name, &L.Position)
		data = append(data, L)
	}
	return
}

//get one
func Get_one(id string) (record common.Record) {
	Sql := "SELECT * FROM staff WHERE id = $1"
	err := common.Db.QueryRow(Sql, id).Scan(&record.Id, &record.Name, &record.Position)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//Insert
func Add_one(record common.Record) (err error) {
	Sql := "INSERT INTO staff (id, name, position) VALUES ($1, $2, $3)"
	stmt, err := common.Db.Prepare(Sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(record.Id, record.Name, record.Position)
	if err != nil {
		return
	}
	return
}

//Edit
func Edit(record common.Record) (err error) {
	Sql := "UPDATE staff SET name=$1, position=$2 WHERE id=$3"
	_, err = common.Db.Exec(Sql, record.Name, record.Position, record.Id)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//Delete
func Delete(id string) (err error) {
	Sql := "DELETE FROM staff WHERE id=$1"
	_, err = common.Db.Exec(Sql, id)
	return
}
