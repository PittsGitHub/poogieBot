package mhwildsdata

import (
	"fmt"

	"github.com/PittsGitHub/poogieBot/internal/services/mhwildservices"
)

func GetSkillIDFromName(skillName string) (string, error) {
	formatted := mhwildservices.FormatTitleCase(skillName)
	mhwildsWackyFormat := mhwildservices.CollapseToPartbreakerStyle(formatted)

	skills, err := LoadSkills()
	if err != nil {
		return "", fmt.Errorf("failed to load skills: %w", err)
	}

	for _, skill := range skills {
		if name, ok := skill.Names["en"]; ok {
			if name == formatted || mhwildservices.CollapseToPartbreakerStyle(name) == mhwildsWackyFormat {
				return fmt.Sprintf("%d", skill.GameID), nil
			}
		}
	}

	return "", fmt.Errorf("skill '%s' not found", skillName)
}
