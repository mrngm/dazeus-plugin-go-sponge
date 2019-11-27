package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/dazeus/dazeus-go"
)

var myCommand string

func main() {
	connStr := "unix:/tmp/dazeus.sock"
	if len(os.Args) > 1 {
		connStr = os.Args[1]
	}
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Paniek! %v\n", p)
			debug.PrintStack()
		}
	}()

	dz, err := dazeus.ConnectWithLoggingToStdErr(connStr)
	if err != nil {
		panic(err)
	}

	if _, hlerr := dz.HighlightCharacter(); hlerr != nil {
		panic(hlerr)
	}

	_, err = dz.SubscribeCommand("spons", dazeus.NewUniversalScope(), func(ev dazeus.Event) {
		ev.Reply(sponsify(ev.Params[0]), true)
	})

	listenerr := dz.Listen()
	panic(listenerr)
}

func sponsify(s string) string {
	ret := ""
	even := true
	for _, c := range s {
		if even {
			ret = ret + strings.ToLower(string(c))
		} else {
			ret = ret + strings.ToUpper(string(c))
		}
		even = !even
	}
	return ret
}
