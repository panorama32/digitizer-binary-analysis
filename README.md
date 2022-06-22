# digitizer-binary-analysis

## How to Use

```
go >= 1.17
```

`go run main.go`, then images are outputed under `/output`.

## For Development

### hexdump

`hexdump $FILE_NAME`

`hexdump $FILE_NAME | grep "5555 aaaa" -n`

### Format

[Spec](https://www.bbtech.co.jp/download-files/16ch-1gsps-digitizer/8ch-5gsps-digitizer-spec-1.0.1.pdf): p.19

One Payload Size: `1024 * 2 * 8 + 4 = 16388`
