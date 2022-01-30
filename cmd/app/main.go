package main

import (
	"fmt"
	randomgenerator "github.com/balram/rest-url-shortner/utils/random_generator"
	"log"
)

func main(){
	log.Println("Starting rest service - url shortner")
	random,_:=randomgenerator.New()
	fmt.Println(random.GenerateRandomString())
}