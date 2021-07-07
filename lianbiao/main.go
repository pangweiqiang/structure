package main

import "fmt"

func main()  {
	res := Hxstart(top)
	if res == nil{
		fmt.Println("null")
	}else{
		fmt.Printf("%+v",res)
	}

}
func Hxstart(p *point) (res *point) {
	//校验 ：略

	//快慢指针判断是否环形链表
	normal := p.adress
	fast := p.adress.adress
	var meet *point
	for {
		if normal == fast {
			meet = normal
			break
		}
		if fast.adress == nil || fast.adress.adress == nil {
			meet = nil
			break
		}
		normal = normal.adress
		fast = fast.adress.adress
	}
	if meet == nil {
		return nil
	}
	s := p.adress
	var ans *point
	for {
		if s == meet {
			ans = s
			break
		}
		s = s.adress
		meet = meet.adress
	}
	return ans
}
type point struct {
	val int
	adress *point
}
var top *point
func init()  {
	var tmp *point

	for i := 1 ; i <= 10 ; i++{
		k := new(point)
		k.val = i
		k.adress = tmp
		tmp = k
	}
	top = tmp
}

func print()  {
	tmp := top
	for {
		fmt.Printf("%p ; %+v \n",tmp,tmp)
		if(tmp.adress == nil){
			break
		}
		tmp = tmp.adress

	}
}
//a->b->c->d
//null<-a<-b<-c<-d

func fanzhuan() {
	var tmp *point
	for{
		v := top.adress
		top.adress = tmp
		tmp = top
		top = v
		if(top == nil){
			top = tmp
			break
		}
	}
	//return top
}

//
func huanxing(p *point) bool  {
	if (p == nil || p.adress == nil || p.adress.adress == nil){
		return false
	}
	slow := p.adress
	fast := p.adress.adress
	for {
		if(slow == fast){
			return true;
		}
		if(fast.adress == nil || fast.adress.adress == nil){
			return false;
		}
		slow = slow.adress
		fast = fast.adress.adress
	}
	return false

}
