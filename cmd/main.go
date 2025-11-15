package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/leaanthony/clir"
	"github.com/tamper000/tor-scanner/internal/checker"

	"github.com/logrusorgru/aurora/v4"
)

var (
	count   = 5
	threads = 50
	maxPing = 250

	output  = ""
	input   = "https://raw.githubusercontent.com/ValdikSS/tor-onionoo-mirror/refs/heads/master/details-running-relays-fingerprint-address-only.json"
	country = ""

	list = []checker.Relays{}
)

func main() {
	cli := clir.NewCli("Tor-scanner", "Easily Find Working Bridges for Tor", "v0.0.1")

	cli.IntFlag("count", "How many relays are needed", &count)
	cli.IntFlag("threads", "How many threads are needed", &threads)
	cli.IntFlag("ping", "Maximum Delay to Relay in ms", &maxPing)
	cli.StringFlag("country", "Including countries. EX: ru,es,us", &country)
	cli.StringFlag("output", "File to save", &output)
	cli.StringFlag("input", "You can set where to get bridges from (including the path to the file)", &input)

	cli.Action(start)

	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}

func start() error {
	fmt.Println(aurora.Cyan("Parsing relays..."))
	relays, err := checker.GetRelays(input)
	if err != nil {
		return err
	}
	chunkSize := len(relays) / threads
	fmt.Println(aurora.Green(fmt.Sprintf("Got %d relays", len(relays))))

	ctx, cancel := context.WithCancel(context.Background())

	alive := make(chan checker.Relays)
	for i := 0; i < threads; i++ {
		start := i * chunkSize
		end := start + chunkSize
		chunk := relays[start:end]

		go checker.Check(ctx, alive, chunk, time.Duration(maxPing)*time.Millisecond, country)
	}

	for len(list) < count {
		addr := <-alive
		list = append(list, addr)
		fmt.Println(aurora.Green("Found"), aurora.Cyan(addr.OrAddresses[0]))
	}
	cancel()

	fmt.Println("\n\nYour relays:")

	var text string
	for _, relay := range list {
		fmt.Println("Bridge", aurora.Cyan(relay.OrAddresses[0]), relay.Fingerprint)
		text += fmt.Sprintf("Bridge %s %s\n", relay.OrAddresses[0], relay.Fingerprint)
	}
	fmt.Println("UseBridges 1")
	text += "UseBridges 1"

	if output != "" {
		fmt.Println()
		if err := os.WriteFile(output, []byte(text), 0644); err != nil {
			return err
		}

		fmt.Println(aurora.Green("Successfully written to a file"), aurora.Underline(output))
	}

	return nil
}
