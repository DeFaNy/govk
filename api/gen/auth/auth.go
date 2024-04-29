// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Auth struct {
	api *api.API
}

func NewAuth(api *api.API) *Auth {
	return &Auth{
		api: api,
	}
}

// AuthRestore Allows to restore account access using a code received via SMS. " This method is only available for apps with [vk.com/dev/auth_direct|Direct authorization] access. "
type AuthRestoreRequest api.Params

func NewAuthRestoreRequest() AuthRestoreRequest {
	params := make(AuthRestoreRequest, 3)
	return params
}

func (a AuthRestoreRequest) WithPhone(a_phone string) AuthRestoreRequest {
	a["phone"] = a_phone
	return a
}

func (a AuthRestoreRequest) WithLastName(a_last_name string) AuthRestoreRequest {
	a["last_name"] = a_last_name
	return a
}

func (a AuthRestoreRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//
//	[ user, open, service ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AuthFloodError ]
//
// https://dev.vk.com/method/auth.restore
func (a *Auth) AuthRestore(params ...api.MethodParams) (resp models.AuthRestoreResponse, err error) {
	req := api.NewRequest[models.AuthRestoreResponse](a.api)

	res, err := req.Execute("auth.restore", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}