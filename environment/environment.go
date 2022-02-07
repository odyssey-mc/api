package environment

import "os"

func GetApiKey() string {
	return os.Getenv("API_KEY")
}

func Test() bool {
	if len(GetApiKey()) == 0 {
		return false
	}
	return true
}