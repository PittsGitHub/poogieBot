package mhwildsdata

import (
	"errors"
	"fmt"
	"sort"

	"github.com/PittsGitHub/poogieBot/internal/data"
	"github.com/PittsGitHub/poogieBot/internal/types/mhwildtypes"
)

var WeaponLoaders = map[string]func() ([]mhwildtypes.Weapon, error){
	"greatsword":     LoadGreatswords,
	"longsword":      LoadLongswords,
	"swordandshield": LoadSwordAndShields,
	"dualblades":     LoadDualBlades,
	"hammer":         LoadHammers,
	"huntinghorn":    LoadHuntingHorns,
	"lance":          LoadLances,
	"gunlance":       LoadGunlances,
	"switchaxe":      LoadSwitchAxes,
	"chargeblade":    LoadChargeBlades,
	"insectglaive":   LoadInsectGlaives,
	"lightbowgun":    LoadLightBowguns,
	"heavybowgun":    LoadHeavyBowguns,
	"bow":            LoadBows,
}

func LoadAllWeapons() ([]mhwildtypes.Weapon, error) {
	// stable iteration over map keys
	keys := make([]string, 0, len(WeaponLoaders))
	for k := range WeaponLoaders {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var all []mhwildtypes.Weapon
	var errs []error

	for _, k := range keys {
		loader := WeaponLoaders[k]
		ws, err := loader()
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", k, err))
			continue
		}
		all = append(all, ws...)
	}

	if len(errs) > 0 {
		return all, errors.Join(errs...)
	}
	return all, nil
}

func LoadAllArmor() ([]mhwildtypes.Armor, error) {
	return data.LoadJSON[mhwildtypes.Armor](CoreItemFiles["armor"])
}

func LoadSkills() ([]mhwildtypes.Skill, error) {
	return data.LoadJSON[mhwildtypes.Skill](CoreItemFiles["skill"])
}

func LoadGreatswords() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["greatsword"])
}

func LoadLongswords() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["longsword"])
}

func LoadSwordAndShields() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["swordandshield"])
}

func LoadDualBlades() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["dualblades"])
}

func LoadHammers() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["hammer"])
}

func LoadHuntingHorns() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["huntinghorn"])
}

func LoadLances() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["lance"])
}

func LoadGunlances() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["gunlance"])
}

func LoadSwitchAxes() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["switchaxe"])
}

func LoadChargeBlades() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["chargeblade"])
}

func LoadInsectGlaives() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["insectglaive"])
}

func LoadLightBowguns() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["lightbowgun"])
}

func LoadHeavyBowguns() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["heavybowgun"])
}

func LoadBows() ([]mhwildtypes.Weapon, error) {
	return data.LoadJSON[mhwildtypes.Weapon](WeaponFileMap["bow"])
}

func LoadTalismans() ([]mhwildtypes.Talisman, error) {
	return data.LoadJSON[mhwildtypes.Talisman](CoreItemFiles["talisman"])
}
