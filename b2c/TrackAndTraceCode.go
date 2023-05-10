package TrackAndTraceCode

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"
)



func GetTrackAndTrace(website string) string {

	req, err := http.NewRequest("GET", website, nil)

	if err != nil {
		log.Fatal()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	str := string(body)

	titleStartIndex := strings.Index(str, `<strong id="text-tracking-number">Your tracking number</strong>`) + 63
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		panic(err)
	}

	trackAndTraceValue := ""

	for i := titleStartIndex; i < len(str); i++ {

		if IsValidChar(string(str[i])) {
			trackAndTraceValue += string(str[i])
		}
		
		if string(str[i]) == "<" {
			break
		}

	}

	return trackAndTraceValue
}


func IsValidChar(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r)   {
			return true
		}

		if unicode.IsNumber(r)   {
			return true
		}
	}

	return false
}