package mhwildcommands

import (
	"fmt"

	"github.com/PittsGitHub/poogieBot/internal/data/mhwildsdata"
	"github.com/PittsGitHub/poogieBot/internal/services/mhwildservices"
	"github.com/bwmarrin/discordgo"
)

func FindArmor(rarityValues []int, s *discordgo.Session, m *discordgo.MessageCreate, skillID string, itemRank string, skillName string, itemType string) {
	foundRarityMatchedArmor, err := mhwildsdata.GetArmorGroupedByRarity(rarityValues)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
	}

	filteredArmorBySkillName := mhwildsdata.FilterArmorBySkillID(foundRarityMatchedArmor, skillID)
	if len(filteredArmorBySkillName) == 0 {
		failedToFindItem(itemRank, s, m, itemType, skillName)
		return
	}

	message := mhwildservices.BuildArmorSkillSummaryMessage(filteredArmorBySkillName)
	s.ChannelMessageSend(m.ChannelID, message)
}

func FindWeapon(rarityValues []int, s *discordgo.Session, m *discordgo.MessageCreate, skillID string, itemRank string, skillName string, itemType string) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("type: %s not implemented yet", itemType))

	//create a map of what functions to load ie weapons, weapon, greatsword, bow etc.

	//load all weapons of desired rarity

	//filter those weapons  based on skill name

	//print the weapon list back to discord

}

func FindHighestRankTalismanWithDesiredSkill(rarityValues []int, s *discordgo.Session, m *discordgo.MessageCreate, skillID string, itemRank string, skillName string, itemType string) {

	filteredTalismanBySkillName, err := mhwildsdata.FilterTalismanBySkill(skillID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
	}
	if len(filteredTalismanBySkillName) == 0 {
		failedToFindItem(itemRank, s, m, itemType, skillName)
		return
	}
	message := mhwildservices.BuildTalismanSkillSummaryMessage(filteredTalismanBySkillName, skillName)
	s.ChannelMessageSend(m.ChannelID, message)
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

	if itemType == "talisman" {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ No %s found with %s", itemType, skillName))
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ No %s %s found with %s", rankValue, itemType, skillName))
	}
}
