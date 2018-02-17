package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//User struct
type User struct {
	Username string
	Password string
}

//APIKey and expiry
type APIKey struct {
	Key    string
	Expiry time.Time
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//Decode JSON body
	dec := json.NewDecoder(r.Body)
	var user User
	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("%+v\n", user)

	if user.Username == "mort" && user.Password == "computer" {
		//Username and password correct, generate key
		data := make([]byte, 64)
		_, err := rand.Read(data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		key := APIKey{
			Key:    base64.StdEncoding.EncodeToString(data),
			Expiry: time.Now().Add(time.Hour * 24 * 7),
		}

		//Return key in json body
		enc := json.NewEncoder(w)
		enc.Encode(key)
	} else {
		//Username or password incorrect
		http.Error(w, "Invalid credentials", http.StatusForbidden)
	}
}
