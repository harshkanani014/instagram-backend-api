package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/harshkanani014/instagram-backend-api/models"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// get user model session from mongodb
type UserController struct {
	session *mgo.Session
}

// get post model session from mongodb
type PostController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

// Function to Get User Details
// API ENDPOINT : /users/:id (id : user id)
// HTTP REQUEST : GET
// params : user id
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	// convert hex to id
	oid := bson.ObjectIdHex(id)

	// call user model from db
	u := models.User{}

	// connect with db session and find user with params id
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// convert object into json
	uj, err := json.Marshal(u)
	if err != nil {
		// error if not able to convert in json
		fmt.Println(err)
	}

	// send response as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

// Function to Create new User Details
// HTTP REQUEST : POST
// API ENDPOINT : /users
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	// get data from body and convert into object
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	// Hash Password obtained from data
	pass := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.Password = string(hashedPassword)

	// Insert New user data into user model
	uc.session.DB("mongo-golang").C("users").Insert(u)

	// convert object to json
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	// send as json response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// Function to Create new Post for given user
// HTTP REQUEST : POST
// API ENDPOINT : /posts
func (pc PostController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Posts{}

	// get data from body and convert into object
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	// get current time
	dt := time.Now()
	u.Timestamp = dt

	// insert new post data into posts model
	pc.session.DB("mongo-golang").C("posts").Insert(u)

	// convert response to json
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	// send response in json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// Function to Get Each Post Details
// API ENDPOINT : /posts/:id (id : pst id)
// HTTP REQUEST : GET
// params : user id
func (pc PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	// convert hex id to int id
	oid := bson.ObjectIdHex(id)

	u := models.Posts{}

	// find post using post id from posts model
	if err := pc.session.DB("mongo-golang").C("posts").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// convert object into json
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// send response in json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

// Function to Get Each Post Details
// API ENDPOINT : /posts/users/:id (id : user id)
// HTTP REQUEST : GET
// params : user id
func (pc PostController) GetAllUserPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.Posts{}

	// Find all post associated with specific user id mentioned in params
	if err := pc.session.DB("mongo-golang").C("posts").Find(oid).All(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// Convert obtained objects into json
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// send obtained json as response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
