package common

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigInfo struct {
	Persons map[int]float64 `yaml: persons`
	Bridges map[int]int `yaml: bridges`
	Problem map[int][]int `yaml: problem`
}

func LoadYamlFile(filePath string) (*ConfigInfo, error){
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var c *ConfigInfo
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err
	}

	for key, value := range c.Persons {
		if key < 0 {
			return nil, errors.New("persons id cannot be negative, please check the yaml file")
		}
		if value < 0 {
			return nil, errors.New("persons speed cannot be negative, please check the yaml file")
		}
	}

	for key, value := range c.Bridges {
		if key < 0 {
			return nil, errors.New("bridge id cannot be negative, please check the yaml file")
		}
		if value <= 0 {
			return nil, errors.New("bridge length cannot be negative or zero, please check the yaml file")
		}
	}

	for key, value := range c.Problem {
		if key < 0 {
			return nil, errors.New("bridge id cannot be negative in the problem map, please check the yaml file")
		} else if c.Bridges[key] <= 0 {
			str := fmt.Sprintf("bridge id %v provided in problem map doesnt exist in bridges list, please check the yaml file", key)
			return nil, errors.New(str)
		}
		for _, val := range value {
			if val < 0 {
				return nil, errors.New("person id cannot be negative in problem map, please check the yaml file")
			}
			if c.Persons[val] <= 0 {
				str := fmt.Sprintf("person id %v provided for bridge id %v in problem map doesnt exist in persons list, please check the yaml file", val, key)
				return nil, errors.New(str)
			}
		}
	}


	return c, nil
}