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
	"os"
)

func main() {
	lioengine.SetProvider("Bing", "API TOKEN")
	project, err := lioengine.FindUpdatesFor("iphone 7")
	// do whatever with the project info
	...
}
```