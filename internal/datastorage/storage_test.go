package datastorage_test

import (
	"encoding/json"
	"errors"
	"pb-dropbox-downloader/internal/datastorage"
	"pb-dropbox-downloader/testing/testutils"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/stretchr/testify/assert"
)

func TestFileStorage_Get(t *testing.T) {
	fs := testutils.FsFromMap(t, map[string]string{
		"config.bin": `{
			"key1": "value1",
			"key2": "value2"
		}`,
		"invalid.bin": `{ "key1": "value1" `,
	})

	tests := []struct {
		name          string
		key           string
		configFile    string
		expected      string
		expectedError string
	}{
		{
			name:          "db file does not exists",
			key:           "test_key",
			configFile:    "notexist.bin",
			expectedError: datastorage.ErrKeyDoesNotExists.Error(),
		},
		{
			name:          "key does not exists",
			key:           "test_key",
			configFile:    "config.bin",
			expectedError: datastorage.ErrKeyDoesNotExists.Error(),
		},
		{
			name:       "key exists",
			key:        "key1",
			configFile: "config.bin",
			expected:   "value1",
		},
		{
			name:          "invalid database",
			key:           "key1",
			configFile:    "invalid.bin",
			expectedError: "unexpected end of JSON input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := datastorage.NewFileStorage(
				datastorage.WithFilesystem(fs),
				datastorage.WithConfigPath(tt.configFile),
				datastorage.WithMarshalFunc(json.Marshal),
				datastorage.WithUnmarshalFunc(json.Unmarshal),
			)

			actual, err := storage.Get(tt.key)

			testutils.AssertError(t, tt.expectedError, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileStorage_ToMap(t *testing.T) {
	fs := testutils.FsFromMap(t, map[string]string{
		"config.bin": `{
			"key1": "value1",
			"key2": "value2"
		}`,
		"invalid.bin": `{ "key1": "value1" `,
	})

	tests := []struct {
		name          string
		configFile    string
		expected      map[string]string
		expectedError string
	}{
		{
			name:       "db file does not exists",
			configFile: "notexist.bin",
			expected:   map[string]string{},
		},
		{
			name:       "key exists",
			configFile: "config.bin",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name:          "invalid database",
			configFile:    "invalid.bin",
			expectedError: "unexpected end of JSON input",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := datastorage.NewFileStorage(
				datastorage.WithFilesystem(fs),
				datastorage.WithConfigPath(tt.configFile),
				datastorage.WithMarshalFunc(json.Marshal),
				datastorage.WithUnmarshalFunc(json.Unmarshal),
			)

			actual, err := storage.ToMap()

			testutils.AssertError(t, tt.expectedError, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileStorage_FromMap(t *testing.T) {
	storage := datastorage.NewFileStorage(
		datastorage.WithFilesystem(memfs.New()),
	)

	expected := map[string]string{
		"key1": "unique test value",
		"key2": "lorem inpsum",
	}

	storage.FromMap(expected)

	actual, _ := storage.ToMap()
	assert.Equal(t, expected, actual)
}

func TestFileStorage_Add(t *testing.T) {
	fs := testutils.FsFromMap(t, map[string]string{
		"storage.bin": `{ "key1": "value1", "key2": "value2" }`,
		"invalid.bin": `{ "key1": "value1" `,
	})

	tests := []struct {
		name          string
		configFile    string
		expected      map[string]string
		expectedError string
	}{
		{
			name:       "add successful",
			configFile: "storage.bin",
			expected: map[string]string{
				"key1": "unique test value",
				"key2": "value2",
			},
		},
		{
			name:          "invalid database",
			configFile:    "invalid.bin",
			expectedError: "unexpected end of JSON input",
			expected:      map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := datastorage.NewFileStorage(
				datastorage.WithUnmarshalFunc(json.Unmarshal),
				datastorage.WithFilesystem(fs),
				datastorage.WithConfigPath(tt.configFile),
			)

			err := storage.Add("key1", "unique test value")

			actual, _ := storage.ToMap()
			testutils.AssertError(t, tt.expectedError, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileStorage_Remove(t *testing.T) {
	fs := testutils.FsFromMap(t, map[string]string{
		"storage.bin": `{ "key1": "value1", "key2": "value2" }`,
		"invalid.bin": `{ "key1": "value1" `,
	})

	tests := []struct {
		name          string
		configFile    string
		expected      map[string]string
		expectedError string
	}{
		{
			name:       "remove successful",
			configFile: "storage.bin",
			expected: map[string]string{
				"key2": "value2",
			},
		},
		{
			name:          "invalid database",
			configFile:    "invalid.bin",
			expectedError: "unexpected end of JSON input",
			expected:      map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := datastorage.NewFileStorage(
				datastorage.WithUnmarshalFunc(json.Unmarshal),
				datastorage.WithFilesystem(fs),
				datastorage.WithConfigPath(tt.configFile),
			)

			err := storage.Remove("key1")

			actual, _ := storage.ToMap()
			testutils.AssertError(t, tt.expectedError, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileStorage_KeyExists(t *testing.T) {
	fs := testutils.FsFromMap(t, map[string]string{
		"storage.bin": `{ "key1": "value1" }`,
		"invalid.bin": `{ "key1": "value1" `,
	})

	tests := []struct {
		name          string
		configFile    string
		key           string
		expected      bool
		expectedError string
	}{
		{
			name:       "key exist",
			key:        "key1",
			configFile: "storage.bin",
			expected:   true,
		},
		{
			name:       "key does not exist",
			key:        "key3",
			configFile: "storage.bin",
			expected:   false,
		},
		{
			name:          "invalid database",
			configFile:    "invalid.bin",
			expectedError: "unexpected end of JSON input",
			expected:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := datastorage.NewFileStorage(
				datastorage.WithUnmarshalFunc(json.Unmarshal),
				datastorage.WithFilesystem(fs),
				datastorage.WithConfigPath(tt.configFile),
			)

			actual, err := storage.KeyExists(tt.key)

			testutils.AssertError(t, tt.expectedError, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileStorage_Commit(t *testing.T) {
	tests := []struct {
		name          string
		expected      map[string]string
		expectedError string
		marshalFunc   datastorage.MarshalFunc
	}{
		{
			name:        "commit correctly",
			expected:    map[string]string{},
			marshalFunc: json.Marshal,
		},
		{
			name:          "marshalling error",
			expected:      map[string]string{},
			expectedError: "test error",
			marshalFunc: func(i interface{}) ([]byte, error) {
				return []byte{}, errors.New("test error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := testutils.FsFromMap(t, map[string]string{
				"storage.bin": "{ }",
			})
			storage := datastorage.NewFileStorage(
				datastorage.WithMarshalFunc(tt.marshalFunc),
				datastorage.WithFilesystem(fs),
			)
			storage.FromMap(tt.expected)

			err := storage.Commit()

			testutils.AssertError(t, tt.expectedError, err)
			assertDatastorageFile(t, fs, tt.expected)
		})
	}
}

func assertDatastorageFile(t *testing.T, fs billy.Filesystem, expected map[string]string) {
	t.Helper()

	data, err := util.ReadFile(fs, "storage.bin")
	assert.NoError(t, err)
	actual := map[string]string{}
	err = json.Unmarshal(data, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
