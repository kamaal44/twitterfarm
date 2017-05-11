package utils

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// Project ...
type Project struct {
	ID                 string    `json:"id" yaml:"id"`
	Name               string    `json:"name" yaml:"name"`
	ConsumerKey        string    `json:"consumer-key" yaml:"consumer-key"`
	ConsumerSecret     string    `json:"consumer-secret" yaml:"consumer-secret"`
	AccessToken        string    `json:"access-token" yaml:"access-token"`
	AccessTokenSecret  string    `json:"access-token-secret" yaml:"access-token-secret"`
	ElasticsearchHost  string    `json:"elasticsearch-host" yaml:"elasticsearch-host"`
	ElasticsearchIndex string    `json:"elasticsearch-index" yaml:"elasticsearch-index"`
	Keyword            string    `json:"keyword" yaml:"keyword"`
	DateCreated        time.Time `json:"date-created" yaml:"date-created"`
}

// ID : Generate Id String
func ID(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

// CreateFile : Save configuration file
func CreateFile(path string, content []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}

// ReadFile ...
func (c *Project) ReadFile(path string) *Project {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil
	}

	return c
}
