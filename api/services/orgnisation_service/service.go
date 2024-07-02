package organization_service

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sarthakjdev/wapikit/api/services"
	"github.com/sarthakjdev/wapikit/database"
	"github.com/sarthakjdev/wapikit/internal"
	"github.com/sarthakjdev/wapikit/internal/api_types"
	"github.com/sarthakjdev/wapikit/internal/interfaces"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/sarthakjdev/wapikit/.db-generated/model"
	table "github.com/sarthakjdev/wapikit/.db-generated/table"
)

type OrganizationService struct {
	services.BaseService `json:"-,inline"`
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{
		BaseService: services.BaseService{
			Name:        "Organization Service",
			RestApiPath: "/api/organization",
			Routes: []interfaces.Route{
				{
					Path:                    "/api/organization",
					Method:                  http.MethodGet,
					Handler:                 GetOrganizations,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Member,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization",
					Method:                  http.MethodPost,
					Handler:                 CreateNewOrganization,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/:id",
					Method:                  http.MethodPost,
					Handler:                 UpdateOrganizationId,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Owner,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/:id",
					Method:                  http.MethodGet,
					Handler:                 GetOrganizationById,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/settings",
					Method:                  http.MethodPost,
					Handler:                 GetOrganizationSettings,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/roles",
					Method:                  http.MethodGet,
					Handler:                 GetOrganizationRoles,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/roles/:id",
					Method:                  http.MethodGet,
					Handler:                 GetRoleById,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/roles/:id",
					Method:                  http.MethodDelete,
					Handler:                 DeleteRoleById,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/roles/:id",
					Method:                  http.MethodPost,
					Handler:                 UpdateRoleById,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/members",
					Method:                  http.MethodGet,
					Handler:                 GetOrganizationMembers,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/members",
					Method:                  http.MethodPost,
					Handler:                 CreateNewOrganizationMember,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/members/:id",
					Method:                  http.MethodPost,
					Handler:                 UpdateOrgMemberById,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/members/:id",
					Method:                  http.MethodGet,
					Handler:                 GetOrgMemberById,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/members/:id",
					Method:                  http.MethodDelete,
					Handler:                 UpdateOrgMemberById,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/members/:id/role",
					Method:                  http.MethodPost,
					Handler:                 UpdateOrganizationMemberRoles,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/syncTemplates",
					Method:                  http.MethodGet,
					Handler:                 syncTemplates,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
				{
					Path:                    "/api/organization/syncMobileNumbers",
					Method:                  http.MethodGet,
					Handler:                 syncMobileNumbers,
					IsAuthorizationRequired: true,
					MetaData: interfaces.RouteMetaData{
						PermissionRoleLevel: api_types.Admin,
						RateLimitConfig: interfaces.RateLimitConfig{
							MaxRequests:    10,
							WindowTimeInMs: 1000 * 60, // 1 minute
						},
					},
				},
			},
		},
	}
}

func CreateNewOrganization(context interfaces.CustomContext) error {
	payload := new(api_types.NewOrganizationSchema)
	if err := context.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var newOrg model.Organization
	var member model.OrganizationMember

	tx, err := context.App.Db.BeginTx(context.Request().Context(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer tx.Rollback()

	// 1. Insert Organization
	err = table.Organization.INSERT().
		MODEL(newOrg).
		RETURNING(table.Organization.AllColumns).
		QueryContext(context.Request().Context(), tx, &newOrg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userUuid, err := uuid.FromBytes([]byte(context.Session.User.UniqueId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// 2. Insert Organization Member
	err = table.OrganizationMember.INSERT().MODEL(model.OrganizationMember{
		AccessLevel:    model.UserPermissionLevel_Owner,
		OrganizationId: newOrg.UniqueId,
		UserId:         userUuid,
	}).RETURNING(table.OrganizationMember.AllColumns).QueryContext(context.Request().Context(), tx, &member)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// 3. Create API key for the organization

	claims := &interfaces.JwtPayload{
		ContextUser: interfaces.ContextUser{
			Username:       context.Session.User.Username,
			Email:          context.Session.User.Email,
			Role:           api_types.UserRoleEnum(api_types.Owner),
			UniqueId:       context.Session.User.UniqueId,
			OrganizationId: newOrg.UniqueId.String(),
			Name:           context.Session.User.Name,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365 * 2).Unix(), // 60-day expiration
			Issuer:    "wapikit",
		},
	}

	//Create the token

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(context.App.Koa.String("app.jwt_secret")))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating token")
	}

	var apiKey model.ApiKey

	err = table.ApiKey.INSERT().MODEL(model.ApiKey{
		MemberId:       member.UniqueId,
		OrganizationId: newOrg.UniqueId,
		Key:            token,
	}).RETURNING(table.ApiKey.AllColumns).QueryContext(context.Request().Context(), tx, &apiKey)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.String(http.StatusOK, "OK")
}

func GetOrganizations(context interfaces.CustomContext) error {
	return context.String(http.StatusOK, "OK")
}

func GetOrganizationById(context interfaces.CustomContext) error {

	organizationId := context.Param("id")
	if organizationId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid organization id")
	}

	orgUuid, _ := uuid.FromBytes([]byte(organizationId))
	hasAccess := VerifyAccessToOrganization(context, context.Session.User.UniqueId, organizationId)

	if !hasAccess {
		return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this organization")
	}

	var dest model.Organization
	organizationQuery := SELECT(table.Organization.AllColumns).
		FROM(table.Organization).
		WHERE(table.Organization.UniqueId.EQ(UUID(orgUuid)))
	err := organizationQuery.QueryContext(context.Request().Context(), context.App.Db, &dest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	uniqueId := dest.UniqueId.String()
	return context.JSON(http.StatusOK, api_types.GetOrganizationByIdResponseSchema{
		Organization: api_types.OrganizationSchema{
			Name:       dest.Name,
			CreatedAt:  dest.CreatedAt,
			UniqueId:   uniqueId,
			FaviconUrl: &dest.FaviconUrl,
			LogoUrl:    dest.LogoUrl,
			WebsiteUrl: dest.WebsiteUrl,
		},
	})
}

func DeleteOrganization(context interfaces.CustomContext) error {

	return context.String(http.StatusInternalServerError, "NOT IMPLEMENTED YET")

	organizationId := context.Param("id")
	if organizationId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid organization id")
	}

	hasAccess := VerifyAccessToOrganization(context, context.Session.User.UniqueId, organizationId)

	if !hasAccess {
		return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this organization")
	}

	return context.String(http.StatusOK, "OK")
}

func UpdateOrganizationId(context interfaces.CustomContext) error {
	organizationId := context.Param("id")
	if organizationId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid organization id")
	}

	hasAccess := VerifyAccessToOrganization(context, context.Session.User.UniqueId, organizationId)

	if !hasAccess {
		return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this organization")
	}
	return context.String(http.StatusOK, "OK")
}

func GetOrganizationRoles(context interfaces.CustomContext) error {

	params := new(api_types.GetOrganizationRolesParams)

	err := internal.BindQueryParams(context, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var dest struct {
		TotalRoles int `json:"totalRoles"`
		roles      []struct {
			model.OrganizationRole
		}
	}

	orgUuid, err := uuid.FromBytes([]byte(context.Session.User.OrganizationId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	whereCondition := table.OrganizationRole.OrganizationId.EQ(UUID(orgUuid))

	organizationRolesQuery := SELECT(table.OrganizationRole.AllColumns).
		FROM(table.OrganizationRole).
		WHERE(whereCondition).
		LIMIT(params.PerPage).
		OFFSET(params.Page * params.PerPage)

	if params.SortBy != nil {
		if *params.SortBy == api_types.Asc {
			organizationRolesQuery.ORDER_BY(table.OrganizationRole.CreatedAt.ASC())
		} else {
			organizationRolesQuery.ORDER_BY(table.OrganizationRole.CreatedAt.DESC())
		}
	}

	err = organizationRolesQuery.QueryContext(context.Request().Context(), context.App.Db, &dest)

	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			roles := make([]api_types.OrganizationRoleSchema, 0)
			total := 0
			return context.JSON(http.StatusOK, api_types.GetOrganizationRolesResponseSchema{
				Roles: roles,
				PaginationMeta: api_types.PaginationMeta{
					Page:    params.Page,
					PerPage: params.PerPage,
					Total:   total,
				},
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	rolesToReturn := make([]api_types.OrganizationRoleSchema, len(dest.roles))

	if len(dest.roles) > 0 {
		for _, role := range dest.roles {
			permissions := make([]api_types.RolePermissionEnum, len(role.Permissions))
			for _, perm := range role.Permissions {
				permissions = append(permissions, api_types.RolePermissionEnum(perm))
			}

			roleId := role.UniqueId.String()

			roleToReturn := api_types.OrganizationRoleSchema{

				Description: &role.Description,
				Name:        role.Name,
				Permissions: permissions,
				UniqueId:    roleId,
			}

			rolesToReturn = append(rolesToReturn, roleToReturn)
		}
	}

	return context.JSON(http.StatusOK, api_types.GetOrganizationRolesResponseSchema{
		Roles: rolesToReturn,
		PaginationMeta: api_types.PaginationMeta{
			Page:    params.Page,
			PerPage: params.PerPage,
			Total:   dest.TotalRoles,
		},
	})

}

func GetOrganizationSettings(context interfaces.CustomContext) error {
	organizationId := context.Param("id")
	if organizationId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid organization id")
	}

	hasAccess := VerifyAccessToOrganization(context, context.Session.User.UniqueId, organizationId)
	if !hasAccess {
		return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this organization")
	}

	return context.String(http.StatusOK, "OK")
}

func GetRoleById(context interfaces.CustomContext) error {
	roleId := context.Param("id")
	if roleId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid role id")
	}

	roleUuid, _ := uuid.FromBytes([]byte(roleId))
	roleQuery := SELECT(table.OrganizationRole.AllColumns).FROM(table.OrganizationRole).WHERE(table.OrganizationRole.UniqueId.EQ(UUID(roleUuid))).LIMIT(1)

	var dest model.OrganizationRole
	err := roleQuery.QueryContext(context.Request().Context(), context.App.Db, &dest)

	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			role := new(api_types.OrganizationRoleSchema)
			return context.JSON(http.StatusOK, api_types.GetRoleByIdResponseSchema{
				Role: *role,
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	if dest.OrganizationId.String() != context.Session.User.OrganizationId {
		return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this resource")
	}

	permissionToReturn := make([]api_types.RolePermissionEnum, len(dest.Permissions))

	for _, perm := range dest.Permissions {
		permissionToReturn = append(permissionToReturn, api_types.RolePermissionEnum(perm))
	}

	role := api_types.OrganizationRoleSchema{
		Description: &dest.Description,
		Name:        dest.Name,
		Permissions: permissionToReturn,
		UniqueId:    roleId,
	}

	return context.JSON(http.StatusOK, role)
}

func DeleteRoleById(context interfaces.CustomContext) error {
	return context.String(http.StatusOK, "OK")
}

func UpdateRoleById(context interfaces.CustomContext) error {
	return context.String(http.StatusOK, "OK")
}

func GetOrganizationMembers(context interfaces.CustomContext) error {

	params := new(api_types.GetOrganizationMembersParams)

	if err := internal.BindQueryParams(context, &params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pageNumber := params.Page
	pageSize := params.PerPage
	sortBy := params.SortBy

	organizationUuid, err := uuid.FromBytes([]byte(context.Session.User.OrganizationId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	organizationMembersQuery := SELECT(table.OrganizationMember.AllColumns,
		table.User.Username,
		table.User.Name,
		table.User.Email,
		table.RoleAssignment.AllColumns,
		table.OrganizationRole.AllColumns,
		COUNT(table.OrganizationMember.OrganizationId.EQ(UUID(organizationUuid))).OVER().AS("totalMembers"),
	).
		FROM(table.OrganizationMember.
			LEFT_JOIN(table.User, table.User.UniqueId.EQ(table.OrganizationMember.UserId)).
			LEFT_JOIN(table.RoleAssignment, table.RoleAssignment.OrganizationMemberId.EQ(table.OrganizationMember.UniqueId)).
			LEFT_JOIN(table.OrganizationRole, table.OrganizationRole.UniqueId.EQ(table.RoleAssignment.OrganizationRoleId))).
		WHERE(table.OrganizationMember.OrganizationId.EQ(UUID(organizationUuid))).
		LIMIT(pageSize).
		OFFSET(pageNumber * pageSize)

	if sortBy != nil {
		if *sortBy == api_types.Asc {
			organizationMembersQuery.ORDER_BY(table.OrganizationMember.CreatedAt.ASC())
		} else {
			organizationMembersQuery.ORDER_BY(table.OrganizationMember.CreatedAt.ASC())
		}
	}

	var dest struct {
		TotalMembers int `json:"totalMembers"`
		members      []struct {
			model.OrganizationMember
			model.User
			Roles []struct {
				model.OrganizationRole
			}
		}
	}

	err = organizationMembersQuery.QueryContext(context.Request().Context(), context.App.Db, &dest)

	if err != nil {

		if err.Error() == "qrm: no rows in result set" {
			members := make([]api_types.OrganizationMemberSchema, 0)
			total := 0
			return context.JSON(http.StatusOK, api_types.GetOrganizationMembersResponseSchema{
				Members: members,
				PaginationMeta: api_types.PaginationMeta{
					Page:    pageNumber,
					PerPage: pageSize,
					Total:   total,
				},
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	membersToReturn := make([]api_types.OrganizationMemberSchema, len(dest.members))

	if len(dest.members) > 0 {
		for _, member := range dest.members {
			memberRoles := make([]api_types.OrganizationRoleSchema, len(member.Roles))
			if len(member.Roles) > 0 {
				for _, role := range member.Roles {
					permissions := make([]api_types.RolePermissionEnum, len(role.Permissions))
					for _, perm := range role.Permissions {
						permissions = append(permissions, api_types.RolePermissionEnum(perm))
					}

					roleId := role.UniqueId.String()

					roleToReturn := api_types.OrganizationRoleSchema{
						Description: &role.Description,
						Name:        role.Name,
						Permissions: permissions,
						UniqueId:    roleId,
					}

					memberRoles = append(memberRoles, roleToReturn)
				}
			}

			accessLevel := api_types.UserRoleEnum(member.OrganizationMember.AccessLevel)
			memberId := member.User.UniqueId.String()
			mmbr := api_types.OrganizationMemberSchema{
				CreatedAt:   member.OrganizationMember.CreatedAt,
				AccessLevel: accessLevel,
				UniqueId:    memberId,
				Email:       member.User.Email,
				Name:        member.User.Name,
				Roles:       memberRoles,
			}
			membersToReturn = append(membersToReturn, mmbr)
		}

	}

	return context.JSON(http.StatusOK, api_types.GetOrganizationMembersResponseSchema{
		Members: membersToReturn,
		PaginationMeta: api_types.PaginationMeta{
			Page:    pageNumber,
			PerPage: pageSize,
			Total:   dest.TotalMembers,
		}})
}

func CreateNewOrganizationMember(context interfaces.CustomContext) error {
	payload := new(interface{})
	if err := context.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	//     return "", fmt.Errorf("error hashing password: %w", err)
	// }
	// return string(hash), nil

	database.GetDbInstance()
	return context.String(http.StatusOK, "OK")
}

func GetOrgMemberById(context interfaces.CustomContext) error {
	memberId := context.Param("id")

	if memberId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid member id")
	}

	memberUuid, _ := uuid.FromBytes([]byte(memberId))
	memberQuery := SELECT(table.OrganizationMember.AllColumns,
		table.User.Username,
		table.User.Name,
		table.User.Email,
		table.RoleAssignment.AllColumns,
		table.OrganizationRole.AllColumns,
	).
		FROM(table.OrganizationMember.
			LEFT_JOIN(table.User, table.User.UniqueId.EQ(table.OrganizationMember.UserId)).
			LEFT_JOIN(table.RoleAssignment, table.RoleAssignment.OrganizationMemberId.EQ(table.OrganizationMember.UniqueId)).
			LEFT_JOIN(table.OrganizationRole, table.OrganizationRole.UniqueId.EQ(table.RoleAssignment.OrganizationRoleId))).
		WHERE(table.OrganizationMember.UniqueId.EQ(UUID(memberUuid))).
		LIMIT(1)

	var dest struct {
		member struct {
			model.OrganizationMember
			model.User
			Roles []struct {
				model.OrganizationRole
			}
		}
	}

	err := memberQuery.QueryContext(context.Request().Context(), context.App.Db, &dest)

	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			member := new(api_types.OrganizationMemberSchema)
			return context.JSON(http.StatusOK, api_types.GetOrganizationMemberByIdResponseSchema{
				Member: *member,
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	memberRoles := make([]api_types.OrganizationRoleSchema, len(dest.member.Roles))
	if len(dest.member.Roles) > 0 {
		for _, role := range dest.member.Roles {
			permissions := make([]api_types.RolePermissionEnum, len(role.Permissions))
			for _, perm := range role.Permissions {
				permissions = append(permissions, api_types.RolePermissionEnum(perm))
			}
			roleId := role.UniqueId.String()
			roleToReturn := api_types.OrganizationRoleSchema{
				Description: &role.Description,
				Name:        role.Name,
				Permissions: permissions,
				UniqueId:    roleId,
			}
			memberRoles = append(memberRoles, roleToReturn)
		}
	}

	accessLevel := api_types.UserRoleEnum(dest.member.OrganizationMember.AccessLevel)

	member := api_types.OrganizationMemberSchema{
		CreatedAt:   dest.member.OrganizationMember.CreatedAt,
		AccessLevel: accessLevel,
		UniqueId:    memberId,
		Email:       dest.member.User.Email,
		Name:        dest.member.User.Name,
		Roles:       memberRoles,
	}

	return context.JSON(http.StatusOK, api_types.GetOrganizationMemberByIdResponseSchema{
		Member: member,
	})
}

func DeleteOrgMemberById(context interfaces.CustomContext) error {
	return context.String(http.StatusOK, "OK")
}

func UpdateOrgMemberById(context interfaces.CustomContext) error {
	memberId := context.Param("id")
	if memberId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid member id")
	}

	memberUuid, _ := uuid.FromBytes([]byte(memberId))

	payload := new(api_types.UpdateOrganizationMemberSchema)
	if err := context.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updateMemberQuery := table.OrganizationMember.
		UPDATE(table.OrganizationMember.AccessLevel).
		SET(payload.AccessLevel).
		WHERE(table.OrganizationMember.UniqueId.EQ(UUID(memberUuid))).
		RETURNING(table.OrganizationMember.AllColumns)

	_, err := updateMemberQuery.ExecContext(context.Request().Context(), context.App.Db)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.String(http.StatusOK, "OK")
}

func UpdateOrganizationMemberRoles(context interfaces.CustomContext) error {

	memberId := context.Param("id")

	if memberId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid member id")
	}

	memberUuid, _ := uuid.FromBytes([]byte(memberId))

	payload := new(api_types.UpdateOrganizationMemberRoleSchema)

	if err := context.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	roleUuid, err := uuid.FromBytes([]byte(*payload.RoleUniqueId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid member id")
	}

	orgRole := SELECT(table.OrganizationRole.AllColumns, table.RoleAssignment.AllColumns).
		FROM(table.OrganizationRole.
			LEFT_JOIN(table.RoleAssignment, table.RoleAssignment.OrganizationRoleId.EQ(table.OrganizationRole.UniqueId)),
		).WHERE(table.OrganizationRole.UniqueId.EQ(UUID(roleUuid))).
		LIMIT(1)

	var dest struct {
		model.OrganizationRole
		Assignment model.RoleAssignment
	}

	err = orgRole.QueryContext(context.Request().Context(), context.App.Db, &dest)

	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return echo.NewHTTPError(http.StatusNotFound, "Role not found")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	switch payload.Action {
	case api_types.Add:
		{

			// Check if the role is already assigned to the member
			if dest.Assignment.UniqueId != uuid.Nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Role already assigned to the member")
			} else {
				// Assign the role to the member
				var roleAssignmentDest model.RoleAssignment

				roleAssignment := model.RoleAssignment{
					OrganizationMemberId: memberUuid,
					OrganizationRoleId:   dest.UniqueId,
				}

				err := table.RoleAssignment.INSERT().
					MODEL(roleAssignment).
					RETURNING(table.RoleAssignment.AllColumns).
					QueryContext(context.Request().Context(), context.App.Db, &roleAssignmentDest)

				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}

				if roleAssignmentDest.UniqueId == uuid.Nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Error assigning role")
				} else {
					return context.String(http.StatusOK, "OK")
				}
			}
		}

	case api_types.Remove:
		{
			if dest.Assignment.UniqueId == uuid.Nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Role not assigned to the member")
			} else {
				_, err := table.RoleAssignment.DELETE().
					WHERE(table.RoleAssignment.UniqueId.EQ(UUID(dest.UniqueId))).
					ExecContext(context.Request().Context(), context.App.Db)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}

				return context.String(http.StatusOK, "OK")
			}
		}
	}

	return context.String(http.StatusOK, "OK")
}

func syncTemplates(context interfaces.CustomContext) error {
	return context.String(http.StatusOK, "OK")
}

func syncMobileNumbers(context interfaces.CustomContext) error {
	return context.String(http.StatusOK, "OK")
}

func VerifyAccessToOrganization(context interfaces.CustomContext, userId, organizationId string) bool {
	orgQuery := SELECT(table.OrganizationMember.AllColumns, table.Organization.AllColumns).
		FROM(table.OrganizationMember.
			LEFT_JOIN(table.Organization, table.Organization.UniqueId.EQ(table.OrganizationMember.OrganizationId)),
		).
		WHERE(table.OrganizationMember.UserId.EQ(String(userId)).
			AND(table.OrganizationMember.OrganizationId.EQ(String(organizationId))))

	var dest struct {
		model.OrganizationMember
		Organization model.Organization
	}

	err := orgQuery.Query(context.App.Db, &dest)

	if err != nil {
		return false
	}

	if dest.Organization.UniqueId.String() == "" {
		return false
	}

	return true
}
