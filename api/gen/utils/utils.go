// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Utils struct {
	api *api.API
}

func NewUtils(api *api.API) *Utils {
	return &Utils{
		api: api,
	}
}

// UtilsCheckLink Checks whether a link is blocked in VK.
type UtilsCheckLinkRequest api.Params

func NewUtilsCheckLinkRequest() UtilsCheckLinkRequest {
	params := make(UtilsCheckLinkRequest, 2)
	return params
}

func (u UtilsCheckLinkRequest) WithUrl(u_url string) UtilsCheckLinkRequest {
	u["url"] = u_url
	return u
}

func (u UtilsCheckLinkRequest) Params() api.Params {
	return api.Params(u)
}

// May execute with listed access token types:
//
//	[ user, group, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/utils.checkLink
func (u *Utils) UtilsCheckLink(params ...api.MethodParams) (resp models.UtilsCheckLinkResponse, err error) {
	req := api.NewRequest[models.UtilsCheckLinkResponse](u.api)

	res, err := req.Execute("utils.checkLink", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// UtilsDeleteFromLastShortened Deletes shortened link from user's list.
type UtilsDeleteFromLastShortenedRequest api.Params

func NewUtilsDeleteFromLastShortenedRequest() UtilsDeleteFromLastShortenedRequest {
	params := make(UtilsDeleteFromLastShortenedRequest, 2)
	return params
}

func (u UtilsDeleteFromLastShortenedRequest) WithKey(u_key string) UtilsDeleteFromLastShortenedRequest {
	u["key"] = u_key
	return u
}

func (u UtilsDeleteFromLastShortenedRequest) Params() api.Params {
	return api.Params(u)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/utils.deleteFromLastShortened
func (u *Utils) UtilsDeleteFromLastShortened(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](u.api)

	res, err := req.Execute("utils.deleteFromLastShortened", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// UtilsGetLastShortenedLinks Returns a list of user's shortened links.
type UtilsGetLastShortenedLinksRequest api.Params

func NewUtilsGetLastShortenedLinksRequest() UtilsGetLastShortenedLinksRequest {
	params := make(UtilsGetLastShortenedLinksRequest, 3)
	return params
}

func (u UtilsGetLastShortenedLinksRequest) WithCount(u_count int) UtilsGetLastShortenedLinksRequest {
	u["count"] = u_count
	return u
}

func (u UtilsGetLastShortenedLinksRequest) WithOffset(u_offset int) UtilsGetLastShortenedLinksRequest {
	u["offset"] = u_offset
	return u
}

func (u UtilsGetLastShortenedLinksRequest) Params() api.Params {
	return api.Params(u)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/utils.getLastShortenedLinks
func (u *Utils) UtilsGetLastShortenedLinks(params ...api.MethodParams) (resp models.UtilsGetLastShortenedLinksResponse, err error) {
	req := api.NewRequest[models.UtilsGetLastShortenedLinksResponse](u.api)

	res, err := req.Execute("utils.getLastShortenedLinks", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// UtilsGetLinkStats Returns stats data for shortened link.
type UtilsGetLinkStatsRequest api.Params

func NewUtilsGetLinkStatsRequest() UtilsGetLinkStatsRequest {
	params := make(UtilsGetLinkStatsRequest, 8)
	return params
}

func (u UtilsGetLinkStatsRequest) WithKey(u_key string) UtilsGetLinkStatsRequest {
	u["key"] = u_key
	return u
}

func (u UtilsGetLinkStatsRequest) WithSource(u_source string) UtilsGetLinkStatsRequest {
	u["source"] = u_source
	return u
}

func (u UtilsGetLinkStatsRequest) WithAccessKey(u_access_key string) UtilsGetLinkStatsRequest {
	u["access_key"] = u_access_key
	return u
}

func (u UtilsGetLinkStatsRequest) WithInterval(u_interval string) UtilsGetLinkStatsRequest {
	u["interval"] = u_interval
	return u
}

func (u UtilsGetLinkStatsRequest) WithIntervalsCount(u_intervals_count int) UtilsGetLinkStatsRequest {
	u["intervals_count"] = u_intervals_count
	return u
}

func (u UtilsGetLinkStatsRequest) WithExtended(u_extended bool) UtilsGetLinkStatsRequest {
	u["extended"] = u_extended
	return u
}

func (u UtilsGetLinkStatsRequest) Params() api.Params {
	return api.Params(u)
}

// May execute with listed access token types:
//
//	[ user, group, service ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_NotFound ]
//
// https://dev.vk.com/method/utils.getLinkStats
func (u *Utils) UtilsGetLinkStats(params ...api.MethodParams) (resp models.UtilsGetLinkStatsResponse, err error) {
	req := api.NewRequest[models.UtilsGetLinkStatsResponse](u.api)

	res, err := req.Execute("utils.getLinkStats", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// UtilsGetLinkStatsExtended Returns stats data for shortened link.
// May execute with listed access token types:
//
//	[ user, group, service ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_NotFound ]
//
// https://dev.vk.com/method/utils.getLinkStats
func (u *Utils) UtilsGetLinkStatsExtended(params ...api.MethodParams) (resp models.UtilsGetLinkStatsExtendedResponse, err error) {
	req := api.NewRequest[models.UtilsGetLinkStatsExtendedResponse](u.api)

	res, err := req.Execute("utils.getLinkStats", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// UtilsGetServerTime Returns the current time of the VK server.
// May execute with listed access token types:
//
//	[ user, group, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/utils.getServerTime
func (u *Utils) UtilsGetServerTime(params ...api.MethodParams) (resp models.UtilsGetServerTimeResponse, err error) {
	req := api.NewRequest[models.UtilsGetServerTimeResponse](u.api)

	res, err := req.Execute("utils.getServerTime", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// UtilsGetShortLink Allows to receive a link shortened via vk.cc.
type UtilsGetShortLinkRequest api.Params

func NewUtilsGetShortLinkRequest() UtilsGetShortLinkRequest {
	params := make(UtilsGetShortLinkRequest, 3)
	return params
}

func (u UtilsGetShortLinkRequest) WithUrl(u_url string) UtilsGetShortLinkRequest {
	u["url"] = u_url
	return u
}

func (u UtilsGetShortLinkRequest) WithPrivate(u_private bool) UtilsGetShortLinkRequest {
	u["private"] = u_private
	return u
}

func (u UtilsGetShortLinkRequest) Params() api.Params {
	return api.Params(u)
}

// May execute with listed access token types:
//
//	[ user, group, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/utils.getShortLink
func (u *Utils) UtilsGetShortLink(params ...api.MethodParams) (resp models.UtilsGetShortLinkResponse, err error) {
	req := api.NewRequest[models.UtilsGetShortLinkResponse](u.api)

	res, err := req.Execute("utils.getShortLink", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// UtilsResolveScreenName Detects a type of object (e.g., user, community, application) and its ID by screen name.
type UtilsResolveScreenNameRequest api.Params

func NewUtilsResolveScreenNameRequest() UtilsResolveScreenNameRequest {
	params := make(UtilsResolveScreenNameRequest, 2)
	return params
}

func (u UtilsResolveScreenNameRequest) WithScreenName(u_screen_name string) UtilsResolveScreenNameRequest {
	u["screen_name"] = u_screen_name
	return u
}

func (u UtilsResolveScreenNameRequest) Params() api.Params {
	return api.Params(u)
}

// May execute with listed access token types:
//
//	[ user, group, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/utils.resolveScreenName
func (u *Utils) UtilsResolveScreenName(params ...api.MethodParams) (resp models.UtilsResolveScreenNameResponse, err error) {
	req := api.NewRequest[models.UtilsResolveScreenNameResponse](u.api)

	res, err := req.Execute("utils.resolveScreenName", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
