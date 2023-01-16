/*
Copyright © 2023 haormj <haormj@gmail.com>
*/
package main

import (
	"log"

	"github.com/haormj/code-generator/cmd"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	cmd.Execute()
}
