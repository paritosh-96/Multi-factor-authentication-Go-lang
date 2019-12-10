package Service

import "time"

type Questions struct {
	QuestionId string
	SerialNo int
	Question string
	Disabled bool

	CreatedBy string
	CreatedOn time.Time
	UpdatedOn time.Time
	UpdatedBy string
}

func (q *Questions) setQuestion(questionId string, serialNo int, question string) {
	q.SerialNo = serialNo
	q.Question = question
	q.QuestionId = questionId
	q.Disabled = false

	q.CreatedOn = time.Now()
	q.CreatedBy = "TEST_USER"
	q.UpdatedBy = "TEST_USER"
	q.UpdatedOn = time.Now()
}

func (q *Questions) setDefaults() {
	q.Disabled = false

	q.CreatedOn = time.Now()
	q.CreatedBy = "TEST_USER"
	q.UpdatedBy = "TEST_USER"
	q.UpdatedOn = time.Now()
}