package collections

type Array struct {
	data   []int
	capacity int
}

func (a *Array) init(n int){
	a.capacity = n;
}

func (a *Array) size() int {
	return len(a.data)
}

func (a * Array) add(num int) {
	a.data = append(a.data, num)
}

func (a *Array) clear() {
	a.data = a.data[:0]
}

func (a *Array) contains(num int) bool {
	for _, a := range a.data{
		if a == num {
			return true
		}
	}
	return false
} 

func (a *Array) get(ind int) int{
	return a.data[ind]

}

func (a *Array) indexOf(num int) int {
	for i, a := range a.data{
		if a == num {
			return i
		}
	}
	return -1
}