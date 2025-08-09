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

	for _, rarity := range rarities {
		armors := grouped[rarity]
		for _, armor := range armors {
			name := armor.Set.Names["en"]
			if name == "" {
				name = "[Unnamed Armor]"
			}

			if armor.SetLevelMatch {
				builder.WriteString(fmt.Sprintf("Rarity %d — %s: %s (set bonus)\n", rarity, name, skillName))
			}

			for _, pieceMatch := range armor.MatchingPieces {
				pieceName := pieceMatch.Piece.Names["en"]
				if pieceName == "" {
					pieceName = strings.Title(pieceMatch.Piece.Kind)
				}
				builder.WriteString(fmt.Sprintf("Rarity %d — %s (%s): %s x%d\n", rarity, name, pieceName, skillName, pieceMatch.SkillLevel))
			}
		}
	}

	if builder.Len() == 0 {
		return "No matching armor pieces found with that skill and rarity."
	}

	return builder.String()
}
