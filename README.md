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
	"os"
)

func main() {
	lioengine.SetProvider("Bing", "API TOKEN")
	project, err := lioengine.FindUpdatesFor("iphone 7")
	// do whatever with the project info
	...
}
```