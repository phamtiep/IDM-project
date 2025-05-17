package jobs

import (
	"context"

	"github.com/phamtiep/IDM-project/internal/logic"
)

type UpdateDownloadingAndFailedDownloadTaskStatusToPending interface {
	Run(context.Context) error
}

type updateDownloadingAndFailedDownloadTaskStatusToPending struct {
	downloadTaskLogic logic.DownloadTask
}

func NewUpdateDownloadingAndFailedDownloadTaskStatusToPending(
	downloadTaskLogic logic.DownloadTask,
) UpdateDownloadingAndFailedDownloadTaskStatusToPending {
	return &updateDownloadingAndFailedDownloadTaskStatusToPending{
		downloadTaskLogic: downloadTaskLogic,
	}
}

func (u updateDownloadingAndFailedDownloadTaskStatusToPending) Run(ctx context.Context) error {
	return u.downloadTaskLogic.UpdateDownloadingAndFailedDownloadTaskStatusToPending(ctx)
}
