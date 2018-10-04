package secrets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type Secrets interface {
	Validate() bool
}

func LoadFromJson(filename string, secrets Secrets) error {
	secrets_file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(secrets_file, &secrets)
	return err
}

func LoadFromEnvironment(secrets Secrets) {
	v := reflect.ValueOf(secrets).Elem()
	t := v.Type()
	
	for index := 0; index < v.NumField(); index++ {
		field := t.Field(index)
		json_value := field.Tag.Get("json")
		value := os.Getenv(json_value)
		if value != "" {
			log.Println(json_value + " has value " + value)
			v.Field(index).SetString(value)
		}
	}
}
