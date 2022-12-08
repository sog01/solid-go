package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sog01/solid-go/internal/search/model"
)

type Repository interface {
	CheckHealth() error
	CreateIndex() error
	InsertData(e *model.Employee) error
	SeedingData(idStart, n int) error
	UpdateData(e *model.Employee) error
	DeleteData(id int) error
	SearchData(keyword model.Keyword) ([]*model.Employee, error)
}

type ElasticSearch struct {
	baseURL string
}

func NewElasticSearch(baseURL string) *ElasticSearch {
	return &ElasticSearch{baseURL}
}

func (c *ElasticSearch) CheckHealth() error {
	response, err := http.Get(c.baseURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed check Elasticsearch health: %v", err)
	}

	return nil
}

func (c *ElasticSearch) CreateIndex() error {
	body := `
	{
		"mappings": {
			"properties": {
				"id": {
					"type": "integer"
				},
				"name": {
					"type": "text"
				},
				"address": {
					"type": "text"
				},
				"salary": {
					"type": "float"
				}
			}
		}
	}
	`

	req, err := http.NewRequest("PUT", c.baseURL+"/employee", strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to make a create index request: %v", err)
	}

	httpClient := http.Client{}
	req.Header.Add("Content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make a http call to create an index: %v", err)
	}
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed read create index response: %v", err)
	}

	return nil
}

func (c *ElasticSearch) InsertData(e *model.Employee) error {
	body, _ := json.Marshal(e)

	id := strconv.Itoa(e.Id)
	req, err := http.NewRequest("PUT", c.baseURL+"/employee/_doc/"+id, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to make a insert data request: %v", err)
	}

	httpClient := http.Client{}
	req.Header.Add("Content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make a http call to insert data: %v", err)
	}
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed read insert data response: %v", err)
	}

	return nil
}

func (c *ElasticSearch) SeedingData(idStart, n int) error {
	for i := idStart; i < n; i++ {
		if err := c.InsertData(&model.Employee{
			Id:      i,
			Name:    "person" + strconv.Itoa(i),
			Address: "address" + strconv.Itoa(i),
			Salary:  float64(i * 100),
		}); err != nil {
			return fmt.Errorf("failed seeding data with id %d: %v", i, err)
		}
	}

	return nil
}

func (c *ElasticSearch) UpdateData(e *model.Employee) error {
	body, _ := json.Marshal(map[string]*model.Employee{
		"doc": e,
	})

	id := strconv.Itoa(e.Id)
	req, err := http.NewRequest("POST", c.baseURL+"/employee/_update/"+id, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to make a update data request: %v", err)
	}

	httpClient := http.Client{}
	req.Header.Add("Content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make a http call to update data: %v", err)
	}
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed read update data response: %v", err)
	}

	return nil
}

func (c *ElasticSearch) DeleteData(id int) error {
	req, err := http.NewRequest("DELETE", c.baseURL+"/employee/_doc/"+strconv.Itoa(id), nil)
	if err != nil {
		return fmt.Errorf("failed to make a delete data request: %v", err)
	}

	httpClient := http.Client{}
	req.Header.Add("Content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make a http call to delete data: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed read delete data response: %v", err)
	}

	log.Println("debug delete data response: ", string(responseBody))

	return nil
}

func (c *ElasticSearch) SearchData(keyword model.Keyword) ([]*model.Employee, error) {
	query := fmt.Sprintf(`
	{
		"query": {
			"match": {
				"name": "%s"
			}
		}
	}
	`, keyword)

	req, err := http.NewRequest("GET", c.baseURL+"/employee/_search", strings.NewReader(query))
	if err != nil {
		return nil, fmt.Errorf("failed to make a search data request: %v", err)
	}

	httpClient := http.Client{}
	req.Header.Add("Content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make a http call to search data: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read insert data response: %v", err)
	}

	var searchHits SearchHits
	if err := json.Unmarshal(responseBody, &searchHits); err != nil {
		return nil, fmt.Errorf("failed read unmarshal data response: %v", err)
	}

	var employees []*model.Employee
	for _, hit := range searchHits.Hits.Hits {
		employees = append(employees, hit.Source)
	}

	return employees, nil
}
