//选择排序
package sort

import "fmt"

func Selection(){
	ar := [...]int {8,3,5,4,3,2,7}
	fmt.Println(ar)

	len := len(ar)
	for i:=0; i < len -1; i++ {
		k := i

		for j := k + 1; j < len; j++{
			if ar[j] < ar[k]{
				k = j
			}
		}

		if i != k{
			temp := ar[i]
			ar[i] = ar[k]
			ar[k] = temp
		}
		fmt.Println(ar)
		fmt.Println("---------")
	}

	fmt.Println(ar)
}
