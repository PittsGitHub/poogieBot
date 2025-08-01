package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PittsGitHub/poogieBot/internal/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Start(token string) {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	dg.AddHandler(onReady)

	dg.AddHandler(onMessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("PoogieBot is up! Press Ctrl+C to shut it down.")

	// Wait for termination
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	fmt.Println("PoogieBot shutting down...")
	dg.Close()
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("Logged in as: %v#%v\n", s.State.User.Username, s.State.User.Discriminator)
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	switch m.Content {
	case "!ping":
		s.ChannelMessageSend(m.ChannelID, "Pong! 🐽")
	case "!oink":
		s.ChannelMessageSend(m.ChannelID, "Oink oink! 🐷")
	case "!beck":
		s.ChannelMessageSend(m.ChannelID, "What a pig! 🐷")
	case "!renn":
		s.ChannelMessageSend(m.ChannelID, "What a pig! 🐷")
	case "!dan":
		s.ChannelMessageSend(m.ChannelID, "What a guy! 😎")
	case "!update-mhwilds":
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

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
}
