// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/store"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyStoreAddStickersToFavoriteRequest(r *requests.StoreAddStickersToFavoriteRequest) {
	lStickerIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithStickerIds(random.RandIntArr(lStickerIds))
}

func TestVKStoreAddStickersToFavoriteSuccess(t *testing.T) {
	params := requests.NewStoreAddStickersToFavoriteRequest()
	fillRandomlyStoreAddStickersToFavoriteRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "store.addStickersToFavorite", params.Params(), expectedJSON))
	resp, err := vk.Api.Store.StoreAddStickersToFavorite(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKStoreGetFavoriteStickersSuccess(t *testing.T) {
	var expected models.StoreGetFavoriteStickersResponse
	fillRandomlyStoreGetFavoriteStickersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "store.getFavoriteStickers", nil, expectedJSON))
	resp, err := vk.Api.Store.StoreGetFavoriteStickers()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyStoreGetProductsRequest(r *requests.StoreGetProductsRequest) {
	r.WithType(random.RandString())
	r.WithMerchant(random.RandString())
	r.WithSection(random.RandString())
	lProductIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithProductIds(random.RandIntArr(lProductIds))
	lFilters := random.RandIntn(random.MaxArrayLength + 1)
	r.WithFilters(random.RandStringArr(lFilters))
	r.WithExtended(random.RandBool())
}

func TestVKStoreGetProductsSuccess(t *testing.T) {
	params := requests.NewStoreGetProductsRequest()
	fillRandomlyStoreGetProductsRequest(&params)
	var expected models.StoreGetProductsResponse
	fillRandomlyStoreGetProductsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "store.getProducts", params.Params(), expectedJSON))
	resp, err := vk.Api.Store.StoreGetProducts(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyStoreGetStickersKeywordsRequest(r *requests.StoreGetStickersKeywordsRequest) {
	lStickersIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithStickersIds(random.RandIntArr(lStickersIds))
	lProductsIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithProductsIds(random.RandIntArr(lProductsIds))
	r.WithAliases(random.RandBool())
	r.WithAllProducts(random.RandBool())
	r.WithNeedStickers(random.RandBool())
}

func TestVKStoreGetStickersKeywordsSuccess(t *testing.T) {
	params := requests.NewStoreGetStickersKeywordsRequest()
	fillRandomlyStoreGetStickersKeywordsRequest(&params)
	var expected models.StoreGetStickersKeywordsResponse
	fillRandomlyStoreGetStickersKeywordsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "store.getStickersKeywords", params.Params(), expectedJSON))
	resp, err := vk.Api.Store.StoreGetStickersKeywords(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyStoreRemoveStickersFromFavoriteRequest(r *requests.StoreRemoveStickersFromFavoriteRequest) {
	lStickerIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithStickerIds(random.RandIntArr(lStickerIds))
}

func TestVKStoreRemoveStickersFromFavoriteSuccess(t *testing.T) {
	params := requests.NewStoreRemoveStickersFromFavoriteRequest()
	fillRandomlyStoreRemoveStickersFromFavoriteRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "store.removeStickersFromFavorite", params.Params(), expectedJSON))
	resp, err := vk.Api.Store.StoreRemoveStickersFromFavorite(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
