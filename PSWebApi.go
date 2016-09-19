package psapi

import (
	"gopkg.in/resty.v0"
	"bytes"
	"fmt"
	"net/url"
	"encoding/xml"
	"strconv"
)

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
	buffer.WriteString(url.Path)
	buffer.WriteString("/")
	
	psService := &prestaService{buffer.String(), apiKey, resty.New()}
	psService.client.SetHeader("Accept", "application/xml")
	psService.client.SetHeader("Content-Type", "text/xml")
	
	return psService
}

func (service prestaService) getUrl(resource string) bytes.Buffer {
	var buffer bytes.Buffer
	buffer.WriteString(service.url)
	buffer.WriteString(resource)
	
	return buffer
}

func (service prestaService) GetResource(resource string, limit int) []byte {
	Url := service.getUrl(resource)
	Url.WriteString("?display=full")
	//Url.WriteString("&output_format=JSON")
	Url.WriteString("&limit=")
	Url.WriteString(strconv.Itoa(limit))
	
	resp, _ := resty.R().Get(Url.String())

	return resp.Body()
}

func (service prestaService) addResource(resource string, Post []byte) []byte {
	Url := service.getUrl(resource)
	resp, _ := resty.R().
	SetBody(Post).
	Post(Url.String())

	return resp.Body()
}

func (service prestaService) GetProducts(Limit int) []Product {
	resource := "products"
	var obj struct {
		Products []Product `xml:"products>product"`
	}
	err := xml.Unmarshal(service.GetResource(resource, Limit), &obj)
	if err != nil {
		fmt.Println(err)
	}

	return obj.Products
}

func (service prestaService) GetCategories(Limit int) []Category {
	resource := "categories"
	var obj struct {
		Categories []Category `xml:"categories>category"`
	}
	err := xml.Unmarshal(service.GetResource(resource, Limit), &obj)
	if err != nil {
		fmt.Println(err)
	}
	
	return obj.Categories
}
/*
func (service prestaService) AddProducts(Product Product) int {
	resource := "products"
	
	
	
	var obj map[string][]Product
	err := xml.Unmarshal(service.getResource(resource, Limit), &obj)
	if err != nil {
		fmt.Println(err)
	}
	//products := sliceToMapById(obj])
	products := obj[strings.ToLower(resource)]

	return products
}*/s