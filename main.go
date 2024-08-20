package main

import (
	"encoding/json"
	"fmt"
	"github.com/WebmiRU/YudolePlatformPackages/message"
	"github.com/WebmiRU/YudolePlatformPackages/module"
	"github.com/gorilla/websocket"
	"goodgame_client/client"
	"log"
	"os"
	"regexp"
	"strings"
)

var channels []string
var smiles = make(map[string]string)
var smileRegexp *regexp.Regexp

func loadSmiles() {
	data, err := os.ReadFile("data/smiles.json")
	if err != nil {
		log.Println("Error loading smiles.json:", err)
	}

	var smilesData []ChatSmile

	err = json.Unmarshal(data, &smilesData)
	if err != nil {
		log.Println("Error parsing smiles.json:", err)
	}

	for _, v := range smilesData {
		if len(v.Images.Gif) > 0 {
			smiles[":"+v.Key+":"] = v.Images.Gif
		} else {
			smiles[":"+v.Key+":"] = v.Images.Big
		}
	}
}

func processSmiles(text string) string {
	for _, v := range smileRegexp.FindAllString(text, -1) {
		if _, ok := smiles[v]; ok {
			fmt.Println(v)
			text = strings.ReplaceAll(text, v, "<img src='"+smiles[v]+"' alt='#'/>")
		}
	}

	return text
}

func ggClient() {
	c, _, err := websocket.DefaultDialer.Dial("wss://chat-1.goodgame.ru/chat2/", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	for {
		_, chatMessage, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
		}

		var msg ChatMessage
		json.Unmarshal(chatMessage, &msg)

		switch msg.Type {
		case "welcome":
			for _, channelId := range channels {

				m := ChatMessageJoin{
					Type: "join",
					Data: ChatMessageJoinData{
						ChannelId: channelId,
						Hidden:    0,
						Mobile:    false,
						Reload:    false,
					},
				}

				err := c.WriteJSON(&m)

				if err != nil {
					fmt.Println("write:", err)
					return
				}
			}

		case "success_join":
			// @TODO SuccessJoin

		case "channel_counters":
			// @TODO ChannelCounters

		case "message":
			var msg ChatMessageMessage
			err := json.Unmarshal(chatMessage, &msg)

			if err != nil {
				log.Println("unmarshal:", err)
			}

			client.WriteChan <- message.StreamChat{
				//Id:      "123",
				Type:    "stream/chat/message",
				Service: "goodgame",
				Module:  "goodgame_client",

				Payload: message.StreamChatPayload{
					//Id: "321",
					User: message.StreamChatUser{
						Login:  msg.Data.UserName,
						Nick:   msg.Data.UserName,
						Badges: nil,
					},
					Channel: "",
					Html:    processSmiles(msg.Data.Text),
					Text:    msg.Data.Text,
					Tags:    nil,
				},
			}

			fmt.Println(msg.Data.ChannelId, msg.Data.UserName, msg.Data.Text)

		default:
			log.Println("unknow type:", msg.Type)
		}

		fmt.Println(msg)
	}
}

func main() {
	var m module.Module
	if err := m.Load("./"); err != nil {
		log.Println(err)
	}

	for _, v := range m.Tabs["channels"].Fields["channels"].Value.([]any) {
		channels = append(channels, v.(string))
	}

	go client.Connect("127.0.0.1", 5127, message.StreamChat{})

	loadSmiles()

	smileRegexp, _ = regexp.Compile(":[A-Za-z0-9_]+:")

	ggClient()
}
