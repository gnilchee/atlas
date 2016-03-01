package main

import(
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    log.Println("Listening on port 4443...")
    http.ListenAndServeTLS("0.0.0.0:4443", "/etc/atlas/ssl/certfile.crt", "/etc/atlas/ssl/keyfile.key", nil)
}
