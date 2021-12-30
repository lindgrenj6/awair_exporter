package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	APIKEY     = os.Getenv("AWAIR_TOKEN")
	httpClient = http.Client{Timeout: 10 * time.Second}
)

func GetDeviceList() (*DeviceList, error) {
	bytes, err := hitAwairAPI("/v1/users/self/devices")
	if err != nil {
		return nil, err
	}

	var devices DeviceList
	err = json.Unmarshal(bytes, &devices)
	if err != nil {
		return nil, err
	}
	return &devices, nil
}

func GetAirDataForDevice(device string, id int64) (*AirDataRaw, error) {
	bytes, err := hitAwairAPI(fmt.Sprintf("/v1/users/self/devices/%s/%d/air-data/latest?fahrenheit=true", device, id))
	if err != nil {
		return nil, err
	}

	var airData AirDataRaw
	err = json.Unmarshal(bytes, &airData)
	if err != nil {
		return nil, err
	}
	return &airData, nil
}

func hitAwairAPI(path string) ([]byte, error) {
	url, err := url.Parse(fmt.Sprintf("https://developer-apis.awair.is%s", path))
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: map[string][]string{"authorization": {fmt.Sprintf("Bearer %v", APIKEY)}},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error in awair api, status code: %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
