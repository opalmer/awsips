package awsips

import (
	"encoding/json"
	"errors"
	"net/http"
)

// AWS_RANGES_URL is url used to pull the ip range information.
const AWS_RANGES_URL = `https://ip-ranges.amazonaws.com/ip-ranges.json`

type AWSRanges struct {
	SyncToken string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes []AWSPrefix `json:"prefixes"`
}

type AWSPrefix struct {
	Prefix string `json:"ip_prefix"`
	Region string `json:"region"`
	Service string `json:"service"`
}

func download() (*AWSRanges, error) {
	request, err := http.Get(AWS_RANGES_URL)
	defer request.Body.Close()

	if err != nil {
		return nil, err
	}

	if request.StatusCode != http.StatusOK {
		return nil, errors.New("Expected 200 OK")
	}
	var ranges AWSRanges
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&ranges)
	return &ranges, nil
}

// Get will download and return a AWSRanges struct representing the
// ip ranges.
func Get() (*AWSRanges, error) {
	return download()
}
