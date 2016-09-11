package psapi

import (
	"gopkg.in/resty.v0"
	"bytes"
	"net/url"
	"encoding/json"
	"strings"
)

var Limit = 50
var Unlimited = 999

type prestaService struct {
	url string
	key string
	client *resty.Client
}

func NewService(apiUrl string, apiKey string) *prestaService {
	url, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}
	
	var buffer bytes.Buffer
	buffer.WriteString(url.Scheme)
	buffer.WriteString("://")
	buffer.WriteString(apiKey)
	buffer.WriteString("@")
	buffer.WriteString(url.Host)
	buffer.WriteString("/")
	buffer.WriteString(url.Path)
	buffer.WriteString("/")
	
	psService := &prestaService{buffer.String(), apiKey, resty.New()}
	psService.client.SetHeader("Accept", "application/xml")
	
	return psService
}

func (service prestaService) getResource(resource string, limit int) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(service.url)
	buffer.WriteString(resource)
	buffer.WriteString("?display=full")
	buffer.WriteString("&output_format=JSON")
	buffer.WriteString("&limit=")
	buffer.WriteString(string(limit))

	resp, _ := resty.R().Get(buffer.String())

	return resp.Body()
}

func (service prestaService) GetProducts() []Product {
	resource := "products"
	obj := make([]Product)
	json.Unmarshal(service.getResource(resource, Limit), &obj)
	products := sliceToMapById(obj[strings.ToLower(resource)])

	return products
}

func (service prestaService) GetCategories() map[string]Category {
	resource := "categories"
	obj := make([]Category)
	json.Unmarshal(service.getResource(resource, Unlimited), &obj)
	
	for k, element := range obj
	
	return obj[strings.ToLower(resource)]
}