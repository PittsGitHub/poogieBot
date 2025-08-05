package mhwildsdata

import (
	"fmt"

	"github.com/PittsGitHub/poogieBot/internal/types/mhwildtypes"
)

func FilterArmorBySkillID(
	rarityGrouped map[int][]mhwildtypes.Armor,
	skillID string,
) map[int][]mhwildtypes.Armor {
	result := make(map[int][]mhwildtypes.Armor)

	for rarity, armorList := range rarityGrouped {
		for _, armor := range armorList {
			// Check if skill is in SetBonus or GroupBonus
			setMatch := armor.SetBonus != nil && fmt.Sprintf("%d", armor.SetBonus.SkillID) == skillID
			groupMatch := armor.GroupBonus != nil && fmt.Sprintf("%d", armor.GroupBonus.SkillID) == skillID

			if setMatch || groupMatch {
				// Keep full armor (including all pieces)
				result[rarity] = append(result[rarity], armor)
				continue
			}

			// Skill not found in bonuses — check pieces
			var matchingPieces []mhwildtypes.ArmorPiece
			for _, piece := range armor.Pieces {
				if level, ok := piece.Skills[skillID]; ok && level > 0 {
					matchingPieces = append(matchingPieces, piece)
				}
			}

			if len(matchingPieces) > 0 {
				// Create a copy of the armor with only matching pieces
				filteredArmor := armor
				filteredArmor.Pieces = matchingPieces
				result[rarity] = append(result[rarity], filteredArmor)
			}
		}
	}

	return result
}

func FilterTalismanBySkill(skillID string) ([]mhwildtypes.TalismanSkillMatch, error) {
	allTalismans, err := LoadTalismans()
	if err != nil {
		return nil, fmt.Errorf("error loading talisman data: %w", err)
	}

	var results []mhwildtypes.TalismanSkillMatch

	for _, talisman := range allTalismans {
		highest := mhwildtypes.TalismanRank{}
		highestLevel := 0

		for _, rank := range talisman.Ranks {
			if level, ok := rank.Skills[skillID]; ok && level > highestLevel {
				highest = rank
				highestLevel = level
			}
		}

		if highestLevel > 0 {
			results = append(results, mhwildtypes.TalismanSkillMatch{
				TalismanName: highest.Names["en"],
				SkillLevel:   highestLevel,
				Rarity:       highest.Rarity,
			})
		}
	}

	return results, nil
}
