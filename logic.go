package logic

import "fmt"

// Add words first name and last name
func AddWords(firstName string, lastName string) string {
	return fmt.Sprintf("My name is %s %s", firstName, lastName)
}
