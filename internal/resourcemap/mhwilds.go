package resourcemap

var WeaponFileMap = map[string]string{
	"greatsword":     "mhwilds/weapons/GreatSword.json",
	"longsword":      "mhwilds/weapons/LongSword.json",
	"swordandshield": "mhwilds/weapons/SwordAndShield.json",
	"dualblades":     "mhwilds/weapons/DualBlades.json",
	"hammer":         "mhwilds/weapons/Hammer.json",
	"huntinghorn":    "mhwilds/weapons/HuntingHorn.json",
	"lance":          "mhwilds/weapons/Lance.json",
	"gunlance":       "mhwilds/weapons/Gunlance.json",
	"switchaxe":      "mhwilds/weapons/SwitchAxe.json",
	"chargeblade":    "mhwilds/weapons/ChargeBlade.json",
	"insectglaive":   "mhwilds/weapons/InsectGlaive.json",
	"lightbowgun":    "mhwilds/weapons/LightBowgun.json",
	"heavybowgun":    "mhwilds/weapons/HeavyBowgun.json",
	"bow":            "mhwilds/weapons/Bow.json",
}

var HuntingHornVariants = []string{
	"mhwilds/weapons/HuntingHornEchoBubbles.json",
	"mhwilds/weapons/HuntingHornEchoWaves.json",
	"mhwilds/weapons/HuntingHornMelodies.json",
	"mhwilds/weapons/HuntingHornSongs.json",
}

var CoreItemFiles = map[string]string{
	"armor":      "mhwilds/armour.json",
	"talisman":   "mhwilds/amulet.json",
	"decoration": "mhwilds/accessory.json",
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
