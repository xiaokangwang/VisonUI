package CoverSimple

import (
	"math"

	"github.com/xiaokangwang/VisonUI/def"
)

type CoverSimple struct{}

func NewCoverSimple() *CoverSimple {
	return &CoverSimple{}
}

func (cs CoverSimple) IsPointinRect(pt def.Point, rect def.Rect) bool {
	if pt.L > rect.L && pt.L < rect.W+rect.L && pt.T > rect.T && pt.T < rect.T+rect.H {
		return true
	}
	return false
}

func (cs CoverSimple) IsPointinCircle(pt def.Point, center def.Point, size float32) bool {
	dist := float32(math.Sqrt(math.Pow(float64(pt.L-center.L), 2) + math.Pow(float64(pt.T-center.T), 2)))
	if dist < size {
		return true
	}
	return false
}
