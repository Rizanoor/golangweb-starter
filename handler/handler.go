package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello ini halaman home"))
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar golang"))
}
func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello mario, saya sedang belajar golang"))
}
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Product page"))

	fmt.Fprintf(w, "Product page: %d", idNumb)
}
