package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/PittsGitHub/poogieBot/internal/handlers/mhwildhandlers"
	"github.com/bwmarrin/discordgo"
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

	if strings.HasPrefix(m.Content, "!find") {
		args := strings.Fields(m.Content)
		mhwildhandlers.FindCommand(s, m, args)
		return
	}

	switch m.Content {
	case "!ping":
		mhwildhandlers.HandlePing(s, m)
	case "!oink":
		mhwildhandlers.HandleOink(s, m)
	case "!update-mhwilds":
		mhwildhandlers.HandleUpdateMHWilds(s, m)
	case "!beck", "!renn", "!dan", "!bilbo":
		mhwildhandlers.HandleMisc(s, m)
	case "!wotd":
		mhwildhandlers.HandleRandomWeapon(s, m)
	}
}
