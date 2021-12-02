package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func GetRate(currency string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest?access_key=cf02f4d5df92d86421daaf722712029a&symbols=%s,RUB", currency))
	if err != nil {
		logrus.Fatalf("error when getting the exchange rate: %s", err.Error())
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("error when getting the exchange rate: %s", err.Error())
		return 0, err
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return 0, err
	}
	rates, ok := jsonData["rates"].(map[string]interface{})
	if ok == false {
		return 0, errors.New("error when getting the exchange rate")
	}
	curRate, ok := rates[currency].(float64)
	if ok == false {
		return 0, errors.New(fmt.Sprintf("no such currency %s", currency))
	}
	rubRate, ok := rates["RUB"].(float64)
	if ok == false {
		return 0, errors.New(fmt.Sprintf("cant't get RUB rate %s", currency))
	}
	return curRate / rubRate, nil

}