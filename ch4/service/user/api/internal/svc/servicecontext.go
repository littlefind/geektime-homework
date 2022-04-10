package svc

import (
	"ch4/service/user/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn,c.CacheRedis),
	}
}
