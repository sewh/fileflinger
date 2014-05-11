# fileflinger

A simple utility for sharing files over a network via HTTP. Utilises the Go standard HTTP library for concurrency and reliability. Built to replace `python -m SimpleHTTPServer .`

## Building
```bash
go get github.com/stevie-holdway/fileflinger
go build github.com/stevie-holdway/fileflinger
```

## Usage
```
% ./fileflinger -h
Usage of ./fileflinger:
  -help=false: Display fileflinger usage.
  -path=".": The path to the directory to serve.
  -port="6606": The port for the server to listen on.
```
