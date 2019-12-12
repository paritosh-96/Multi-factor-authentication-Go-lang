package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	. "github.com/paritosh-96/RestServer/Service"
	_ "github.com/paritosh-96/RestServer/startup"
	"github.com/paritosh-96/RestServer/util"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func listAllBankQuestions(w http.ResponseWriter, r *http.Request) {
	var questions, err = RestListOfQuestions()
	if err != nil {
		fmt.Fprintf(w, "No questions found")
	} else {
		log.Println("questions found: ", len(questions))
		json.NewEncoder(w).Encode(questions)
	}
}

func addNewBankQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	bodyParts := strings.Split(string(body), ":")
	RestAddQuestion(bodyParts[0], bodyParts[1])
}

func deleteBankQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var id int
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	id, _ = strconv.Atoi(string(body))
	RestDeleteQuestion(id)
}

func updateSerialNo(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	bodyParts := strings.Split(string(body), ":")
	id, _ := strconv.Atoi(bodyParts[0])
	serialNo, _ := strconv.Atoi(bodyParts[1])
	RestUpdateSerialNo(id, serialNo)
}

func listAllCustomerQuestion(w http.ResponseWriter, r *http.Request) {
	var questions, err = RestListAll()
	if err != nil {
		fmt.Fprintf(w, "No questions found")
	} else {
		log.Println("questions found: ", len(questions))
		json.NewEncoder(w).Encode(questions)
	}
}

func addCustomerAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var answer []Answer
	_ = json.Unmarshal([]byte(body), &answer)
	RestAdd(answer)
}

func bankHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Bank Questionnaire!")
}

func customerHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the customer Page")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/bank", bankHomePage)
	router.HandleFunc("/bank/listAll", listAllBankQuestions)
	router.HandleFunc("/bank/addQuestion", addNewBankQuestion)
	router.HandleFunc("/bank/disableQuestion", deleteBankQuestion)
	router.HandleFunc("/bank/updateSerialNo", updateSerialNo)
	router.HandleFunc("/customer", customerHomePage)
	router.HandleFunc("/customer/listAll", listAllCustomerQuestion)
	router.HandleFunc("/customer/addCustomerAnswer", addCustomerAnswer)
	log.Fatal(http.ListenAndServe(":9080", router))
}

func main() {
	handleRequests()
	defer util.Close()
}
