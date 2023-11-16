package jobs

import (
	"context"
	"fmt"
	"porkast-server/internal/consts"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/celery"
	"porkast-server/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

func UpdateUserSubKeywordJobs(ctx context.Context) {
	var (
		err error
	)

	_, err = gcron.Add(ctx, "0 0 */3 * * *", func(ctx context.Context) {
		var (
			randomSleepTime time.Duration
		)
		randomSleepTime = celery.GetRandomStartTime()
		g.Log().Line().Info(ctx, "start channel update jobs, sleep random time : ", randomSleepTime)
		time.Sleep(randomSleepTime)
		if !celery.IsJobStarted(ctx, consts.USER_SUB_KEYWORD_UPDATE) {
			celery.JobIsStarted(ctx, consts.USER_SUB_KEYWORD_UPDATE)
			assignUserSubKeywordUpdateJob(ctx)
		}

	}, consts.USER_SUB_KEYWORD_UPDATE)

	if err != nil {
		g.Log().Line().Error(ctx, "The UserSubKeyword job start failed : ", err)
	}
}

func assignUserSubKeywordUpdateJob(ctx context.Context) {
	var (
		err                error
		userSubKeywordList []entity.UserSubscription
	)

	userSubKeywordList, err = dao.GetAllKindSubKeywordList(ctx, 0, 0)
	if err != nil {
		g.Log().Line().Errorf(ctx, "Get user sub keyword list failed %s", err)
		return
	}

	for _, userSubKeywordItem := range userSubKeywordList {
		var (
			keyword       string
			country       string
			excludeFeedId string
			source        string
		)

		keyword = userSubKeywordItem.Keyword
		country = userSubKeywordItem.Country
		excludeFeedId = userSubKeywordItem.ExcludeFeedId
		source = userSubKeywordItem.Source
		_, err = celery.GetClient().Delay(consts.USER_SUB_KEYWORD_UPDATE, keyword, country, excludeFeedId, source)
		if err != nil {
			g.Log().Line().Error(ctx, fmt.Sprintf("Assign USER_SUB_KEYWORD_UPDATE failed : %s\n", err))
		}
	}

}
