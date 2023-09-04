# Welcome to URLShortener!

<img src="https://camo.githubusercontent.com/94761affed6454156a526a0fcab454ed4a432d9472087a9d330598a38ffe56cd/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67" width="175px" />

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
