package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"mongo-client/client"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

var mh *client.MongoHandler

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/clients", func(r chi.Router) {
		r.Get("/", getAllLogins)          //GET /clients
		r.Get("/{login}", getLogin)       //GET /client/login
		r.Post("/", addLogin)             //POST /client
		r.Put("/{email}", updateData)     //PUT /client/email
		r.Delete("/{login}", deleteLogin) //DELETE /client/login
	})
	return r
}

func main() {
	mongoDbConnection := "mongodb://localhost:27017"
	mh = client.NewHandler(mongoDbConnection)
	r := registerRoutes()
	log.Fatal(http.ListenAndServe(":3060", r))
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")
	if login == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	login := &client.Client{}
	err := mh.GetOne(client, bson.M{"login": login})
	if err != nil {
		http.Error(w, fmt.Sprintf("Client with login: %s not found", login), 404)
		return
	}
	json.NewEncoder(w).Encode(login)
}

func getAllLogins(w http.ResponseWriter, r *http.Request) {
	logins := mh.Get(bson.M{})
	json.NewEncoder(w).Encode(logins)
}

func addLogin(w http.ResponseWriter, r *http.Request) {
	existingLogin := &clent.Client{}
	var login client.Client
	json.NewDecoder(r.Body).Decode(&client)
	login.CreatedOn = time.Now()
	err := mh.GetOne(existingLogin, bson.M{"login": client.Login})
	if err == nil {
		http.Error(w, fmt.Sprintf("Client with login: %s already exist", client.Login), 400)
		return
	}
	_, err = mh.AddOne(&client)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Client created successfully"))
	w.WriteHeader(201)
}

func deleteLogin(w http.ResponseWriter, r *http.Request) {
	existingLogin := &contact.Contact{}
	login := chi.URLParam(r, "login")
	if login == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	err := mh.GetOne(existingLogin, bson.M{"login": login})
	if err != nil {
		http.Error(w, fmt.Sprintf("Client with login: %s does not exist", login), 400)
		return
	}
	_, err = mh.RemoveOne(bson.M{"login": login})
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Client deleted"))
	w.WriteHeader(200)
}

func updateData(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "email")
	if data == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	oldData := &client.Client{}
	json.NewDecoder(r.Body).Decode(oldData)
	_, err := mh.Update(bson.M{"email": data}, oldData)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Client update successful"))
	w.WriteHeader(200)
}
