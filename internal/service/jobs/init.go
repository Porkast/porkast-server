package jobs

import "context"

func InitJobs(ctx context.Context) {
	UpdateUserSubKeywordJobs(ctx)
}
