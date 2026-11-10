package loaders

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/sdk/utils"
)

type SoundLoader struct {
	sounds map[AppSoundName]rl.Sound
}

func NewSoundLoader() *SoundLoader {
	loader := &SoundLoader{sounds: make(map[AppSoundName]rl.Sound)}
	loader._load()
	return loader
}

func (sl *SoundLoader) GetSound(soundName AppSoundName) rl.Sound {
	return sl.sounds[soundName]
}

func (sl *SoundLoader) _load() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(app_sounds))

	for _, soundPath := range app_sounds {
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
