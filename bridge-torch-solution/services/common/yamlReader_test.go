package common

import (
	"fmt"
	"testing"
)

func TestConfigInfo_LoadYamlFile(t *testing.T) {
	c, err := LoadYamlFile("../resources/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		panic(1)
	}
	fmt.Println(c)
}
