package checker

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type RelaysInfo struct {
	Relays []Relays `json:"relays"`
}

type Relays struct {
	Fingerprint string   `json:"fingerprint"`
	OrAddresses []string `json:"or_addresses"`
	Country     string   `json:"country"`
}

func GetRelays(input string) ([]Relays, error) {
	var data []byte
	if strings.HasPrefix(input, "https://") {
		client := &http.Client{
			Timeout: time.Second * 5,
		}

		resp, err := client.Get(input)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		data, err = os.ReadFile(input)
		if err != nil {
			return nil, err
		}
	}

	var info RelaysInfo
	json.Unmarshal(data, &info)
	return info.Relays, nil
}
