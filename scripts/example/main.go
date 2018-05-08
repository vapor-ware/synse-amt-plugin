package main

import (
	"fmt"
	"os/exec"
)

func main() {

	fmt.Println("======================")
	fmt.Println("Running Script #1 (OK)")
	fmt.Println("======================")

	cmd := exec.Command("python", "test.py", "foobar") // nolint: gas
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Output (Raw)    --> %v\n", out)
	fmt.Printf("Output (String) --> %s\n", out)

	fmt.Println("=======================")
	fmt.Println("Running Script #2 (ERR)")
	fmt.Println("=======================")

	cmd = exec.Command("python", "err.py") // nolint: gas
	out, err = cmd.Output()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Output (Raw)    --> %v\n", out)
	fmt.Printf("Output (String) --> %s\n", out)
}
