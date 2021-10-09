package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
  	"os"
  	"encoding/json"
)

type ResponseTest struct {
	Status		string
	Text string
}


func main() {
	port := os.Getenv("PORT")
	log.Println("server start with port: " +port)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":" +port, nil)
}


func api(w http.ResponseWriter, r *http.Request) {
	 ErrUrl := [...]string{
		"Error 101:\nBad title ( Empty )",
		"Error 102:\nBad token",
	}

	log.Print("new request")
	query := r.URL.Query()

	TitleByte := query.Get("title")
	TokenByte := query.Get("token")
	EmailByte := query.Get("email")
	UserKeyByte := query.Get("key")

	title := string(TitleByte[:])
	token := string(TokenByte[:])
	email := string(EmailByte[:])
	userkey := string(UserKeyByte[:])

	if title == "" {
		log.Println(ErrUrl[0])
		fmt.Fprintf(w, ErrUrl[0])
	}

	if token == "" {
		log.Println(ErrUrl[1])
		fmt.Fprintf(w, ErrUrl[1])
	}



	if title == "test" {
			if token != "cardinal" {
				OtherErrors := [...]string{
					"Error 201:\rIncorrect Token Cardinal",
				}
				log.Println(OtherErrors[0])
				fmt.Fprintf(w, OtherErrors[0])
			}else{
				//send req cardinal
        responsetest := ResponseTest{"Status: OK",  "Code 1 TRUE"}

	      js, err := json.Marshal(responsetest)
	      if err != nil {
	      	http.Error(w, err.Error(), http.StatusInternalServerError)
		      return
	      }

	      w.Header().Set("Content-Type", "application/json")
	      w.Write(js)
			  }
	}

}
