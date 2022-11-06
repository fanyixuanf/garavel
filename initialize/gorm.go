/*
@Time : 2022/10/3 20:15
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : gorm
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package initialize

import (
	"context"
	"garavel/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func GormWriteHandler (config *config.Server, log *zap.Logger) *gorm.DB {
	dsn := config.Mysql.Username + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Path + ")/" + config.Mysql.Dbname + "?" + config.Mysql.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	gormConfig := logConfig(config.Mysql.LogMode)
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		log.Error("database connect failed(write), err:", zap.Any("err", err))
		os.Exit(0)
	} else {
		log.Info("database connect success(write)")
		return db
	}
	return nil
}

func GormReadHandler (config *config.Server, log *zap.Logger) *ReadDb {
	dsn := config.MysqlRead.Username + ":" + config.MysqlRead.Password + "@tcp(" + config.MysqlRead.Path + ")/" + config.MysqlRead.Dbname + "?" + config.MysqlRead.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	gormConfig := logConfig(config.MysqlRead.LogMode)
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		log.Error("database connect failed, err(read):", zap.Any("err", err))
		os.Exit(0)
	} else {
		log.Info("database connect success(read)")
		return &ReadDb{db:db}
	}
	return nil
}

// config 根据配置决定是否开启日志
func logConfig(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	return
}

type ReadDb struct {
	db *gorm.DB
}

type WriteDb struct {
	db *gorm.DB
}


type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}

// NewTransaction .
func NewTransaction(w *WriteDb) Transaction {
	return w
}
type contextTxKey struct{}

func (d *WriteDb) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *WriteDb) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}
