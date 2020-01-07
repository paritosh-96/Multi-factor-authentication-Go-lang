package Service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paritosh-96/RestServer/startup"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ListAllBankQuestions(c *gin.Context) {
	var questions, err = RestListOfQuestions()
	if err != nil {
		fmt.Fprintf(c.Writer, "No questions found")
	} else {
		log.Println("questions found: ", len(questions))
		json.NewEncoder(c.Writer).Encode(questions)
	}
}

func AddNewBankQuestion(c *gin.Context) {
	reqBody := c.Request.Body
	if reqBody == nil {
		http.Error(c.Writer, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(reqBody)

	type newQuestion struct {
		Question string
		UserId   string
	}
	var newQues newQuestion
	if err != nil {
		http.Error(c.Writer, err.Error(), 400)
		return
	}
	_ = json.Unmarshal(body, &newQues)
	if err = RestAddQuestion(newQues.Question, newQues.UserId); err != nil {
		log.Println(err.Error())
		http.Error(c.Writer, err.Error(), 400)
	}
}

func DeleteBankQuestion(c *gin.Context) {
	reqBody := c.Request.Body
	if reqBody == nil {
		http.Error(c.Writer, "Please send a request body", 400)
		return
	}
	type Delete struct {
		QuestionId int
		UserId     string
	}
	var deleteJson Delete
	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		http.Error(c.Writer, err.Error(), 400)
		return
	}
	_ = json.Unmarshal(body, &deleteJson)
	if err = RestDeleteQuestion(deleteJson.QuestionId, deleteJson.UserId); err != nil {
		http.Error(c.Writer, err.Error(), 400)
	}
}

func UpdateSerialNo(c *gin.Context) {
	reqBody := c.Request.Body
	if reqBody == nil {
		http.Error(c.Writer, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(reqBody)
	type question struct {
		Id       int
		SerialNo int
	}
	var ques question
	if err != nil {
		http.Error(c.Writer, err.Error(), 400)
		return
	}
	_ = json.Unmarshal(body, &ques)
	RestUpdateSerialNo(ques.Id, ques.SerialNo)
}

func CustomerQuestionsCount(c *gin.Context) {
	fmt.Fprintf(c.Writer, strconv.Itoa(startup.ConfigParameters.QuestionsPerUser))
}

func ListAllCustomerQuestion(c *gin.Context) {
	var questions, err = RestListAll()
	if err != nil {
		fmt.Fprintf(c.Writer, "No questions found")
	} else {
		log.Println("questions found: ", len(questions))
		json.NewEncoder(c.Writer).Encode(questions)
	}
}

func AddCustomerAnswer(c *gin.Context) {
	reqBody := c.Request.Body
	if reqBody == nil {
		http.Error(c.Writer, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(reqBody)

	if err != nil {
		http.Error(c.Writer, err.Error(), 400)
		return
	}
	var answers []Answer
	_ = json.Unmarshal([]byte(body), &answers)
	if _, error := RestAdd(answers); error != nil {
		fmt.Fprintf(c.Writer, "Error in adding: "+error.Error())
	}
}

func ListAnsweredQuestion(c *gin.Context) {
	custId := c.Request.URL.Query().Get("customerId")
	var questions, er = RestListAnsweredQuestions(custId)
	if er != nil {
		fmt.Fprintf(c.Writer, "")
	} else {
		log.Println("Answers found: ", len(questions))
		json.NewEncoder(c.Writer).Encode(questions)
	}
}

func ResetAnswers(c *gin.Context) {
	custId := c.Request.URL.Query().Get("customerId")
	RestReset(custId)
}

func ModifyAnswer(c *gin.Context) {
	reqBody := c.Request.Body
	if reqBody == nil {
		http.Error(c.Writer, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(reqBody)
	var answers []Answer
	if err != nil {
		http.Error(c.Writer, err.Error(), 400)
		return
	}
	_ = json.Unmarshal(body, &answers)
	json.NewEncoder(c.Writer).Encode(RestModify(answers))
}

func DeleteAnswer(c *gin.Context) {
	reqBody := c.Request.Body
	if reqBody == nil {
		http.Error(c.Writer, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(reqBody)
	var answer Answer
	if err != nil {
		http.Error(c.Writer, err.Error(), 400)
		return
	}

	_ = json.Unmarshal(body, &answer)
	RestDelete(answer)
}

func GetChallengeQuesCount(c *gin.Context) {
	fmt.Fprintf(c.Writer, strconv.Itoa(startup.ConfigParameters.NoOfQuestionsForChallenger))
}

func ChallengeUser(c *gin.Context) {
	customerId := c.Request.URL.Query().Get("customerId")
	fmt.Println(customerId)
	questions, restError := RestChallenge(customerId)
	if restError != nil {
		http.Error(c.Writer, restError.Error(), 400)
		return
	}
	json.NewEncoder(c.Writer).Encode(questions)
}

func ValidateAnswers(c *gin.Context) {
	reqBody := c.Request.Body
	if reqBody == nil {
		http.Error(c.Writer, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(reqBody)
	var answer []Answer
	if err != nil {
		http.Error(c.Writer, err.Error(), 400)
		return
	}

	_ = json.Unmarshal(body, &answer)
	if err2 := RestValidateAnswers(answer); err2 != nil {
		http.Error(c.Writer, err2.Error(), 400)
	}
}
