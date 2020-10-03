# passwords_bot

This codebase powers [DoamPasswordsBot](https://t.me/DoamPasswordsBot). It's simple chat bot, that sends you randomly generated password on each request. I used to generate passwords in my console terminal, but chat bot is faster and handy way to do this.

## Build and test

```Bash
$ git clone https://github.com/tonymadbrain/passwords_bot.git
$ cd passwords_bot
$ go build
$ TG_BOT_TOKEN=1234567890:xxxYYYzzz ./passwords_bot
```

## Run on server

```Bash
$ cp passwords_bot /opt/bin/passwords_bot
$ cp passwords_bot.service /etc/systemd/system/passwords_bot.service
$ vim /etc/systemd/system/passwords_bot.service # Set proper TG_BOT_TOKEN
$ systemctl daemon-reload
$ systemctl start passwords_bot
```

## Run on Heroku

```Bash
$ heroku create --region eu my-passwords-bot
$ git push heroku main
$ heroku ps:scale worker=1
```

## TODO

- [ ] Generate binaries ans store them in Github releases
