package mhwildsdata

var WeaponFileMap = map[string]string{
	"greatsword":     "data/mhwilds/weapons/GreatSword.json",
	"longsword":      "data/mhwilds/weapons/LongSword.json",
	"swordandshield": "data/mhwilds/weapons/SwordShield.json",
	"dualblades":     "data/mhwilds/weapons/DualBlades.json",
	"hammer":         "data/mhwilds/weapons/Hammer.json",
	"huntinghorn":    "data/mhwilds/weapons/HuntingHorn.json",
	"lance":          "data/mhwilds/weapons/Lance.json",
	"gunlance":       "data/mhwilds/weapons/Gunlance.json",
	"switchaxe":      "data/mhwilds/weapons/SwitchAxe.json",
	"chargeblade":    "data/mhwilds/weapons/ChargeBlade.json",
	"insectglaive":   "data/mhwilds/weapons/InsectGlaive.json",
	"lightbowgun":    "data/mhwilds/weapons/LightBowgun.json",
	"heavybowgun":    "data/mhwilds/weapons/HeavyBowgun.json",
	"bow":            "data/mhwilds/weapons/Bow.json",
}

var HuntingHornVariants = []string{
	"data/mhwilds/weapons/HuntingHornEchoBubbles.json",
	"data/mhwilds/weapons/HuntingHornEchoWaves.json",
	"data/mhwilds/weapons/HuntingHornMelodies.json",
	"data/mhwilds/weapons/HuntingHornSongs.json",
}

var CoreItemFiles = map[string]string{
	"armor":      "data/mhwilds/Armor.json",
	"talisman":   "data/mhwilds/Amulet.json",
	"decoration": "data/mhwilds/Accessory.json",
	"skill":      "data/mhwilds/Skill.json",
}

// GetAllItemFiles returns a full list of all .json file paths relevant to item lookups
func GetAllItemFiles() []string {
	files := []string{}

	// Core files
	for _, path := range CoreItemFiles {
		files = append(files, path)
	}

	// Weapon files
	for _, path := range WeaponFileMap {
		files = append(files, path)
	}

	// Hunting Horn variants
	files = append(files, HuntingHornVariants...)

	return files
}
