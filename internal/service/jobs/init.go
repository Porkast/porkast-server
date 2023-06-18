package jobs

import "context"

func InitJobs(ctx context.Context) {
	UpdateChannelTotalCountJob(ctx)
	UpdateItemTotalCountJob(ctx)
}
