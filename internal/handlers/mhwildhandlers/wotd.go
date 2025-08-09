package mhwildhandlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PittsGitHub/poogieBot/internal/types/mhwildtypes"
	"github.com/bwmarrin/discordgo"
)

func HandleRandomWeapon(s *discordgo.Session, m *discordgo.MessageCreate) {
	dir := "data/mhwilds/weapons"

	// Load skills
	skillData, err := os.ReadFile("data/mhwilds/Skill.json")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "❌ Couldn't read skill data.")
		return
	}

	var skills []mhwildtypes.Skill
	if err := json.Unmarshal(skillData, &skills); err != nil {
		s.ChannelMessageSend(m.ChannelID, "❌ Failed to parse skill data.")
		return
	}

	skillMap := make(map[int]mhwildtypes.Skill)
	for _, sk := range skills {
		skillMap[sk.GameID] = sk
	}

	// Load weapon files (only .json)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "❌ Couldn't load weapon directory.")
		return
	}

	var jsonFiles []os.FileInfo
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			jsonFiles = append(jsonFiles, file)
		}
	}
	if len(jsonFiles) == 0 {
		s.ChannelMessageSend(m.ChannelID, "❌ No valid weapon files found.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomFile := jsonFiles[rand.Intn(len(jsonFiles))]
	fullPath := filepath.Join(dir, randomFile.Name())

	data, err := os.ReadFile(fullPath)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "❌ Failed to read weapon file.")
		return
	}

	var weapons []mhwildtypes.Weapon
	if err := json.Unmarshal(data, &weapons); err != nil || len(weapons) == 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("❌ Failed to parse or empty file: %s", randomFile.Name()))
		return
	}

	randomWeapon := weapons[rand.Intn(len(weapons))]
	name := randomWeapon.Names["en"]
	if name == "" {
		s.ChannelMessageSend(m.ChannelID, "❌ Weapon entry missing English name.")
		return
	}

	// Format skills
	skillStrings := []string{}
	for rawID, level := range randomWeapon.Skills {
		id, err := strconv.Atoi(rawID)
		if err != nil {
			continue
		}
		skill, found := skillMap[id]
		if !found {
			continue
		}
		skillStrings = append(skillStrings,
			fmt.Sprintf("%s: %d/%d", skill.Names["en"], level, len(skill.Ranks)))
	}
	skillsLine := "Skills.\nnone."
	if len(skillStrings) > 0 {
		skillsLine = "Skills.\n" + strings.Join(skillStrings, "\n")
	}

	// Format URL
	wikiName := strings.ReplaceAll(name, " ", "+")
	url := fmt.Sprintf("https://monsterhunterwilds.wiki.fextralife.com/%s", wikiName)

	// Send final message
	message := fmt.Sprintf(
		"🗡️ **%s** (Rarity %d)\nAttack: %d\nAffinity: %d%%\n%s\n🔗 %s",
		name,
		randomWeapon.Rarity,
		randomWeapon.AttackRaw,
		randomWeapon.Affinity,
		skillsLine,
		url,
	)

	s.ChannelMessageSend(m.ChannelID, message)
}
