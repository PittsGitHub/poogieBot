package mhwildhandlers

import (
	"fmt"
	"strings"

	"log"

	"github.com/PittsGitHub/poogieBot/internal/services/mhwildservices"
	"github.com/bwmarrin/discordgo"
)

func FindCommand(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
	// Step 1. arrange sanitise input

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
	itemRank := mhwildservices.Normalise(parts[0])
	itemType := mhwildservices.Normalise(parts[1])
	skillName := mhwildservices.FormatTitleCase(parts[2])

	// Step 2. obtain item collections and Skill name Id

	//collect the file locations for all item types requested
	itemCollectionFilePaths, err := mhwildservices.GetItemTypeDataFilePaths(itemType)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	//collect the skill id from skill name
	skillID, err := mhwildservices.GetSkillIDFromName(skillName)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	// Step 3. obtain a collection of matching items of the required rarity that have the desired skill

	// obtain a collection of each rarity value as a string
	rarityValues, err := mhwildservices.ResolveRarityValues(itemRank)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	//DEBUG output rarity values to logs
	log.Printf("Resolved rarity values for '%s':", itemRank)
	for _, val := range rarityValues {
		log.Println(" -", val)
	}

	//we need to refine below

	// obtain a collection of desired item types of the rarity value
	//STUB
	// we write a json parser service, we take in a collection of strings that are file paths and a collection of rarity values
	//we loop through each json file found in the file path collection
	//each file we loop through we check it's rarity value and see if it exists in our collection of raritys
	//if a match is found we add that object to our collection of items
	//we continue looping through all files adding objects based on rarity

	//we now have a collection of all requested item types with the requested rank or rarity value
	//this will be our itemsWithDesiredRarity collection

	//we now want to create a new collection itemswithDesiredRarityAndSkill objects from our itemsWithDesiredRarity collection
	//STUB we write a service to return a collection of objects itemswithDesiredRarityAndSkill
	//we loop  through the itemsWithDesiredRarity collection and put any objects found matching skill id into itemswithDesiredRarityAndSkill
	// itemswithDesiredRarityAndSkill collection is then returned

	// Step 4. sort the itemswithDesiredRarityAndSkill to be formatted and returned as a message

	//STUB service....
	//break the itemswithDesiredRarityAndSkill collection into smaller collections each grouped by their rarity
	//so in effect... we end up with a collection of collections with unique rarity but all have the desired skill
	//we then go through each of these sub collections sorting them
	//so the items with the most skill level of the desired skill appear first in the collection or ordered list
	//this is the skills object note the :1 denotes skill level
	//   "skills": {
	//   "-481419552": 1,
	//   "850626240": 1
	// },

	//STUB service...
	//now we have our collection with sub collections we need to build the returned message(s)
	//we build a message with highest rarity collection grouped items at the top of the message
	// each line has an item it's rarity, item name, skill name, skill level
	// so it's looks like:
	// Rarity 8, Steves Shoes, Part Breaker, 2
	// Rarity 8, Steves Fist, Part Breaker, 1
	// Rarity 7, MonsterMon Tooth, Part Breaker, 3
	// Rarity 5, Piddly Helmet, Part Breaker 5

	// place holder return value and logging
	// actual logic to follow
	log.Printf("Files to search: %+v", itemCollectionFilePaths)
	rarity := mhwildservices.Normalise(itemRank)
	rank := mhwildservices.Normalise(itemRank)

	if rank == "high" || rank == "low" {
		log.Printf("Parsed input — Rank: %s | Type: %s | Skill: %s | Skill ID %s", rank, itemType, skillName, skillID)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Searching for:\nRank: %s | Type: %s | Skill: %s ", rank, itemType, skillName))
		return
	} else {
		log.Printf("Parsed input — Rarity: %s | Type: %s | Skill: %s | Skill ID %s", rarity, itemType, skillName, skillID)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Searching for:\nRarity: %s | Type: %s | Skill: %s", rarity, itemType, skillName))
		return
	}
}
