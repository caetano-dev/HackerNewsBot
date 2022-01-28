# HackerNews Telegram bot - Golang version

A Telegram bot that serves you with personalized HackerNews articles. You can self host it and make it filter only the news that are relevant to you.

## How to set up

In order to run the bot you will need a [Telegram](https://core.telegram.org/api) API key.

Rename `example.env` to `.env` and put the key in the file.
```env
TOKEN=<TOKEN>
```
Now, you can go to `api.js` and change the relevant topics in the array. By default, the topics are: `privacy, hack, linux, golang, hacker, malware, exploit, leak, CIA, NSA, hacked, breaches, breached, security, OSINT, leaked, GNU, free and open source, open source`.

install packages:
```
$ go get
```
run the application:
```
$ go run main.go
```
