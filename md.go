package main

import (
	"os"
	"html/template"
	"log"
	"net/http"
)
var (
	root=""
	view_Dir="/views"
)
var 引擎=template.New("引擎")
func initTmpl(){
	引擎, _ = template.ParseFiles(root+"/"+view_Dir+"/"+"header.tmpl",root+"/"+view_Dir+"/"+ "writingspace.tmpl",root+"/"+ view_Dir+"/"+"footer.tmpl")
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
func staticDirHandler(mux *http.ServeMux,prefix string,staticDir string,flags int){
	mux.HandleFunc(prefix,
		func(w http.ResponseWriter,r *http.Request){
			log.Println(r.URL.Path)
			file:=root+r.URL.Path
			log.Println(file)
			http.ServeFile(w,r,file)
		})
}
func main(){
	//检查根目录
	root,_=os.Getwd()
	initTmpl()
	var mux=http.NewServeMux()
	staticDirHandler(mux,"/assets/",root+"/assets",0)
	mux.HandleFunc("/writingspace",writingspaceHandler)
	err:=http.ListenAndServe(":8090",mux)
	log.Println("http.ListenAndServe(:8090)")
	if err!=nil{
		log.Fatal("http.ListenAndServe:",err.Error())
	}
}