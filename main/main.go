package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("jimxshaw"),
		os.Getenv(""),
		os.Getenv("test_db"))

	a.Run(":8080")
}
