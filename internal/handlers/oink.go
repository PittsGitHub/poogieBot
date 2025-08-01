package handlers

import "github.com/bwmarrin/discordgo"

func HandleOink(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Oink oink! 🐷")
}
