# go-groovecoaster
[![CircleCI](https://circleci.com/gh/lycoris0731/go-groovecoaster.svg?style=svg&circle-token=64c5cb9f75f93df522eecfd16ddb0d1e517e1b42)](https://circleci.com/gh/lycoris0731/go-groovecoaster)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)  

===

## **Attention**
- go-groovecoaster is now developping...
- Unofficial library. So don't abuse this. Please take full responsibility for your action.

## Description  
Go-GrooveCoaster is an unofficial client library for GrooveCoaster.  

## Equipments
- Go

## Installation
``` sh
$ go get github.com/lycoris0731/go-groovecoaster/groovecoaster
```

## Usage
``` go
gc, err := groovecoaster.New()

gc.Personal()
gc.MusicList()
gc.Music(509) // Music ID
```

## License
Please see LICENSE.
