package util

import (
	"dbo-test/config"
	"fmt"
	"log"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	var err error
	config.Configure, err = config.GetConfig("../config/config.yml")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(GenerateToken("test name", 1))
}
