package sort

import "fmt"

func Insertion()  {
	ar := [...]int {4,2,5,3,6,11,3}
	fmt.Println(ar)

	len := len(ar)

	for i := 1; i<len; i++ {
		get := ar[i]

		j := i - 1
		for ;j >= 0 && ar[j] > get; j--{
			ar[j + 1] = ar[j]
		}

		ar[j + 1] = get
	}

	fmt.Println(ar)

}
