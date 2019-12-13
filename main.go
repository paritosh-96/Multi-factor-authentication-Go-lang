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

	type newQuestion struct {
		Question string
		UserId   string
	}
	var newQues newQuestion
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = json.Unmarshal(body, &newQues)
	if err = RestAddQuestion(newQues.Question, newQues.UserId); err != nil {
		http.Error(w, err.Error(), 400)
	}
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
	type question struct {
		Id       int
		SerialNo int
	}
	var ques question
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = json.Unmarshal(body, &ques)
	RestUpdateSerialNo(ques.Id, ques.SerialNo)
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
	var answers []Answer
	_ = json.Unmarshal([]byte(body), &answers)
	json.NewEncoder(w).Encode(RestAdd(answers))
}

func listAnsweredQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	var custId Answer
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = json.Unmarshal(body, &custId)
	var questions, er = RestListAnsweredQuestions(custId.CustomerId)
	if er != nil {
		fmt.Fprintf(w, "No Answers found")
	} else {
		log.Println("Answers found: ", len(questions))
		json.NewEncoder(w).Encode(questions)
	}
}

func resetAnswers(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	var answer Answer
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = json.Unmarshal(body, &answer)
	RestReset(answer)
}

func modifyAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	var answer Answer
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_ = json.Unmarshal(body, &answer)
	RestModify(answer)
}

func deleteAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	var answer Answer
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_ = json.Unmarshal(body, &answer)
	RestDelete(answer)
}

func challengeUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	var answer Answer
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_ = json.Unmarshal(body, &answer)
	questions, restError := RestChallenge(answer.CustomerId)
	if restError != nil {
		http.Error(w, restError.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(questions)
}

func validateAnswers(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	var answer []Answer
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_ = json.Unmarshal(body, &answer)
	successStatus := RestValidateAnswers(answer)
	json.NewEncoder(w).Encode(successStatus)
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
	router.HandleFunc("/customer/add", addCustomerAnswer)
	router.HandleFunc("/customer/list", listAnsweredQuestion)
	router.HandleFunc("/customer/reset", resetAnswers)
	router.HandleFunc("/customer/modify", modifyAnswer)
	router.HandleFunc("/customer/delete", deleteAnswer)
	router.HandleFunc("/event/challenge", challengeUser)
	router.HandleFunc("/event/validate", validateAnswers)
	log.Fatal(http.ListenAndServe(":9080", router))
}

func main() {
	handleRequests()
	defer util.Close()
}
