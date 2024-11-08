package main

import (
	// the JSON encoding library will be used to process inbound data and format outbound JSON outputs
	"encoding/json"
	"os"

	// the fmt standard library formats string output
	"fmt"
	// the HTTP library is used to manage our HTTP server connection
	"net/http"
	// the os/exec library allows Go to call other applications installed on the machine
	"os/exec"
)

// custom Go type called cmdresult using struct keyword to instruct Go how to format the JSON emitted from a REST endpoint
type cmdresult struct {
	// contains tow attributes indicating whether or not the function call was successful and and any specific message to be transmitted as part of the success result
	Success bool	`json:"success"`
	Message string 	`json:"message"`	
}

//function to resond when the root level of th eserver path is called from a browser. Not essential, but offers a quick check to see of the server is running
func homepage(write http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(write, "Go Home Simple REST API Server")
}

//primary function to handle an inbound request that will query and return the server's current date, transmitted in our defined JSON-formatted response. This is an example of how to use a system command to obtain information you can return in a RESTful way. (This is similar to how we will run other commands later, such as to activate the Raspberry Pi camera.)
func getdate(write http.ResponseWriter, _ *http.Request) {
	// declare result variable and initialize to an empty cmdresult struct.
	result := cmdresult{}

	// using exec.Command function from the os/exec package, call the external date program and capture its output. 
	out, err :=exec.Command("date").Output()
	
	// If the call  succeeds, set Success boolean attribute value to result to true and the Message to a sentence containing the command output. otherwise these attributes will remain with the zero values of false and empty string.
	if err == nil {
		result.Success = true
		result.Message = "The date is " + string(out)
	}
	// Encode the result struct to JSON and return the result via HTTP to the caller
	json.NewEncoder(write).Encode(result)
}

// initialize the HTTP server, set the port number, and declare the relative paths that will be used to call the respective hompage and getdate functions.
func main() {
	fmt.Println("Hello from main.go: main")
	http.HandleFunc("/", homepage)
	http.HandleFunc("/api/v1/getdate", getdate)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
	fmt.Println("main() after the ListenAndServe")

}

