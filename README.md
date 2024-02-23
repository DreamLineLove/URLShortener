# URL Shortener

<img src="https://cdn.worldvectorlogo.com/logos/gopher.svg" width="175px" />

http.Handler that forwards paths to other URLs (similar to Bitly)

## Usage

- First, clone the repository to your local machine:

```
    $ git clone https://github.com/DreamLineLove/URLShortener.git
```

- Run the binary or build it yourself (<a href="https://go.dev/learn/" target="_blank">Go toolchain required</a>):

```go
    // This will run the binary with default flags
    $ ./URLShortener

    // Or build the binary yourself
    $ go build .
```

## Flags
- By default -filename is "source.yml".
- However, you can set the flags anyway you like:
```go
    // The data source is set as "mydata.json" from the same directory
    $ ./URLShortener -filename mydata.json
```
- Both ".json" and ".yml" file extensions are supported.
