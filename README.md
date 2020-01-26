# Proxy
[![CircleCI](https://circleci.com/gh/VolticFroogo/proxy.svg?style=svg)](https://circleci.com/gh/VolticFroogo/proxy)

A simple way to use the [proxy-list.download API](https://www.proxy-list.download/api/v1) in Go.
The API (and consequently this client) can find HTTP and HTTPS proxies filtered by anonymity and country.

## Usage

```go
func Find(protocol, anonymity, country string) (proxies []url.URL, err error)
```

The protocol can be either of the protocols below.

```go
ProtocolHTTP   = "http"
ProtocolHTTPS  = "https"
```

The anonymity can be one of the anonymity settings below, with "all" not filtering based on anonymity.

```go
AnonymityAll         = "all"
AnonymityTransparent = "transparent"
AnonymityAnonymous   = "anonymous"
AnonymityElite       = "elite"
```

The country code can either be "all" or an [ISO 3166 Alpha-2 code](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes).

```go
CountryAll = "all"
```
