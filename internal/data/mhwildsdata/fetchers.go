package mhwildsdata

import (
	"fmt"

	"github.com/PittsGitHub/poogieBot/internal/services"
)

func GetSkillIDFromName(skillName string) (string, error) {
	formatted := services.FormatTitleCase(skillName)
	mhwildsWackyFormat := services.CollapseToPartbreakerStyle(formatted)

	skills, err := LoadSkills()
	if err != nil {
		return "", fmt.Errorf("failed to load skills: %w", err)
	}

	for _, skill := range skills {
		if name, ok := skill.Names["en"]; ok {
			if name == formatted || services.CollapseToPartbreakerStyle(name) == mhwildsWackyFormat {
				return fmt.Sprintf("%d", skill.GameID), nil
			}
		}
	}

	return "", fmt.Errorf("❌ skill '%s' not found", skillName)
}
