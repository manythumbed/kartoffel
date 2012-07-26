package main

import (
	"kartoffel"
)

func main()	{
	status := make(chan kartoffel.ServerStatus)
	kartoffel.Listen(0, status)
}
