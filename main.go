package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	proxy      string
	target     string
	subcommand string
	verbose    bool
)

func init() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'clone' or 'wget' subcommands")
		os.Exit(1)
	}

}
func main() {
	subcommand = os.Args[1]
	switch subcommand {
	case "clone":
		clone := flag.NewFlagSet("clone", flag.ExitOnError)
		clone.StringVar(&proxy, "proxy", "https://ghproxy.com/", "github proxy")
		clone.StringVar(&target, "target", "", "github repo or file")
		clone.BoolVar(&verbose, "verbose", false, "show debug log")
		clone.Parse(os.Args[2:])
		if verbose {
			fmt.Printf("runing %s with proxy %s and target %s\n", subcommand, proxy, target)
		}
		newTarget := Join(proxy, target)
		origin := target
		Run([]func() error{Clone(newTarget), SetOrigin(origin)})
	default:
		if len(os.Args) < 2 {
			fmt.Println("expected 'clone' or 'wget' subcommands")
			os.Exit(1)
		}
	}
}

func Run(fs []func() error) {
	for _, f := range fs {
		if err := f(); err != nil {
			fmt.Printf("%s failed, %q", subcommand, err)
			os.Exit(1)
		}
	}

}
