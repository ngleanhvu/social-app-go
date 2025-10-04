package middleware

import (
	"crud-go/common"
	"crud-go/component/appctx"
	"crud-go/component/tokenprovider/jwt"
	userstorage "crud-go/module/user/storage"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomErrorResponse(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

// extractTokenFromHeaderString returns the access token in authorization field in request header
func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//Authorization : Bearn{token}
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB

func RequireAuth(appCtx appctx.AppContext) func(ctx *gin.Context) {

	tokenProvider := jwt.NewJWTProvider(common.SecretKey)

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()

		store := userstorage.NewSqlStore(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := store.Find(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			if err == common.RecordNotFound {
				panic(common.ErrNoPermission(errors.New("user not found")))
			}
			//c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}

}
