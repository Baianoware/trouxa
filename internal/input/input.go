package input

import "strings"

func GetData(input string) ([]byte, error) {
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		data, err := getPackagesFromURL(input)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return getPackagesFromFile(input)
}
