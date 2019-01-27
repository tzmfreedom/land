package main

import (
	"html/template"
	"net/http"
)

func main() {
	//file := "../sobjects.yml"
	//loader := builtin.NewMetaFileLoader(file)
	//sobjects, err := loader.Load()
	//if err != nil {
	//	panic(err)
	//}
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	uri := r.RequestURI
	//	paths := strings.Split(uri[1:], "/")
	//	// /Account GET, POST
	//	if len(paths) == 1 {
	//		switch r.Method {
	//		case http.MethodGet:
	//			objects := SearchObjects(paths[0])
	//			Render("index", w, objects)
	//		case http.MethodPost:
	//			id := CreateObject(paths[0])
	//			http.Redirect(w, r, fmt.Sprintf("/%s/%s", paths[0], id), http.StatusFound)
	//			return
	//		}
	//	}
	//	// /Account/{id} GET, PUT, DELETE
	//	// /Account/new GET
	//	if len(paths) == 2 {
	//		if paths[1] == "new" && r.Method == "GET" {
	//			sobj := sobjects[paths[0]]
	//			Render("show", w, sobj)
	//		}
	//		switch r.Method {
	//		case "GET":
	//			sobj := FindObject(paths[0], paths[1])
	//			Render("new", w, sobj)
	//		case "PUT":
	//			id := UpdateObject(paths[0], paths[1])
	//			http.Redirect(w, r, fmt.Sprintf("/%s/%s", paths[0], id), http.StatusFound)
	//			return
	//		case "DELETE":
	//			DeleteObject(paths[0], paths[1])
	//			http.Redirect(w, r, fmt.Sprintf("/%s", paths[0]), http.StatusFound)
	//			return
	//		}
	//	}
	//	// /Account/{id}/edit GET
	//	if len(paths) == 3 && paths[2] == "edit" {
	//		FindObject(paths[0], paths[1])
	//		Render(paths[1], w, nil)
	//	}
	//	w.WriteHeader(404)
	//})
	//http.ListenAndServe(":8080", nil)
}

func Render(name string, w http.ResponseWriter, bind interface{}) {
	t, err := template.ParseFiles(name)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, bind)
	if err != nil {
		panic(err)
	}
}

func FindObject(name, id string) {

}

func SearchObjects(name, id string) {

}

func CreateObject(name, id string) {

}

func UpdateObject(name, id string, data map[string]string) {

}

func DeleteObject(name, id string) {

}
