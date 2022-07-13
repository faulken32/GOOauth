/*
 * Copyright (c) 2022.
 *
 *   canicatti.nicolas@gmail.com
 *
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package Utils

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

var configPath = "config.yml"

var AppConfig *Configuration

type Configuration struct {
	//Database info
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		DbName   string `yaml:"db_name"`
	}
	// secret config
	Secret struct {
		SecretKey string `yaml:"secret_key"`
	}
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags(testEnv bool) (string, error) {
	// String that contains the configured configuration path

	var name string
	var path string
	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	if testEnv {
		name = "configForTest"
		path = "../config_test.yml"
	} else {
		name = "config"
		path = "config.yml"
	}
	flag.StringVar(&configPath, name, path, "path to config file")

	flag.Parse()

	return configPath, nil
}

// ReadConfig file override value for test env to config_test.yml
func ReadConfig(testEnv bool) error {

	flags, err := ParseFlags(testEnv)

	if err != nil {
		return err
	}

	if flags != configPath && testEnv {
		configPath = flags
	}

	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		return err
	}
	return nil
}
