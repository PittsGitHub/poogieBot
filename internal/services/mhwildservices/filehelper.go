package mhwildservices

import (
	"errors"

	"github.com/PittsGitHub/poogieBot/internal/resourcemap"
)

// GetItemTypeDataFilePaths returns a list of file paths based on the given item type.
func GetItemTypeDataFilePaths(rawItemType string) ([]string, error) {

	//we don't assume data has been sanitised
	itemType := Normalise(rawItemType)

	//sanitise weapon to weapons for all weapon subtypes
	if itemType == "weapon" {
		itemType = "weapons"
	}

	var dataFilePaths []string

	switch itemType {
	case "armor", "decoration", "talisman":
		if path, ok := resourcemap.CoreItemFiles[itemType]; ok {
			dataFilePaths = append(dataFilePaths, path)
		} else {
			return nil, errors.New("unknown core type: " + itemType)
		}

	case "weapons":
		for _, path := range resourcemap.WeaponFileMap {
			dataFilePaths = append(dataFilePaths, path)
		}
		dataFilePaths = append(dataFilePaths, resourcemap.HuntingHornVariants...)

	default:
		if path, ok := resourcemap.WeaponFileMap[itemType]; ok {
			dataFilePaths = append(dataFilePaths, path)
			if itemType == "huntinghorn" {
				dataFilePaths = append(dataFilePaths, resourcemap.HuntingHornVariants...)
			}
		} else {
			return nil, errors.New("unknown weapon type: " + itemType)
		}
	}

	return dataFilePaths, nil
}
