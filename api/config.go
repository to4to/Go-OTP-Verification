package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal("ACCOUNT_SID"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("err")
		log.Fatal("Error in Handling .env")

	}

	return os.Getenv("TWILIO_ACCOUNT_SID")

}

func envAUTHTOKEN() string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("err")
		log.Fatal("Error in Handling .env")

	}

	return os.Getenv("TWILIO_AUTHTOKEN")

}

func envSERVICESID() string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("err")
		log.Fatal("Error in Handling .env")

	}

	return os.Getenv("TWILIO_SERVICES_ID")

}
