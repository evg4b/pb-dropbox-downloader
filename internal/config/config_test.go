package config_test

import (
	"pb-dropbox-downloader/internal/config"
	"pb-dropbox-downloader/testing/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	fs := testutils.FsFromMap(t, map[string]string{
		"/mnt/ext1/system/config/pb-dropbox-downloader.bin": `
{
	"access_token": "test access token",
	"folder": "./dropbox",
	"allow_delete_files": true,
	"on_sd_card": false
}
		`,
		"/mnt/ext1/system/config/empty.bin":   "{}",
		"/mnt/ext1/system/config/invalid.bin": "{ this is invalid",
	})

	tests := []struct {
		name          string
		configPath    string
		expected      config.PbSyncConfig
		expectedError string
	}{
		{
			name:       "valid config loading",
			configPath: "/mnt/ext1/system/config/pb-dropbox-downloader.bin",
			expected: config.PbSyncConfig{
				AccessToken:      "test access token",
				Folder:           "./dropbox",
				AllowDeleteFiles: true,
				OnSdCard:         false,
			},
			expectedError: "",
		},
		{
			name:       "empty config loading",
			configPath: "/mnt/ext1/system/config/empty.bin",
			expected: config.PbSyncConfig{
				AccessToken:      "",
				Folder:           "",
				AllowDeleteFiles: false,
				OnSdCard:         false,
			},
			expectedError: "",
		},
		{
			name:          "config not exist",
			configPath:    "/mnt/ext1/system/config/not_exist.bin",
			expected:      config.PbSyncConfig{},
			expectedError: "failed to loaded config /mnt/ext1/system/config/not_exist.bin: file does not exist",
		},
		{
			name:          "invalid config",
			configPath:    "/mnt/ext1/system/config/invalid.bin",
			expected:      config.PbSyncConfig{},
			expectedError: "failed to unmarshal config file invalid character 't' looking for beginning of object key string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := config.LoadConfig(fs, tt.configPath)

			testutils.CheckError(t, tt.expectedError, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
