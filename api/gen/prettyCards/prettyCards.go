// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type PrettyCards struct {
	api *api.API
}

func NewPrettyCards(api *api.API) *PrettyCards {
	return &PrettyCards{
		api: api,
	}
}

// PrettyCardsCreate ...
type PrettyCardsCreateRequest api.Params

func NewPrettyCardsCreateRequest() PrettyCardsCreateRequest {
	params := make(PrettyCardsCreateRequest, 8)
	return params
}

func (p PrettyCardsCreateRequest) WithOwnerId(p_owner_id int) PrettyCardsCreateRequest {
	p["owner_id"] = p_owner_id
	return p
}

func (p PrettyCardsCreateRequest) WithPhoto(p_photo string) PrettyCardsCreateRequest {
	p["photo"] = p_photo
	return p
}

func (p PrettyCardsCreateRequest) WithTitle(p_title string) PrettyCardsCreateRequest {
	p["title"] = p_title
	return p
}

func (p PrettyCardsCreateRequest) WithLink(p_link string) PrettyCardsCreateRequest {
	p["link"] = p_link
	return p
}

func (p PrettyCardsCreateRequest) WithPrice(p_price string) PrettyCardsCreateRequest {
	p["price"] = p_price
	return p
}

func (p PrettyCardsCreateRequest) WithPriceOld(p_price_old string) PrettyCardsCreateRequest {
	p["price_old"] = p_price_old
	return p
}

func (p PrettyCardsCreateRequest) WithButton(p_button string) PrettyCardsCreateRequest {
	p["button"] = p_button
	return p
}

func (p PrettyCardsCreateRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_PrettyCardsTooManyCards ]
//
// https://dev.vk.com/method/prettyCards.create
func (p *PrettyCards) PrettyCardsCreate(params ...api.MethodParams) (resp models.PrettyCardsCreateResponse, err error) {
	req := api.NewRequest[models.PrettyCardsCreateResponse](p.api)

	res, err := req.Execute("prettyCards.create", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PrettyCardsDelete ...
type PrettyCardsDeleteRequest api.Params

func NewPrettyCardsDeleteRequest() PrettyCardsDeleteRequest {
	params := make(PrettyCardsDeleteRequest, 3)
	return params
}

func (p PrettyCardsDeleteRequest) WithOwnerId(p_owner_id int) PrettyCardsDeleteRequest {
	p["owner_id"] = p_owner_id
	return p
}

func (p PrettyCardsDeleteRequest) WithCardId(p_card_id int) PrettyCardsDeleteRequest {
	p["card_id"] = p_card_id
	return p
}

func (p PrettyCardsDeleteRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_PrettyCardsCardNotFound, Error_PrettyCardsCardIsConnectedToPost ]
//
// https://dev.vk.com/method/prettyCards.delete
func (p *PrettyCards) PrettyCardsDelete(params ...api.MethodParams) (resp models.PrettyCardsDeleteResponse, err error) {
	req := api.NewRequest[models.PrettyCardsDeleteResponse](p.api)

	res, err := req.Execute("prettyCards.delete", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PrettyCardsEdit ...
type PrettyCardsEditRequest api.Params

func NewPrettyCardsEditRequest() PrettyCardsEditRequest {
	params := make(PrettyCardsEditRequest, 9)
	return params
}

func (p PrettyCardsEditRequest) WithOwnerId(p_owner_id int) PrettyCardsEditRequest {
	p["owner_id"] = p_owner_id
	return p
}

func (p PrettyCardsEditRequest) WithCardId(p_card_id int) PrettyCardsEditRequest {
	p["card_id"] = p_card_id
	return p
}

func (p PrettyCardsEditRequest) WithPhoto(p_photo string) PrettyCardsEditRequest {
	p["photo"] = p_photo
	return p
}

func (p PrettyCardsEditRequest) WithTitle(p_title string) PrettyCardsEditRequest {
	p["title"] = p_title
	return p
}

func (p PrettyCardsEditRequest) WithLink(p_link string) PrettyCardsEditRequest {
	p["link"] = p_link
	return p
}

func (p PrettyCardsEditRequest) WithPrice(p_price string) PrettyCardsEditRequest {
	p["price"] = p_price
	return p
}

func (p PrettyCardsEditRequest) WithPriceOld(p_price_old string) PrettyCardsEditRequest {
	p["price_old"] = p_price_old
	return p
}

func (p PrettyCardsEditRequest) WithButton(p_button string) PrettyCardsEditRequest {
	p["button"] = p_button
	return p
}

func (p PrettyCardsEditRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_PrettyCardsCardNotFound ]
//
// https://dev.vk.com/method/prettyCards.edit
func (p *PrettyCards) PrettyCardsEdit(params ...api.MethodParams) (resp models.PrettyCardsEditResponse, err error) {
	req := api.NewRequest[models.PrettyCardsEditResponse](p.api)

	res, err := req.Execute("prettyCards.edit", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PrettyCardsGet ...
type PrettyCardsGetRequest api.Params

func NewPrettyCardsGetRequest() PrettyCardsGetRequest {
	params := make(PrettyCardsGetRequest, 4)
	return params
}

func (p PrettyCardsGetRequest) WithOwnerId(p_owner_id int) PrettyCardsGetRequest {
	p["owner_id"] = p_owner_id
	return p
}

func (p PrettyCardsGetRequest) WithOffset(p_offset int) PrettyCardsGetRequest {
	p["offset"] = p_offset
	return p
}

func (p PrettyCardsGetRequest) WithCount(p_count int) PrettyCardsGetRequest {
	p["count"] = p_count
	return p
}

func (p PrettyCardsGetRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/prettyCards.get
func (p *PrettyCards) PrettyCardsGet(params ...api.MethodParams) (resp models.PrettyCardsGetResponse, err error) {
	req := api.NewRequest[models.PrettyCardsGetResponse](p.api)

	res, err := req.Execute("prettyCards.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PrettyCardsGetById ...
type PrettyCardsGetByIdRequest api.Params

func NewPrettyCardsGetByIdRequest() PrettyCardsGetByIdRequest {
	params := make(PrettyCardsGetByIdRequest, 3)
	return params
}

func (p PrettyCardsGetByIdRequest) WithOwnerId(p_owner_id int) PrettyCardsGetByIdRequest {
	p["owner_id"] = p_owner_id
	return p
}

func (p PrettyCardsGetByIdRequest) WithCardIds(p_card_ids []int) PrettyCardsGetByIdRequest {
	p["card_ids"] = p_card_ids
	return p
}

func (p PrettyCardsGetByIdRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/prettyCards.getById
func (p *PrettyCards) PrettyCardsGetById(params ...api.MethodParams) (resp models.PrettyCardsGetByIdResponse, err error) {
	req := api.NewRequest[models.PrettyCardsGetByIdResponse](p.api)

	res, err := req.Execute("prettyCards.getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PrettyCardsGetUploadURL ...
// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/prettyCards.getUploadURL
func (p *PrettyCards) PrettyCardsGetUploadURL(params ...api.MethodParams) (resp models.PrettyCardsGetUploadURLResponse, err error) {
	req := api.NewRequest[models.PrettyCardsGetUploadURLResponse](p.api)

	res, err := req.Execute("prettyCards.getUploadURL", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
