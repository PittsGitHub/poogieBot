package mhwildsdata

import (
	"github.com/PittsGitHub/poogieBot/internal/data"
	"github.com/PittsGitHub/poogieBot/internal/types/mhwildtypes"
)

func LoadAllArmor() ([]mhwildtypes.Armor, error) {
	return data.LoadJSON[mhwildtypes.Armor](CoreItemFiles["armor"])
}

func LoadSkills() ([]mhwildtypes.Skill, error) {
	return data.LoadJSON[mhwildtypes.Skill](CoreItemFiles["skill"])
}

func LoadGreatswords() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["greatsword"])
}
