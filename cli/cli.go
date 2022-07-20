package cli

import (
	"flag"
	"fmt"
	"git_with_filecoin/util"
	"log"
	"os"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  apply -token TOKEN -cid CID - apply patch")
	fmt.Println("  commit -token TOKEN - Commit changes to w3s")
	fmt.Println("  get -token TOKEN -cid CID - Commit changes in LOCATION")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()
	applyCmd := flag.NewFlagSet("apply", flag.ExitOnError)
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	applyToken := applyCmd.String("token", "", "w3s api token")
	applyCid := applyCmd.String("cid", "", "CID of patch")

	commitToken := commitCmd.String("token", "", "w3s api token")

	getToken := getCmd.String("token", "", "w3s api token")
	getCid := getCmd.String("cid", "", "CID of your file")

	switch os.Args[1] {
	case "apply":
		err := applyCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "commit":
		err := commitCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "get":
		err := getCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	}

	if applyCmd.Parsed() {
		cli.apply(*applyToken, *applyCid)
	}

	if commitCmd.Parsed() {
		cli.commit(*commitToken)
	}

	if getCmd.Parsed() {
		util.GetFile(*getToken, *getCid)
	}

}
