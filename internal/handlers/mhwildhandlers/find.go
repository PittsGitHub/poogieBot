package mhwildhandlers

import (
	"fmt"
	"strings"

	"github.com/PittsGitHub/poogieBot/internal/data/mhwildsdata"
	"github.com/PittsGitHub/poogieBot/internal/services"
	"github.com/PittsGitHub/poogieBot/internal/services/mhwildservices"
	"github.com/bwmarrin/discordgo"
)

func FindCommand(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
	//#
	// Step 1. sanitise user inputs
	//#

	//remove !find from the message and use , as the delimeter
	raw := strings.TrimPrefix(m.Content, "!find")
	parts := strings.Split(raw, ",")

	//assert we have 3 parts if not return with message
	if len(parts) != 3 {
		s.ChannelMessageSend(m.ChannelID, "Usage: !find [rank or rarity], [type or subtype], [skill name]")
		s.ChannelMessageSend(m.ChannelID, "Example: !find high, armor, part breaker")
		return
	}

	//santise our parts
	itemRank := services.Normalise(parts[0])
	//itemType := mhwildservices.Normalise(parts[1]) might not be needed yet with the new lazy loader... maybe?
	skillName := services.FormatTitleCase(parts[2])

	//#
	// Step 2. obtain item collection file paths, desired skill names Id, collection of rarity containers
	//#

	//collect the skill id from skill name
	skillID, err := mhwildsdata.GetSkillIDFromName(skillName)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	// create a collection of each rarity value as int
	rarityValues, err := mhwildservices.ResolveRarityValues(itemRank)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	//#
	// Step 3. obtain a collection of matching items of the required rarity that have the desired skill
	//#
	foundArmor, err := mhwildsdata.GetArmorGroupedByRarity(rarityValues)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	filteredArmor := mhwildsdata.FilterArmorBySkillID(foundArmor, skillID)

	var rankValue string

	switch itemRank {
	case "high":
		rankValue = "high rank"
	case "low":
		rankValue = "low rank"
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10":
		rankValue = fmt.Sprintf("rarity %s", itemRank)
	default:
		rankValue = ""
	}

	if len(filteredArmor) == 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ No %s armor found with %s", rankValue, skillName))
		return
	}

	message := mhwildservices.BuildArmorSkillSummaryMessage(filteredArmor)
	s.ChannelMessageSend(m.ChannelID, message)

}
