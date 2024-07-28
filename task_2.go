package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func SolutionTask2() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var inputDataCount, goodsCount int
	var comm, goodsPrice, goodsSum float64
	fmt.Fscan(in, &inputDataCount)
	for i := 0; i < inputDataCount; i++ {
		goodsSum = 0.00
		fmt.Fscan(in, &goodsCount, &comm)
		for i := 0; i < goodsCount; i++ {
			fmt.Fscan(in, &goodsPrice)
			a := goodsPrice + (goodsPrice * (0.01 * comm))
			_, c := math.Modf(a)
			goodsSum += math.Round(c*100) / 100
		}
		fmt.Fprintf(out, "%.2f\n", float64(math.Round(goodsSum*100)/100))
	}
}
