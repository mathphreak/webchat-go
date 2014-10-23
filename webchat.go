package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type handler func(w http.ResponseWriter, req *http.Request)

func decompress(gzip64 string) string {
	compressed, _ := base64.StdEncoding.DecodeString(gzip64)
	b := bytes.NewReader(compressed)
	reader, _ := gzip.NewReader(b)
	data, _ := ioutil.ReadAll(reader)
	return string(data)
}

func asset(compressedData string) handler {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, decompress(compressedData))
	}
}

func main() {

	//homeTempl := template.Must(template.ParseFiles("home.html"))

	r := mux.NewRouter()
	/*r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		template.Must(template.ParseFiles("home.html")).Execute(w, req.Host)
		_ = homeTempl // for production, embed home.html and don't live-reload it
	})*/
	r.HandleFunc("/", asset(html))
	r.HandleFunc("/ws", wsHandler)
	r.HandleFunc("/main.js", asset(js))
	r.HandleFunc("/main.css", asset(css))

	// Use the default middleware.
	n := negroni.Classic()
	// ... Add any other middlware here

	// add the router as the last handler in the stack
	n.UseHandler(r)

	go h.run()

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "3000"
	}
	n.Run(":" + portString)
}
