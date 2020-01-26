package proxy

import (
	"errors"
	"testing"
)

var (
	errNoProxies = errors.New("no proxies returned")
)

func TestFind(t *testing.T) {
	proxies, err := Find(ProtocolHTTP, AnonymityAll, CountryAll)
	if err != nil {
		t.Error(err)
		return
	}

	if len(proxies) == 0 {
		t.Error(errNoProxies)
		return
	}

	t.Logf("returned %d proxies", len(proxies))
}
