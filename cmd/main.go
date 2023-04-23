package main

import "github.com/Hayao0819/go-distro"

func main(){
	d := distro.Get()
	println(d.String())
}
