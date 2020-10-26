package handler

import "github.com/bwmarrin/discordgo"

// MessageCreate -> This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "คะเมียวตําปรู๊ช" {
		s.ChannelMessageSend(m.ChannelID, "คะเมียวตําปร๊าส")
	}

	if m.Content == "สาวๆเดินผ่านมา" {
		s.ChannelMessageSend(m.ChannelID, "อะชิชิป่ะเห้ย์~")
	}
}
