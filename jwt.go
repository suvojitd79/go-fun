package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type UserD struct {
	User string	`json:"user"`
	Email string `json:"email"`
}

type Auth struct {
	UserD
	jwt.StandardClaims
}


var UserData  = make(map[string]UserD)


func signup(res http.ResponseWriter, req *http.Request){

	//var user UserD
	//
	//err := json.NewDecoder(req.Body).Decode(&user)
	//
	//if err != nil || user.Email == "" || user.User == ""{
	//	res.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//UserData[user.Email] = user

	ti := time.Now().Add(5 * time.Hour)

	claim := &Auth{
		UserD: UserD{
			User: "ss",
			Email: "suv",
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ti.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	jwtToken, err := token.SignedString([]byte("ABC"))

	if err != nil{
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(res, &http.Cookie{
		Name: "AT",
		Value: jwtToken,
		Expires: ti,
		Path: "/",
	})
	res.Write([]byte("Auth Token has been set.."))
}

func search(res http.ResponseWriter, req *http.Request){

	defer func() {
		if data := recover(); data != nil{
			log.Println(data)
		}
	}()


	var user Auth

	query := mux.Vars(req)["query"]

	cook, e := req.Cookie("AT")

	if e != nil{
		log.Println(e.Error())
		res.Write([]byte(e.Error()))
		return
	}

	token, e := jwt.ParseWithClaims(cook.Value, &user,func(token *jwt.Token) (interface{}, error) {
		return []byte("ABC"), nil
	})

	if e != nil{
		log.Println(e.Error())
		return
	}

	if token.Valid {
		res.Write([]byte(fmt.Sprintf("Welcome, %s, your query is %s", user.User, query)))
	}
}
func Logout(res http.ResponseWriter, req *http.Request){

	http.SetCookie(res, &http.Cookie{
		Name: "AT",
		MaxAge: -1,
	})

	http.Redirect(res, req, "https://www.google.com", http.StatusTemporaryRedirect)
}

func Handle404(res http.ResponseWriter, req *http.Request){
	log.Println(req.URL)
}

func Middle1(handler http.Handler)  http.Handler{
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("log-1")
		handler.ServeHTTP(writer, request)
	})
}

func main(){

	router := mux.NewRouter()
	router.Use(Middle1, Middle1)
	router.HandleFunc("/auth/check", signup).Methods("GET")
	router.HandleFunc("/suggest/{query:[0-9]+}", search)
	router.HandleFunc("/auth/out", Logout)
	router.NotFoundHandler = http.HandlerFunc(Handle404)


	log.Fatal(http.ListenAndServe(":9000", router))

}
