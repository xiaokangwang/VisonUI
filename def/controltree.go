package def

type Control interface {
	OnInputEvent()  //TODO
	OnOutputEvent() //TODO
	OnDraw(area Rect)
	TerritoryTest(pt Point)
}

type ControlTree interface {
	EnumChildElement() []ControlTree
	PutChildElement(ct ControlTree)
	RemoveChildElement(ct ControlTree)
	GetRootElement() Control
	SetRootLayoutPosition(val Rect)
	GetRootLayoutPosition() Rect
}
