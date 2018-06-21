package main

import "os"
import "fmt"


func main() {
	path, ok := os.LookupEnv("DSPA_TTS_COMMAND")

	if !ok {
		panic("DSPA_TTS_COMMAND not set")
	}

	fmt.Printf("Path is %v\n", path)
}
