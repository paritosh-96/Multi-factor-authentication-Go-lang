package Service

import (
	"database/sql"
	"errors"
	"github.com/paritosh-96/RestServer/startup"
	"github.com/paritosh-96/RestServer/util"
	"log"
)

type Question struct {
	QuestionId string
	SerialNo   int
	Question   string
}

func RestListOfQuestions() (questions []Question, err error) {
	rows, err := startup.Db.Query("[SP_QUESTION_BANK_GET]")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	questions = []Question{}
	for rows.Next() {
		var ques Question
		if err := rows.Scan(&ques.QuestionId, &ques.SerialNo, &ques.Question); err != nil {
			log.Fatal(err)
		}
		questions = append(questions, ques)
	}
	if len(questions) == 0 {
		return nil, errors.New("ERROR: No Question found")
	}
	return questions, nil
}

func RestAddQuestion(question string, userId string) {
	if util.Empty(question) || util.Empty(userId) {
		log.Fatal("User Id or the question can not be left blank")
		return
	}
	_, err := startup.Db.Query("[SP_QUESTION_BANK_ADD]", sql.Named("text", question), sql.Named("userId", userId))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Added a new Question successfully by [", userId, "]")
}

func RestDeleteQuestion(id int) {
	_, err := startup.Db.Query("[SP_QUESTION_BANK_DELETE]", sql.Named("id", id))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Deleted Question [", id, "] successfully")
}

func RestUpdateSerialNo(id int, serialNo int) {
	_, err := startup.Db.Query("[SP_QUESTION_BANK_UPDATE_ORDER]", sql.Named("id", id), sql.Named("serialNo", serialNo))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Updated serial number")
}
