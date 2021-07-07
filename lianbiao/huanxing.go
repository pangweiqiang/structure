package main

/**
寻找环形链表的起点
*/
func Hxstart1(p *point) (res *point) {
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
