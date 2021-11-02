package episode

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const dataPath = "data/stvoy-episodes.json"

type EpisodeData struct {
	Episodes []Details
}

type Details struct {
	Name    string `json:"name"`
	Season  int    `json:"season"`
	Episode int    `json:"episode"`
	Number  int    `json:"number"`
}

// Info returns the title of an episode given an episode number.
func Info(n int) (Details, error) {
	var details Details

	if err := validateNumber(n); err != nil {
		return details, fmt.Errorf("validation: %w", err)
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return details, err
	}

	var epData EpisodeData
	if err := json.Unmarshal(data, &epData); err != nil {
		return details, err
	}
	return epData.Episodes[n-1], nil
}

// Random returns a random episode title.
func Random() (Details, error) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(173)

	details, err := Info(n)
	if err != nil {
		return details, err
	}

	return details, nil
}

func validateNumber(n int) error {
	if n < 1 {
		return fmt.Errorf("number must be greater than 0; %q was provided", n)
	}

	if n > 172 {
		return fmt.Errorf("number must be less than than `73; %q was provided", n)
	}

	return nil
}

func (d Details) String() string {
	return fmt.Sprintf("%q\nSeason: %d, Episode: %d, Number: %d", d.Name, d.Season, d.Episode, d.Number)
}
