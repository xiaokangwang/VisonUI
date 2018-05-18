package def

type ScreenPointEvent interface {
	Point() Point
	Type() int
	Extra() string
}
