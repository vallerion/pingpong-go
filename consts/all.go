package consts

const (
	ScreenHeight             = 768
	ScreenWidth              = 1366
	GameFieldGap             = 20
	GameZoneLeft             = GameFieldGap
	GameZoneRight            = ScreenWidth - GameFieldGap
	GameZoneTop              = GameFieldGap + 50
	GameZoneBottom           = ScreenHeight - GameFieldGap
	GameZoneVerticalCenter   = (GameZoneTop + GameZoneBottom) / 2
	GameZoneHorizontalCenter = (GameZoneRight + GameZoneLeft) / 2
	PlayerWidth              = 16
	PlayerHeight             = 48
	BallMaxSpeed             = 30
	BallSize                 = 15
	GateHeight               = 120
	SideBorderHeight         = (GameZoneBottom - GameZoneTop - GateHeight) / 2
	GateTop                  = GameZoneTop + (GameZoneBottom-GameZoneTop-GateHeight)/2
	GateBottom               = GameZoneBottom - (GameZoneBottom-GameZoneTop-GateHeight)/2
)
