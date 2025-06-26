package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type BankingDetails struct {
	Name    string `yaml:"name"`
	Bank    string `yaml:"bank"`
	Type    string `yaml:"type"`
	Account string `yaml:"account"`
}

type CompanyInfo struct {
	Name     string         `yaml:"name"`
	Address  string         `yaml:"address"`
	Phone    string         `yaml:"phone"`
	Facebook string         `yaml:"facebook"`
	Email    string         `yaml:"email"`
	Logo     string         `yaml:"logo"`
	Banking  BankingDetails `yaml:"banking"`
}

type Config struct {
	Company CompanyInfo `yaml:"company"`
}

var Global Config

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &Global)
}
