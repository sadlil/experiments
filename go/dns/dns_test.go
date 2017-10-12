package main

import (
	"context"
	"fmt"
	"github.com/sanity-io/litter"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestDNS(t *testing.T) {
	//go runDNS()

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		fmt.Println("address original =", addr)
		if addr == "test.service:8080" {
			addr = "127.0.0.1:8080"
			fmt.Println("address modified =", addr)
		}
		return dialer.DialContext(ctx, network, addr)
	}

	resp, err := http.Get("http://test.service:8080")
	litter.Dump(err)
	litter.Dump(resp)
}

func TestDNSLookUP(t *testing.T) {
	x, y := net.LookupHost("http.appscode.test")
	litter.Dump(x)
	litter.Dump(y)
}
