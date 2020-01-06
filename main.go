package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	. "github.com/paritosh-96/RestServer/Service"
	_ "github.com/paritosh-96/RestServer/startup"
	"github.com/paritosh-96/RestServer/util"
)

func start() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("github.com/paritosh-96/RestServer/client/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/bank/listAll", ListAllBankQuestions)
		api.POST("/bank/addQuestion", AddNewBankQuestion)
		api.POST("/bank/disableQuestion", DeleteBankQuestion)
		api.POST("/bank/updateSerialNo", UpdateSerialNo)

		api.GET("/customer/customerQuestionCount", CustomerQuestionsCount)
		api.GET("/customer/listAll", ListAllCustomerQuestion)
		api.POST("/customer/add", AddCustomerAnswer)
		api.GET("/customer/list", ListAnsweredQuestion)
		api.POST("/customer/reset", ResetAnswers)
		api.POST("/customer/modify", ModifyAnswer)
		api.POST("/customer/delete", DeleteAnswer)

		api.GET("/event/challengeQuestionsCount", GetChallengeQuesCount)
		api.GET("/event/challenge", ChallengeUser)
		api.POST("/event/validate", ValidateAnswers)
	}

	// Start and run the server
	router.Run(":6080")
}

func main() {
	start()
	defer util.Close()
}
