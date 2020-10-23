# attendant

[![Build Status](https://travis-ci.com/yusufsyaifudin/attendant.svg?branch=master)](https://travis-ci.com/yusufsyaifudin/attendant)
[![codecov](https://codecov.io/gh/yusufsyaifudin/attendant/branch/master/graph/badge.svg)](https://codecov.io/gh/yusufsyaifudin/attendant)
[![Go Report Card](https://goreportcard.com/badge/github.com/yusufsyaifudin/attendant)](https://goreportcard.com/report/github.com/yusufsyaifudin/attendant)


Attendant just contains abstraction layer required by server to run.

So, every app inside `internal/app` which use REST API as interface, must implement
every abstraction in this `attendant` directory. This make us easy to call handler in `cmd`
when we need to build the binary.

## Usage

```
go get -u github.com/yusufsyaifudin/attendant
```

See `/example` directory for usage.