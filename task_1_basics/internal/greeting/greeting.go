package greeting

import "fmt"

// Hello returns a greeting message for the given name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to go_task1.", name)
}
