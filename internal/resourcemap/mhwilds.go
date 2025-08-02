package resourcemap

var WeaponFileMap = map[string]string{
	"Greatsword":       "mhwilds/weapons/GreatSword.json",
	"Longsword":        "mhwilds/weapons/LongSword.json",
	"Sword and Shield": "mhwilds/weapons/SwordAndShield.json",
	"Dual Blades":      "mhwilds/weapons/DualBlades.json",
	"Hammer":           "mhwilds/weapons/Hammer.json",
	"Hunting Horn":     "mhwilds/weapons/HuntingHorn.json", // base file
	"Lance":            "mhwilds/weapons/Lance.json",
	"Gunlance":         "mhwilds/weapons/Gunlance.json",
	"Switch Axe":       "mhwilds/weapons/SwitchAxe.json",
	"Charge Blade":     "mhwilds/weapons/ChargeBlade.json",
	"Insect Glaive":    "mhwilds/weapons/InsectGlaive.json",
	"Light Bowgun":     "mhwilds/weapons/LightBowgun.json",
	"Heavy Bowgun":     "mhwilds/weapons/HeavyBowgun.json",
	"Bow":              "mhwilds/weapons/Bow.json",
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
