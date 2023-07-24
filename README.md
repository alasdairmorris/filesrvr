# filesrvr

A super-simple file server.

## Installation

`filesrvr` will run on most Linux, MacOS and Windows systems.

To install it, just `cd` into the directory in which you wish to install it and then copy-paste the appropriate one-liner from below (based on the destination O/S and architecture).

### Linux (32-bit)

```
curl -s -L -o filesrvr https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-linux-386 && chmod +x filesrvr
```

### Linux (64-bit)

```
curl -s -L -o filesrvr https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-linux-amd64 && chmod +x filesrvr
```

### Mac OS X (Intel)

```
curl -s -L -o filesrvr https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-darwin-amd64 && chmod +x filesrvr
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o filesrvr https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-darwin-arm64 && chmod +x filesrvr
```

### Windows (32-bit)

```
curl -s -L -o filesrvr.exe https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-windows-386.exe
```

### Windows (64-bit)

```
curl -s -L -o filesrvr.exe https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-windows-amd64.exe
```


### Build From Source

If you have Go installed and would prefer to build the app yourself, you can do:

```
go install github.com/alasdairmorris/filesrvr@latest
```


## Usage

```
A super-simple file server.

Usage:
  filesrvr -r ROOTDIR -p PORT [-a USER:PASS]
  filesrvr -h | --help
  filesrvr --version

Options:
  -h, --help              Show this screen.
  --version               Show version.
  -r, --rootdir DIR       Root directory for files.
  -p, --port PORT         Port to listen on.
  -a, --auth USER:PASS    Enable basic auth protection, using user/pass combo.

Homepage: https://github.com/alasdairmorris/filesrvr
```

## Examples

```
filesrvr -r ~/public -p 8081
```

## License

[MIT](LICENSE)

