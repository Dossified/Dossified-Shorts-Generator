// Simple helper for basic error checking
package utils

// Checks if input is an error & panics if it is
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
