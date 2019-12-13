package Service

import (
	"database/sql"
	"errors"
	"github.com/paritosh-96/RestServer/startup"
	"github.com/paritosh-96/RestServer/util"
	"log"
	"math/rand"
	"time"
)

func RestChallenge(customerId string) ([]Question, error) {
	rows, err := startup.Db.Query("[SP_CUSTOMER_QUESTIONS_GET_QUESTION_ID]", sql.Named("customerId", customerId))
	util.Check(err, "Error while loading all the question ids for the customer")
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var quesId int
		err := rows.Scan(&quesId)
		util.Check(err, "")
		ids = append(ids, quesId)
	}
	if ids == nil {
		log.Fatal("No question answered for the customer [" + customerId + "]")
		return nil, errors.New("No question answered for the customer [" + customerId + "]")
	}
	rand.Seed(time.Now().UnixNano())
	var questions []Question
	for i := 0; i < startup.ConfigParameters.NoOfQuestionsForChallenger; i++ {
		ind := rand.Intn(len(ids))
		rows, err := startup.Db.Query("[SP_QUESTION_BANK_GET]", sql.Named("questionId", ids[ind+1]))
		util.Check(err, "Error while loading question for challenge")
		for rows.Next() {
			var question Question
			err := rows.Scan(&question.QuestionId, &question.Question)
			util.Check(err, "Error while loading the question from the resultSet")
			questions = append(questions, question)
		}
	}
	return questions, nil
}

func RestValidateAnswers(answers []Answer) map[int]string {
	if len(answers) == 0 {
		log.Fatal("No answers found to validate ")
		return nil
	}
	status := map[int]string{}
	for _, answer := range answers {
		rows, err := startup.Db.Query("[SP_CUSTOMER_QUESTIONS_VALIDATE]",
			sql.Named("customerId", answer.CustomerId),
			sql.Named("questionId", answer.QuestionId),
			sql.Named("answer", answer.Answer))
		util.Check(err, "Error while validating answers")
		for rows.Next() {
			var success int
			err := rows.Scan(&success)
			util.Check(err, "Error while loading the success status of validation of answers")
			if success == 0 {
				status[answer.QuestionId] = "Correct"
			} else {
				status[answer.QuestionId] = "Wrong"
			}
		}
	}
	return status
}
