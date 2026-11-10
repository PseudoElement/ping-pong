1. **DONE** Создать структ Renderer, в нем на каждую итерацию цикла рендеринга вызывать renderer.Render(),
   и пройтись по всем объектам в пулах и для каждого вызвать rl.DrawText / rl.DrawRectangle / rl.DrawTexture и тд в зависимости от типа игрового объекта

2. Написать ImageLoader с предзагрузкой всех картинок приложения

3. struct NextFrameEventEmiter {} // calls Emit() on every frame
   struct GameEventEmiter {} // manually calls Emit() on game event

struct WasdKeysMovementStrategy(object IBaseObject, injector \*injector.Injector){}
struct ArrowKeysMovementStrategy(object IBaseObject, injector injector.Injector){}
struct TopBottomMovementStrategy(object IBaseObject, injector injector.Injector){}
struct BallMirroringMovementStrategy(object IBaseObject, injector injector.Injector){}

IMovementStrategy для игрока(-ов) и для бота(-ов) хранится в:
• GameSettings.LeftPlayerMovementStrategy()
• GameSettings.SetLeftPlayerMovementStrategy(IMovementStrategy)
• GameSettings.RightPlayerMovementStrategy()
• GameSettings.SetRightPlayerMovementStrategy(IMovementStrategy)

На нажатие Play with bot -> Dump вызывать:
• GameSettings.SetLeftPlayerMovementStrategy(WasdKeysMovementStrategy)
• GameSettings.SetRightPlayerMovementStrategy(TopBottomMovementStrategy)

На нажатие Play with bot -> Pro вызывать:
• GameSettings.SetLeftPlayerMovementStrategy(WasdKeysMovementStrategy)
• GameSettings.SetRightPlayerMovementStrategy(BallMirroringMovementStrategy)

- `СТРАТЕГИЮ ПЕРЕДАВАТЬ КАЖДОМУ ОБЪЕКТУ ЧЕРЕЗ КОНСТРУКТОР`
- `ПРИ СМЕНЕ СТРАТЕГИИ ВЫЗЫВАТЬ prevStrategy.Stop(); newStrategy.Setup()`

interface IMovementStrategy {
Setup()
Stop()
}
