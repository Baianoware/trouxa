package input

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func isURL(input string) bool {
	return strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://")
}

func getPackagesFromURL(packagesURL string) ([]byte, error) {
	response, err := http.Get(packagesURL)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("could not get the packages from url")
	}
	if ok, _ := regexp.MatchString("text/plain", response.Header.Get("content-type")); !ok {
		return nil, errors.New("content-type invalid: it must be a text/plain one")
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
