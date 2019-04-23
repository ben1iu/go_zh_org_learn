package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {

	
	a := make([][]unit8,dy)
	for x:=range a{
		b:=make([]unit8,dx)
		for y:=range b{
			b[y] =unit8(x*y-1)
		}
		a[x]=b
	}
	return a
}

func main() {
	pic.Show(Pic)
}

