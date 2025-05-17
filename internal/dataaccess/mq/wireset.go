package mq

import (
	"github.com/google/wire"

	"github.com/phamtiep/IDM-project/internal/dataaccess/mq/consumer"
	"github.com/phamtiep/IDM-project/internal/dataaccess/mq/producer"
)

var WireSet = wire.NewSet(
	consumer.WireSet,
	producer.WireSet,
)
