package interface11

type people struct {
}

func (p people) hello(name string) string {
	return "hello " + name
}

func GetI() i {
	return &people{}
}
