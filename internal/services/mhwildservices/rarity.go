package mhwildservices

import (
	"fmt"
	"strconv"
	"strings"
)

func ResolveRarityValues(input string) ([]int, error) {
	input = strings.ToLower(strings.TrimSpace(input))

	// 🆕 Default to "high" if input is empty
	if input == "" {
		input = "high"
	}

	switch input {
	case "low":
		return []int{0, 1, 2, 3, 4}, nil
	case "high":
		return []int{5, 6, 7, 8, 9, 10, 11, 12}, nil
	default:
		if strings.HasSuffix(input, "+") {
			base := strings.TrimSuffix(input, "+")
			if start, err := strconv.Atoi(base); err == nil && start >= 0 && start <= 12 {
				var result []int
				for i := start; i <= 12; i++ {
					result = append(result, i)
				}
				return result, nil
			}
		}

		if strings.HasSuffix(input, "-") {
			base := strings.TrimSuffix(input, "-")
			if end, err := strconv.Atoi(base); err == nil && end >= 0 && end <= 12 {
				var result []int
				for i := end; i >= 0; i-- {
					result = append(result, i)
				}
				return result, nil
			}
		}

		if num, err := strconv.Atoi(input); err == nil && num >= 0 && num <= 12 {
			return []int{num}, nil
		}
	}

	return nil, fmt.Errorf("invalid rarity or rank input: '%s'", input)
}
