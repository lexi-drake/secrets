package secrets

import (
	"os"
	"testing"
)

type MySecrets struct {
	KeyOne string `json:"key_one"`
	KeyTwo string `json:"key_two"`
}

func (s MySecrets) Validate() bool {
	return true
}

func TestFileDoesNotExist(t *testing.T) {
	var s MySecrets
	
	err := LoadFromJson("not a real path", s)
	if err == nil {
		t.Errorf("file read error not caught!")
	}
}

func TestLoadFromJson(t *testing.T) {
	var s MySecrets
	
	dir := "./test_secrets.json"	
	file, err := os.Create(dir)
	if err != nil {
		t.Errorf("Error creating file: %s", err.Error())
	}
	defer os.Remove(dir)

	text := "{\n\"key_one\":\"golang unit tests!\", \n\"key_two\": \"4501\"\n}"
	_, err = file.WriteString(text)
	if err != nil {
		t.Errorf("Error writing to file: %s", err.Error())
	}
	file.Sync()
	
	err = LoadFromJson(dir, &s)
	if err != nil {
		t.Errorf("Error loading secrets: %s", err.Error())
	}

	if s.KeyOne != "golang unit tests!" {
		t.Errorf("%s != %s", s.KeyOne, "golang unit tests!")
	}
	if s.KeyTwo != "4501" {
		t.Errorf("%s != %s", s.KeyTwo, "4501")
	}
}

func TestLoadFromEnvironment(t *testing.T) {
	prev_one := os.Getenv("key_one")
	prev_two := os.Getenv("key_two")
	defer func() {
		os.Setenv("key_one", prev_one)
		os.Setenv("key_two", prev_two)
	}()
	
	key_one := "golang unit tests!"
	key_two := "4501"
	os.Setenv("key_one", key_one)
	os.Setenv("key_two", key_two)

	var s MySecrets
	LoadFromEnvironment(&s)

	if s.KeyOne != os.Getenv("key_one") {
		t.Errorf("%s != %s", s.KeyOne, key_one)
	}
	if s.KeyTwo != os.Getenv("key_two") {
		t.Errorf("%s != %s", s.KeyTwo, key_two)
	}
}
