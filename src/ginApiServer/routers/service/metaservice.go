package service

import (
	"io/ioutil"
	"net/http"
)

// DailySentence 每日一句
type DailySentence struct {
	ReqURL   string      // Request URL
	RespData interface{} // Response Data
}

// NewDailySentence new DailySentence
func NewDailySentence(url string) *DailySentence {
	return &DailySentence{
		ReqURL: url,
	}
}

func (d *DailySentence) GetSentence() (*DailySentence, error) {
	resp, err := http.Get(d.ReqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	d.RespData = string(data)

	return d, nil
}
