package def

type Drawing interface {
	DrawRect(area Rect, color Color)
	DrawTriangle(vect [3]Point, color Color)
	DrawText(area Rect, maxsize float32, color Color)
	DrawImage(area Rect, img UploadedImage)
	DrawCircle(center Point, size float32)
}

type UploadedImage interface {
	Unload()
}

type ImageUploader interface {
	UploadImage() //TODO
}
