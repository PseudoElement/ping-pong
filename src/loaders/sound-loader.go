package loaders

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
	sdkconstants "github.com/pseudoelement/ping-pong/sdk/constants"
	"github.com/pseudoelement/ping-pong/sdk/utils"
	"github.com/pseudoelement/ping-pong/src/constants"
)

type SoundLoader struct {
	sounds map[sdkconstants.AppSoundName]rl.Sound
}

func NewSoundLoader() *SoundLoader {
	loader := &SoundLoader{sounds: make(map[sdkconstants.AppSoundName]rl.Sound)}
	loader._load()
	return loader
}

func (sl *SoundLoader) GetSound(soundName sdkconstants.AppSoundName) rl.Sound {
	return sl.sounds[soundName]
}

func (sl *SoundLoader) _load() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(constants.APP_SOUNDS))

	for _, soundPath := range constants.APP_SOUNDS {
		go func() {
			defer wg.Done()
			path := utils.AppPath(soundPath)
			sound := rl.LoadSound(path.Normalize())

			mu.Lock()
			sl.sounds[soundPath] = sound
			mu.Unlock()
		}()
	}

	wg.Wait()
}
