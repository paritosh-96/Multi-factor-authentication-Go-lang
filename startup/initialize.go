package startup

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"io/ioutil"
	"log"
	"net/url"
)

/**
Structure to store all the configuration parameters
*/
type Parameters struct {
	MaxQuestions               int            `json:"max_questions"`
	QuestionsPerUser           int            `json:"questions_per_user"`
	NoOfQuestionsForChallenger int            `json:"no_of_questions_for_challenger"`
	DbParams                   DatabaseParams `json:"db_params"`
}

/**
Structure to store all the database parameters
*/
type DatabaseParams struct {
	DbType   string `json:"db_type"`
	DbUser   string `json:"db_user"`
	DbPwd    string `json:"db_pwd"`
	DbHost   string `json:"db_host"`
	DbPort   int    `json:"db_port"`
	DbSchema string `json:"db_schema"`
}

var (
	Db               *sql.DB
	ConfigParameters Parameters
)

/**
Load all the configuration parameters
Setup the database handler using the config parameters
*/
func init() {
	loadParams()
	log.Println("Setting up con")
	u := &url.URL{
		Scheme: ConfigParameters.DbParams.DbSchema,
		User:   url.UserPassword(ConfigParameters.DbParams.DbUser, ConfigParameters.DbParams.DbPwd),
		Host:   fmt.Sprintf("%s:%d", ConfigParameters.DbParams.DbHost, ConfigParameters.DbParams.DbPort),
	}
	dbCon, err := sql.Open(ConfigParameters.DbParams.DbType, u.String())

	if err != nil {
		log.Fatal("Could not open Database handler: ", err)
	} else {
		Db = dbCon
		log.Println("Successfully created Database con")
		log.Println("Db handler valid: ", IsConOk())
	}
}

/**
Verify if the database connection handler is valid or not
*/
func IsConOk() bool {
	if err := Db.Ping(); err != nil {
		log.Fatal("Database connection is not valid, error: ", err)
		return false
	} else {
		return true
	}
}

/**
Load all the configuration parameters from the json file
*/
func loadParams() {
	params, _ := ioutil.ReadFile("/home/paritosh/go/src/github.com/paritosh-96/RestServer/config/setupParameters.json")
	ConfigParameters = Parameters{}
	_ = json.Unmarshal([]byte(params), &ConfigParameters)
	log.Println("Config parameters loaded...")
}
