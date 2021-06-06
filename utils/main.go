package main

import "fmt"

func main()  {
	a := []interface{}{
		1,2,3,
	}
	b := []interface{}{
		6,6,8,
	}
	c:= []interface{}{2,3,9666}
 fmt.Printf("%+v",array_merge(a,b,c))
}

func array_merge(slice ...[]interface{}) (res []interface{}){
	if (len(slice) == 0){
		return
	}
	if(len(slice) == 1){
		res = slice[0];
		return
	}
	tmp1 := slice[0]
	tmp2 := array_merge(slice[1:]...)
	res = make([]interface{},len(tmp1) + len(tmp2))
	copy(res,tmp1)
	copy(res[len(tmp1):],tmp2)
	return
}
