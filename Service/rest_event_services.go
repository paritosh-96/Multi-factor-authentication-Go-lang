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
		log.Println("No question answered for the customer [" + customerId + "]")
		return nil, errors.New("No question answered for the customer [" + customerId + "]")
	}
	rand.Seed(time.Now().UnixNano())
	var questions []Question
	var visitedQuestionIds []int
	for len(visitedQuestionIds) < startup.ConfigParameters.NoOfQuestionsForChallenger {
		ind := rand.Intn(len(ids))
		if util.Contains(visitedQuestionIds, ind) {
			continue
		}
		rows, err := startup.Db.Query("[SP_QUESTION_BANK_GET]", sql.Named("questionId", ids[ind]))
		util.Check(err, "Error while loading question for challenge")
		for rows.Next() {
			var question Question
			err := rows.Scan(&question.QuestionId, &question.Question)
			util.Check(err, "Error while loading the question from the resultSet")
			if err == nil {
				questions = append(questions, question)
				visitedQuestionIds = append(visitedQuestionIds, ind)
			}
		}
	}
	return questions, nil
}

func RestValidateAnswers(answers []Answer) error {
	if len(answers) == 0 {
		log.Fatal("No answers found to validate ")
		return nil
	}
	for _, answer := range answers {
		rows, err := startup.Db.Query("[SP_CUSTOMER_QUESTIONS_VALIDATE]",
			sql.Named("customerId", answer.CustomerId),
			sql.Named("questionId", answer.QuestionId))
		util.Check(err, "Error while validating answers")
		for rows.Next() {
			var correctAnswer string
			err := rows.Scan(&correctAnswer)
			util.Check(err, "Error while loading the success status of validation of answers")
			if answer.Answer != correctAnswer {
				return errors.New("Invalid answer: ")
			}
		}
	}
	return nil
}
