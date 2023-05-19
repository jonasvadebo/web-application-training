package main

import (
	"fmt"

	"github.com/jonasvadebo/web-application-training/domain"
)

func main() {

	u := domain.User{
		ID:        2,
		FirstName: "Jonas",
		LastName:  "Vadebo",
	}

	fmt.Println(u)

	/*
		 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello World!"))
			})

			http.ListenAndServe(":8080", nil)
	*/
}
