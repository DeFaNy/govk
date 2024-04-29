// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Streaming struct {
	api *api.API
}

func NewStreaming(api *api.API) *Streaming {
	return &Streaming{
		api: api,
	}
}

// StreamingGetServerUrl Allows to receive data for the connection to Streaming API.
// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/streaming.getServerUrl
func (s *Streaming) StreamingGetServerUrl(params ...api.MethodParams) (resp models.StreamingGetServerUrlResponse, err error) {
	req := api.NewRequest[models.StreamingGetServerUrlResponse](s.api)

	res, err := req.Execute("streaming.getServerUrl", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// StreamingSetSettings ...
type StreamingSetSettingsRequest api.Params

func NewStreamingSetSettingsRequest() StreamingSetSettingsRequest {
	params := make(StreamingSetSettingsRequest, 2)
	return params
}

func (s StreamingSetSettingsRequest) WithMonthlyTier(s_monthly_tier string) StreamingSetSettingsRequest {
	s["monthly_tier"] = s_monthly_tier
	return s
}

func (s StreamingSetSettingsRequest) Params() api.Params {
	return api.Params(s)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/streaming.setSettings
func (s *Streaming) StreamingSetSettings(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](s.api)

	res, err := req.Execute("streaming.setSettings", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
