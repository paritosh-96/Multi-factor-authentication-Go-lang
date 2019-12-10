package main

import (
	"encoding/json"
	"fmt"
	. "github.com/paritosh-96/RestServer/Service"
	"log"
	"net/http"
)



func listAllQuestions(w http.ResponseWriter, r *http.Request){
	var questions, err = RestListOfQuestions()
	if err != nil {
		fmt.Fprintf(w, "No questions found")
	} else {
		json.NewEncoder(w).Encode(questions)
	}
}

func addNewQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var question Questions
	err := json.NewDecoder(r.Body).Decode(question)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if RestAddQuestionFrom(&question) {
		fmt.Fprintf(w, "Question added successfully")
	} else {
		fmt.Fprintf(w, "Question could not be added")
	}
}

func disableQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var question Questions
	err := json.NewDecoder(r.Body).Decode(question)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if RestDisableQuestion(question.Question) {
		fmt.Fprintf(w, "Question deleted successfully")
	} else  {
		fmt.Fprintf(w, "Question could not be deleted")
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Bank Questionnaire!")
}

func handleRequests() {
	http.HandleFunc("/bank", homePage)
	http.HandleFunc("/bank/listAll", listAllQuestions)
	http.HandleFunc("/bank/add", addNewQuestion)
	http.HandleFunc("/bank/disable", disableQuestion)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
