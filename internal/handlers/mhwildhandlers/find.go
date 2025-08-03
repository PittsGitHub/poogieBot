package mhwildhandlers

import (
	"fmt"
	"strings"

	"log"

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
		s.ChannelMessageSend(m.ChannelID, "Example: !find high, greatsword, part breaker")
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
	log.Printf("LOL PRINT skillID: %s ", skillID)

	// create a collection of each rarity value as int
	rarityValues, err := mhwildservices.ResolveRarityValues(itemRank)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	log.Printf("Resolved rarity values for '%s':", itemRank)
	for _, val := range rarityValues {
		log.Println(" -", val)
	}

	//#
	// Step 3. obtain a collection of matching items of the required rarity that have the desired skill
	//#

	//We collect all the armor sets that match the desired rarity value
	//The rarity value is determined by each int entry in the rarityValues var which will be a passed to this func as a int[]
	//This function needs to return a collection for each passed rarity value, inside each collection is each armor object matching the rarity
	//We don't need the armor pieces of each armor yet I assume we don't need to provide it to our armors to create an object of that type?

	foundArmor, err := mhwildsdata.GetArmorGroupedByRarity(rarityValues)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	// Optional debug output
	for rarity, armorList := range foundArmor {
		fmt.Printf("Rarity %d:\n", rarity)
		for _, armor := range armorList {
			name := armor.Names["en"]
			if name == "" {
				name = "[Unnamed]"
			}
			fmt.Printf("  - %s\n", name)

			if len(armor.Pieces) > 0 {
				firstPiece := armor.Pieces[0]
				pieceName := firstPiece.Names["en"]
				if pieceName == "" {
					pieceName = "[Unnamed Piece]"
				}
				fmt.Printf("    First Piece: %s (%s)\n", pieceName, firstPiece.Kind)
			} else {
				fmt.Println("    No pieces found.")
			}
		}
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

	// Debug output
	for rarity, armorList := range filteredArmor {
		fmt.Printf("Rarity %d:\n", rarity)
		for _, armor := range armorList {
			armorName := armor.Names["en"]
			if armorName == "" {
				armorName = "[Unnamed Armor]"
			}
			fmt.Printf("  - %s\n", armorName)

			if len(armor.Pieces) > 0 {
				for _, piece := range armor.Pieces {
					pieceName := piece.Names["en"]
					if pieceName == "" {
						pieceName = "[Unnamed Piece]"
					}
					fmt.Printf("    Piece: %s (%s)\n", pieceName, piece.Kind)
				}
			} else {
				fmt.Println("    [Set/Group bonus match]")
			}
		}
	}

	message := mhwildservices.BuildArmorSkillSummaryMessage(filteredArmor)
	s.ChannelMessageSend(m.ChannelID, message)

}
