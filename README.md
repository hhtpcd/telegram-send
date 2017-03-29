### telegram-send

Tiny go project to send telegram message from command line

```bash
go build
./telegram-send --help                                                                                                                                                                                                                  ⏎
Usage of ./telegram-send:
  -m string
        Message to send
```

Reads API token and Recipient ID from env vars.

```bash
TELEGRAM_API_TOKEN=super-secret-token
TELEGRAM_API_RECIPIENT=00000000
```

> example send message

```bash
telegram-send -m "hello this is Turbert, hope you are well"
```
