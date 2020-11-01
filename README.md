# Multi Factor Authentication
Basic Rest apis in GO LANG

The server holds a list of questions in it. The rest services allow a user to :
  1) Admin
    
		a) List the questions available
    
		b) Add a new question
    
		c) Delete an existing question (Disable the question)
  2) Customer
    
		a) Select questions to answer
    
		b) Answer the selected questions and submit.
    
		c) Update answers
    
		d) reset selected questions and answers to them.
  3) Validation
    
		a) Pop up random questions out of the answered questions for the given user.
    
		b) Verify if the answers to the popped up questions are correct or not.
  
	
	The client code is also present which is written using react Js. When the main.go binary is run, the client code is built already will be hosted over the configured port. The details will be printed in the information of the run.
