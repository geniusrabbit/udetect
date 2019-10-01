# UDetect protocol

[![Build Status](https://travis-ci.org/sspserver/udetect.svg?branch=master)](https://travis-ci.org/sspserver/udetect)
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

* [ ] Provide minimal realisation with geo and device detection without user information
* [ ] Extend Golan objects with extra methods to manipulate objects
* [ ] Add support UUID converter
