# UDetect protocol

[![Build Status](https://github.com/geniusrabbit/udetect/workflows/Tests/badge.svg)](https://github.com/geniusrabbit/udetect/actions?workflow=Tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/geniusrabbit/udetect)](https://goreportcard.com/report/github.com/geniusrabbit/udetect)
[![GoDoc](https://godoc.org/github.com/geniusrabbit/udetect?status.svg)](https://godoc.org/github.com/geniusrabbit/udetect)
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
* [x] Add support UUID converter## License

[LICENSE](LICENSE)

Copyright 2024 Dmitry Ponomarev & Geniusrabbit

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

<http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
