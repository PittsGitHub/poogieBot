package mhwildhandlers

import "github.com/bwmarrin/discordgo"

func HandlePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong! 🐽")
}
