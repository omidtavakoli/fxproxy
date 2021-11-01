package app

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type client struct {
	Validate      *validator.Validate
	HttpClient    *http.Client
	UpdaterClient *http.Client
	Ready         bool
	Config        *Config
}

type Config struct {
	Url string `yaml:"URL"`
}

func NewClient(config *Config, validator *validator.Validate) (*client, error) {
	httpClient := &http.Client{
		Timeout: 100 * time.Millisecond,
	}
	return &client{
		Validate:   validator,
		HttpClient: httpClient,
		Config:     config,
	}, nil
}

func (r client) IsReady() bool {
	return r.Ready
}

func (r client) apiCall(req *http.Request, data interface{}) error {
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		if fail, ok := err.(net.Error); ok && fail.Timeout() {
			return errors.New("request timeout")
		}
		return errors.New("request unknown error")
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(responseData))
	}
	return nil
}

func (r client) CompanyAPI() (string, error) {
	urlString := r.Config.Url + "/company"
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	//q := req.URL.Query()
	var resp string
	err = r.apiCall(req, &resp)
	if err != nil {
		return "", err
	}
	return resp, nil
}
