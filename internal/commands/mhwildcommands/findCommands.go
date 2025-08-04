package mhwildcommands

import (
	"fmt"

	"github.com/PittsGitHub/poogieBot/internal/data/mhwildsdata"
	"github.com/PittsGitHub/poogieBot/internal/services/mhwildservices"
	"github.com/bwmarrin/discordgo"
)

func FindArmor(rarityValues []int, s *discordgo.Session, m *discordgo.MessageCreate, skillID string, itemRank string, skillName string, itemType string) {
	foundArmor, err := mhwildsdata.GetArmorGroupedByRarity(rarityValues)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
	}

	filteredArmor := mhwildsdata.FilterArmorBySkillID(foundArmor, skillID)
	if len(filteredArmor) == 0 {
		failedToFindItem(itemRank, s, m, itemType, skillName)
	}

	message := mhwildservices.BuildArmorSkillSummaryMessage(filteredArmor)
	s.ChannelMessageSend(m.ChannelID, message)
}

func FindWeapon(rarityValues []int, s *discordgo.Session, m *discordgo.MessageCreate, skillID string, itemRank string, skillName string, itemType string) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("type: %s not implemented yet", itemType))

}

func FindTalisman(rarityValues []int, s *discordgo.Session, m *discordgo.MessageCreate, skillID string, itemRank string, skillName string, itemType string) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("type: %s not implemented yet", itemType))
}

func FindDecoration(rarityValues []int, s *discordgo.Session, m *discordgo.MessageCreate, skillID string, itemRank string, skillName string, itemType string) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("type: %s not implemented yet", itemType))
}

func failedToFindItem(itemRank string, s *discordgo.Session, m *discordgo.MessageCreate, itemType string, skillName string) {
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

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ No %s %s found with %s", rankValue, itemType, skillName))
}
