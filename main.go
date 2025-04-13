package main

import "fmt"

// NOTE:
// 1. If we import a package which is not included in the go.mod
// go automatically adds it to the go.mod file (if go test is run)
// otherwise, we need to add it via go mod tidy
// 2. go list -m all = list all modules
// 3. go.sum = contains checksums of all dependencies, lists all dependencies/sub-dependencies
// 4. go get golang.org/x/text = download the package
func main() {
	cmd.Init()
	fmt.Println("Hello, World!")
}