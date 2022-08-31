package main

import (
	"flag"
	"time"

	"github.com/lokidb/cli"
	"github.com/lokidb/server/client"
)

var (
	addr    = flag.String("addr", "localhost:50051", "Server address")
	timeout = flag.Int("timeout", 5, "Timeout in seconds")
)

func main() {
	client := client.New(*addr, time.Duration(*timeout*int(time.Second)))
	defer client.Close()

	cli.ShellLoop(client)
}
