package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleMisc(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch strings.ToLower(m.Content) {
	case "!beck", "!renn":
		s.ChannelMessageSend(m.ChannelID, "What a pig! 🐷")
	case "!dan":
		s.ChannelMessageSend(m.ChannelID, "What a guy! 😎")
	case "!bilbo":
		s.ChannelMessageSend(m.ChannelID, "Fine boy! 🐶")
	default:
		s.ChannelMessageSend(m.ChannelID, "Oink?")
	}
}
