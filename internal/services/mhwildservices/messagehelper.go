package mhwildservices

import (
	"fmt"
	"sort"
	"strings"

	"github.com/PittsGitHub/poogieBot/internal/types/mhwildtypes"
)

func FormatArmorSkillMessage(grouped map[int][]mhwildtypes.ArmorMatchResult, skillName string) string {
	var builder strings.Builder

	var rarities []int
	for rarity := range grouped {
		rarities = append(rarities, rarity)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rarities)))

	fmt.Println("Formatting message for skill:", skillName)

	for _, rarity := range rarities {
		armors := grouped[rarity]
		fmt.Printf("Rarity %d — %d items\n", rarity, len(armors))
		for _, armor := range armors {
			name := armor.Set.Names["en"]
			level, ok := armor.Set.Skills[skillName]
			fmt.Printf("Checking %s, has skill? %v\n", name, ok)

			if ok {
				line := fmt.Sprintf("Rarity %d — %s: %s %d\n", rarity, name, skillName, level)
				builder.WriteString(line)
			}
		}
	}

	if builder.Len() == 0 {
		fmt.Println("No matches found for skill:", skillName)
		return "No matching armor pieces found with that skill and rarity."
	}

	return builder.String()
}
