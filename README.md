# go-groovecoaster
[![GoDoc](https://godoc.org/github.com/lycoris0731/go-groovecoaster/groovecoaster?status.svg)](https://godoc.org/github.com/lycoris0731/go-groovecoaster/groovecoaster)
[![CircleCI](https://circleci.com/gh/lycoris0731/go-groovecoaster.svg?style=svg&circle-token=64c5cb9f75f93df522eecfd16ddb0d1e517e1b42)](https://circleci.com/gh/lycoris0731/go-groovecoaster)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)  

===

## **Attention**
Unofficial library. So don't abuse this. Please take full responsibility for your action.

## Description  
Go-GrooveCoaster is an unofficial client library for Groove Coaster.  
You can get information about `mypage.groovecoaster.jp/sp/`

## Installation
``` sh
$ go get github.com/lycoris0731/go-groovecoaster/groovecoaster
```

## Usage
You will need to set environment variables (or create `.env`)  
That's values are id and password for login to NESiCA and mypage in Groove Coaster.
``` sh
$ export NESICA_CARD_ID
$ export NESICA_PASSWORD
```

``` go
gc, err := groovecoaster.New()

gc.Personal()
gc.MusicList()
gc.Music(509) // Music ID
```

## License
Please see LICENSE.
