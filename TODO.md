1. **DONE** Создать структ Renderer, в нем на каждую итерацию цикла рендеринга вызывать renderer.Render(),
   и пройтись по всем объектам в пулах и для каждого вызвать rl.DrawText / rl.DrawRectangle / rl.DrawTexture и тд в зависимости от типа игрового объекта

2. **DONE** Написать ImageLoader с предзагрузкой всех картинок приложения

3. **DONE** struct NextFrameEventEmiter {} // calls Emit() on every frame
   struct GameEventEmiter {} // manually calls Emit() on game event

**DONE** struct WasdKeysMovementStrategy(object IBaseObject, injector \*injector.Injector){}
**DONE** struct ArrowKeysMovementStrategy(object IBaseObject, injector injector.Injector){}
**DONE** struct TopBottomMovementStrategy(object IBaseObject, injector injector.Injector){}
**DONE** struct BallMirroringMovementStrategy(object IBaseObject, injector injector.Injector){}

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

4. **DONE** Добавить новое окно после Play with bot с выбором сложности бота Easy/Professional
5. ОТРЕФАКТОРИТЬ папку sdk, оставить там только интерфейсы объектов, IScreen, ивенты
6. **DONE** Добавить рандомность velocity для движения мяча по Y оси на каждую отбивку
7. **DONE** Изменить шрифт
8. **DONE** Добавить звук рикошета мяча, добавить для этого GameStorage.KickBallEventEmitter()
9. **DONE** Добавить белый враппер для счета
10. **DONE** Пофиксить коллизию мяча и ракетки
