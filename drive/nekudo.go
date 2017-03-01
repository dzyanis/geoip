package drive

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DriveNekudo struct {
	Ip string
}

func (d *DriveNekudo) GetIp() string {
	return d.Ip
}

func (d *DriveNekudo) GetCountryCode() (string, error) {
	url := fmt.Sprintf("http://geoip.nekudo.com/api/%s", d.GetIp())
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var ipMeta struct {
		Country map[string]string `json:"country"`
	}
	if err := dec.Decode(&ipMeta); err != nil {
		return "", err
	}

	s, _ := ipMeta.Country["code"]
	return s, nil
}
