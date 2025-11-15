package checker

import (
	"context"
	"net"
	"slices"
	"strings"
	"time"
)

func Check(ctx context.Context, alive chan Relays, relays []Relays, timeout time.Duration, country string) {
	countries := strings.Split(country, ",")

	dialer := &net.Dialer{
		Timeout: timeout,
	}

	for _, relay := range relays {
		select {
		case <-ctx.Done():
			return
		default:
			if countries[0] != "" && !slices.Contains(countries, relay.Country) {
				continue
			}
			addr := relay.OrAddresses[0]
			splittedAddr := strings.Split(addr, ":")
			address := splittedAddr[0] + ":" + splittedAddr[1]
			conn, err := dialer.Dial("tcp", address)
			if err == nil {
				alive <- relay
				conn.Close()
			}
		}
	}
}
