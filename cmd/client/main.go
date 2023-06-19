package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"

	completer "secretKeeper/internal/client/prompt/completer"
	executor "secretKeeper/internal/client/prompt/executor"
)

var (
	buildVersion string
	buildDate    string
)

func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\n", buildVersion, buildDate)

	p := prompt.New(
		executor.NewExecutor().Execute,
		completer.NewCompleter().Complete,
		prompt.OptionTitle("Gophkeeper"),
		prompt.OptionPrefix(">>>"),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
