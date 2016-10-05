[![Visual Studio Team services](https://img.shields.io/vso/build/larsbrinkhoff/953a34b9-5966-4923-a48a-c41874cfb5f5/1.svg?maxAge=2592000)]()

# lioengine
liovinci's machine learning ai that finds updates for projects

## Installation
```
go get github.com/Shixzie/lioengine
```

## Usage
```
package main

import (
	"github.com/Shixzie/lioengine"
	"bufio"
	"os"
	"fmt"
)

func main() {
	
	// Sets Bing as our news/updates provider.
	lioengine.SetProvider("Bing", "API TOKEN")

	// Creates a reader so we can read from the console
	// Instead of using the os.Stdin we could use a JSON request to get
	// the project name or get the value from a html form.
	reader := bufio.NewReader(os.Stdin)

	// We ask for the project name we want to get updates for.
    fmt.Print("Enter project name: ")

    // Wait until the users press Enter.
    text, _ := reader.ReadString('\n')

    // Search for updates.
	updates, _ := lioengine.FindUpdatesFor(text)

	for _, update := range updates {
		// do whatever with the update info
	}
	...
}
```