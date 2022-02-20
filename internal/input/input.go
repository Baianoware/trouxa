package input

// GetData returns the data from the packages' list selected
func GetData(input string) ([]byte, error) {
	if isURL(input) {
		return getPackagesFromURL(input)
	} else {
		return getPackagesFromFile(input)
	}
}
