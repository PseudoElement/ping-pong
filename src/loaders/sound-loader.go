package loaders

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/utils"
)

type SoundLoader struct {
	sounds map[AppSound]rl.Sound
}

func NewSoundLoader() *SoundLoader {
	loader := &SoundLoader{sounds: make(map[AppSound]rl.Sound)}
	loader._load()
	return loader
}

func (sl *SoundLoader) GetSound(soundName AppSound) (rl.Sound, bool) {
	sound, ok := sl.sounds[soundName]
	return sound, ok
}

func (sl *SoundLoader) _load() {
	var wg sync.WaitGroup
	wg.Add(len(ALL_SOUNDS))

	for _, soundPath := range ALL_SOUNDS {
		go func() {
			defer wg.Done()
			path := utils.AppPath(soundPath)
			sound := rl.LoadSound(path.Normalize())
			sl.sounds[soundPath] = sound
		}()
	}

	wg.Wait()
}
