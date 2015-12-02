package main

import (
	"time"
	"MongoHelper"
	"StringHelper"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	root    = ""
	viewDir = "/views"
)
var 引擎 = template.New("引擎")

func initTmpl() {
	引擎, _ = template.ParseFiles(root+"/"+viewDir+"/"+"header.htm", root+"/"+viewDir+"/"+"writingspace.htm", root+"/"+viewDir+"/"+"footer.htm")
}
func writingspaceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("GET")
		if err := 引擎.ExecuteTemplate(w, "writingspace", nil); err != nil {
			log.Fatalln(err.Error())
		}
	case "POST":
		data, _ := ioutil.ReadAll(r.Body)
		datastr := string(data)
		log.Printf(datastr)
		datastr="{\"createtime\":\""+StringHelper.Substr(time.Now().String(),0,19)+"\",\"text\":"+datastr+"}"
		log.Println(datastr)
		MongoHelper.Insert("test", "article",datastr)
		w.Write([]byte(""))
	}
}
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix,
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			file := root + r.URL.Path
			log.Println(file)
			http.ServeFile(w, r, file)
		})
}
func main() {
	//检查根目录
	root, _ = os.Getwd()
	initTmpl()
	var mux = http.NewServeMux()
	staticDirHandler(mux, "/assets/", root+"/assets", 0)
	mux.HandleFunc("/writingspace", writingspaceHandler)
	err := http.ListenAndServe(":8090", mux)
	log.Println("http.ListenAndServe(:8090)")
	if err != nil {
		log.Fatal("http.ListenAndServe:", err.Error())
	}
}
