package rusprofile

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"wrapper_for_rusprofile/pkg/value"
)

type Client struct {
	url string
}

type InnInformation struct {
	Inn     string `json:"inn"`
	Kpp     string `json:"raw_ogrn"`
	Fio     string `json:"ceo_name"`
	Company string `json:"name"`
}

type GetFullInformationByInnResponse struct {
	UlCount int              `json:"ul_count"`
	Ul      []InnInformation `json:"ul"`
}

func New() Client {
	return Client{url: value.URL}
}

func (c *Client) GetFullInformationByInn(inn string) ([]InnInformation, error) {
	var innInfoResponse GetFullInformationByInnResponse
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("query", inn)
	q.Add("action", "search")
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []InnInformation{}, err
	}
	err = json.Unmarshal(body, &innInfoResponse)
	if err != nil {
		return []InnInformation{}, err
	}
	log.Println(innInfoResponse)
	return innInfoResponse.Ul, nil
}
