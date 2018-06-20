package main

import "fmt"

type TT struct {
	A string
	B int64
	C struct {
		Cd string
		Ce string
		Cf []float64
	}
	G bool
	H float32
}

func main() {
	var intI int
	intI = 10
	fmt.Println(intI)

	var pInt *int = &intI
	//pInt = new(int)
	fmt.Println(*pInt)

	var i *[]int
	i = new([]int)
	fmt.Println(i, len(*i), cap(*i))
	*i = append(*i, 1)
	fmt.Println(i, len(*i), cap(*i))


	var iArr *[3]int = &[3]int{}
	//iArr = new([3]int)
	fmt.Println(iArr, len(*iArr), cap(*iArr))
	iArr[0] = 1
	fmt.Println(iArr, len(*iArr), cap(*iArr))

	var istring *[]string
	istring = new([]string)
	fmt.Println(istring, len(*istring), cap(*istring))

	var t *TT = &TT{}
	fmt.Println(t)

	var tt = new(TT)
	fmt.Println(tt)

	var ttt *TT
	ttt = &TT{}
	fmt.Println(ttt)

	fmt.Println("#####")
	j := make([]int, 1)
	fmt.Println(j, len(j), cap(j))

	m := make(map[string]string)
	fmt.Println(m, len(m))

	mp := &map[string]string{}
	fmt.Println(mp, len(*mp))

	mn := new(map[string]string)
	fmt.Println(mn, len(*mn))

	c := make(chan bool)
	fmt.Println(c, len(c))

	cn := new(chan bool)
	fmt.Println(cn, len(*cn))

}

