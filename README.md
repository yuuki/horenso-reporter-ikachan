horenso-reporter-ikachan
========================

[horenso](https://github.com/Songmu/horenso) reporter plugin for [ikachan](https://github.com/yappo/p5-App-Ikachan).

## Usage

```
horenso-reporter-ikachan -c/--channel CHANNEL [-t/--type notice|privmsg] [-T/--tag TAG] [-e/--error-only]
```

```
$ horenso -r 'horenso-reporter-ikachan -c #cron -t notice -e' -- command 2>&1 | logger -t clear_cache
```

## License

MIT

## Author

[y_uuki](https://github.com/yuuki)
