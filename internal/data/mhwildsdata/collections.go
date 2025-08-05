package mhwildsdata

import (
	"fmt"

	"github.com/PittsGitHub/poogieBot/internal/types/mhwildtypes"
)

func GetArmorGroupedByRarity(rarities []int) (map[int][]mhwildtypes.Armor, error) {
	result := make(map[int][]mhwildtypes.Armor)

	allArmors, err := LoadAllArmor()
	if err != nil {
		return result, fmt.Errorf("error loading armor data: %w", err)
	}

	// Prepare rarity filter set
	raritySet := make(map[int]bool)
	for _, r := range rarities {
		raritySet[r] = true
		result[r] = []mhwildtypes.Armor{}
	}

	// Filter and group
	for _, armor := range allArmors {
		if raritySet[armor.Rarity] {
			result[armor.Rarity] = append(result[armor.Rarity], armor)
		}
	}

	return result, nil
}

func GetWeaponsGroupedByRarity(rarities []int) (map[int][]mhwildtypes.Armor, error) {
	result := make(map[int][]mhwildtypes.Armor)

	allArmors, err := LoadAllArmor()
	if err != nil {
		return result, fmt.Errorf("error loading armor data: %w", err)
	}

	// Prepare rarity filter set
	raritySet := make(map[int]bool)
	for _, r := range rarities {
		raritySet[r] = true
		result[r] = []mhwildtypes.Armor{}
	}

	// Filter and group
	for _, armor := range allArmors {
		if raritySet[armor.Rarity] {
			result[armor.Rarity] = append(result[armor.Rarity], armor)
		}
	}

	return result, nil
}
