package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	a := "start"
	defer fmt.Println(a)
	a = "end"

	// c, d := 1, 0
	// ans := c / d
	// fmt.Println(ans)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		panic(err2.Error())
	}

}
