package enricher

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Enricher struct {
}

type EnrichedParam struct {
	City  string
	State string
}

type jsonResponse struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func NewEnricher() *Enricher {
	return &Enricher{}
}

func (e *Enricher) Enrich(param string) (EnrichedParam, error) {

	// submit request to auth service
	req, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer([]byte(param)))
	if err != nil {
		return EnrichedParam{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return EnrichedParam{}, err
	}
	defer resp.Body.Close()

	// respond with appropriate code
	if resp.StatusCode == http.StatusUnauthorized {
		return EnrichedParam{}, errors.New("failed to enrich parameters: Unauthorized")
	} else if resp.StatusCode != http.StatusAccepted {
		return EnrichedParam{}, errors.New(fmt.Sprintf("failed to enrich parameters: %s", resp.Status))
	}

	var jsonFromSvc jsonResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonFromSvc)
	if err != nil {
		return EnrichedParam{}, err
	}

	return EnrichedParam(jsonFromSvc), nil

}
