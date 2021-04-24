package apod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Apod struct {
		Apikey   string `json:"apikey"`
		Apod_end string `json:"apod_end"`
	} `json:"apod"`
}

func (c Config) LoadConfig(file string) {
	configfile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configfile)
	jsonParser.Decode(&c)
}

type APODResponce struct {
	Title           string `json:"title"`
	Date            string `json:"date"`
	Explanation     string `json:"explanation"`
	Hdurl           string `json:"hdurl"`
	Media_type      string `json:"media_type"`
	Service_version string `json:"service_version"`
	Url             string `json:"url"`
}

func (a APODResponce) LookUpAPOD(d string) {
	var c Config
	c.LoadConfig("settings.json")
	req, _ := http.NewRequest("GET", config.Apod.Apod_end, nil)
	q := req.URL.Query()
	q.Add("api_key", config.Apod.Apikey)
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

	json.Unmarshal(data, &a)

	return jsonresponce

}
