package main

import (
	"fmt"
	"goproject/ch16/usepkg/custompkg" //모듈 내 패키지 custompkg

	"github.com/guptarohit/asciigraph"           //외부 저장소 패키지 asciigraph
	"github.com/tuckersGo/musthaveGo/ch16/expkg" //외부 저장소 패키지 expkg
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
