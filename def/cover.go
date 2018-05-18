package def

type CoverDetectorSimple interface {
	IsPointinRect(pt Point, rect Rect) bool
	IsPointinCircle(pt Point, center Point, size float32) bool
}

type CoverDetector interface {
	IsPointinTrangle(pt Point, vect [3]Point) bool
	IsPointinPoly(pt Point, vect []Point) bool
}
