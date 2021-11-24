<h1 align="center">PocketBook Dropbox downloader</h1>
<p align="center">
    <img width="80%" src="./.github/assets/feature-image.png">
</p>
<p align="center">
    DropBox client for PocketBook reader written on Go.
</p>
<p align="center">
    <a href="https://github.com/evg4b/pb-dropbox-downloader/actions?query=workflow%3AGo+branch%3Amaster">
        <img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/evg4b/pb-dropbox-downloader/Go?label=Build">
    </a>
    <a href="https://github.com/evg4b/pb-dropbox-downloader/blob/master/LICENSE">
        <img alt="GitHub license" src="https://img.shields.io/github/license/evg4b/pb-dropbox-downloader?label=License">
    </a>
    <a href="https://github.com/evg4b/pb-dropbox-downloader/blob/main/go.mod">
        <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/evg4b/pb-dropbox-downloader">
    </a>
    <a href="https://goreportcard.com/report/github.com/evg4b/pb-dropbox-downloader">
        <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/evg4b/pb-dropbox-downloader">
    </a>
    <a href="https://app.codecov.io/gh/evg4b/pb-dropbox-downloader">
        <img alt="Codecov" src="https://img.shields.io/codecov/c/gh/evg4b/pb-dropbox-downloader">
    </a>
</p>

## How to install

1. Copy `pb-dropbox-downloader.app` to `/applications` folder on your reader.
2. Fill configuration from `config.example.json`, and save it as `pb-dropbox-downloader-config.json`.
3. Copy `pb-dropbox-downloader-config.json` to `/system/config` folder on reader.
4. Turn on your book reader, go to application > pb-dropbox-downloader

## How to build

**Requirements**: [task v3](https://taskfile.dev/), [golang](https://golang.org/), [docker](https://www.docker.com/). [golang-ci-lint](https://golangci-lint.run/)

Use task for run, build and test application:

``` bash
task # to run application

task lint # to lint code

task build # to build .app file for reader

task test # to run all tests in docker container

task test_local # to run all tests on local machine
```

### Custom build

You can build application with custom [ldflags flags](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications-ru). 

powershell:
``` powershell
$env:GOOS = 'linux'
$env:GOARCH = 'arm'
$env:GOARM = '5'
go build -ldflags="-s -w -X <your custom fdflegs>" -o pb-dropbox-downloader.app .
```

bash:
``` bash
GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w -X <your custom fdflegs>" -o pb-dropbox-downloader.app .
```

Flags :
- `main.parallelism` - Number of goroutines used for downloading files (default value `3`)
- `main.logFileName` - Name of log file (default value `pb-dropbox-downloader.log`)
- `main.databaseFileName` - Name of file for data storage (default value `pb-dropbox-downloader.bin`)
- `main.configFileName` - Name of configuration file (default  value`pb-dropbox-downloader-config.json`)

## Testing: 

Currently this application testes only on next devices:
1. Pocketbook 624
2. Reader book 2 (this device has no application item in menu, but you can find  and run program from `gallery`)
