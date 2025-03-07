package models

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var AvaliableIntegrations map[string]IntegrationDefinition

var AvaliableIntegrationFields = map[string]IntegrationField{
	"access_token": {Type: "text", Key: "access_token"},
	"key":          {Type: "text", Key: "key"},
	"secret":       {Type: "text", Key: "secret"},
	"url":          {Type: "text", Key: "url"},
}

func init() {
	AvaliableIntegrations = make(map[string]IntegrationDefinition)
	address := "integrations"
	files, err := ioutil.ReadDir(address)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		var yamlFile IntegrationFile
		yamlData, err := ioutil.ReadFile(address + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlData, &yamlFile)
		if err != nil {
			panic(err)
		}
		integrationDefinition := IntegrationDefinition{Type: yamlFile.Type, Fields: make([]string, 0)}
		if yamlFile.NeedsAccessToken {
			integrationDefinition.Fields = append(integrationDefinition.Fields, "access_token")
		}
		if yamlFile.NeedsKey {
			integrationDefinition.Fields = append(integrationDefinition.Fields, "key")
		}
		if yamlFile.NeedsSecret {
			integrationDefinition.Fields = append(integrationDefinition.Fields, "secret")
		}
		if yamlFile.NeedsUrl {
			integrationDefinition.Fields = append(integrationDefinition.Fields, "url")
		}
		AvaliableIntegrations[integrationDefinition.Type] = integrationDefinition
	}
	fmt.Println(AvaliableIntegrations)
}
