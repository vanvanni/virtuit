package support

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func SendFirecrackerRequest(client *http.Client, method, path string, body interface{}) error {
	url := "http://unix" + path
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return fmt.Errorf("json encode failed: %w", err)
		}
	}
	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		dump, _ := httputil.DumpResponse(resp, true)
		return fmt.Errorf("Firecracker error: %s", dump)
	}

	return nil
}
