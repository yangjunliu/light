package light

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type environment struct {
	serverName string `yaml:"SERVER_NAME"`
	iP         string `yaml:"IP"`
	port       int    `yaml:"PORT"`
}

var env environment = environment{}

func init() {
	file, err := ioutil.ReadFile("./env.yaml")
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(file, &env)

	fmt.Println(env.serverName)
	fmt.Printf("%+v", env)
}
