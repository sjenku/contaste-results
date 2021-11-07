package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

func (manager *ContasteManager) GetCompetitionInfo(url string) (Competition, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do the get request,err: %v", err.Error())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the body,err: %v", err.Error())
	}
	defer res.Body.Close()

	comperoBody, err := manager.getComperoVarBodyFrom(body)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	compero, err := manager.parseComperoBody(comperoBody)
	if err != nil {
		return nil, fmt.Errorf("erorr: %v", err.Error())
	}
	return compero, nil
}

func (manager *ContasteManager) getComperoVarBodyFrom(body []byte) ([]byte, error) {
	bodyStr := string(body)

	comperoAndNext := strings.Split(bodyStr, "var compero =")
	if len(comperoAndNext) < 2 {
		return nil, errors.New("not info for var compero")
	}
	afterVarCompero := comperoAndNext[1]

	splitedAfterVarCompero := strings.Split(afterVarCompero, "; var dcrd =")
	if len(splitedAfterVarCompero) < 1 {
		return nil, errors.New("not info for var compero")
	}
	comperoOnly := splitedAfterVarCompero[0]

	return []byte(comperoOnly), nil

}

func (manager *ContasteManager) CreateResultsString(compero Competition) string {
	var text string
	for _, obj := range compero {
		if obj.Achivments != nil {
			if obj.StoredContestTitle != "" {
				text = fmt.Sprintf("%v----------------------%v---------------------\n", text, obj.StoredContestTitle)
			} else {
				text = fmt.Sprintf("%v---------- %v-%v %v %v---------------------\n", text, obj.AgeFrom, obj.AgeTill, obj.DancingLevel, obj.Group)
			}
			results := manager.getDancersResultsPlaceCouple(obj)
			sortedPlaces := make([]string, 0, len(results))
			for place := range results {
				sortedPlaces = append(sortedPlaces, place)
			}
			sort.Strings(sortedPlaces)
			for _, place := range sortedPlaces {
				text = fmt.Sprintf("%v%v: %v\n", text, place, results[place])
			}
		}
	}
	return text
}

func (manager *ContasteManager) PrintCompetition(compero Competition) {
	text := manager.CreateResultsString(compero)
	fmt.Printf("results:\n%v", text)
}

func (magaer *ContasteManager) getDancersResultsPlaceCouple(compObj CompetitionObject) map[string]string {

	// create new map holding as a key the place and as value the dancers names
	placesMap := map[string]string{}
	for key, achivment := range compObj.Achivments {
		placesMap[achivment.Award] = compObj.Dancers[key].Title
	}

	return placesMap
}

func (manager *ContasteManager) parseComperoBody(body []byte) (Competition, error) {
	var compero Competition
	err := json.Unmarshal(body, &compero)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshel body to type compero,error: %v", err)
	}
	return compero, nil
}
