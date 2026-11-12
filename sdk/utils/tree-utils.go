package utils

import "github.com/pseudoelement/ping-pong/sdk/interfaces"

func ForEachInObject(rootObj interfaces.IBaseObject, callback func(child interfaces.IBaseObject)) {
	childrenQueue := append([]interfaces.IBaseObject{}, rootObj.Children()...)
	for len(childrenQueue) > 0 {
		firstChild := childrenQueue[0]
		childrenQueue = childrenQueue[1:]

		callback(firstChild)
		childrenQueue = append(childrenQueue, firstChild.Children()...)
	}
}

func ForEachInScreen(screen interfaces.IScreenBase, callback func(child interfaces.IBaseObject)) {
	childrenQueue := make([]interfaces.IBaseObject, 0)

	for _, child := range screen.ImageObjects() {
		childrenQueue = append(childrenQueue, child)
	}
	for _, child := range screen.RectObjects() {
		childrenQueue = append(childrenQueue, child)
	}
	for _, child := range screen.TextObjects() {
		childrenQueue = append(childrenQueue, child)
	}

	for len(childrenQueue) > 0 {
		firstChild := childrenQueue[0]
		childrenQueue = childrenQueue[1:]

		callback(firstChild)
		childrenQueue = append(childrenQueue, firstChild.Children()...)
	}
}

func GetChildByIdScreen(screen interfaces.IScreen, id string) (interfaces.IBaseObject, bool) {
	childrenQueue := make([]interfaces.IBaseObject, 0)

	for _, child := range screen.ImageObjects() {
		childrenQueue = append(childrenQueue, child)
	}
	for _, child := range screen.RectObjects() {
		childrenQueue = append(childrenQueue, child)
	}
	for _, child := range screen.TextObjects() {
		childrenQueue = append(childrenQueue, child)
	}

	for len(childrenQueue) > 0 {
		firstChild := childrenQueue[0]
		childrenQueue = childrenQueue[1:]

		if firstChild.Id() == id {
			return firstChild, true
		}

		for _, child := range firstChild.Children() {
			childrenQueue = append(childrenQueue, child)
		}
	}

	return nil, false
}
