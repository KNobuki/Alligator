package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func showTop(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "TopPage!\n")
}

func showSignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "SignUp!\n")
}

func showLogIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Login!\n")
}

func showMyPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sのページ\n", p.ByName("userid"))
}

func registUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sを登録しました!\n", p.ByName("userid"))
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sを削除しました!\n", p.ByName("userid"))
}

func updateUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sの情報を更新しました!\n", p.ByName("userid"))
}

func showItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sの%sのページ\n", p.ByName("userid"), p.ByName("itemid"))
}

func uploadItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sが%sを登録しました\n", p.ByName("userid"), p.ByName("itemid"))
}

func deleteItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sが%sを削除しました\n", p.ByName("userid"), p.ByName("itemid"))
}

func updateItemInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sの%sのページ\n", p.ByName("userid"), p.ByName("itemid"))
}

func loginUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%sがログインしました\n", p.ByName("userid"))
}

func main() {

	mux := httprouter.New()

	mux.GET("/", showTop)
	mux.GET("/signup", showSignUp)
	mux.GET("/login", showLogIn)
	mux.GET("/UserPage/:userid", showMyPage)
	mux.POST("/:userid", registUser)
	mux.DELETE("/:userid", deleteUser)
	mux.PUT("/:userid/", updateUserInfo)
	mux.GET("/UserPage/:userid/:itemid", showItem)
	mux.POST("/:userid/:itemid", uploadItem)
	mux.DELETE("/:userid/:itemid", deleteItem)
	mux.PUT("/:userid/:itemid", updateItemInfo)
	mux.POST("/login/:userid", loginUser)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
