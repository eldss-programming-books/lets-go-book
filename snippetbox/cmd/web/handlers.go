package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// home displays the home page for the site.
func home(w http.ResponseWriter, r *http.Request) {
	// Ensure only "/" is matched
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Files to parse
	files := []string{
		"./ui/html/home.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.partial.gohtml",
	}

	// Attempt to parse the files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Render and send the resulting html
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

// showSnippet returns a specific snippet to display.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// createSnippet creates a new snippet.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		// * Note: The following methods must be called in this
		// * order for them to work. Otherwise defaults will be
		// * called which cannot be changed later.
		// Inform user of allowed methods
		w.Header().Set("Allow", http.MethodPost)
		// If response is JSON, must set that in the Header map or it
		// will be sent as plaintext.

		// Send response
		http.Error(w, "Method Not Allowed", 405)
		// Above error is a shortcut for the following two lines
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		return
	}

	// Success
	w.Write([]byte("Create a new snippet..."))
}
