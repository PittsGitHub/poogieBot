package handlers

import (
	"strings"

	"log"

	"github.com/PittsGitHub/poogieBot/internal/resourcemap"
	"github.com/bwmarrin/discordgo"
	"github.com/google/shlex"
)

func FindCommand(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
	args, err := shlex.Split(m.Content)
	if err != nil || len(args) < 4 {
		s.ChannelMessageSend(m.ChannelID, "Usage: !find [rank] [type] \"Skill Name\"")
		return
	}

	rank := strings.ToLower(args[1])
	itemType := strings.Trim(args[2], `"`)
	skillName := strings.Join(args[3:], " ") // Join in case skill name has spaces
	skillName = strings.Trim(skillName, `"`)

	// Step 1: Normalize the item type
	itemTypeLower := strings.ToLower(itemType)

	// Step 2: Determine what files to search
	var filesToSearch []string

	switch itemTypeLower {
	case "armor", "decoration", "talisman":
		if path, ok := resourcemap.CoreItemFiles[itemTypeLower]; ok {
			filesToSearch = append(filesToSearch, path)
		}
	case "weapon":
		filesToSearch = resourcemap.GetAllItemFiles() // Or split out GetAllWeaponFiles if preferred
	default:
		// Maybe it's a specific weapon type
		if path, ok := resourcemap.WeaponFileMap[itemType]; ok {
			filesToSearch = append(filesToSearch, path)
			if itemType == "Hunting Horn" {
				filesToSearch = append(filesToSearch, resourcemap.HuntingHornVariants...)
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Unknown item type: "+itemType)
			return
		}
	}

	// Step 3: Stub — search the `filesToSearch` for `skillName` and apply `rank` filter
	log.Printf("Searching for skill: %s in rank: %s within: %v", skillName, rank, filesToSearch)
	s.ChannelMessageSend(m.ChannelID, "Found item type: "+itemType)
	// TODO: Load each JSON file in filesToSearch
	// TODO: Match entries that include the specified skill
	// TODO: Filter by rarity based on rank (e.g., high = rarity 5+)
	// TODO: Group and send results to channel
}
