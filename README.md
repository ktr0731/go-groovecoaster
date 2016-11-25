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
client := groovecoaster.New()

// Your information in /
client.Personal()
client.Statistics()

// Outline of currently held event
client.EventSummary()

// All played musics summary
client.MusicSummary()
// The detail of a music by MusicID (client.MusicSummary() contains all MusicID)
client.Music(161)
// Total number of ranking pages by MusicID and difficulty
client.MusicRankingPageCount(161, groovecoaster.Normal)
// The ranking of a music by MusicID and difficulty
client.MusicRanking(161, groovecoaster.Normal, 0)

// Shop information
client.ShopSummary()
// All items that can be purchased
client.ShopAvatars()
client.ShopItems()
client.ShopSkins()
client.ShopMusics()
client.ShopMessages()

// All played online battle summary
client.OnlineBattleSummary()
// The detail of a result of online battle by EID and MID (client.OnlineBattleSummary() contains all EID and MID)
client.OnlineBattle(34, 6448)

// All events summary that has been held until now
client.EventArchiveSummary()
// The detail of a result of an event by EventID (client.EventArchiveSummary() contains all EventID)
client.EventArchive()
```

## License
Please see LICENSE.
