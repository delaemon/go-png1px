package main

import (
	"encoding/base64"
	"log"
	"fmt"
	"net/http"
	"time"
	"os"
)

var (
	info *log.Logger
	error *log.Logger
)

func handler(w http.ResponseWriter, r *http.Request) {
	png1px := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAABGdBTUEAAK/INwWK6QAAAAtJREFUGFdj+A8EAAn7A/0r1QhFAAAAAElFTkSuQmCC"
	data, err := base64.StdEncoding.DecodeString(png1px)
	if err != nil {
		error.Println(err)
		return
	}

	row := fmt.Sprintf("%s %s \"%s %s %s %s\" %d %s %s",
		r.RemoteAddr,
		r.Header.Get("Content-Type"),
		r.Method,
		r.URL.Path,
		r.Proto,
		r.URL.RawQuery,
		200,
		time.Now(),
		r.UserAgent())
	info.Println(row)
	w.Header().Add("Content-Type", "image/png")
	fmt.Fprintf(w, "%s", data)
}

func main() {
	info = log.New(os.Stdout,"[Info] ", log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	error = log.New(os.Stderr, "[Error] ", log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
