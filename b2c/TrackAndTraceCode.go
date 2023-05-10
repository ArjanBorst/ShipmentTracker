package TrackAndTraceCode

import (
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

const (
	trackingStr       string = "Your tracking number: "
	offsetTrackingStr int    = len(trackingStr)
	notFoundMsg              = "Not found"
)

/*
Scrape the track and trace code from this website https://www.trackyourparcel.eu/
Return empty string if track and trace could not be found and err if website could not be read
*/
func GetTrackAndTrace(website string) (string, error) {
	websiteBody, err := readWebsiteBody(website)
	if err != nil {
		return "", err
	}

	startIndex := strings.Index(websiteBody, trackingStr)
	if startIndex == -1 {
		return "", nil
	}

	var trackAndTraceCode strings.Builder
	for _, r := range websiteBody[startIndex+offsetTrackingStr:] {
		if isValidCharForTrackAndTrace(string(r)) {
			trackAndTraceCode.WriteByte(byte(r))
		} else {
			break
		}
	}

	return trackAndTraceCode.String(), nil
}

func isValidCharForTrackAndTrace(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return true
		}
	}

	return false
}

func readWebsiteBody(website string) (string, error) {
	req, err := http.NewRequest("GET", website, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
