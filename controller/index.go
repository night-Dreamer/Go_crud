package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/solenovex/web-tuborial/common"
	"github.com/solenovex/web-tuborial/funcs"
	"github.com/solenovex/web-tuborial/model"
)

func registerIndexRoutes() {
	http.HandleFunc("/index", Lists)
	http.HandleFunc("/add", add_list)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/delete/", delete)
}

//查询所有记录
func Lists(w http.ResponseWriter, r *http.Request) {
	data, err := model.Get_all()
	if err != nil {
		log.Fatalln(err)
	}
	funcMap := template.FuncMap{"add": funcs.Add}
	t := template.New("companies").Funcs(funcMap)
	t, err = t.ParseFiles("templates/index.html", "templates/list.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "lyout", data)
}

//添加一条记录
func add_list(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles("templates/index.html", "templates/add.html")
		if err != nil {
			fmt.Println(err)
		}
		t.ExecuteTemplate(w, "lyout", nil)

	case http.MethodPost:
		record := common.Record{}
		record.Id = r.FormValue("id")
		record.Name = r.FormValue("name")
		record.Position = r.FormValue("position")
		err := model.Add_one(record)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalln(err)
		} else {
			http.Redirect(w, r, "/index", http.StatusFound)
		}
	}
}

//编辑记录
func edit(w http.ResponseWriter, r *http.Request) {
	var id string
	switch r.Method {

	case http.MethodGet:
		query := r.URL.Query()
		id = query.Get("id")
		list := model.Get_one(id)
		t, err := template.ParseFiles("templates/index.html", "templates/edit.html")
		if err != nil {
			log.Fatalln(err)
		}
		t.ExecuteTemplate(w, "lyout", list)

	case http.MethodPost:
		record := common.Record{}
		record.Id = r.FormValue("id")
		record.Name = r.FormValue("name")
		record.Position = r.FormValue("position")
		fmt.Println(record)
		err := model.Edit(record)
		if err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

//删除一条记录
func delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	Id := query.Get("id")
	err := model.Delete(Id)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	http.Redirect(w, r, "/index", http.StatusFound)
}
