package dataaccess

import (
	"github.com/google/wire"

	"github.com/phamtiep/IDM-project/internal/dataaccess/cache"
	"github.com/phamtiep/IDM-project/internal/dataaccess/database"
	"github.com/phamtiep/IDM-project/internal/dataaccess/file"
	"github.com/phamtiep/IDM-project/internal/dataaccess/mq"
)

var WireSet = wire.NewSet(
	cache.WireSet,
	database.WireSet,
	mq.WireSet,
	file.WireSet,
)
