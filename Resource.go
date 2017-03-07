package dsc

type Resource interface {
	Check() bool
	Apply() error
}
