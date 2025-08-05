package mhwildservices

import (
	"fmt"
	"log"
	"strings"

	"github.com/PittsGitHub/poogieBot/internal/data/mhwildsdata"
	"github.com/PittsGitHub/poogieBot/internal/types/mhwildtypes"
)

func BuildArmorSkillSummaryMessage(filteredArmor map[int][]mhwildtypes.Armor) string {
	skillNameMap, err := mhwildsdata.GetSkillNameMap()
	if err != nil {
		log.Printf("Warning: failed to load skill name map: %v", err)
	}

	var sb strings.Builder

	for rarity, armorList := range filteredArmor {
		sb.WriteString(fmt.Sprintf("**Rarity %d**:\n", rarity))

		for _, armor := range armorList {
			armorName := armor.Names["en"]
			if armorName == "" {
				armorName = "[Unnamed Armor]"
			}

			var armorSkillNames []string

			if armor.SetBonus != nil {
				id := fmt.Sprintf("%d", armor.SetBonus.SkillID)
				if name := skillNameMap[id]; name != "" {
					armorSkillNames = append(armorSkillNames, name)
				}
			}

			if armor.GroupBonus != nil {
				id := fmt.Sprintf("%d", armor.GroupBonus.SkillID)
				if name := skillNameMap[id]; name != "" {
					armorSkillNames = append(armorSkillNames, name)
				}
			}

			if len(armorSkillNames) > 0 {
				sb.WriteString(fmt.Sprintf(" 🛡️ %s:\n", armorName))

				sb.WriteString(fmt.Sprintf("  %s\n", strings.Join(armorSkillNames, ", ")))

			} else {
				sb.WriteString(fmt.Sprintf("• %s\n", armorName))
			}

			if len(armor.Pieces) > 0 {
				for _, piece := range armor.Pieces {
					pieceName := strings.Title(piece.Kind)

					var skills []string
					for id, level := range piece.Skills {
						name := skillNameMap[id]
						if name == "" {
							name = fmt.Sprintf("[Unknown Skill %s]", id)
						}
						skills = append(skills, fmt.Sprintf("%s x%d", name, level))
					}

					skillLine := strings.Join(skills, ", ")
					sb.WriteString(fmt.Sprintf("    ◦ *%s*: %s\n", pieceName, skillLine))

				}
			} else {
				sb.WriteString("    ◦ [Matched via set or group bonus]\n")
			}

			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func BuildTalismanSkillSummaryMessage(talismans []mhwildtypes.TalismanSkillMatch, skillName string) string {

	var sb strings.Builder

	for _, talisman := range talismans {
		sb.WriteString(fmt.Sprintf("📿 %s:\n (Rarity %d)\n %s x%d \n ",
			talisman.TalismanName,
			talisman.Rarity,

			skillName,
			talisman.SkillLevel,
		))
	}

	if len(talismans) == 0 {
		sb.WriteString("No matching talismans found.\n")
	}

	return sb.String()
}
