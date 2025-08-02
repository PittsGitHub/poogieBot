package mhwildservices

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PittsGitHub/poogieBot/internal/types/mhwildtypes"
)

func GetSkillIDFromName(skillName string) (string, error) {
	formatted := FormatTitleCase(skillName)
	mhwildsWackyFormat := CollapseToPartbreakerStyle(formatted)

	file, err := os.Open("./data/mhwilds/Skill.json")
	if err != nil {
		return "", fmt.Errorf("failed to open Skill.json: %w", err)
	}
	defer file.Close()

	var skills []mhwildtypes.Skill
	if err := json.NewDecoder(file).Decode(&skills); err != nil {
		return "", fmt.Errorf("failed to parse Skill.json: %w", err)
	}

	for _, skill := range skills {
		if name, ok := skill.Names["en"]; ok {
			if name == formatted || CollapseToPartbreakerStyle(name) == mhwildsWackyFormat {
				return fmt.Sprintf("%d", skill.GameID), nil
			}
		}
	}

	return "", fmt.Errorf("skill '%s' not found", skillName)
}
