package proxy

// "github.com/joho/godotenv"

var PRIVATE_KEY string
var CERTIFICATE string
var ENABLE_TLS string
var FORBIDDEN_HOSTS_PATH string = "proxy/data/forbidden-hosts.txt"
var BANNED_WORDS_PATH string = "proxy/data/banned-words.txt"

// func LoadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	PRIVATE_KEY = os.Getenv("PRIVATE_KEY")
// 	CERTIFICATE = os.Getenv("CERTIFICATE")
// 	ENABLE_TLS = os.Getenv("ENABLE_TLS")
// }
