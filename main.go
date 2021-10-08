package main

import (
	"net/http"

	"github.com/harshkanani014/instagram-backend-api/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())

	// API endpoints
	r.POST("/users", uc.CreateUser)
	r.GET("/users/:id", uc.GetUser)
	r.POST("/posts", pc.CreatePost)
	r.GET("/posts/:id", pc.GetPost)
	r.GET("/posts/users/:id", pc.GetAllUserPost)

	// server start at localhost:9000
	http.ListenAndServe("localhost:9000", r)
}

// connect with mongodb
func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		// if connection error
		panic(err)
	}

	// return connection session
	return s
}
