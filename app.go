package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"
)

var brands = []string{
	"google",
	"microsoft",
	"bloomreach",
	"apple",
	"facebook",
}

var people = []string{
	"suvojit.das",
	"suvojit.singh",
}

var extensions = []string{
	"in",
	"com",
	"ca",
	"fr",
	"de",
}

type UserName struct {
	User string	`json:"user"`
	Email string `json:"email"`
}

func GetSearch(res http.ResponseWriter, req *http.Request){

	res.Header()

	for k,v := range req.Header {
		fmt.Println(k, v)
	}
	d, _ := json.Marshal(UserName{})
	res.Write(d)
}

type r struct{
	PageTitle string
	Data []UserName
}


func Home(res http.ResponseWriter, req *http.Request){

	num,_ := strconv.ParseInt(req.URL.Query().Get("num"),10, 64)

	title := req.URL.Query().Get("title")



	if token := req.Header.Get("Cookie"); true{
		fmt.Println(req.Cookie("sample"), utf8.RuneCountInString(token))
	}


	data := r{
		PageTitle: title,
		Data: []UserName{},
	}

	for i:=int64(0);i<num;i+=1{

		p := people[rand.Intn(len(people))]
		b := brands[rand.Intn(len(brands))]
		e := extensions[rand.Intn(len(extensions))]

		email := fmt.Sprintf("%s@%s.%s", p, b, e)
		data.Data = append(data.Data, UserName{User: p, Email: email})
	}


	http.SetCookie(res, &http.Cookie{
		Name: "sample",
		Value: "120",
		Expires: time.Now().Add(5*time.Second),
	})


	tem,_ := template.ParseFiles("home.html")
	tem.Execute(res, data)

}

func main(){

	http.HandleFunc("/", Home)

	log.Fatalln(http.ListenAndServe(":9000", nil))
}