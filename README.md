# UDetect protocol

[![Build Status](https://github.com/sspserver/udetect/workflows/run%20tests/badge.svg)](https://github.com/sspserver/udetect/actions?workflow=run%20tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/sspserver/udetect)](https://goreportcard.com/report/github.com/sspserver/udetect)
[![GoDoc](https://godoc.org/github.com/sspserver/udetect?status.svg)](https://godoc.org/github.com/sspserver/udetect)
[![Coverage Status](https://coveralls.io/repos/github/sspserver/udetect/badge.svg)](https://coveralls.io/github/sspserver/udetect)

> License Apache 2.0

The package describes the protocol of personal information detection to the integration of third-party service.
This protocol designed for the common purposes using valuable for advertisers data at the same time there is
no very sensitive for security reasons information about user person.

> In the package, you will find the *proto3* protocol description and prepared version for Golang and build script.
> You can use this protocol definitions to create your own user information service.

## TODO

* [ ] Add docker client example
* [x] Add minimal implementation of detector server
* [x] Provide minimal realisation with geo and device detection without user information
* [x] Add support UUID converter
