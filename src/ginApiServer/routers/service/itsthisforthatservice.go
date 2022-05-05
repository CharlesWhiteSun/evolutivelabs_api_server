package service

import (
	"io/ioutil"
	"net/http"
)

// ItsThisForThatDailySentence 帶參數每日一句
type ItsThisForThatDailySentence struct {
	ReqURL   string      // Request URL
	RespData interface{} // Response Data
}

// NewItsThisForThatDailySentence new ItsThisForThatDailySentence
func NewItsThisForThatDailySentence(url string) *ItsThisForThatDailySentence {
	return &ItsThisForThatDailySentence{
		ReqURL: url,
	}
}

func (sentence *ItsThisForThatDailySentence) GetItsThisForThatSentence() (*ItsThisForThatDailySentence, error) {
	resp, err := http.Get(sentence.ReqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sentence.RespData = string(data)

	return sentence, nil
}
