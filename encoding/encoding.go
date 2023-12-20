package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	jsonData, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("Json file reading error: %s", err.Error())
		return err
	}

	err = json.Unmarshal(jsonData, &j.DockerCompose)
	if err != nil {
		fmt.Printf("Json deserialise error: %s", err.Error())
		return err
	}
	b, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		fmt.Printf("Yaml serialise error: %s", err.Error())
		return err
	}
	os.WriteFile(j.FileOutput, b, 0755)
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlData, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("Yaml file reading error: %s", err.Error())
		return err
	}

	err = yaml.Unmarshal(yamlData, &y.DockerCompose)
	if err != nil {
		fmt.Printf("Yaml deserialise error: %s", err.Error())
		return err
	}
	b, err := json.Marshal(y.DockerCompose)
	if err != nil {
		fmt.Printf("Json serialise error: %s", err.Error())
		return err
	}
	os.WriteFile(y.FileOutput, b, 0755)
	return nil
}
