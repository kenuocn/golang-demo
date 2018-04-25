package main

import "fmt"

func main() {
	const (
		cpp = iota
		_
		python
		php
		javascript
	)

	fmt.Println(cpp,javascript,python,php)
}
