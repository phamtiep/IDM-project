package handler

import (
	"github.com/google/wire"

	"github.com/phamtiep/IDM-project/internal/handler/consumers"
	"github.com/phamtiep/IDM-project/internal/handler/grpc"
	"github.com/phamtiep/IDM-project/internal/handler/http"
	"github.com/phamtiep/IDM-project/internal/handler/jobs"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
	consumers.WireSet,
	jobs.WireSet,
)
