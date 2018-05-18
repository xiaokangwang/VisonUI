package def

type OSTextInput interface {
	AskForInput(area Rect, origtext string) string
}

type KeyboardInputEventDispatcher interface {
}

type MouseInputEventDispatcher interface {
}

type ScrollInputEventDispatcher interface {
}
