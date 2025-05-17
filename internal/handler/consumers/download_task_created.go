package consumers

import (
	"context"

	"go.uber.org/zap"

	"github.com/phamtiep/IDM-project/internal/dataaccess/mq/producer"
	"github.com/phamtiep/IDM-project/internal/logic"
	"github.com/phamtiep/IDM-project/internal/utils"
)

type DownloadTaskCreated interface {
	Handle(ctx context.Context, event producer.DownloadTaskCreated) error
}

type downloadTaskCreated struct {
	downloadTaskLogic logic.DownloadTask
	logger            *zap.Logger
}

func NewDownloadTaskCreated(
	downloadTaskLogic logic.DownloadTask,
	logger *zap.Logger,
) DownloadTaskCreated {
	return &downloadTaskCreated{
		downloadTaskLogic: downloadTaskLogic,
		logger:            logger,
	}
}

func (d downloadTaskCreated) Handle(ctx context.Context, event producer.DownloadTaskCreated) error {
	logger := utils.LoggerWithContext(ctx, d.logger).With(zap.Any("event", event))
	logger.Info("download task created event received")

	if err := d.downloadTaskLogic.ExecuteDownloadTask(ctx, event.ID); err != nil {
		logger.With(zap.Error(err)).Error("failed to handle download task created event")
		return err
	}

	return nil
}
