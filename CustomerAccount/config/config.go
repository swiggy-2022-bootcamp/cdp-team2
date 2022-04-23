package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var Server = map[string]string{}
var AWS = map[string]string{}



func FromEnv()  {

	cfgFilename := "../../etc/config.localhost.env"
	if err := godotenv.Load(cfgFilename); err != nil {
		fmt.Println("No config files found to load to env.",err.Error())
	} 
	Server["PORT"]="8092"
	os.Getenv("PORT")
	Server["CART_PORT"]="9000"
	//os.Getenv("CART_PORT")
	Server["REWARD_PORT"]="9000"
	//os.Getenv("REWARD_PORT")
	Server["AUTH_PORT"]="9000"
	//os.Getenv("AUTH_PORT")
	AWS["TABLE_NAME"]="team-2-Customers"
	//os.Getenv("TABLE_NAME")
	AWS["region"]="us-west-2"
	//os.Getenv("region")
	AWS["endpoint"]="https://dynamodb.us-west-2.amazonaws.com"
	//os.Getenv("endpoint")

}


// func main(){
// 	FromEnv()
// 	fmt.Println(Server)
// 	fmt.Println(AWS)
// }