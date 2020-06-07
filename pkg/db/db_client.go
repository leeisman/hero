package db

import (
	"context"
	"hero/configs"
	"hero/database/ent"
	"hero/pkg/logger"
	"time"
)

func Client() *ent.Client {
	localDataSourceName := configs.Get("database.url")
	client, err := ent.Open("mysql", localDataSourceName)
	if err != nil {
		logger.Error("failed connecting to mysql: " + err.Error())
	}
	// Add a global hook that runs on all types and all operations.
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				logger.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})
	return client
}
