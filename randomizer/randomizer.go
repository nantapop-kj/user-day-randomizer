package randomizer

import (
	"errors"
	"math"
	"math/rand"
)

// RandomFromList returns a random element from a slice
func RandomFromList(items []string) string {
	return items[rand.Intn(len(items))]
}

// calculateMaxUsersPerDay determines how many users can share the same day
func calculateMaxUsersPerDay(userCount int, dayCount int) int {
	return int(math.Ceil(float64(userCount) / float64(dayCount)))
}

// GenerateDayUserMap returns map[day][]users
func GenerateDayUserMap(users []string, days []string) (map[string][]string, error) {
	result := make(map[string][]string)

	maxUsersPerDay := calculateMaxUsersPerDay(len(users), len(days))
	maxCapacity := len(days) * maxUsersPerDay
	if len(users) > maxCapacity {
		return nil, errors.New("number of users exceeds day capacity")
	}

	dayUsageCount := make(map[string]int)

	for _, user := range users {
		for {
			randomDay := RandomFromList(days)

			if dayUsageCount[randomDay] < maxUsersPerDay {
				result[randomDay] = append(result[randomDay], user)
				dayUsageCount[randomDay]++
				break
			}
		}
	}

	return result, nil
}