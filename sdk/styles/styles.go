package styles

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/models"
)

type FlexDirection int

const (
	FLEX_DIR_ROW FlexDirection = iota
	FLEX_DIR_COLUMN
)

/**
 * Use only between rl.BeginDrawing() and rl.EndDrawing() calls
 */
func Flex(objects []ITextObject, flexDirection FlexDirection, startX, startY int32, gap int32) {
	switch flexDirection {
	case FLEX_DIR_ROW:
		for idx := range len(objects) {
			object := objects[idx]

			var posX int32
			if idx == 0 {
				posX = startX
			} else {
				prevObject := objects[idx-1]
				startX = prevObject.Coords().X + prevObject.Width() + gap
			}

			rl.DrawText(object.Value(), posX, startY, object.FontSize(), object.Color())
			object.SetCoords(models.Coords{X: posX, Y: startY})
		}
	case FLEX_DIR_COLUMN:
		var largestObject ITextObject = &EmptyTextObject{}
		for _, object := range objects {
			if object.Width() > largestObject.Width() {
				largestObject = object
			}
		}

		for idx := range len(objects) {
			object := objects[idx]

			var posY int32
			if idx == 0 {
				posY = startY
			} else {
				prevObject := objects[idx-1]
				posY = prevObject.Coords().Y + prevObject.Height() + gap
			}

			var posX int32
			if object == largestObject {
				posX = startX
			} else {
				posX = startX + int32((largestObject.Width()-object.Width())/2)
			}

			rl.DrawText(object.Value(), posX, posY, object.FontSize(), object.Color())
			object.SetCoords(models.Coords{X: posX, Y: posY})
		}
	default:
		fmt.Println("[Flex] unknown flexDirection: ", flexDirection)
	}
}
