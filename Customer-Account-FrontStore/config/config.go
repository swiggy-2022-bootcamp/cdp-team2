package config


var Server = map[string]string{
"PORT":"8091",
"CART_PORT":"9010",
"REWARD_PORT":"9011",
"AUTH_PORT":"9012",
"GRPC_PORT":"8092",
}
var AWS = map[string]string{
"region":"us-west-2",
"endpoint":"https://dynamodb.us-west-2.amazonaws.com",
"TABLE_NAME":"team-2-Customers",
}



// func FromEnv()  {

// 	cfgFilename := "../../etc/config.localhost.env"
// 	if err := godotenv.Load(cfgFilename); err != nil {
// 		fmt.Println("No config files found to load to env.",err.Error())
// 	} 
// 	Server["PORT"]="8092"
// 	os.Getenv("PORT")
// 	Server["CART_PORT"]="9000"
// 	//os.Getenv("CART_PORT")
// 	Server["REWARD_PORT"]="9000"
// 	//os.Getenv("REWARD_PORT")
// 	Server["AUTH_PORT"]="9000"
// 	//os.Getenv("AUTH_PORT")
// 	//os.Getenv("endpoint")
// }


// func main(){
// 	FromEnv()
// 	fmt.Println(Server)
// 	fmt.Println(AWS)
// }