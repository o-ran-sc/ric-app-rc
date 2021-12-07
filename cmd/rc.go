package main

import (
	"gerrit.o-ran-sc.org/r/scp/ric-app/control"
)

func main() {
	c := control.NewControl()
	c.Run()
}
