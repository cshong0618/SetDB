package internal

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

type CLIInterface interface {
	parse(command string)
	Execute(command string)
	Run()
}

type CLI struct {
	db *DB
}

var COMMANDS = map[string]struct{}{
	"put": {},
	"find": {},
	"count": {},
}

func InitCLI() *CLI {
	cli := CLI{}
	cli.db = InitDB()

	return &cli
}

func (cli CLI) parse(command string) ([2]string, error){
	commands := [2]string{}
	buffer := strings.Builder{}

	// parse command
	currentIdx := 0
	for i, c := range command {
		currentIdx = i
		if c == ' '{
			break
		}

		ch := byte(c)
		buffer.WriteByte(ch)
	}

	cmd := buffer.String()
	if _, ok := COMMANDS[buffer.String()]; !ok {
		return commands, errors.New("Invalid command")
	}

	commands[0] = cmd
	commands[1] = command[currentIdx + 1:]

	return commands, nil
}

func (cli *CLI) Execute(command string) {
	tokens, err := cli.parse(command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch tokens[0] {
	case "put":
		if len(tokens[1]) == 0 {
			fmt.Printf("No params given. Expected put [value]\n")
			return
		}

		start := time.Now()
		cli.db.PutString(tokens[1])
		duration := time.Since(start).Nanoseconds()
		fmt.Printf("Put [%s] done in %dns.\n", tokens[1], duration)
		break
	case "find":
		if len(tokens[1]) == 0 {
			fmt.Printf("No params given. Expected find [value]\n")
			return
		}
		start := time.Now()
		result := cli.db.FindString(tokens[1])
		duration := time.Since(start).Nanoseconds()
		fmt.Printf("Find [%s]\n\tResult [%v]\n\tTime taken: %dns\n", tokens[1], result, duration)
		break
	case "count":
		count := cli.db.Items()
		fmt.Printf("Item count in DB: %d\n", count)
		break
	}
}

func (cli *CLI) Run() {
	run := true
	scanner := bufio.NewScanner(os.Stdin)
	for run {
		scanner.Scan()
		command := scanner.Text()

		if command == "exit" {
			run = false
		} else {
			cli.Execute(command)
		}
	}
}