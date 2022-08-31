package cli

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"

	"github.com/lokidb/server/client"
)

type empty struct{}

var keysList = make(map[string]empty, 1000)
var valuesList = make(map[string]empty, 1000)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("get", readline.PcItemDynamic(listKeys())),
	readline.PcItem("set",
		readline.PcItemDynamic(listKeys(),
			readline.PcItemDynamic(listValues()),
		),
	),
	readline.PcItem("del", readline.PcItemDynamic(listKeys())),
	readline.PcItem("keys"),
	readline.PcItem("flush"),

	readline.PcItem("bye"),
	readline.PcItem("help"),
)

func help() string {
	msg := "commands:\n"
	msg += completer.Tree("    ")
	return msg
}

func listKeys() readline.DynamicCompleteFunc {
	return func(s string) []string {
		keys := make([]string, len(keysList))

		i := 0
		for k := range keysList {
			keys[i] = k
			i++
		}
		return keys
	}
}

func listValues() readline.DynamicCompleteFunc {
	return func(s string) []string {
		values := make([]string, len(valuesList))

		i := 0
		for k := range valuesList {
			values[i] = k
			i++
		}
		return values
	}
}

func saveKey(key string) {
	keysList[key] = empty{}
}

func saveValue(value string) {
	valuesList[value] = empty{}
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func handleLine(reader *readline.Instance, line string, c *client.Client) bool {
	line = strings.TrimSpace(line)
	switch {

	case line == "bye" || line == "exit":
		return true
	case line == "":

	case line == "help":
		println(help())

	case line == "get":
		println("expecting 'get <key>'\n")
	case strings.HasPrefix(line, "get "):
		println(get(c, line[4:]) + "\n")

	case line == "set":
		println("expecting 'set <key> <value>'\n")
	case strings.HasPrefix(line, "set "):
		keyValue := strings.Split(line[4:], " ")
		if len(keyValue) != 2 {
			println("expecting 'set <key> <value>'\n")
		} else {
			println(set(c, keyValue[0], keyValue[1]))
		}

	case line == "del":
		println("expecting 'del <key>'\n")
	case strings.HasPrefix(line, "del "):
		println(del(c, line[4:]))

	case line == "keys":
		println(keys(c))
	case line == "flush":
		println(flush(c))
	default:
		fmt.Println("invalid command try 'help'")
	}

	return false
}

func ShellLoop(client *client.Client) {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          ">>> ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	l.CaptureExitSignal()

	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		stop := handleLine(l, line, client)

		if stop {
			break
		}
	}
}
