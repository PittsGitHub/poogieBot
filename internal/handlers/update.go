package handlers

import (
	"os"

	"github.com/PittsGitHub/poogieBot/internal/commands"
	"github.com/bwmarrin/discordgo"
)

func HandleUpdateMHWilds(s *discordgo.Session, m *discordgo.MessageCreate) {
	ownerID := os.Getenv("OWNER_ID")
	if m.Author.ID != ownerID {
		s.ChannelMessageSend(m.ChannelID, "🚫 Oink!? You are not permitted to do that.")
		return
	}

	output, err := commands.RunUpdateScript("./scripts/update-mhwilds.sh")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "❌ Rip. Update failed:\n"+err.Error())
		return
	}

	s.ChannelMessageSend(m.ChannelID, "✅ Oink! Update complete:\n```"+output+"```")
}
