package utils

var Arrays = &arrays{}

type arrays struct {
}

func (a *arrays) ContainsStr(arr []string, ele string) bool {
	for _, item := range arr {
		if item == ele {
			return true
		}

	}
	return false
}

func (a *arrays) ContainsInt(arr []int, ele int) bool {
	for _, item := range arr {
		if item == ele {
			return true
		}

	}
	return false
}
