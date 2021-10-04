package input

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func getPackagesFromURL(packagesURL string) ([]byte, error) {
	if strings.HasPrefix(packagesURL, "http://") || strings.HasPrefix(packagesURL, "https://") {
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
	return nil, errors.New("this parameter is not a link")
}
