package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println(r)
}

// go mod init example.com/a
// go get github.com/gorilla/mux
// import
