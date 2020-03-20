package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var users map

type User struct {
	UserID string
	// Password string
	// Mail     string
	Items []Item
}

type Item struct {
	ItemID    string
	ImagePath string
	User      *User
}

func retrieve (userid string, users []User) (User, error) {
	return users[userid], nil
}

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
	user := retrieve()
	fmt.Fprintf(w, "hello, %s !\n", p.ByName("userid"))
}

func showItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %sの%s !\n", p.ByName("userid"), p.ByName("itemid"))
}

func main() {

	users = make(map[string] User)
	user := User{}
	item := Item{}
	user.UserID = "testID"
	item.ImagePath = "/example"
	item.ItemID = "testItemID"
	item.User = &user
	user.Items = append(user.Items, item)
	users[user.UserID] = user

	mux := httprouter.New()

	mux.GET("/", showTop)
	mux.GET("/signup", showSignUp)
	mux.GET("/login", showLogIn)
	mux.GET("/UserPage/:userid", showMyPage)
	// mux.POST("/:userid", registUser)
	// mux.DELETE("/:userid", deleteUser)
	// mux.PUT("/:userid/", updateUserInfo)
	mux.GET("/UserPage/:userid/:itemid", showItem)
	// mux.POST("/:userid/:itemid", uploadItem)
	// mux.DELETE("/:userid/:itemid", deleteItem)
	// mux.PUT("/:userid/:itemid", updateItemInfo)

	server := http.Server{
		Addr:    "127.0.0.1:8083",
		Handler: mux,
	}
	server.ListenAndServe()
}
