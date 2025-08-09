package main

import (
	"log"
	"os"

	"github.com/PittsGitHub/poogieBot/bot"
	"github.com/PittsGitHub/poogieBot/internal/commands/mhwildcommands"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN not set in environment")
	}

	// Ensure data is present and fresh
	log.Println("🐷 Checking for latest mhwilds data...")
	output, err := mhwildcommands.RunUpdateScript("./scripts/update-mhwilds.sh")
	if err != nil {
		log.Fatalf("❌ Failed to update mhwilds data:\n%s\n%s", output, err)
	}
	log.Println("✅ Data update complete.")

	bot.Start(token)
}
