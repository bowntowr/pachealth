package check

var registry []Checker

func Register(c Checker) {
	registry = append(registry, c)
}

func Registered() []Checker {
	return registry
}
