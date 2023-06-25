package mysql

import (
	"fmt"
	"github.com/marmotedu/errors"
	"github.com/ngsin/iam-learning/internal/apiserver/store"
	genericoptions "github.com/ngsin/iam-learning/internal/pkg/options"
	"github.com/ngsin/iam-learning/pkg/db"
	"gorm.io/gorm"
	"sync"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Secrets() store.SecretStore {
	return newSecrets(ds)
}

func (ds *datastore) Policies() store.PolicyStore {
	return newPolicies(ds)
}

func (ds *datastore) PolicyAudits() store.PolicyAuditStore {
	return newPolicyAudits(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			//Logger:                logger.New(opts.LogLevel),
		}
		dbIns, err = db.New(options)

		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		// migrateDatabase(dbIns)

		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}
