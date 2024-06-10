package migration

import "embed"

//go:embed *
var FS embed.FS

const (
	//MysqlPath      = "mysql"
	//MongoPath      = "mongo"
	PostgresPath = "postgres"
	//ClickhousePath = "clickhouse"
)
