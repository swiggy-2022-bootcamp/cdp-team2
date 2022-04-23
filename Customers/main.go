package main

import (
    // Import godotenv
  "github.com/joho/godotenv"
  "fmt"
	"os"
	"log"
)


// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  PORT= os.Getenv("PORT")
  GRPCPORT=os.Getenv("GRPCPORT")
  CARTPORT=os.Getenv("CARTPORT")
  REWARDPORT=os.Getenv("REWARDPORT")
}

func main() {
    // os package


  // godotenv package
  dotenv := goDotEnvVariable("STRONGEST_AVENGER")

  fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", dotenv)
}
