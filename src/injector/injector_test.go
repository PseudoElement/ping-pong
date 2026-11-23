package injector

import (
	"testing"

	"github.com/pseudoelement/ping-pong/sdk/interfaces"
	"github.com/pseudoelement/ping-pong/src/events"
	appinterfaces "github.com/pseudoelement/ping-pong/src/interfaces"
	"github.com/pseudoelement/ping-pong/src/loaders"
	"github.com/pseudoelement/ping-pong/src/settings"
)

type testCase[T appinterfaces.IInjector] struct {
	name  string
	gotFn func(res T) bool
	want  bool
}

type TestRenderer struct{}

var _ appinterfaces.IRenderer = (*TestRenderer)(nil)

func (r *TestRenderer) Render()                           {}
func (r *TestRenderer) CurrentScreen() interfaces.IScreen { return nil }
func (r *TestRenderer) SetScreen(interfaces.IScreen)      {}

func TestInjector(t *testing.T) {
	inj := NewInjector()

	inj.SetEventManager(events.NewEventManager())
	inj.SetFontLoader(loaders.NewFontLoader())
	inj.SetImageLoader(loaders.NewImageLoader(1920, 1080))
	inj.SetSoundLoader(loaders.NewSoundLoader())
	inj.SetGameSettings(settings.NewGameSettings())
	inj.SetRenderer(&TestRenderer{})

	cases := []testCase[*Injector]{
		{
			name:  "NotNilEventManager",
			gotFn: func(res *Injector) bool { return res.EventManager() != nil },
			want:  true,
		},
		{
			name:  "NotNilFontLoader",
			gotFn: func(res *Injector) bool { return res.FontLoader() != nil },
			want:  true,
		},
		{
			name:  "NotNilImageLoader",
			gotFn: func(res *Injector) bool { return res.ImageLoader() != nil },
			want:  true,
		},
		{
			name:  "NotNilSoundLoader",
			gotFn: func(res *Injector) bool { return res.SoundLoader() != nil },
			want:  true,
		},
		{
			name:  "NotNilGameSettings",
			gotFn: func(res *Injector) bool { return res.GameSettings() != nil },
			want:  true,
		},
		{
			name:  "NotNilRenderer",
			gotFn: func(res *Injector) bool { return res.Renderer() != nil },
			want:  true,
		},
	}

	for _, test := range cases {
		res := test.gotFn(inj)
		if res != test.want {
			t.Errorf("Test %s Failed. Got - %v, want - %v \n", test.name, res, test.want)
		}
	}
}
