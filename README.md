# URL Shortener

<img src="https://cdn.worldvectorlogo.com/logos/gopher.svg" width="175px" />

A program written in Go that forwards paths to other URLs (similar to Bitly)

## Usage

- First, clone the repository to your local machine:

```
    $ git clone https://github.com/DreamLineLove/URLShortener.git
```

- Run the binary or build it yourself (<a href="https://go.dev/learn/" target="_blank">Go toolchain required</a>):

```go
    // The following each runs the binary with default flags
    // Run the one for your operating system
    $ ./bin/apple-silicon-mac
    $ ./bin/intel-mac
    $ ./bin/windows
    $ ./bin/linux

    // Or build the binary yourself
    $ go build .
```

## Flags
- By default -path is "redirect.yml".
- However, you can set the flags anyway you like:
```go
    // The data source is set as "mydata.json" from the same directory
    // It is run on a Linux machine
    $ ./bin/linux -path mydata.json
```
- Both ".json" and ".yml" file extensions are supported.
