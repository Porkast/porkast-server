// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"guoshao-fm-web/internal/service/internal/dao/internal"
)

// feedItemDao is the data access object for table feed_item.
// You can define custom methods on it to extend its functionality as you wish.
type feedItemDao struct {
	*internal.FeedItemDao
}

var (
	// FeedItem is globally public accessible object for table feed_item operations.
	FeedItem = feedItemDao{
		internal.NewFeedItemDao(),
	}
)

// Fill with you ideas below.
