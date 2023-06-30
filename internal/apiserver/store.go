package apiserver

import (
	"github.com/marmotedu/iam/pkg/log"
	"github.com/ngsin/iam-learning/internal/apiserver/store"
	"github.com/ngsin/iam-learning/internal/apiserver/store/mysql"
)

func (s *Server) initStore() {
	initMySQL(s)
}

func initMySQL(s *Server) {
	storeIns, err := mysql.GetMySQLFactoryOr(s.extraConfig.mysqlOptions)
	if err != nil {
		log.Fatalf("get mysql factory failed: %s", err.Error())
	}

	store.SetClient(storeIns)

}
