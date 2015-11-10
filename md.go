package main

import (
	"html/template"
	"log"
	"net/http"
)
var (
	view_Dir="./views"
)
var 引擎=template.New("引擎")
func initTmpl(){
	引擎, _ = template.ParseFiles(view_Dir+"/"+"header.tmpl",view_Dir+"/"+ "writingspace.tmpl", view_Dir+"/"+"footer.tmpl")
}
func writingspaceHandler(w http.ResponseWriter,r *http.Request){
	switch r.Method{
		case "GET":
			log.Println("GET")
			if err:=引擎.ExecuteTemplate(w,"writingspace",nil);err!=nil{
				log.Fatalln(err.Error())
			}
		case "POST":
					
		}
}
func main(){
	initTmpl()
	var mux=http.NewServeMux()
	
	mux.HandleFunc("/writingspace",writingspaceHandler)
	err:=http.ListenAndServe(":8090",mux)
	log.Println("http.ListenAndServe(:8090)")
	if err!=nil{
		log.Fatal("http.ListenAndServe:",err.Error())
	}
}