module github.com/jasonbronson/crawlmonster-api

go 1.16

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/pprof v1.3.0 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-redis/redis/v7 v7.4.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/hashicorp/go-uuid v1.0.1
	github.com/joho/godotenv v1.3.0
	github.com/newrelic/go-agent v3.12.0+incompatible // indirect
	github.com/newrelic/go-agent/v3 v3.12.0
	github.com/newrelic/go-agent/v3/integrations/nrgin v1.1.0
	github.com/newrelic/go-agent/v3/integrations/nrlogrus v1.0.0 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrpq v1.1.0 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrredis-v7 v1.0.0
	github.com/newrelic/go-agent/v3/integrations/nrsqlite3 v1.1.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/sumit-tembe/gin-requestid v0.0.0-20191217132119-618fbd2c6306
	github.com/xo/dburl v0.7.0
	go.uber.org/zap v1.13.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	gorm.io/driver/mysql v1.0.6 // indirect
	gorm.io/driver/postgres v1.1.0 // indirect
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.9
)
