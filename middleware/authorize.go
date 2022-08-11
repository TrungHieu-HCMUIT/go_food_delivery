package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_restaurant/common"
	"go_restaurant/component"
	"go_restaurant/component/tokenprovider/jwt"
	"go_restaurant/modules/user/userstorage"
	"strings"
)

type jwtMiddleware struct {
	TokenType string
}

func NewJwtMiddleware(tokType string) *jwtMiddleware {
	return &jwtMiddleware{
		TokenType: tokType,
	}
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authentication header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeader(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization": "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(errors.New("wrong authentication header"))
	}

	return parts[1], nil
}

func RequiredAuth(appCtx component.AppContext) gin.HandlerFunc {
	tokProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {

		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlStorage(db)

		payload, err := tokProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
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
