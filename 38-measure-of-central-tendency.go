package main

import "fmt"

func meanAlgorithm(datas ...int) []float32 {
	result := 0
	for i := 0; i < len(datas); i++ {
		result += datas[i]
	}
	mean:=[]float32{}
	mean = append(mean,float32(result )/ float32(len(datas)))
	return mean
}
 
func sortfunction(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left := []int{}
	mid := []int{}
	right := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] == pivot {
			mid = append(mid, arr[i])
		} else {
			if arr[i] > pivot {
				right = append(right, arr[i])
			}
		}
	}

	left = sortfunction(left)
	right = sortfunction(right)
	sort := append(left, mid...)
	sort = append(sort, right...)
	return sort

}

func medianAlgorithm(x ...int) interface{} {
	arr := sortfunction(x)
	if len(arr)%2 == 0 {
		//  jumlah array genap
		median := []float32{}
		leftmid := []int{len(arr)/2 - 1}
		rightmid := []int{len(arr)/2 + 1}
		slicearr := arr[leftmid[0]:rightmid[0]]
		jumlahmedian := 0
		for i := 0; i < len(slicearr); i++ {
			jumlahmedian += slicearr[i]
		}
		median = append(median, float32(jumlahmedian)/2)
		return median
	} else {
		// jumlah array ganjil
		median := []int{}
		// berisi penjumlahan array ketika ada dua array [10]+[10] = [20]
		array := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] == arr[len(arr)/2] {
				array += arr[i]
				if len(median) < 1 {
					median = append(median, arr[i])
				} else {
					median = []int{array / 2}
				}

			}
		}
		return median
	}

}

func main() {
	df := []int{
		50, 11, 50, 50,
	}
	fmt.Println("mean dari", df, "adalah :", meanAlgorithm(df...))

	df2 := []int{
		10, 21, 30, 30, 40,
	}

	fmt.Println("median dari", df2, " adalah :", medianAlgorithm(df2...))
}
