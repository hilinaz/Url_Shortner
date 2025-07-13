package main

import (
	"Url_short/db"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	db.CreateTable()

	router := mux.NewRouter()
	router.HandleFunc("/", shortenUrl).Methods("POST")
	router.HandleFunc("/{url}", redirectUrl).Methods("GET")
	http.ListenAndServe(":5500", router)
	log.Println("Listening on http://localhost:5500")
}

func shortenUrl( w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	err:= r.ParseForm()
	if err!=nil{
		http.Error(w,"not found",http.StatusBadRequest)
		
	}
	url:=r.FormValue("Url")

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	value := make([]byte, 6)
	for ind := range value {
		value[ind] = charset[rand.Intn(len(charset))]
	}
	shorturl:=string(value)
	UrlData:= `INSERT INTO UrlPaths (OriginalUrl,ShortUrl) VALUES (?,?)`
	_,err = db.DB.Exec(UrlData,url,shorturl)
	if err!=nil{
		http.Error(w, "Failed to save shortened URL", http.StatusInternalServerError)
	}
	w.Write([]byte(shorturl))
	
}

func redirectUrl(w http.ResponseWriter,r *http.Request)  {
	vars:=mux.Vars(r)
	shorturl:=vars["url"]
	var originalUrl string
	
	if shorturl==""{
		http.Error(w,"not found",http.StatusBadRequest)
	}
	
	query:= `SELECT OriginalUrl  FROM UrlPaths WHERE ShortUrl = ?`

	err:=db.DB.QueryRow(query,shorturl).Scan(&originalUrl)
	if err!=nil{
		http.Error(w,"url not found",http.StatusNotFound)
	}
	
	http.Redirect(w,r,originalUrl,http.StatusFound)

}


