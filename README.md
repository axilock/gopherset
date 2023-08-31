# goperset


Golang Apache Superset API Package

This project structure is based on: [gocloak](https://github.com/Nerzal/gocloak)

For Questions either raise an issue, or come to the [gopher-slack](https://invite.slack.golangbridge.org/) into the channel [#goperset](https://gophers.slack.com/app_redirect?channel=goperset)

__SOON__: Benchmarks can be found [here](https://justmike1.github.io/goperset/dev/bench/)

## Contribution

(WIP) <https://github.com/justmike1/goperset/wiki/Contribute>

## Changelog

For release notes please consult the specific releases [here](https://github.com/justmike1/goperset/releases)


## Usage

### Installation

```shell
go get github.com/justmike1/goperset@v0.1.0
```

### Importing

```go
 import "github.com/justmike1/goperset@v0.1.0"
```

### Connecting to Superset Client

```go
func main() {
    client := goperset.NewClient("https://superset.domain.net/") // Client also has a context
    authToken, csrfToken, err := goperset.GetAccessTokens(client, "admin", "admin")
    if err != nil {
        t.Errorf("Error getting access tokens: %v", err)
        return
    }
}
```
 