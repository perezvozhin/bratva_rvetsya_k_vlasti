package main

import (
	"JOB_FINDER/caller"
	"JOB_FINDER/gem_service"
	loggersystem "JOB_FINDER/internals/logger"
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger := loggersystem.Init()
	c := caller.Init(logger, "./")
	ctx := context.Background()
	g := gem_service.Init(ctx, c, os.Getenv("GEMINI_API_KEY"))

	askNQuestions(ctx, g, 5)

}

func askNQuestions(ctx context.Context, g *gem_service.GeminiService, n int) {
	sc := bufio.NewScanner(os.Stdin)

	for range n {
		fmt.Print("> ")
		if !sc.Scan() {
			break // Ctrl+D or EOF
		}
		text := sc.Text()
		if text == "" {
			continue
		}

		ch, err := g.Ask(ctx, "test", text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for r := range ch {
			if r.Error != nil {
				fmt.Println(r.Error)
				break
			}
			fmt.Print(r.Chunk)
		}
		fmt.Println()
	}
}
