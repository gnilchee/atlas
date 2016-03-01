package main

import(
    "net/http"
)

func redirector(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://hostname.com"+req.RequestURI, http.StatusMovedPermanently)
}

func main() {
	http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(redirector))
}
