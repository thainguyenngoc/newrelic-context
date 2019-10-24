package nrcontext

import (
	"bitbucket.org/snapmartinc/newrelic-context/nrgorm"
	"bitbucket.org/snapmartinc/newrelic-context/nrredis"
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/newrelic/go-agent"
)

type contextKey int

const txnKey contextKey = 0

// Set NewRelic transaction to context
func ContextWithTxn(c context.Context, txn newrelic.Transaction) context.Context {
	return context.WithValue(c, txnKey, txn)
}

// Get NewRelic transaction from context anywhere
func GetTnxFromContext(c context.Context) newrelic.Transaction {
	if tnx := c.Value(txnKey); tnx != nil {
		fmt.Println("transaction is not nill")
		return tnx.(newrelic.Transaction)
	}
	return nil
}

// Sets transaction from Context to gorm settings, returns cloned DB
func SetTxnToGorm(ctx context.Context, db *gorm.DB) *gorm.DB {
	txn := GetTnxFromContext(ctx)
	return nrgorm.SetTxnToGorm(txn, db)
}

func SetTnxToGormMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			txn := GetTnxFromContext(ctx)
			db = nrgorm.SetTxnToGorm(txn, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Gets transaction from Context and applies RedisWrapper, returns cloned client
func WrapRedisClient(ctx context.Context, c *redis.Client) *redis.Client {
	txn := GetTnxFromContext(ctx)
	return nrredis.WrapRedisClient(txn, c)
}
