package config

import (
	"encoding/json"
	"fmt"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
)

type PbSyncConfig struct {
	AccessToken      string `json:"access_token"`
	Folder           string `json:"folder"`
	AllowDeleteFiles bool   `json:"allow_delete_files"`
	OnSdCard         bool   `json:"on_sd_card"`
}

func LoadConfig(files billy.Filesystem, configPath string) (PbSyncConfig, error) {
	config := PbSyncConfig{}

	configData, err := util.ReadFile(files, configPath)
	if err != nil {
		return config, fmt.Errorf("failed to loaded config %v: %w", configPath, err)
	}

	err = json.Unmarshal(configData, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config file %w", err)
	}

	return config, nil
}
