# filesrvr

A super-simple file server.

## Installation

`filesrvr` will run on most Linux and Mac OS X systems.

To install it, just find the appropriate one-liner below - based on the destination O/S and architecture - and copy-paste it into your terminal.

Feel free to change the install dir - `$HOME/bin` in the examples below - to be something more appropriate for your needs.

### Linux (32-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-linux-386.tar.gz | tar -zxf - -C $HOME/bin
```

### Linux (64-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-linux-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Intel)

```
curl -s -L -o - https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-darwin-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o - https://github.com/alasdairmorris/filesrvr/releases/latest/download/filesrvr-darwin-arm64.tar.gz | tar -zxf - -C $HOME/bin
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
  filesrvr -r ROOTDIR -p PORT
  filesrvr -h | --help
  filesrvr --version

Options:
  -h, --help              Show this screen.
  --version               Show version.
  -r, --rootdir DIR       Root directory for files.
  -p, --port PORT         Port to listen on.

Homepage: https://github.com/alasdairmorris/filesrvr
```

## Examples

```
filesrvr -r ~/public -p 8081
```
