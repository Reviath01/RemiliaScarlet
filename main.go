package main

import (
	"git.randomchars.net/Reviath/RemiliaScarlet/client"
	"git.randomchars.net/Reviath/RemiliaScarlet/web"
)

func main() {
	go web.Listen()
	client.Start()
}
