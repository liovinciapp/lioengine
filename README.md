[![Wercker](https://img.shields.io/wercker/ci/wercker/docs.svg?maxAge=2592000)]()  [![AUR](https://img.shields.io/aur/license/yaourt.svg?maxAge=2592000)]() [![Waffle.io](https://img.shields.io/waffle/label/evancohen/smart-mirror/in%20progress.svg?maxAge=2592000)]()

# lioengine
liovinci's machine learning ai that finds updates for projects. This still on a very early stage of development so don't spect to have it working soon.

## Supported providers

1. Bing

## Installation
```
go get github.com/Shixzie/lioengine
```

## Usage
```go
package main

import (
	"github.com/Shixzie/lioengine"
	"bufio"
	"os"
	"fmt"
)

func main() {
	
	// Sets Bing as our news/updates provider.
	lioengine.AddUpdatesProvider("Bing", "API TOKEN")

	// Creates a reader so we can read from the console.
	// Instead of using the os.Stdin we could use a JSON request to get
	// the project name or get the value from a html form.
	reader := bufio.NewReader(os.Stdin)

	// We ask for the project name we want to get updates for.
    fmt.Print("Enter project name: ")

    // Wait until the user press Enter.
    project, _ := reader.ReadString('\n')

    // Search for updates.
	updates, _ := lioengine.FindUpdatesFor(project)

	// Iterates through updates
	for _, update := range updates {
		// do whatever with the update info
	}
	...
}
```

##Documentation
[Godoc:](http://godoc.org/github.com/Shixzie/lioengine)