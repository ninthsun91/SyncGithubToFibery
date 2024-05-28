package request

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ninthsun91/SyncGithubToFibery/utils"
)

type SyncState string

const (
	IN_PROGRESS SyncState = "IN_PROGRESS"
	COMPLETED   SyncState = "COMPLETED"
)

type StatusResponse struct {
	State   SyncState `json:"state"`
	Message string    `json:"message"`
}

func Status() (StatusResponse, error) {
	token := os.Getenv("FIBERY_API_TOKEN")

	url := createUrl() + "/status"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("error creating request: %v\n", err)
		return StatusResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error calling %s: %v\n", url, err)
		return StatusResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("status request failed with status code %d", resp.StatusCode)
		fmt.Println(err)
		return StatusResponse{}, err
	}

	body, err := utils.DecodeResponseBody[StatusResponse](resp)
	if err != nil {
		fmt.Printf("error reading response body: %v\n", err)
		return StatusResponse{}, err
	}

	return body, nil
}
