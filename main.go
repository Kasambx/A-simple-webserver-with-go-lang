package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
	// Define the root route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to HTML
		w.Header().Set("Content-Type", "text/html")

		// Read the HTML file from disk
		htmlBytes, err := ioutil.ReadFile("static/video.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the HTML to the response writer
		w.Write(htmlBytes)
	})

	// Serve the video file from the "static" directory
	http.Handle("/unwrapped.mp4", http.FileServer(http.Dir("static")))

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
	
	//print server listening
	fmt.Println("Server running on port 8080")
}
