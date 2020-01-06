package Service

import (
	"database/sql"
	"errors"
	"github.com/paritosh-96/RestServer/startup"
	"github.com/paritosh-96/RestServer/util"
	"log"
)

type Question struct {
	QuestionId int
	SerialNo   int
	Question   string
}

func RestListOfQuestions() (questions []Question, err error) {
	rows, err := startup.Db.Query("[SP_QUESTION_BANK_GET_ALL]")
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

func RestAddQuestion(question string, userId string) error {
	if util.Empty(question) || util.Empty(userId) {
		log.Fatal("User Id or the question can not be left blank")
		return errors.New("User Id or the question can not be left blank ")
	}

	questions, _ := RestListOfQuestions()
	if len(questions) >= startup.ConfigParameters.MaxQuestions {
		log.Fatal("Maximum questions limit already reached, Can not add a new question")
		return errors.New("Maximum questions limit already reached, Can not add a new question ")
	}
	_, err := startup.Db.Query("[SP_QUESTION_BANK_ADD]", sql.Named("text", question), sql.Named("userId", userId))
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Added a new Question successfully by [", userId, "]")
	return nil
}

func RestDeleteQuestion(id int, userId string) error {
	if !isIdValid(id) {
		return errors.New("Question id does not exists! ")
	}
	_, err := startup.Db.Query("[SP_QUESTION_BANK_DELETE]", sql.Named("id", id), sql.Named("userId", userId))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Deleted Question [", id, "] successfully")
	return nil
}

func RestUpdateSerialNo(id int, serialNo int) {
	_, err := startup.Db.Query("[SP_QUESTION_BANK_UPDATE_ORDER]", sql.Named("id", id), sql.Named("serialNo", serialNo))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Updated serial number")
}

func isIdValid(id int) bool {
	found := false
	questions, _ := RestListOfQuestions()
	for _, v := range questions {
		if v.QuestionId == id {
			found = true
		}
	}
	return found
}
