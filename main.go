package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/nantapop-kj/user-day-randomizer/randomizer"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// printResultByDay displays users grouped by day in fixed order
func printResultByDay(days []string, result map[string][]string) {
	fmt.Println("User Schedule")
	fmt.Println("-------------")

	for _, day := range days {
		usersInDay, exists := result[day]
		if !exists || len(usersInDay) == 0 {
			fmt.Printf("%s: (no users)\n", day)
			continue
		}

		// copy slice before sorting (avoid mutating map value)
		sortedUsers := make([]string, len(usersInDay))
		copy(sortedUsers, usersInDay)

		sort.Strings(sortedUsers)

		fmt.Printf("%s: ", day)
		for i, user := range sortedUsers {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(user)
		}
		fmt.Println()
	}
}

func main() {
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri"}
	users := []string{
		"Alice", "Bob", "Charlie", "David", "Emma", "Frank",
		"Grace", "Henry", "Ivy", "Jack", "Kevin", "Lily",
		"Mike", "Nina", "Oscar", "Paul",
	}

	result, err := randomizer.GenerateDayUserMap(users, days)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printResultByDay(days, result)
}
