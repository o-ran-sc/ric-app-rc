package main

import (
	"gerrit.o-ran-sc.org/r/ric-app/rc/control"
)

func main() {
	c := control.NewControl()
	c.Run()
}
