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

			// 🔗 Add wiki link per armor entry
			url := buildArmorUrlFromName(armorName)
			sb.WriteString(fmt.Sprintf("    🔗 %s\n\n", url))
		}
	}

	return sb.String()
}

func BuildTalismanSkillSummaryMessage(talismans []mhwildtypes.TalismanSkillMatch, skillName string) string {
	var sb strings.Builder

	for _, talisman := range talismans {
		sb.WriteString(fmt.Sprintf("📿 %s:\n (Rarity %d)\n %s x%d\n",
			talisman.TalismanName,
			talisman.Rarity,
			skillName,
			talisman.SkillLevel,
		))

		url := buildUrlFromName(talisman.TalismanName)
		sb.WriteString(fmt.Sprintf("🔗 %s\n\n", url))
	}

	if len(talismans) == 0 {
		sb.WriteString("No matching talismans found.\n")
	}

	return sb.String()
}

func BuildDecorationSkillSummaryMessage(decorations []mhwildtypes.DecorationSkillMatch, skillName string) string {
	var sb strings.Builder

	for _, decoration := range decorations {
		sb.WriteString(fmt.Sprintf("💎 %s:\n Deco Level: %d\n Used in: %s\n %s x%d\n",
			decoration.DecorationName,
			decoration.DecorationLevel,
			decoration.AllowedOn,
			skillName,
			decoration.SkillLevel,
		))

		url := buildUrlFromName(decoration.DecorationName)
		sb.WriteString(fmt.Sprintf("🔗 %s\n\n", url))
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

			// Add wiki link for this weapon
			url := buildUrlFromName(name)
			sb.WriteString(fmt.Sprintf("    🔗 %s\n\n", url))
		}
	}

	return sb.String()
}

func buildUrlFromName(name string) string {
	// Clean name: remove [ and ], then replace spaces with +
	cleanName := strings.ReplaceAll(name, "[", "")
	cleanName = strings.ReplaceAll(cleanName, "]", "")
	cleanName = strings.ReplaceAll(cleanName, "/", "-")
	wikiName := strings.ReplaceAll(cleanName, " ", "+")

	url := fmt.Sprintf("https://monsterhunterwilds.wiki.fextralife.com/%s", wikiName)
	return url
}

func buildArmorUrlFromName(name string) string {
	// Step 1: Clean name
	cleanName := strings.ReplaceAll(name, "[", "")
	cleanName = strings.ReplaceAll(cleanName, "]", "")
	cleanName = strings.ReplaceAll(cleanName, "/", "-")

	// Step 2: Replace Alpha/Beta symbols with full names
	cleanName = strings.ReplaceAll(cleanName, "α", "Alpha")
	cleanName = strings.ReplaceAll(cleanName, "β", "Beta")

	// Step 3: Append " Set" if not already there
	if !strings.HasSuffix(cleanName, "Set") {
		cleanName += " Set"
	}

	// Step 4: Format as wiki URL
	wikiName := strings.ReplaceAll(cleanName, " ", "+")
	url := fmt.Sprintf("https://monsterhunterwilds.wiki.fextralife.com/%s", wikiName)

	return url
}
