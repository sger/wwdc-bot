[![Build Status](https://travis-ci.org/sger/wwdc-bot.svg?branch=master)](https://travis-ci.org/sger/wwdc-bot)
[![Coverage Status](https://coveralls.io/repos/github/sger/wwdc-bot/badge.svg?branch=master)](https://coveralls.io/github/sger/wwdc-bot?branch=master)

# WWDC Slack Bot

Searchable full-text transcripts of WWDC sessions

WWDC Slack Bot uses [https://github.com/ASCIIwwdc/asciiwwdc.com](https://github.com/ASCIIwwdc/asciiwwdc.com) API

Usage:

First create a new bot in Slack more info here [https://api.slack.com/bot-users](https://api.slack.com/bot-users) then get project with go get command:

```
$ go get https://github.com/sger/wwdc-bot
$ ./wwdc-bot my-bot-token
```
open up your slack client and send a message to you bot 
for example find a session with id 101 and year 2015

my-bot-name: 101 2015 

or search for uiview

my-bot-name: uiview

see all sessions here [http://asciiwwdc.com](http://asciiwwdc.com)

maybe you want to deploy this go app to heroku from command line type:

```
$ heroku create
$ git add .
$ git commit -m "adding project to heroku service"
$ git push heroku master
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

