//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire
package wiring

import (
	"github.com/google/wire"

	"github.com/phamtiep/IDM-project/internal/app"
	"github.com/phamtiep/IDM-project/internal/configs"
	"github.com/phamtiep/IDM-project/internal/dataaccess"
	"github.com/phamtiep/IDM-project/internal/handler"
	"github.com/phamtiep/IDM-project/internal/logic"
	"github.com/phamtiep/IDM-project/internal/utils"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	utils.WireSet,
	dataaccess.WireSet,
	logic.WireSet,
	handler.WireSet,
	app.WireSet,
)

func InitializeStandaloneServer(configFilePath configs.ConfigFilePath) (*app.StandaloneServer, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}
