package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// WebHooker ...
type WebHooker interface {
	GetData() interface{}
	GetURL() string
}

// Send ...
func Send(w WebHooker) error {
	jsonData, err := json.Marshal(w.GetData())
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(jsonData)
	res, err := http.Post(w.GetURL(), "application/json", buffer)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return errors.New(res.Status)
	}

	return nil
}
