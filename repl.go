package main

import "setdb/internal"

func main() {
	cli := internal.InitCLI()
	cli.Run()
}
