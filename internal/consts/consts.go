package consts

import "github.com/gogf/gf/v2/frame/g"

// template consts
const APP_NAME_KEY = "appname"
const APP_NAME = "锅烧FM"

const TOTAL_PAGE = "totalPage"
const SEARCH_RESULT_COUNT_TEXT = "searchResultCountText"
const SEARCH_RESULT_COUNT_TEXT_VALUE = " %d 搜索结果"
const SEARCH_TOOK_TIME_TEXT = "searchTookTime"
const SEARCH_TOOK_TIME_TEXT_VALUE = "(%s 秒)"
const FEED_ITEMS = "feedItems"
const CHANNEL_INFO = "channelInfo"
const ITEM_INFO = "itemInfo"
const TOTAL_CHANNE_ITEMS_COUNT = "total_channe_items_count"

const SEARCH_KEY_WORD = "searchKeyword"
const CURRENT_PAGE = "currentPage"

const LISTEN_LATER_HEADER_TAG_TEXT = "稍后再听"
const LISTEN_LATER_HEADER_TAG = "listen_later_header_tag"
const MY_ACCOUNT_TAG = "my_account_tag"
const MY_ACCOUNT_TAG_VALUE = "我的账号"
const LOGIN_TAG = "login_tag"
const LOGIN_TAG_VALUE = "登录"
const LOGOUT_TAG = "logout_tag"
const LOGOUT_TAG_VALUE = "登出"

const ADD_ON = "add_on"
const ADD_ON_TEXT = "添加于"

const PAST_FEED_ITEMS = "past_feed_items"

func GetCommonTplMap() (tplMap g.Map) {

	tplMap = g.Map{
		APP_NAME_KEY:            APP_NAME,
		LISTEN_LATER_HEADER_TAG: LISTEN_LATER_HEADER_TAG_TEXT,
		MY_ACCOUNT_TAG:          MY_ACCOUNT_TAG_VALUE,
		LOGOUT_TAG:              LOGOUT_TAG_VALUE,
		LOGIN_TAG:               LOGIN_TAG_VALUE,
	}

	return
}
