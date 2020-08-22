package main

import "fmt"

func main() {
	hen1 := 1.0
	hen2 := 2.3
	hen3 := 3.5
	hen4 := 32.4
	hen5 := 54.3
	hen6 := 56.3

	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)
	fmt.Printf("totalweight=%v,avgweight=%v\n", totalWeight, avgWeight)

	var hens [6]float64
	hens[0] = 1.0
	hens[1] = 2.3
	hens[2] = 3.5
	hens[3] = 32.4
	hens[4] = 54.3
	hens[5] = 56.3

	totalWeight = 0.0

	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	avgWeight = fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("totalweight=%v,avgweight=%v\n", totalWeight, avgWeight)
}
