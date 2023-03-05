package collections

type Vector struct {
	data     []int
	capacity int
}

func (v *Vector) init(n int) {
	v.capacity = n
}

func (v *Vector) size() int {
	return len(v.data)
}

func (v *Vector) add(num int) {
	v.data = append(v.data, num)
}

func (v *Vector) clear() {
	v.data = v.data[:0]
}

func (v *Vector) contains(num int) bool {
	for _, v := range v.data {
		if v == num {
			return true
		}
	}
	return false
}
func (v *Vector) get(idx int) int {
	return v.data[idx]
}

func (v *Vector) indexOf(num int) int {
	for i, v := range v.data {
		if v == num {
			return i
		}
	}
	return -1
}
