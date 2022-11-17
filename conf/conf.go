package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
)

type (
	configT[T any] struct {
		Data T
	}
)

var config interface{}

// loads "config.json" in the main directory and returns it value of type T

func Load[T any]() T {
	c := configT[T]{}
	config = c
	data, err := os.ReadFile(GetPath() + "config.json")

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		lvn.Logger.Panicf("Error while reading config: %s, looking for file: %s", err, GetPath()+"config.json")
	}
	if err != nil && errors.Is(err, os.ErrNotExist) {
		lvn.Logger.Error("file not found, creating new: %s", err, GetPath()+"config.json")
		saveConf[T]()
		return c.Data
	}

	err = json.Unmarshal(data, &c.Data)
	if err != nil {
		lvn.Logger.Panicf("Error while parsing config: %s, file is:", err, GetPath()+"config.json")
	}
	config = c
	saveConf[T]()
	return c.Data
}

func Get[T any]() T {
	return config.(configT[T]).Data
}

func Set[T any](c T) {
	config = configT[T]{Data: c}
	saveConf[T]()
}

func saveConf[T any]() {
	data, _ := json.MarshalIndent(config.(configT[T]).Data, "", "    ")
	err := os.WriteFile(GetPath()+"config.json", data, 0644)
	if err != nil {
		fmt.Println(err)
		lvn.Logger.Panicf("Error while saving config: %s", err)
	}
}

func GetPath() string {
	path, _ := os.Executable()
	path = path[:strings.LastIndex(path, "/")+1]
	return path
}
