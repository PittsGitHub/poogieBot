package mhwildsdata

import (
	"fmt"
	"sync"
)

var (
	skillNameMap     map[string]string
	skillNameMapOnce sync.Once
	skillNameMapErr  error
)

// GetSkillNameMap returns a cached map of skillID (string) → skillName (English).
func GetSkillNameMap() (map[string]string, error) {
	skillNameMapOnce.Do(func() {
		skills, err := LoadSkills()
		if err != nil {
			skillNameMapErr = fmt.Errorf("failed to load skills: %w", err)
			return
		}

		skillNameMap = make(map[string]string)
		for _, s := range skills {
			id := fmt.Sprintf("%d", s.GameID)
			skillNameMap[id] = s.Names["en"]
		}
	})

	return skillNameMap, skillNameMapErr
}
