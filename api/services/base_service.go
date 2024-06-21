package services

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarthakjdev/wapikit/internal/interfaces"
)

type BaseService struct {
	Name        string `json:"name"`
	RestApiPath string `json:"rest_api_path"`
	Routes      []interfaces.Route
}

func (s *BaseService) GetServiceName() string {
	return s.Name

}

func (s *BaseService) GetRoutes() []interfaces.Route {
	return s.Routes
}

func (s *BaseService) GetRestApiPath() string {
	return s.RestApiPath
}

func isAuthorized(role interfaces.PermissionRole, routerPermissionLevel interfaces.PermissionRole) bool {
	switch role {
	case interfaces.OwnerRole:
		return true
	case interfaces.AdminRole:
		return routerPermissionLevel == interfaces.AdminRole || routerPermissionLevel == interfaces.MemberRole
	case interfaces.MemberRole:
		return routerPermissionLevel == interfaces.MemberRole
	default:
		return false
	}
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		app := ctx.Get("app").(*interfaces.App)
		headers := ctx.Request().Header
		authToken := headers.Get("x-access-token")
		origin := headers.Get("origin")

		if authToken == "" {
			return echo.ErrUnauthorized
		}

		// verify the token
		// get the user
		// inject user to context
		// consider user permission level too
		// return error if token is invalid
		fmt.Println("authMiddleware: ", authToken, origin)
		metadata := ctx.Get("metadata").(interfaces.Route)

		// ! TODO: fetch the user from db and check role here
		mockRole := interfaces.OwnerRole

		if isAuthorized(mockRole, metadata.PermissionRoleLevel) {

			// ! TODO : try if this works anyway, the only confusion here is even if I create this parent and child context thing,
			// ! still would i be getting intellisense in the context consumer to access the new props added such as app and session.
			// appContext := context.WithValue(ctx.Request().Context(), "App", *app)
			// sessionContext := context.WithValue(ctx.Request().Context(), "App", *app)

			return next(interfaces.CustomContext{
				Context: ctx,
				App:     *app,
				Session: interfaces.ContextSession{
					Token: authToken,
					User: interfaces.ContextUser{
						UniqueId: "",
						Username: "",
						Email:    "",
						Role:     mockRole,
					},
				},
			})
		} else {
			fmt.Println("authMiddleware: ", authToken, origin, metadata.PermissionRoleLevel, mockRole)
			return echo.ErrForbidden
		}
	}
}

func rateLimiter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		// rate limit the request
		// return error if rate limit is exceeded
		return next(context)
	}
}

// Register function now uses the Routes field
func (service *BaseService) Register(server *echo.Echo) {
	for _, route := range service.Routes {
		handler := interfaces.CustomHandler(route.Handler).Handle
		if route.IsAuthorizationRequired {
			handler = authMiddleware(interfaces.CustomHandler(route.Handler).Handle)
		}
		switch route.Method {
		case http.MethodGet:
			server.GET(route.Path, handler)
		case http.MethodPost:
			server.POST(route.Path, handler)
		case http.MethodPut:
			server.PUT(route.Path, handler)
		case http.MethodDelete:
			server.DELETE(route.Path, handler)
		}
	}
}
