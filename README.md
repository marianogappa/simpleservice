# simpleservice - [![GoDoc](https://godoc.org/github.com/marianogappa/simpleservice?status.svg)] [![Build Status](https://img.shields.io/travis/marianogappa/simpleservice.svg)](https://travis-ci.org/marianogappa/simpleservice)  [![Go Report Card](https://goreportcard.com/badge/github.com/marianogappa/simpleservice?style=flat-square)](https://goreportcard.com/report/github.com/marianogappa/simpleservice) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/marianogappa/simpleservice/master/LICENSE)

an implementation of a service in the MESG Network

## Check that builds

```
make
```

## Test

```
make test
```

## Service configuration

[mesg.yml](mesg.yml)

## Test results

[RESULTS.md](RESULTS.md)

## Features

- Tasks as subpackages implementing the task interface
- Tasks fully tested
- `ExecuteMany` task implements conditional parallel POST requests
- Dockerfile relies on vendored dependencies: builds faster and more reliably
- UUID generator generates monotonically increasing UUIDs: ExecutionID string order is meaningful
