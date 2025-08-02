package handlers

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	findType "github.com/PittsGitHub/poogieBot/internal/types/find"
	"github.com/bwmarrin/discordgo"
)

func FindCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// Split the raw message to extract rank, type, subtype (optional)
	parts := strings.Fields(m.Content)
	if len(parts) < 4 {
		s.ChannelMessageSend(m.ChannelID, "Usage: !find [rank] [type] [optional subtype] \"skill name\"")
		return
	}

	rankArg := strings.ToLower(parts[1])
	itemType := strings.ToLower(parts[2])
	subType := ""
	if len(parts) >= 5 && !strings.HasPrefix(parts[4], "\"") {
		subType = strings.ToLower(parts[3])
	}

	// Handle rarity/rank filter
	rarityFilter := 5 // default to high rank
	switch rankArg {
	case "low":
		rarityFilter = 0
	case "high":
		rarityFilter = 5
	default:
		rarity, err := strconv.Atoi(rankArg)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Invalid rarity or rank. Use `low`, `high`, or a rarity number.")
			return
		}
		rarityFilter = rarity
	}

	// Extract quoted skill name
	skillName := extractQuotedSkill(m.Content)
	if skillName == "" {
		s.ChannelMessageSend(m.ChannelID, "Please include a skill in quotes. Example: `\"critical boost\"`")
		return
	}

	// Resolve skill ID
	skillID, matchedSkillName, err := resolveSkillID(skillName)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Skill not found: %s", skillName))
		return
	}

	var weapons []findType.Weapon
	var armors []findType.Armor

	switch itemType {
	case "weapon":
		rawWeapons, err := findWeaponsWithSkill(skillID, subType, 0) // no filter yet
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error while searching weapons: "+err.Error())
			return
		}
		convertedWeapons := convertWeapons(rawWeapons)
		weapons = filterWeaponsBySkillAndRarity(convertedWeapons, skillID, rarityFilter)

	case "armor":
		rawArmors, err := findArmorWithSkill(skillID, 0) // no filter yet
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error finding armor: %v", err))
			return
		}
		convertedArmors := convertArmors(rawArmors)
		armors = filterArmorsBySkillAndRarity(convertedArmors, skillID, rarityFilter)

	default:
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unknown item type: %s", itemType))
		return
	}

	// Format and send result
	output := formatFindResults(weapons, armors, matchedSkillName, itemType)
	s.ChannelMessageSend(m.ChannelID, output)
}

func HandleFindComm5and(input string) string {
	parts := strings.Fields(input)

	if len(parts) < 4 {
		return "❌ Invalid format. Try: `!find high weapon greatsword \"skill name\"`"
	}

	rankKeyword := parts[0]
	itemType := strings.ToLower(parts[1])
	subType := strings.ToLower(parts[2])
	skillNameQuoted := strings.Join(parts[3:], " ")

	skillName := strings.Trim(skillNameQuoted, `"'`)
	minRarity := getRarityFromKeyword(rankKeyword)

	if itemType != "weapon" && itemType != "armor" {
		return "❌ Only `weapon` and `armor` types are supported right now."
	}

	skillID, err := findSkillIDByName(skillName)
	if err != nil {
		return fmt.Sprintf("❌ Skill '%s' not found.", skillName)
	}

	return fmt.Sprintf("🔍 Searching for %s %s with skill **%s** (ID: %d, rarity ≥ %d)",
		strings.Title(subType), itemType, skillName, skillID, minRarity)
}

func findSkillIDByName(name string) (int64, error) {
	file, err := os.Open("./data/mhwilds/Skill.json")
	if err != nil {
		return 0, fmt.Errorf("failed to open Skill.json: %w", err)
	}
	defer file.Close()

	var skills []findType.Skill
	if err := json.NewDecoder(file).Decode(&skills); err != nil {
		return 0, fmt.Errorf("failed to decode Skill.json: %w", err)
	}

	nameLower := strings.ToLower(name)
	for _, skill := range skills {
		for _, localized := range skill.Names {
			if strings.ToLower(localized) == nameLower {
				return skill.GameID, nil
			}
		}
	}
	return 0, fmt.Errorf("skill not found")
}

func getRarityFromKeyword(keyword string) int {
	switch strings.ToLower(keyword) {
	case "low":
		return 1
	case "high":
		return 5
	case "master":
		return 9
	default:
		return 1 // Default to lowest rarity if unrecognized
	}
}

func loadWeaponsFromPath(path string) ([]findType.Weapon, error) {
	var weapons []findType.Weapon

	err := filepath.Walk(path, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			file, err := os.Open(p)
			if err != nil {
				return err
			}
			defer file.Close()

			var w []findType.Weapon
			if err := json.NewDecoder(file).Decode(&w); err != nil {
				return err
			}

			weapons = append(weapons, w...)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error loading weapons: %w", err)
	}
	return weapons, nil
}

func filterWeaponsBySkillAndRarity(weapons []findType.Weapon, skillID string, minRarity int) []findType.Weapon {
	var results []findType.Weapon
	for _, w := range weapons {
		if w.Rarity >= minRarity {
			for sid := range w.Skills {
				if sid == skillID {
					results = append(results, w)
					break
				}
			}
		}
	}
	return results
}

func loadArmor() ([]findType.Armor, error) {
	file, err := os.Open("./data/mhwilds/findType.Armor.json")
	if err != nil {
		return nil, fmt.Errorf("error opening findType.Armor.json: %w", err)
	}
	defer file.Close()

	var armors []findType.Armor
	if err := json.NewDecoder(file).Decode(&armors); err != nil {
		return nil, fmt.Errorf("error decoding findType.Armor.json: %w", err)
	}

	return armors, nil
}

func filterArmorsBySkillAndRarity(armors []findType.Armor, skillID string, minRarity int) []findType.Armor {
	var results []findType.Armor
	for _, a := range armors {
		if a.Rarity >= minRarity {
			for _, piece := range a.Pieces {
				for sid := range piece.Skills {
					if sid == skillID {
						results = append(results, a)
						break
					}
				}
			}
		}
	}
	return results
}

func formatFindResults(weapons []findType.Weapon, armors []findType.Armor, skillName string, itemType string) string {
	var sb strings.Builder

	title := fmt.Sprintf("**Items with skill: %s**", skillName)
	sb.WriteString(title + "\n")

	switch itemType {
	case "weapon":
		sb.WriteString("**Weapons:**\n")
		if len(weapons) == 0 {
			sb.WriteString("_No matching weapons found._\n")
		} else {
			grouped := groupWeaponsByRarity(weapons)
			for rarity := 10; rarity >= 1; rarity-- {
				if group, exists := grouped[rarity]; exists {
					sb.WriteString(fmt.Sprintf("🟨 Rarity %d:\n", rarity))
					for _, w := range group {
						sb.WriteString(fmt.Sprintf("- %s\n", w.Names["en"]))
					}
				}
			}
		}
	case "armor":
		sb.WriteString("**findType.Armor Sets:**\n")
		if len(armors) == 0 {
			sb.WriteString("_No matching armor sets found._\n")
		} else {
			grouped := groupArmorByRarity(armors)
			for rarity := 10; rarity >= 1; rarity-- {
				if group, exists := grouped[rarity]; exists {
					sb.WriteString(fmt.Sprintf("🟦 Rarity %d:\n", rarity))
					for _, a := range group {
						sb.WriteString(fmt.Sprintf("- %s\n", a.Names["en"]))
					}
				}
			}
		}
	default:
		sb.WriteString("_Unsupported item type._")
	}

	return sb.String()
}

func groupWeaponsByRarity(items []findType.Weapon) map[int][]findType.Weapon {
	grouped := make(map[int][]findType.Weapon)
	for _, i := range items {
		grouped[i.Rarity] = append(grouped[i.Rarity], i)
	}
	return grouped
}

func groupArmorByRarity(items []findType.Armor) map[int][]findType.Armor {
	grouped := make(map[int][]findType.Armor)
	for _, i := range items {
		grouped[i.Rarity] = append(grouped[i.Rarity], i)
	}
	return grouped
}

func extractQuotedSkill(input string) string {
	re := regexp.MustCompile(`["']([^"']+)["']`)
	match := re.FindStringSubmatch(input)
	if len(match) > 1 {
		return strings.ToLower(match[1]) // normalize for matching
	}
	return ""
}
func resolveSkillID(name string) (string, string, error) {
	data, err := os.ReadFile("./data/mhwilds/Skill.json")
	if err != nil {
		return "", "", err
	}

	var skills []struct {
		GameID int               `json:"game_id"`
		Names  map[string]string `json:"names"`
	}
	if err := json.Unmarshal(data, &skills); err != nil {
		return "", "", err
	}

	for _, skill := range skills {
		for _, localized := range skill.Names {
			if strings.EqualFold(localized, name) {
				return fmt.Sprint(skill.GameID), localized, nil
			}
		}
	}

	return "", "", fmt.Errorf("skill not found")
}
func findWeaponsWithSkill(skillID string, weaponType string, rarityFilter int) ([]map[string]interface{}, error) {
	var weaponFiles []string

	// Decide what to scan
	if weaponType == "" {
		files, err := filepath.Glob("./data/mhwilds/weapons/*.json")
		if err != nil {
			return nil, err
		}
		weaponFiles = files
	} else {
		weaponFiles = []string{fmt.Sprintf("./data/mhwilds/weapons/%s.json", weaponType)}
	}

	var results []map[string]interface{}

	for _, file := range weaponFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			continue // skip broken files
		}

		var weapons []map[string]interface{}
		if err := json.Unmarshal(data, &weapons); err != nil {
			continue
		}

		for _, weapon := range weapons {
			// Match skill ID
			skills, ok := weapon["skills"].(map[string]interface{})
			if !ok {
				continue
			}
			if _, hasSkill := skills[skillID]; !hasSkill {
				continue
			}

			// Optional rarity match
			if rarityFilter > 0 {
				if weaponRarity, ok := weapon["rarity"].(float64); ok && int(weaponRarity) != rarityFilter {
					continue
				}
			}

			results = append(results, weapon)
		}
	}

	return results, nil
}
func findArmorWithSkill(skillID string, rarityFilter int) ([]map[string]interface{}, error) {
	data, err := os.ReadFile("./data/mhwilds/Armor.json")
	if err != nil {
		return nil, err
	}

	var allArmor []map[string]interface{}
	if err := json.Unmarshal(data, &allArmor); err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for _, armorSet := range allArmor {
		// Optional rarity match
		if rarityFilter > 0 {
			if r, ok := armorSet["rarity"].(float64); !ok || int(r) != rarityFilter {
				continue
			}
		}

		pieces, ok := armorSet["pieces"].([]interface{})
		if !ok {
			continue
		}

		for _, p := range pieces {
			piece, ok := p.(map[string]interface{})
			if !ok {
				continue
			}

			skills, ok := piece["skills"].(map[string]interface{})
			if !ok {
				continue
			}

			if _, hasSkill := skills[skillID]; hasSkill {
				results = append(results, piece)
			}
		}
	}

	return results, nil
}

func convertWeapons(raw []map[string]interface{}) []findType.Weapon {
	var result []findType.Weapon
	for _, item := range raw {
		names := map[string]string{}
		if nameMap, ok := item["names"].(map[string]interface{}); ok {
			for lang, val := range nameMap {
				if str, ok := val.(string); ok {
					names[lang] = str
				}
			}
		}

		rarity := intFromAny(item["rarity"])

		skills := map[string]int{}
		if rawSkills, ok := item["skills"].(map[string]interface{}); ok {
			for id, val := range rawSkills {
				skills[id] = intFromAny(val)
			}
		}

		result = append(result, findType.Weapon{
			Names:  names,
			Rarity: rarity,
			Skills: skills,
		})
	}
	return result
}

func convertArmors(raw []map[string]interface{}) []findType.Armor {
	var result []findType.Armor
	for _, item := range raw {
		names := map[string]string{}
		if nameMap, ok := item["names"].(map[string]interface{}); ok {
			for lang, val := range nameMap {
				if str, ok := val.(string); ok {
					names[lang] = str
				}
			}
		}

		rarity := intFromAny(item["rarity"])

		// Optional: Parse pieces if needed in the future
		result = append(result, findType.Armor{
			Names:  names,
			Rarity: rarity,
			Pieces: nil, // Stubbed for now
		})
	}
	return result
}

func intFromAny(val interface{}) int {
	switch v := val.(type) {
	case float64:
		return int(v)
	case int:
		return v
	default:
		return 0
	}
}
