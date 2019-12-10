package Service

import (
	"errors"
	"strconv"
	"time"
)

var (
	questions []Questions
	count = 1
)


func RestListOfQuestions() (ques []string , err error){
	if questions == nil {
		return nil, errors.New("No questions found")
	}

	for _, q := range questions {
		ques = append(ques, q.Question)
	}

	return ques, nil
}

func RestAddQuestion(question string, serialNo int) bool {
	questionObj := Questions{}
	questionObj.setQuestion(strconv.Itoa(count), serialNo, question)
	count++

	questions = append(questions, questionObj)
	return true
}

func RestAddQuestionFrom(ques *Questions) bool {
	ques.SerialNo = count
	count++

	ques.setDefaults()
	return true;
}

func RestDisableQuestion(question string) bool {
	for _, q := range questions {
		if q.Question == question {
			q.Disabled = true
			q.UpdatedOn = time.Now()
			q.UpdatedBy = "UPDATE_USER"
			return true
		}
	}
	return false
}