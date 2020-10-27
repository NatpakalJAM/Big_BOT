package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/bwmarrin/discordgo"
)

// message -> struct for message list
type message struct {
	Prefix   string `json:"prefix"`
	Response struct {
		Type string   `json:"type"`
		List []string `json:"list"`
	} `json:"response"`
}

// MessageCreate -> This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// get message list from file
	msgList := readMessageList()

	// loop list message
	for _, msg := range msgList {
		// check regex match message string
		match, _ := regexp.MatchString(msg.Prefix, m.Content)
		if match {
			// check message response type (accept: `normal` | `random`)
			switch msg.Response.Type {
			case "normal":
				for _, res := range msg.Response.List {
					s.ChannelMessageSend(m.ChannelID, res)
				}
			case "random":
				rand.Seed(time.Now().UnixNano())
				ranNum := rand.Intn(len(msg.Response.List))
				s.ChannelMessageSend(m.ChannelID, msg.Response.List[ranNum])
			}
			// break loop ...
			break
		}
	}
}

func readMessageList() (messageList []message) {
	fileName := "./message/message.json"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	byteValue, _ := ioutil.ReadAll(f)
	json.Unmarshal(byteValue, &messageList)
	f.Close()
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	return messageList
}
