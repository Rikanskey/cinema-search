package main

import "cinema-search/internal/runner"

const configDir = "./config/"

func main() {
	runner.Start(configDir)
}
