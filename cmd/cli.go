package main

import (
	"github.com/lokidb/cli"
	client "github.com/lokidb/server/clients/go"
)

func main() {
	client := client.New()
	defer client.Close()

	cli.ShellLoop(client)
}
