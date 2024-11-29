package routes

import "go.uber.org/fx"

type Route interface {
	SetUp()
}
type RoutesGateWay []Route

func (rgw *RoutesGateWay) SetUp() {
	for _, r := range *rgw {
		r.SetUp()
	}
}
func NewRoutesGateWay(ur *UserRouting, ar *AuthRouting) *RoutesGateWay {
	return &RoutesGateWay{
		ur,
		ar,
	}
}

var RoutesGateWayModule = fx.Options(fx.Provide(NewUserRouting), fx.Provide(NewAuthRouting), fx.Provide(NewRoutesGateWay))
