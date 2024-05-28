package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ninthsun91/SyncGithubToFibery/utils"
)

type FiberySyncResponse struct {
	State SyncState `json:"state"`
}

func Sync() error {
	token := os.Getenv("FIBERY_API_TOKEN")

	url := createUrl() + "/sync"
	body := createBody()
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		fmt.Printf("error creating request: %v\n", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error calling %s: %v\n", url, err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("sync request failed with status code %d", resp.StatusCode)
		fmt.Println(err)
		return err
	}

	_, err = utils.DecodeResponseBody[FiberySyncResponse](resp)
	if err != nil {
		fmt.Printf("error reading response body: %v\n", err)
		return err
	}

	return nil
}

func createBody() io.Reader {
	name := os.Getenv("FIBERY_SPACE_NAME")
	id := os.Getenv("FIBERY_SPACE_ID")

	body := fmt.Sprintf(
		`{"app":{"name":"%s","id":"%s"},"options":{}}`,
		name,
		id,
	)
	return bytes.NewBufferString(body)

}
