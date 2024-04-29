// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Pages struct {
	api *api.API
}

func NewPages(api *api.API) *Pages {
	return &Pages{
		api: api,
	}
}

// PagesClearCache Allows to clear the cache of particular 'external' pages which may be attached to VK posts.
type PagesClearCacheRequest api.Params

func NewPagesClearCacheRequest() PagesClearCacheRequest {
	params := make(PagesClearCacheRequest, 2)
	return params
}

func (p PagesClearCacheRequest) WithUrl(p_url string) PagesClearCacheRequest {
	p["url"] = p_url
	return p
}

func (p PagesClearCacheRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/pages.clearCache
func (p *Pages) PagesClearCache(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](p.api)

	res, err := req.Execute("pages.clearCache", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PagesGet Returns information about a wiki page.
type PagesGetRequest api.Params

func NewPagesGetRequest() PagesGetRequest {
	params := make(PagesGetRequest, 8)
	return params
}

func (p PagesGetRequest) WithOwnerId(p_owner_id int) PagesGetRequest {
	p["owner_id"] = p_owner_id
	return p
}

func (p PagesGetRequest) WithPageId(p_page_id int) PagesGetRequest {
	p["page_id"] = p_page_id
	return p
}

func (p PagesGetRequest) WithGlobal(p_global bool) PagesGetRequest {
	p["global"] = p_global
	return p
}

func (p PagesGetRequest) WithSitePreview(p_site_preview bool) PagesGetRequest {
	p["site_preview"] = p_site_preview
	return p
}

func (p PagesGetRequest) WithTitle(p_title string) PagesGetRequest {
	p["title"] = p_title
	return p
}

func (p PagesGetRequest) WithNeedSource(p_need_source bool) PagesGetRequest {
	p["need_source"] = p_need_source
	return p
}

func (p PagesGetRequest) WithNeedHtml(p_need_html bool) PagesGetRequest {
	p["need_html"] = p_need_html
	return p
}

func (p PagesGetRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/pages.get
func (p *Pages) PagesGet(params ...api.MethodParams) (resp models.PagesGetResponse, err error) {
	req := api.NewRequest[models.PagesGetResponse](p.api)

	res, err := req.Execute("pages.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PagesGetHistory Returns a list of all previous versions of a wiki page.
type PagesGetHistoryRequest api.Params

func NewPagesGetHistoryRequest() PagesGetHistoryRequest {
	params := make(PagesGetHistoryRequest, 4)
	return params
}

func (p PagesGetHistoryRequest) WithPageId(p_page_id int) PagesGetHistoryRequest {
	p["page_id"] = p_page_id
	return p
}

func (p PagesGetHistoryRequest) WithGroupId(p_group_id int) PagesGetHistoryRequest {
	p["group_id"] = p_group_id
	return p
}

func (p PagesGetHistoryRequest) WithUserId(p_user_id int) PagesGetHistoryRequest {
	p["user_id"] = p_user_id
	return p
}

func (p PagesGetHistoryRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AccessPage, Error_ParamPageId ]
//
// https://dev.vk.com/method/pages.getHistory
func (p *Pages) PagesGetHistory(params ...api.MethodParams) (resp models.PagesGetHistoryResponse, err error) {
	req := api.NewRequest[models.PagesGetHistoryResponse](p.api)

	res, err := req.Execute("pages.getHistory", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PagesGetTitles Returns a list of wiki pages in a group.
type PagesGetTitlesRequest api.Params

func NewPagesGetTitlesRequest() PagesGetTitlesRequest {
	params := make(PagesGetTitlesRequest, 2)
	return params
}

func (p PagesGetTitlesRequest) WithGroupId(p_group_id int) PagesGetTitlesRequest {
	p["group_id"] = p_group_id
	return p
}

func (p PagesGetTitlesRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AccessPage ]
//
// https://dev.vk.com/method/pages.getTitles
func (p *Pages) PagesGetTitles(params ...api.MethodParams) (resp models.PagesGetTitlesResponse, err error) {
	req := api.NewRequest[models.PagesGetTitlesResponse](p.api)

	res, err := req.Execute("pages.getTitles", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PagesGetVersion Returns the text of one of the previous versions of a wiki page.
type PagesGetVersionRequest api.Params

func NewPagesGetVersionRequest() PagesGetVersionRequest {
	params := make(PagesGetVersionRequest, 5)
	return params
}

func (p PagesGetVersionRequest) WithVersionId(p_version_id int) PagesGetVersionRequest {
	p["version_id"] = p_version_id
	return p
}

func (p PagesGetVersionRequest) WithGroupId(p_group_id int) PagesGetVersionRequest {
	p["group_id"] = p_group_id
	return p
}

func (p PagesGetVersionRequest) WithUserId(p_user_id int) PagesGetVersionRequest {
	p["user_id"] = p_user_id
	return p
}

func (p PagesGetVersionRequest) WithNeedHtml(p_need_html bool) PagesGetVersionRequest {
	p["need_html"] = p_need_html
	return p
}

func (p PagesGetVersionRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AccessPage ]
//
// https://dev.vk.com/method/pages.getVersion
func (p *Pages) PagesGetVersion(params ...api.MethodParams) (resp models.PagesGetVersionResponse, err error) {
	req := api.NewRequest[models.PagesGetVersionResponse](p.api)

	res, err := req.Execute("pages.getVersion", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PagesParseWiki Returns HTML representation of the wiki markup.
type PagesParseWikiRequest api.Params

func NewPagesParseWikiRequest() PagesParseWikiRequest {
	params := make(PagesParseWikiRequest, 3)
	return params
}

func (p PagesParseWikiRequest) WithText(p_text string) PagesParseWikiRequest {
	p["text"] = p_text
	return p
}

func (p PagesParseWikiRequest) WithGroupId(p_group_id int) PagesParseWikiRequest {
	p["group_id"] = p_group_id
	return p
}

func (p PagesParseWikiRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/pages.parseWiki
func (p *Pages) PagesParseWiki(params ...api.MethodParams) (resp models.PagesParseWikiResponse, err error) {
	req := api.NewRequest[models.PagesParseWikiResponse](p.api)

	res, err := req.Execute("pages.parseWiki", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PagesSave Saves the text of a wiki page.
type PagesSaveRequest api.Params

func NewPagesSaveRequest() PagesSaveRequest {
	params := make(PagesSaveRequest, 6)
	return params
}

func (p PagesSaveRequest) WithText(p_text string) PagesSaveRequest {
	p["text"] = p_text
	return p
}

func (p PagesSaveRequest) WithPageId(p_page_id int) PagesSaveRequest {
	p["page_id"] = p_page_id
	return p
}

func (p PagesSaveRequest) WithGroupId(p_group_id int) PagesSaveRequest {
	p["group_id"] = p_group_id
	return p
}

func (p PagesSaveRequest) WithUserId(p_user_id int) PagesSaveRequest {
	p["user_id"] = p_user_id
	return p
}

func (p PagesSaveRequest) WithTitle(p_title string) PagesSaveRequest {
	p["title"] = p_title
	return p
}

func (p PagesSaveRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AccessPage, Error_ParamPageId, Error_ParamTitle ]
//
// https://dev.vk.com/method/pages.save
func (p *Pages) PagesSave(params ...api.MethodParams) (resp models.PagesSaveResponse, err error) {
	req := api.NewRequest[models.PagesSaveResponse](p.api)

	res, err := req.Execute("pages.save", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PagesSaveAccess Saves modified read and edit access settings for a wiki page.
type PagesSaveAccessRequest api.Params

func NewPagesSaveAccessRequest() PagesSaveAccessRequest {
	params := make(PagesSaveAccessRequest, 6)
	return params
}

func (p PagesSaveAccessRequest) WithPageId(p_page_id int) PagesSaveAccessRequest {
	p["page_id"] = p_page_id
	return p
}

func (p PagesSaveAccessRequest) WithGroupId(p_group_id int) PagesSaveAccessRequest {
	p["group_id"] = p_group_id
	return p
}

func (p PagesSaveAccessRequest) WithUserId(p_user_id int) PagesSaveAccessRequest {
	p["user_id"] = p_user_id
	return p
}

func (p PagesSaveAccessRequest) WithView(p_view int) PagesSaveAccessRequest {
	p["view"] = p_view
	return p
}

func (p PagesSaveAccessRequest) WithEdit(p_edit int) PagesSaveAccessRequest {
	p["edit"] = p_edit
	return p
}

func (p PagesSaveAccessRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AccessPage, Error_ParamPageId ]
//
// https://dev.vk.com/method/pages.saveAccess
func (p *Pages) PagesSaveAccess(params ...api.MethodParams) (resp models.PagesSaveAccessResponse, err error) {
	req := api.NewRequest[models.PagesSaveAccessResponse](p.api)

	res, err := req.Execute("pages.saveAccess", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}