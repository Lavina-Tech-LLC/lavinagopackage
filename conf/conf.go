package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
)

type (
	configT[T any] struct {
		Data T
		Path string
	}
)

var config interface{}

// loads "config.json" in the main directory and returns it value of type T
func Load[T any](args ...string) T {
	path := ""
	if len(args) == 0 || args[0] == "" {
		path = GetPath() + "config.json"
	} else {
		path = args[0]
		reg := regexp.MustCompile(`\.\.\/`)
		matches := reg.FindAllStringIndex(path, -1)
		path = reg.ReplaceAllString(path, "")

		path = GetPath(len(matches)) + path + "config.json"
	}

	c := configT[T]{
		Path: path,
	}
	data, err := os.ReadFile(path)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		lvn.Logger.Panicf("Error while reading file: %s, looking for file: %s", err, path)
	}

	if err != nil && errors.Is(err, os.ErrNotExist) {
		lvn.Logger.Error("file not found, creating new: %s", err, path)
		config = c
		saveConf[T]()
		return c.Data
	}

	err = json.Unmarshal(data, &c.Data)
	if err != nil {
		lvn.Logger.Panicf("Error while parsing config: %s, file is:", err, path)
	}
	c.Path = path
	config = c
	saveConf[T]()
	return c.Data
}

// Reads json file and returns it value of type T
func LoadFile[T any](path string) T {
	reg := regexp.MustCompile(`\.\.\/`)
	matches := reg.FindAllStringIndex(path, -1)
	path = reg.ReplaceAllString(path, "")

	path = GetPath(len(matches)) + path

	c := configT[T]{
		Path: path,
	}
	data, err := os.ReadFile(path)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		lvn.Logger.Panicf("Error while reading config: %s, looking for file: %s", err, path)
	}
	if err != nil && errors.Is(err, os.ErrNotExist) {
		lvn.Logger.Error("file not found, creating new: %s", err, path)
		config = c
		saveConf[T]()
		return c.Data
	}

	err = json.Unmarshal(data, &c.Data)
	if err != nil {
		lvn.Logger.Panicf("Error while parsing config: %s, file is:", err, path)
	}
	c.Path = path
	config = c
	saveConf[T]()
	return c.Data
}

// args optional: path to load config
func Get[T any](args ...string) T {
	if config == nil {
		if len(args) == 0 {
			panic("config is not loaded and path is not provided")
		}
		Load[T](args[0])
	}
	return config.(configT[T]).Data
}

func Set[T any](c T) {
	config = configT[T]{Data: c, Path: config.(configT[T]).Path}
	saveConf[T]()
}

func saveConf[T any]() {
	data, _ := json.MarshalIndent(config.(configT[T]).Data, "", "    ")
	path := config.(configT[T]).Path
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println(err)
		lvn.Logger.Panicf("Error while saving config: %s", err)
	}
}

// can receive arg as integer of times to go up in subdirectories
func GetPath(args ...any) string {
	up := 0
	if len(args) > 0 {
		up = args[0].(int)
	}
	d := string(rune(os.PathSeparator))

	path, _ := os.Executable()
	// if strings.LastIndex(path, d) < 0 {
	// 	return path
	// }

	path = path[:strings.LastIndex(path, d)]

	pattern := fmt.Sprintf(`%s[^%s]+$`, d, d)
	regRes := regexp.MustCompile(pattern)

	for i := 0; i < up; i++ {
		path = regRes.ReplaceAllString(path, "")
	}
	return path + d
}
