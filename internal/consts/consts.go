package consts

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// template consts
const APP_NAME_KEY = "appname"
const APP_NAME = "锅烧FM"

const TOTAL_PAGE = "totalPage"
const SEARCH_RESULT_COUNT_TEXT = "searchResultCountText"
const SEARCH_RESULT_COUNT_TEXT_VALUE = "search_total_count"
const SEARCH_TOOK_TIME_TEXT = "searchTookTime"
const SEARCH_TOOK_TIME_TEXT_VALUE = "search_took_time"
const FEED_ITEMS = "feedItems"
const FEED_CHANNELS = "feedChannels"
const CHANNEL_INFO = "channelInfo"
const ITEM_INFO = "itemInfo"
const TOTAL_CHANNE_ITEMS_COUNT = "total_channe_items_count"

const SEARCH_KEYWORD = "searchKeyword"
const CURRENT_PAGE = "currentPage"
const SEARCH_CHANNEL = "search_channel"
const SEARCH_CHANNEL_SCOPE = "channel"
const SEARCH_ORDER_BY_DATE = "search_order_by_date"
const TOTAL_CHANNEL_COUNT = "total_channe_count"
const TOTAL_ITEM_COUNT = "total_item_count"
const SEARCH_CN_FEED_ITEM_CHANNEL_TOTAL_WITH_COUNT = "search_cn_feed_item_channel_total_with_count"
const SEARCH_PODCAST_BY_KEYWORD = "search_podcast_by_keyword"
const SEARCH_ONLY_MATCH_TITLE = "search_only_match_title"

const LISTEN_LATER_HEADER_TAG_TEXT = "稍后再听"
const LISTEN_LATER_HEADER_TAG = "listen_later_header_tag"
const USER_SUB_LIST = "user_sub_list"
const USER_SUB_LIST_TEXT = "订阅列表"
const MY_ACCOUNT_TAG = "my_account_tag"
const MY_ACCOUNT_TAG_VALUE = "我的账号"
const LOGIN_TAG = "login_tag"
const LOGIN_TAG_VALUE = "登录"
const LOGOUT_TAG = "logout_tag"
const LOGOUT_TAG_VALUE = "登出"

const ADD_ON = "add_on"
const ADD_ON_TEXT = "添加于"

const PAST_FEED_ITEMS = "past_feed_items"

const SUBSCRIPTION = "subscription"
const COPY = "copy"

func GetCommonTplMap(ctx context.Context) (tplMap g.Map) {

	tplMap = g.Map{
		APP_NAME_KEY:            APP_NAME,
		LISTEN_LATER_HEADER_TAG: LISTEN_LATER_HEADER_TAG_TEXT,
		USER_SUB_LIST:           USER_SUB_LIST_TEXT,
		MY_ACCOUNT_TAG:          MY_ACCOUNT_TAG_VALUE,
		LOGOUT_TAG:              LOGOUT_TAG_VALUE,
		LOGIN_TAG:               LOGIN_TAG_VALUE,
	}
	tplMap[SEARCH_ONLY_MATCH_TITLE] = g.I18n().T(ctx, SEARCH_ONLY_MATCH_TITLE)

	return
}
