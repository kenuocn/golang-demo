package main

import "fmt"

func main() {
	const (
		cpp = iota
		java
		python
		php
		javascript
	)

	fmt.Println(cpp,javascript,java,python,php)
}
