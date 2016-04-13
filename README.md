horenso-reporter-ikachan
========================

[horenso](https://github.com/Songmu/horenso) reporter plugin for [ikachan](https://github.com/yappo/p5-App-Ikachan).

## Usage

```
horenso-reporter-ikachan --host HOSTNAME --channel '#CHANNEL' [--port=PORT] [--type=MSGTYPE] [--error-only]

Application Options:
  -H, --host=hostname         ikachan hostname
  -p, --port=port             ikachan port (4979)
  -c, --channel='#channel'    destination channel
  -t, --type=msgtype          message type notice/privmsg) (notice)
  -e, --error-only            report only when error ocurrs (false)

Help Options:
  -h, --help                  Show this help message
```

```
$ horenso -r 'horenso-reporter-ikachan -c #cron -t notice -e' -- command 2>&1 | logger -t clear_cache
```

## License

MIT

## Author

[y_uuki](https://github.com/yuuki)
