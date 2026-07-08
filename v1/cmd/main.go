package main

import "githum.com/fehmicorp/v1/cmd/internal"

func main() {
	println(internal.Current.BuildType)
	println("Version: ", internal.Current.Version)
}
