package DAG_Rider

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  start - Start DAG Rider main process")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run(){
	cli.validateArgs()

	startCmd := flag.NewFlagSet("start", flag.ExitOnError)

	switch os.Args[1] {
	case "start":
		err := startCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if startCmd.Parsed() {
		cli.startNode()
	}
}
