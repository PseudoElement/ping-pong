package utils

import "github.com/pseudoelement/ping-pong/sdk/interfaces"

func ForEach(rootObj interfaces.IBaseObject, callback func(child interfaces.IBaseObject)) {
	childrenQueue := append([]interfaces.IBaseObject{}, rootObj.Children()...)
	for len(childrenQueue) > 0 {
		firstChild := childrenQueue[0]
		childrenQueue = childrenQueue[1:]

		callback(firstChild)
		childrenQueue = append(childrenQueue, firstChild.Children()...)
	}
}
