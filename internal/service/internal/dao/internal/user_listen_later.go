// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-05-11 22:32:54
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserListenLaterDao is the data access object for table user_listen_later.
type UserListenLaterDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns UserListenLaterColumns // columns contains all the column names of Table for convenient usage.
}

// UserListenLaterColumns defines and stores column names for table user_listen_later.
type UserListenLaterColumns struct {
	Id        string //
	UserId    string //
	ItemId    string //
	ChannelId string //
	RegDate   string //
}

//  userListenLaterColumns holds the columns for table user_listen_later.
var userListenLaterColumns = UserListenLaterColumns{
	Id:        "id",
	UserId:    "user_id",
	ItemId:    "item_id",
	ChannelId: "channel_id",
	RegDate:   "reg_date",
}

// NewUserListenLaterDao creates and returns a new DAO object for table data access.
func NewUserListenLaterDao() *UserListenLaterDao {
	return &UserListenLaterDao{
		group:   "default",
		table:   "user_listen_later",
		columns: userListenLaterColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserListenLaterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserListenLaterDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserListenLaterDao) Columns() UserListenLaterColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserListenLaterDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserListenLaterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserListenLaterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
