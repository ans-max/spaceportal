package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apikey   = "CTowbyv9K3mxAdPCrTU4c7mdytwgiia9iS6oFsC9"
	apod_end = "https://api.nasa.gov/planetary/apod"
)

type APODResponce struct {
	Title           string `json:"title"`
	Date            string `json:"date"`
	Explanation     string `json:"explanation"`
	Hdurl           string `json:"hdurl"`
	Media_type      string `json:"media_type"`
	Service_version string `json:"service_version"`
	Url             string `json:"url"`
}

func LookUpAPOD(d string) (jresp APODResponce) {
	var jsonresponce APODResponce
	req, _ := http.NewRequest("GET", apod_end, nil)
	q := req.URL.Query()
	q.Add("api_key", apikey)
	q.Add("date", d)
	req.URL.RawQuery = q.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error : ", err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error :", err)
	}

	json.Unmarshal(data, &jsonresponce)

	return jsonresponce

}
