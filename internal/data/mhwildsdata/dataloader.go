package mhwildsdata

import (
	"errors"
	"fmt"
	"math"
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

// shared helper: load JSON and apply modifier to calculate AttackCalculated
func loadWeaponsWithModifier(file string, mod float64) ([]mhwildtypes.Weapon, error) {
	ws, err := data.LoadJSON[mhwildtypes.Weapon](file)
	if err != nil {
		return nil, err
	}
	for i := range ws {
		ws[i].AttackCalculated = int(math.Round(float64(ws[i].AttackRaw) * mod))
	}
	return ws, nil
}

// Great Sword
func LoadGreatswords() ([]mhwildtypes.Weapon, error) {
	const greatSwordModifier = 4.8
	return loadWeaponsWithModifier(WeaponFileMap["greatsword"], greatSwordModifier)
}

// Long Sword
func LoadLongswords() ([]mhwildtypes.Weapon, error) {
	const longSwordModifier = 3.3
	return loadWeaponsWithModifier(WeaponFileMap["longsword"], longSwordModifier)
}

// Sword & Shield
func LoadSwordAndShields() ([]mhwildtypes.Weapon, error) {
	const swordAndShieldModifier = 1.4
	return loadWeaponsWithModifier(WeaponFileMap["swordandshield"], swordAndShieldModifier)
}

// Dual Blades
func LoadDualBlades() ([]mhwildtypes.Weapon, error) {
	const dualBladesModifier = 1.4
	return loadWeaponsWithModifier(WeaponFileMap["dualblades"], dualBladesModifier)
}

// Hammer
func LoadHammers() ([]mhwildtypes.Weapon, error) {
	const hammerModifier = 5.2
	return loadWeaponsWithModifier(WeaponFileMap["hammer"], hammerModifier)
}

// Hunting Horn
func LoadHuntingHorns() ([]mhwildtypes.Weapon, error) {
	const huntingHornModifier = 4.2
	return loadWeaponsWithModifier(WeaponFileMap["huntinghorn"], huntingHornModifier)
}

// Lance
func LoadLances() ([]mhwildtypes.Weapon, error) {
	const lanceModifier = 2.3
	return loadWeaponsWithModifier(WeaponFileMap["lance"], lanceModifier)
}

// Gunlance
func LoadGunlances() ([]mhwildtypes.Weapon, error) {
	const gunlanceModifier = 2.3
	return loadWeaponsWithModifier(WeaponFileMap["gunlance"], gunlanceModifier)
}

// Switch Axe
func LoadSwitchAxes() ([]mhwildtypes.Weapon, error) {
	const switchAxeModifier = 3.5
	return loadWeaponsWithModifier(WeaponFileMap["switchaxe"], switchAxeModifier)
}

// Charge Blade
func LoadChargeBlades() ([]mhwildtypes.Weapon, error) {
	const chargeBladeModifier = 3.6
	return loadWeaponsWithModifier(WeaponFileMap["chargeblade"], chargeBladeModifier)
}

// Insect Glaive
func LoadInsectGlaives() ([]mhwildtypes.Weapon, error) {
	const insectGlaiveModifier = 3.1
	return loadWeaponsWithModifier(WeaponFileMap["insectglaive"], insectGlaiveModifier)
}

// Light Bowgun
func LoadLightBowguns() ([]mhwildtypes.Weapon, error) {
	const lightBowgunModifier = 1.3
	return loadWeaponsWithModifier(WeaponFileMap["lightbowgun"], lightBowgunModifier)
}

// Heavy Bowgun
func LoadHeavyBowguns() ([]mhwildtypes.Weapon, error) {
	const heavyBowgunModifier = 1.5
	return loadWeaponsWithModifier(WeaponFileMap["heavybowgun"], heavyBowgunModifier)
}

// Bow
func LoadBows() ([]mhwildtypes.Weapon, error) {
	const bowModifier = 1.2
	return loadWeaponsWithModifier(WeaponFileMap["bow"], bowModifier)
}

func LoadTalismans() ([]mhwildtypes.Talisman, error) {
	return data.LoadJSON[mhwildtypes.Talisman](CoreItemFiles["talisman"])
}

func LoadDecorations() ([]mhwildtypes.Decoration, error) {
	return data.LoadJSON[mhwildtypes.Decoration](CoreItemFiles["decoration"])
}
