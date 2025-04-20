# go-window

A tiny and generic Go package for producing sliding windows over slices, inspired by Rust’s `.windows()` method.

## ✨ Features

- Minimal API: just one function — `Slide`
- Generic: works with any slice type (`[]int`, `[]string`, etc.)
- Lazy and efficient: uses a goroutine and channel to stream windows

## 📦 Installation


```
go get github.com/johan-scriptdrift/go-window
```

## 🚀 Usage

```go
package main

import (
	"fmt"
	"github.com/johan-scriptdrift/go-window"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for win := range window.Slide(nums, 3) {
		fmt.Println(win)
	}
}
```

### Output

```
[1 2 3]
[2 3 4]
[3 4 5]
```