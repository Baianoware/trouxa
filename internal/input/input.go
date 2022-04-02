package input

// GetData returns the data from the packages' list selected. input is a path to a local file or a URL to a remote one.
func GetData(input string) ([]byte, error) {
	if isURL(input) {
		return getPackagesFromURL(input)
	} else {
		return getPackagesFromFile(input)
	}
}
