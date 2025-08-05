package mhwildhandlers

import (
	"fmt"
	"strings"

	"github.com/PittsGitHub/poogieBot/internal/commands/mhwildcommands"
	"github.com/PittsGitHub/poogieBot/internal/data/mhwildsdata"
	"github.com/PittsGitHub/poogieBot/internal/services"
	"github.com/PittsGitHub/poogieBot/internal/services/mhwildservices"
	"github.com/bwmarrin/discordgo"
)

func HandleFind(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
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
	itemType := services.Normalise(parts[1])
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
	// these rarity value collections need to be moved further down
	// as the current method is correct and works for Weapons / Armours
	// for decorations and talismans the ruleset doesn't apply
	rarityValues, err := mhwildservices.ResolveRarityValues(itemRank)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	//#
	// Step 3. obtain a collection of matching items of the required rarity that have the desired skill
	//#
	switch itemType {
	case "armor", "armour":
		mhwildcommands.FindArmor(rarityValues, s, m, skillID, itemRank, skillName, itemType)
		return
	case "weapon", "weapons",
		"greatsword", "longsword", "swordandshield", "dualblades",
		"hammer", "huntinghorn", "lance", "gunlance",
		"switchaxe", "chargeblade", "insectglaive",
		"lightbowgun", "heavybowgun", "bow":
		mhwildcommands.FindWeapon(rarityValues, s, m, skillID, itemRank, skillName, itemType)
		return

	case "talisman":
		mhwildcommands.FindHighestRankTalismanWithDesiredSkill(rarityValues, s, m, skillID, itemRank, skillName, itemType)
		return
	case "decoration":
		mhwildcommands.FindDecoration(rarityValues, s, m, skillID, itemRank, skillName, itemType)
		return
	default:
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ item type: %s not recognised", itemType))
		return
	}

}
