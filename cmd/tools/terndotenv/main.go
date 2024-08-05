package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os/exec"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"internal/store/pgstore/migrations",
		"--config",
		"internal/store/pgstore/migrations/tern.conf",
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\nOutput: %s\n", err, out)
		panic(err)
	}
}
