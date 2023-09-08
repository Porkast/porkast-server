package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/celery"
	"guoshao-fm-web/internal/service/internal/dao"
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
		userSubKeywordList []entity.UserSubKeyword
	)

	userSubKeywordList, err = dao.GetAllKindSubKeywordList(ctx, 0, 0)
	if err != nil {
		g.Log().Line().Errorf(ctx, "Get user sub keyword list failed %s", err)
		return
	}

	for _, userSubKeywordItem := range userSubKeywordList {
		var (
			keyword     string
			lang        string
			orderBydate int
		)

		keyword = userSubKeywordItem.Keyword
		lang = userSubKeywordItem.Lang
		orderBydate = userSubKeywordItem.OrderByDate
		_, err = celery.GetClient().Delay(consts.USER_SUB_KEYWORD_UPDATE, keyword, lang, orderBydate)
		if err != nil {
			g.Log().Line().Error(ctx, fmt.Sprintf("Assign USER_SUB_KEYWORD_UPDATE failed : %s\n", err))
		}
	}

}
