package feed

import (
	"context"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

func SubFeedByKeyword(ctx context.Context, userId, keyword, lang string, sortByDate int) (err error) {

	var (
		userSubKeyword entity.UserSubKeyword
	)

	userSubKeyword = entity.UserSubKeyword{
		Id:          guid.S(),
		UserId:      userId,
		Keyword:     keyword,
		Lang:        lang,
		OrderByDate: sortByDate,
		CreateTime:  gtime.Now(),
	}

	err = dao.CreateUserSubKeyword(ctx, userSubKeyword)
	if err != nil {
		return
	}

	return err
}
