package Service

import (
	"database/sql"
	"errors"
	"github.com/paritosh-96/RestServer/startup"
	"github.com/paritosh-96/RestServer/util"
	"log"
)

type Answer struct {
	QuestionId string
	Answer     string
	CustomerId string
	UserId     string
}

func RestListAll() (questions []Question, err error) {
	rows, err := startup.Db.Query("[SP_QUESTION_BANK_GET]")
	util.Check(err, "")

	defer rows.Close()
	questions = []Question{}
	for rows.Next() {
		var ques Question
		err := rows.Scan(&ques.QuestionId, &ques.SerialNo, &ques.Question)
		util.Check(err, "")
		questions = append(questions, ques)
	}

	if len(questions) == 0 {
		return nil, errors.New("ERROR: No Question found")
	}

	return questions, nil
}

func RestAdd(answers []Answer) {
	for _, answer := range answers {
		_, err := startup.Db.Query("[SP_CUSTOMER_QUESTIONS_ADD]",
			sql.Named("questionId", answer.QuestionId),
			sql.Named("customerId", answer.CustomerId),
			sql.Named("answer", answer.Answer),
			sql.Named("userId", answer.UserId))
		if err != nil {
			log.Println("Error while adding answer for [", answer.QuestionId, " ]: ", err)
			continue
		}
		log.Println("Answer for question [", answer.QuestionId, "] successfully added")
	}
}

func RestListAnsweredQuestions(customerId string) (answers []Answer, err error) {
	rows, err := startup.Db.Query("[SP_CUSTOMER_QUESTIONS_GET]", sql.Named("customerId", customerId))
	util.Check(err, "")

	defer rows.Close()

	answers = []Answer{}
	for rows.Next() {
		var ans Answer
		err := rows.Scan(&ans.QuestionId, &ans.CustomerId, &ans.Answer)
		util.Check(err, "")
		answers = append(answers, ans)
	}

	if len(answers) == 0 {
		log.Println("No answers found for the given customer id")
		return nil, errors.New("ERROR: No answers found")
	}

	return answers, nil
}

func RestReset(answer Answer) {
	_, err := startup.Db.Query("[SP_CUSTOMER_QUESTIONS_RESET]",
		sql.Named("customerId", answer.CustomerId),
		sql.Named("userId", answer.UserId))
	if err != nil {
		log.Fatal("Could not reset answers for customer [", answer.CustomerId, "] Error: ", err)
		return
	}
	log.Println("All answers reset for customer  [", answer.CustomerId, "] successful")

}

func RestModify(answer Answer) {
	_, err := startup.Db.Query("[SP_CUSTOMER_QUESTIONS_UPDATE]",
		sql.Named("customerId", answer.CustomerId),
		sql.Named("questionId", answer.QuestionId),
		sql.Named("answer", answer.Answer),
		sql.Named("userId", answer.UserId))
	if err != nil {
		log.Fatal("Could not update the answer for quetion [", answer.QuestionId, "] Error: ", err)
		return
	}
	log.Println("Answer modified successfully")
}

func RestDelete(answer Answer) {
	_, err := startup.Db.Query("[SP_CUSTOMER_QUESTION_DELETE]",
		sql.Named("customerId", answer.CustomerId),
		sql.Named("questionId", answer.QuestionId),
		sql.Named("userId", answer.UserId))
	if err != nil {
		log.Fatal("Could not delete answer for the custoemr [", answer.CustomerId, "] Error: ", err)
		return
	}
	log.Println("Answer deleted from the customer [" + answer.CustomerId + "]")
}
