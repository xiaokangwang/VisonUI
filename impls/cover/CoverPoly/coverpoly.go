package coverpoly

import (
	"github.com/JamesMilnerUK/pip-go"
	"github.com/xiaokangwang/VisonUI/def"
)

type CoverPoly struct{}

func NewCoverPoly() *CoverPoly {
	return &CoverPoly{}
}

func (cs CoverPoly) IsPointinTrangle(pt def.Point, vect [3]def.Point) bool {
	poly := pip.Polygon{
		Points: []pip.Point{
			pip.Point{X: float64(vect[0].L), Y: float64(vect[0].T)},
			pip.Point{X: float64(vect[1].L), Y: float64(vect[1].T)},
			pip.Point{X: float64(vect[2].L), Y: float64(vect[2].T)},
		},
	}
	return pip.PointInPolygon(pip.Point{X: float64(pt.L), Y: float64(pt.T)}, poly)
}
