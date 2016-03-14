package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var filename string

func captureMessageHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	check(err)
	fmt.Println("Received Message.")

	defer f.Close()
	fmt.Fprintf(f, "<-------------------------Begin Message------------------------>\n")
	for k, v := range r.Header {
			fmt.Fprintf(f, k)
			fmt.Fprintf(f, ":  ")
			for _, value := range v {
					fmt.Fprintf(f, value)
			}
			fmt.Fprintf(f, "\n")
	}
	fmt.Fprint(f, r.Header)
	fmt.Fprintf(f, "\n")
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(f, "<-------------Begin Body------------>\n")
	fmt.Fprint(f, string(body))
	fmt.Fprintf(f, "\n")
	fmt.Fprintf(f, "<-------------------------End Message------------------------>\n")

	w.Write([]byte("Success!"))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var port = flag.Int("port", 8080, "The port number you want the server running on. Default is 8080")
	var logLocation = flag.String("loc", "./messages.log", "Where you would like the messages logged to.")

	filename = *logLocation

	flag.Parse()

	http.HandleFunc("/", captureMessageHandler)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)

	check(err)

}
