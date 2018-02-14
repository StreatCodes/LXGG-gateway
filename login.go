package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, req *http.Request) {
	//TODO parse body

	data := make([]byte, 64)
	_, err := rand.Read(data)
	if err != nil {
		http.Error(w, "error:"+err.Error(), 500)
		return
	}

	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)

	fmt.Fprintf(w, "New key: "+str)
}
