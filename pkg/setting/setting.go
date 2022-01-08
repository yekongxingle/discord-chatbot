package setting

import (
	"encoding/json"
	"log"
	"os"
)

type ResponseDiscord struct {
	Chaturl string `json:"chaturl"`
	Posturl string `json:"posturl"`
	ChannelID string `json:"channel_id"`
	Authorization string `json:"Authorization"`
	Username string `json:"username"`
}
var Conf ResponseDiscord
func InitConfig()  {
	file, err := os.Open("config/conf.json")
	if err!=nil{
		log.Println(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Conf=ResponseDiscord{}
	err = decoder.Decode(&Conf)
	if err!=nil{
		log.Println(err)
	}
}
