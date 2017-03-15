package configuration

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Host     string `json:"Host,required"`
	Port     uint   `json:"Port,required"`
	Database string `json:"Database, required"`
	User     string `json:"User,required"`
	Password string `json:"Password,required"`
}

func Usage() {

}

func ReadCmd(args []string) (*Config, error) {
	fileCommand := flag.NewFlagSet("file", flag.ExitOnError)
	cmdCommand := flag.NewFlagSet("cmd", flag.ExitOnError)

	fileTextPtr := fileCommand.String("from", "", "Path for configuration file.")
	cmdHostPtr := cmdCommand.String("host", "", "Hostname")
	cmdPortPtr := cmdCommand.Uint("port", 0, "Port")
	cmdDBPtr := cmdCommand.String("dbname", "", "Database name")
	cmdUserPtr := cmdCommand.String("user", "", "Username")
	cmdPassPtr := cmdCommand.String("pass", "", "Password")

	if len(args) < 2 {
		return &Config{
			Host:     "127.0.0.1",
			Port:     5432,
			Database: "movies",
			User:     "postgres",
			Password: "postgres",
		}, nil
	}

	switch args[1] {
	case "file":
		fileCommand.Parse(args[2:])
	case "cmd":
		cmdCommand.Parse(args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if fileCommand.Parsed() {
		if fileTextPtr == nil {
			fileCommand.PrintDefaults()
			os.Exit(1)
		}
		return ReadConfig(*fileTextPtr)
	}

	if cmdCommand.Parsed() {
		if cmdHostPtr != nil && cmdPortPtr != nil && cmdDBPtr != nil && cmdUserPtr != nil && cmdPassPtr != nil {
			return &Config{
				Host:     *cmdHostPtr,
				Port:     *cmdPortPtr,
				Database: *cmdDBPtr,
				User:     *cmdUserPtr,
				Password: *cmdPassPtr,
			}, nil
		}
	}

	flag.PrintDefaults()
	return nil, fmt.Errorf("Read usage!")
}

func ReadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
