package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type myData struct { // struct for formatting decoded json
	URL string `json:"url"`
}

type printJSON struct { // struct for encoded json sent back to client
	ShortURL string `json:"short_url"`
}

var characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // characters used for generating random URL ID

func main() {
	rand.Seed(time.Now().UnixNano()) // time for Random() to generate different ID at every startup

	urlMap := map[string]string{} // map used to track urls/IDs

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			if path, ok := urlMap[r.URL.Path]; ok { // if Path matches an ID in the map, it redirects original link
				http.Redirect(w, r, path, http.StatusFound)
				fmt.Println("URL ID Redirected!")
			} else {
				fmt.Fprintf(w, "ID does not exist")
				fmt.Println("ID does not exist.")
			}
		case http.MethodPost:
			jsonMap := map[string]string{} // declares temporary map for json
			var data myData                // declares data variable for storing decoded json

			decoder := json.NewDecoder(r.Body)

			err := decoder.Decode(&data) // decodes sent json
			if err != nil {
				panic(err)
			}

			x := fmt.Sprintf(data.URL) // converts the sent url into a string

			if isURL(x) == false { // validates the sent url
				fmt.Fprintf(w, "URL not valid.")
				return
			}

			ran := random() // created random ID to assign

			urlMap["/"+ran] = x // ID is set at the index of  urlMap with x as the original link

			jsonMap[x] = "http://127.0.0.1:8080/" + ran // assign new link to the jsonMap to be sent back to user
			p := printJSON{jsonMap[x]}                  // makes the map into a struct

			js, err := json.Marshal(p) // marshalize p into json
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(js)) //new link sent back as json
			fmt.Println("Url Successfuly Shortened!")
		default:
			fmt.Fprintf(w, "Accepts only GET and POST requests.") // if recieved invalid requests
		}
	})

	http.ListenAndServe(":8080", nil)
}

func isURL(str string) bool { // function to check for URL validity
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func random() string { // function used to generate random string for URL ID
	b := make([]byte, 8)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	sURL := string(b)
	return sURL
}
