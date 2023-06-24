package apiserver

import (
	"context"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github/ngsin/iam-learning/internal/apiserver/store"
	"github/ngsin/iam-learning/internal/pkg/middleware"
	"github/ngsin/iam-learning/internal/pkg/middleware/auth"
	"time"
)

const (
	// APIServerAudience defines the value of jwt audience field.
	APIServerAudience = "iam.api.marmotedu.com"

	// APIServerIssuer defines the value of jwt issuer field.
	APIServerIssuer = "iam-apiserver"
)

type loginInfo struct {
	Username string `form:"username" json:"username" binding:"required,username"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func newBasicAuth() middleware.AuthStrategy {
	return auth.NewBasicStrategy(func(username string, password string) bool {
		// fetch user from database
		user, err := store.Client().Users().Get(context.TODO(), username, metav1.GetOptions{})
		if err != nil {
			return false
		}

		// Compare the login password with the user password.
		if err := user.Compare(password); err != nil {
			return false
		}

		user.LoginedAt = time.Now()
		_ = store.Client().Users().Update(context.TODO(), user, metav1.UpdateOptions{})
		return true
	})
}
