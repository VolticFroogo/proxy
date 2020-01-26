package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// Protocol constants.
	ProtocolHTTP   = "http"
	ProtocolHTTPS  = "https"

	// Anonymity constants.
	AnonymityAll         = "all"
	AnonymityTransparent = "transparent"
	AnonymityAnonymous   = "anonymous"
	AnonymityElite       = "elite"

	// Country constants.
	CountryAll = "all"

	// URL constants.
	urlBase      = "https://www.proxy-list.download/api/v1/get"
	urlProtocol  = "type"
	urlAnonymity = "anon"
	urlCountry   = "country"
)

var (
	Client http.Client
)

func Find(protocol, anonymity, country string) (proxies []url.URL, err error) {
	req, err := http.NewRequest(http.MethodGet, urlBase, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add(urlProtocol, protocol)

	if anonymity != AnonymityAll {
		query.Add(urlAnonymity, anonymity)
	}

	if country != CountryAll {
		query.Add(urlCountry, country)
	}

	req.URL.RawQuery = query.Encode()

	res, err := Client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("got status %s", res.Status)
		return
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	body := string(bodyBytes)
	proxyStrings := strings.Split(body, "\r\n")
	proxyStrings = proxyStrings[:len(proxyStrings)-1]

	for _, proxy := range proxyStrings {
		proxyURL, err := url.Parse(protocol + "://" + proxy)
		if err != nil {
			return nil, err
		}

		proxies = append(proxies, *proxyURL)
	}

	return
}
