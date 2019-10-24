package nrcontext

import (
	"context"

	"bitbucket.org/snapmartinc/newrelic-context/nrgorm"
	"bitbucket.org/snapmartinc/newrelic-context/nrredis"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/newrelic/go-agent"
)

// Sets transaction from Context to gorm settings, returns cloned DB
func SetTxnToGorm(ctx context.Context, db *gorm.DB) *gorm.DB {
	txn := newrelic.FromContext(ctx)
	return nrgorm.SetTxnToGorm(txn, db)
}


// Gets transaction from Context and applies RedisWrapper, returns cloned client
func WrapRedisClient(ctx context.Context, c *redis.Client) *redis.Client {
	txn := newrelic.FromContext(ctx)
	return nrredis.WrapRedisClient(txn, c)
}
