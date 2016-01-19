package useless

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	// Remote server to push encrypted key to
	HTTPServer string

	// Time to wait in minutes before encrypting user files
	Wait int64
}

func ParseConfig(filename string) (cfg Config, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &cfg); err != nil {
		return
	}
	return
}