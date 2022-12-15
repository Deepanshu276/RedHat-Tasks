package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Deepanshu276/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	//sending the request to start a server
	r := router.Router()
	fmt.Println("Server is getting started...")
	//log.Fatal---->It logs an error message with a priority of FATAL and immediately exits the program (by calling os.Exit(1))

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to Port 4000")
}
