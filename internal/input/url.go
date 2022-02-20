package input

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func isURL(input string) bool {
	return strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://")
}

func getPackagesFromURL(packagesURL string) ([]byte, error) {
	log.Debugln("Trying to download the remote package.txt")
	response, err := http.Get(packagesURL)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("could not get the packages from url")
	}
	log.Debugln("Download successfully")
	if ok, _ := regexp.MatchString("text/plain", response.Header.Get("content-type")); !ok {
		return nil, errors.New("content-type invalid: it must be a text/plain one")
	}
	log.Debugln("File is a text/plain")
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	log.Debugln("Successful download of the remote packages.txt")
	return data, nil
}
