package main

import(
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    log.Println("Server is listening on 4443 and 8080...")
    go http.ListenAndServeTLS("0.0.0.0:4443", "/etc/atlas/ssl/certfile.crt", "/etc/atlas/ssl/keyfile.key", nil)
    http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(redirector))
}

func redirector(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://hostname.com"+req.RequestURI, http.StatusMovedPermanently)
}
