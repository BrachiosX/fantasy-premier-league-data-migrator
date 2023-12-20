package fplapiconnector

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/zest97/fpl-player-data-migration/models"
)

func GetFPLBoostrapData() models.BootStrapData {
	apiUrl := "https://fantasy.premierleague.com/api/bootstrap-static/"
	req, err := http.NewRequest("GET", apiUrl, nil)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// map response body to struct type Bootstrap
	var body models.BootStrapData
	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func GetFixtures() []models.Fixture {
	apiUrl := "https://fantasy.premierleague.com/api/fixtures/"
	req, err := http.NewRequest("GET", apiUrl, nil)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// map response body to struct type Fixture
	var body []models.Fixture
	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func GetMatchHistory(playerId int) models.PlayerDetail {
	apiUrl := "https://fantasy.premierleague.com/api/element-summary/" + strconv.Itoa(playerId) + "/"
	req, err := http.NewRequest("GET", apiUrl, nil)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// map response body to struct type PlayerDetail
	var body models.PlayerDetail
	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
