// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/apps"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVKAppsDeleteAppRequestsSuccess(t *testing.T) {
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.deleteAppRequests", nil, expectedJSON))
	resp, err := vk.Api.Apps.AppsDeleteAppRequests()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsGetRequest(r *requests.AppsGetRequest) {
	r.WithAppId(random.RandInt())
	lAppIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithAppIds(random.RandStringArr(lAppIds))
	r.WithPlatform(random.RandString())
	r.WithExtended(random.RandBool())
	r.WithReturnFriends(random.RandBool())
	Fields := new([]models.UsersFields)
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
	r.WithNameCase(random.RandString())
}

func TestVKAppsGetSuccess(t *testing.T) {
	params := requests.NewAppsGetRequest()
	fillRandomlyAppsGetRequest(&params)
	var expected models.AppsGetResponse
	fillRandomlyAppsGetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsGetCatalogRequest(r *requests.AppsGetCatalogRequest) {
	r.WithSort(random.RandString())
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	r.WithPlatform(random.RandString())
	r.WithExtended(random.RandBool())
	r.WithReturnFriends(random.RandBool())
	Fields := new([]models.UsersFields)
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
	r.WithNameCase(random.RandString())
	r.WithQ(random.RandString())
	r.WithGenreId(random.RandInt())
	r.WithFilter(random.RandString())
}

func TestVKAppsGetCatalogSuccess(t *testing.T) {
	params := requests.NewAppsGetCatalogRequest()
	fillRandomlyAppsGetCatalogRequest(&params)
	var expected models.AppsGetCatalogResponse
	fillRandomlyAppsGetCatalogResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.getCatalog", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGetCatalog(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsGetFriendsListRequest(r *requests.AppsGetFriendsListRequest) {
	r.WithExtended(random.RandBool())
	r.WithCount(random.RandInt())
	r.WithOffset(random.RandInt())
	r.WithType(random.RandString())
	Fields := new([]models.UsersFields)
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
}

func TestVKAppsGetFriendsListSuccess(t *testing.T) {
	params := requests.NewAppsGetFriendsListRequest()
	fillRandomlyAppsGetFriendsListRequest(&params)
	params.WithExtended(false)
	var expected models.AppsGetFriendsListResponse
	fillRandomlyAppsGetFriendsListResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.getFriendsList", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGetFriendsList(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKAppsGetFriendsListExtendedSuccess(t *testing.T) {
	params := requests.NewAppsGetFriendsListRequest()
	fillRandomlyAppsGetFriendsListRequest(&params)
	params.WithExtended(true)
	var expected models.AppsGetFriendsListExtendedResponse
	fillRandomlyAppsGetFriendsListExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.getFriendsList", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGetFriendsListExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsGetLeaderboardRequest(r *requests.AppsGetLeaderboardRequest) {
	r.WithType(random.RandString())
	r.WithGlobal(random.RandBool())
	r.WithExtended(random.RandBool())
}

func TestVKAppsGetLeaderboardSuccess(t *testing.T) {
	params := requests.NewAppsGetLeaderboardRequest()
	fillRandomlyAppsGetLeaderboardRequest(&params)
	params.WithExtended(false)
	var expected models.AppsGetLeaderboardResponse
	fillRandomlyAppsGetLeaderboardResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.getLeaderboard", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGetLeaderboard(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKAppsGetLeaderboardExtendedSuccess(t *testing.T) {
	params := requests.NewAppsGetLeaderboardRequest()
	fillRandomlyAppsGetLeaderboardRequest(&params)
	params.WithExtended(true)
	var expected models.AppsGetLeaderboardExtendedResponse
	fillRandomlyAppsGetLeaderboardExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.getLeaderboard", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGetLeaderboardExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsGetMiniAppPoliciesRequest(r *requests.AppsGetMiniAppPoliciesRequest) {
	r.WithAppId(random.RandInt())
}

func TestVKAppsGetMiniAppPoliciesSuccess(t *testing.T) {
	params := requests.NewAppsGetMiniAppPoliciesRequest()
	fillRandomlyAppsGetMiniAppPoliciesRequest(&params)
	var expected models.AppsGetMiniAppPoliciesResponse
	fillRandomlyAppsGetMiniAppPoliciesResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.getMiniAppPolicies", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGetMiniAppPolicies(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsGetScopesRequest(r *requests.AppsGetScopesRequest) {
	r.WithType(random.RandString())
}

func TestVKAppsGetScopesSuccess(t *testing.T) {
	params := requests.NewAppsGetScopesRequest()
	fillRandomlyAppsGetScopesRequest(&params)
	var expected models.AppsGetScopesResponse
	fillRandomlyAppsGetScopesResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.getScopes", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGetScopes(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsGetScoreRequest(r *requests.AppsGetScoreRequest) {
	r.WithUserId(random.RandInt())
}

func TestVKAppsGetScoreSuccess(t *testing.T) {
	params := requests.NewAppsGetScoreRequest()
	fillRandomlyAppsGetScoreRequest(&params)
	var expected models.AppsGetScoreResponse
	fillRandomlyAppsGetScoreResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.getScore", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsGetScore(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsPromoHasActiveGiftRequest(r *requests.AppsPromoHasActiveGiftRequest) {
	r.WithPromoId(random.RandInt())
	r.WithUserId(random.RandInt())
}

func TestVKAppsPromoHasActiveGiftSuccess(t *testing.T) {
	params := requests.NewAppsPromoHasActiveGiftRequest()
	fillRandomlyAppsPromoHasActiveGiftRequest(&params)
	var expected models.BaseBoolResponse
	fillRandomlyBaseBoolResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.promoHasActiveGift", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsPromoHasActiveGift(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsPromoUseGiftRequest(r *requests.AppsPromoUseGiftRequest) {
	r.WithPromoId(random.RandInt())
	r.WithUserId(random.RandInt())
}

func TestVKAppsPromoUseGiftSuccess(t *testing.T) {
	params := requests.NewAppsPromoUseGiftRequest()
	fillRandomlyAppsPromoUseGiftRequest(&params)
	var expected models.BaseBoolResponse
	fillRandomlyBaseBoolResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.promoUseGift", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsPromoUseGift(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppsSendRequestRequest(r *requests.AppsSendRequestRequest) {
	r.WithUserId(random.RandInt())
	r.WithText(random.RandString())
	r.WithType(random.RandString())
	r.WithName(random.RandString())
	r.WithKey(random.RandString())
	r.WithSeparate(random.RandBool())
}

func TestVKAppsSendRequestSuccess(t *testing.T) {
	params := requests.NewAppsSendRequestRequest()
	fillRandomlyAppsSendRequestRequest(&params)
	var expected models.AppsSendRequestResponse
	fillRandomlyAppsSendRequestResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "apps.sendRequest", params.Params(), expectedJSON))
	resp, err := vk.Api.Apps.AppsSendRequest(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
