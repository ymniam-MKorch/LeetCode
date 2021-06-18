type OrderedStream struct {
	currentIndex int
	stringList   []string
}

func Constructor(n int) OrderedStream {
	stringList := make([]string, n+1)
	res := OrderedStream{
		currentIndex: 1,
		stringList:   stringList,
	}
	return res
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.stringList[id] = value
	res := []string{}
	if this.currentIndex == id || this.stringList[this.currentIndex] != "" {
		res = append(res, this.stringList[this.currentIndex])
		for i := id + 1; i < len(this.stringList); i++ {
			if this.stringList[i] != "" {
				res = append(res, this.stringList[i])
			} else {
				this.currentIndex = i
				return res
			}
		}
	}
	if len(res) > 0 {
		return res
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */