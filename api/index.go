package handler

import (
	"fmt"
	"net/http"
)

// Handler is the exported function that Vercel looks for
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
