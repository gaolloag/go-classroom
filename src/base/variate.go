package base

import "fmt"

func D1(){
	x := 1;
	p := &x; //& 取 x 内存地址即：指针
	fmt.Println(p)
	fmt.Println(*p) // * 获取内存地址的值
	*p = 2
	fmt.Println(x)

	var y int;

	fmt.Println(&y)

	fmt.Println(d1() == p)
}

func d1() *int {
	v := 2;
	fmt.Println(&v)
	return &v;
}
