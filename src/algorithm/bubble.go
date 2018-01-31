//冒泡法
package algorithm

import "fmt"

func BubbleSort(){
	ar := [10]int{9, 8, 6, 4, 2, 7, 1, 3, 0, 5}
	num := len(ar)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i] //相比某语言不需要临时交换变量
			}
		}
	}
	fmt.Print(ar)
}