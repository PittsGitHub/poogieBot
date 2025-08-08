package mhwildservices

import (
	"fmt"
	"log"
	"sort"
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

func BuildDecorationSkillSummaryMessage(decorations []mhwildtypes.DecorationSkillMatch, skillName string) string {
	var sb strings.Builder

	for _, decoration := range decorations {
		sb.WriteString(fmt.Sprintf("💎 %s:\n Deco Level: %d\n Used in: %s\n %s x%d\n\n",
			decoration.DecorationName,
			decoration.DecorationLevel,
			decoration.AllowedOn, // Properly display the AllowedOn field
			skillName,
			decoration.SkillLevel,
		))
	}

	if len(decorations) == 0 {
		sb.WriteString("No matching decorations found.\n")
	}

	return sb.String()
}
func BuildWeaponSkillSummaryMessage(filteredWeapons map[int][]mhwildtypes.Weapon) string {

	skillNameMap, err := mhwildsdata.GetSkillNameMap()
	if err != nil {
		log.Printf("Warning: failed to load skill name map: %v", err)
	}

	var sb strings.Builder

	for rarity, weaponList := range filteredWeapons {
		sb.WriteString(fmt.Sprintf("**Rarity %d**:\n", rarity))

		for _, w := range weaponList {
			name := w.Names["en"]
			if name == "" {
				name = "[Unnamed Weapon]"
			}

			// Weapon name
			sb.WriteString(fmt.Sprintf(" 🗡️ %s\n", name))
			// Stats on their own line
			sb.WriteString(fmt.Sprintf("    ATK %d, Affinity %d%%\n", w.AttackCalculated, w.Affinity))

			// Skills block
			if len(w.Skills) > 0 {
				type kv struct {
					name string
					lvl  int
				}
				items := make([]kv, 0, len(w.Skills))
				for id, lvl := range w.Skills {
					sn := skillNameMap[id]
					if sn == "" {
						sn = fmt.Sprintf("[Unknown Skill %s]", id)
					}
					items = append(items, kv{name: sn, lvl: lvl})
				}
				sort.Slice(items, func(i, j int) bool { return items[i].name < items[j].name })

				sb.WriteString("    Skills:\n")
				for i, it := range items {
					trailing := ""
					if i < len(items)-1 {
						trailing = ","
					}
					sb.WriteString(fmt.Sprintf("    %s x%d%s\n", it.name, it.lvl, trailing))
				}
			} else {
				sb.WriteString("    Skills: [none]\n")
			}

			sb.WriteString("\n")
		}
	}

	return sb.String()
}
