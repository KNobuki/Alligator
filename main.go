package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s !\n", p.ByName("id"))
}

func showTop(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
}


func main(){
	mux := httprouter.New()
	mux.GET("/hello/:id", hello)

	mux.GET("/", showTop)
	mux.GET("/signup", showSignUp)
	mux.GET("/login", showLogIn)
	mux.GET("/:userid", showMyPage)
	mux.POST("/:userid", registUser)
	mux.DELETE("/:userid", deleteUser)
	mux.PUT("/:userid/", updateUserInfo)
	mux.GET("/:userid/:itemid", showItem)
	mux.POST("/:userid/:itemid", uploadItem)
	mux.DELETE("/:userid/:itemid", deleteItem)
	mux.PUT("/:userid/:itemid", updateItemInfo)


	server := http.Server{
		Addr: "127.0.0.1:8083",
		Handler: mux,
	}
	server.ListenAndServe()
}