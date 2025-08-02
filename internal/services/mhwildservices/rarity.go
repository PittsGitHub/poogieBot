package mhwildservices

import (
	"fmt"
	"strconv"
	"strings"
)

func ResolveRarityValues(input string) ([]string, error) {
	input = strings.ToLower(strings.TrimSpace(input))

	switch input {
	case "low":
		return []string{"0", "1", "2", "3", "4"}, nil
	case "high":
		return []string{"5", "6", "7", "8", "9", "10", "11", "12"}, nil
	default:
		// Handle N+ pattern (e.g. 7+)
		if strings.HasSuffix(input, "+") {
			base := strings.TrimSuffix(input, "+")
			if start, err := strconv.Atoi(base); err == nil && start >= 0 && start <= 12 {
				var result []string
				for i := start; i <= 12; i++ {
					result = append(result, strconv.Itoa(i))
				}
				return result, nil
			}
		}

		// Handle N- pattern (e.g. 4-)
		if strings.HasSuffix(input, "-") {
			base := strings.TrimSuffix(input, "-")
			if end, err := strconv.Atoi(base); err == nil && end >= 0 && end <= 12 {
				var result []string
				for i := end; i >= 0; i-- {
					result = append(result, strconv.Itoa(i))
				}
				return result, nil
			}
		}

		// Handle single numeric value
		if num, err := strconv.Atoi(input); err == nil && num >= 0 && num <= 12 {
			return []string{strconv.Itoa(num)}, nil
		}
	}

	return nil, fmt.Errorf("invalid rarity or rank input: '%s'", input)
}
