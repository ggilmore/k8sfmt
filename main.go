package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var tmp interface{}
	err = yaml.Unmarshal(input, &tmp)
	if err != nil {
		log.Fatal(err)
	}

	k8sFmtYaml, err := yaml.Marshal(tmp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(k8sFmtYaml))
}
