package drive

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DriveFreeGeoIp struct {
	Ip string
}

func (d *DriveFreeGeoIp) GetIp() string {
	return d.Ip
}

func (d *DriveFreeGeoIp) GetCountryCode() (string, error) {
	url := fmt.Sprintf("http://freegeoip.net/json/%s", d.GetIp())
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var ipMeta struct {
		CountryCode string `json:"country_code"`
	}
	if err := dec.Decode(&ipMeta); err != nil {
		return "", err
	}

	return ipMeta.CountryCode, nil
}