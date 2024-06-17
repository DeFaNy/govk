// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/newsfeed"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyNewsfeedAddBanRequest(r *requests.NewsfeedAddBanRequest) {
	lUserIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithUserIds(random.IntArr(lUserIds))
	lGroupIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithGroupIds(random.IntArr(lGroupIds))
}

func TestVKNewsfeedAddBanSuccess(t *testing.T) {
	params := requests.NewNewsfeedAddBanRequest()
	fillRandomlyNewsfeedAddBanRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.addBan", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedAddBan(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedDeleteBanRequest(r *requests.NewsfeedDeleteBanRequest) {
	lUserIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithUserIds(random.IntArr(lUserIds))
	lGroupIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithGroupIds(random.IntArr(lGroupIds))
}

func TestVKNewsfeedDeleteBanSuccess(t *testing.T) {
	params := requests.NewNewsfeedDeleteBanRequest()
	fillRandomlyNewsfeedDeleteBanRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.deleteBan", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedDeleteBan(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedDeleteListRequest(r *requests.NewsfeedDeleteListRequest) {
	r.WithListId(random.Int())
}

func TestVKNewsfeedDeleteListSuccess(t *testing.T) {
	params := requests.NewNewsfeedDeleteListRequest()
	fillRandomlyNewsfeedDeleteListRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.deleteList", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedDeleteList(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedGetRequest(r *requests.NewsfeedGetRequest) {
	Filters := new([]models.NewsfeedNewsfeedItemType)
	lFilters := random.IntDiapason(random.MaxArrayLength + 1)
	*Filters = make([]models.NewsfeedNewsfeedItemType, lFilters)
	for i0 := 0; i0 < lFilters; i0++ {
		fillRandomlyNewsfeedNewsfeedItemType(&(*Filters)[i0])
	}
	r.WithFilters(*Filters)
	r.WithReturnBanned(random.Bool())
	r.WithStartTime(random.Int())
	r.WithEndTime(random.Int())
	r.WithMaxPhotos(random.Int())
	r.WithSourceIds(random.String())
	r.WithStartFrom(random.String())
	r.WithCount(random.Int())
	Fields := new([]models.BaseUserGroupFields)
	lFields := random.IntDiapason(random.MaxArrayLength + 1)
	*Fields = make([]models.BaseUserGroupFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyBaseUserGroupFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
	r.WithSection(random.String())
}

func TestVKNewsfeedGetSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetRequest()
	fillRandomlyNewsfeedGetRequest(&params)
	var expected models.NewsfeedGenericResponse
	fillRandomlyNewsfeedGenericResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedGetBannedRequest(r *requests.NewsfeedGetBannedRequest) {
	r.WithExtended(random.Bool())
	Fields := new([]models.UsersFields)
	lFields := random.IntDiapason(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
	r.WithNameCase(random.String())
}

func TestVKNewsfeedGetBannedSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetBannedRequest()
	fillRandomlyNewsfeedGetBannedRequest(&params)
	params.WithExtended(false)
	var expected models.NewsfeedGetBannedResponse
	fillRandomlyNewsfeedGetBannedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.getBanned", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGetBanned(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKNewsfeedGetBannedExtendedSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetBannedRequest()
	fillRandomlyNewsfeedGetBannedRequest(&params)
	params.WithExtended(true)
	var expected models.NewsfeedGetBannedExtendedResponse
	fillRandomlyNewsfeedGetBannedExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.getBanned", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGetBannedExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedGetCommentsRequest(r *requests.NewsfeedGetCommentsRequest) {
	r.WithCount(random.Int())
	Filters := new([]models.NewsfeedCommentsFilters)
	lFilters := random.IntDiapason(random.MaxArrayLength + 1)
	*Filters = make([]models.NewsfeedCommentsFilters, lFilters)
	for i0 := 0; i0 < lFilters; i0++ {
		fillRandomlyNewsfeedCommentsFilters(&(*Filters)[i0])
	}
	r.WithFilters(*Filters)
	r.WithReposts(random.String())
	r.WithStartTime(random.Int())
	r.WithEndTime(random.Int())
	r.WithLastCommentsCount(random.Int())
	r.WithStartFrom(random.String())
	Fields := new([]models.BaseUserGroupFields)
	lFields := random.IntDiapason(random.MaxArrayLength + 1)
	*Fields = make([]models.BaseUserGroupFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyBaseUserGroupFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
}

func TestVKNewsfeedGetCommentsSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetCommentsRequest()
	fillRandomlyNewsfeedGetCommentsRequest(&params)
	var expected models.NewsfeedGetCommentsResponse
	fillRandomlyNewsfeedGetCommentsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.getComments", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGetComments(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedGetListsRequest(r *requests.NewsfeedGetListsRequest) {
	lListIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithListIds(random.IntArr(lListIds))
	r.WithExtended(random.Bool())
}

func TestVKNewsfeedGetListsSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetListsRequest()
	fillRandomlyNewsfeedGetListsRequest(&params)
	params.WithExtended(false)
	var expected models.NewsfeedGetListsResponse
	fillRandomlyNewsfeedGetListsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.getLists", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGetLists(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKNewsfeedGetListsExtendedSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetListsRequest()
	fillRandomlyNewsfeedGetListsRequest(&params)
	params.WithExtended(true)
	var expected models.NewsfeedGetListsExtendedResponse
	fillRandomlyNewsfeedGetListsExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.getLists", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGetListsExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedGetMentionsRequest(r *requests.NewsfeedGetMentionsRequest) {
	r.WithOwnerId(random.Int())
	r.WithStartTime(random.Int())
	r.WithEndTime(random.Int())
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
}

func TestVKNewsfeedGetMentionsSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetMentionsRequest()
	fillRandomlyNewsfeedGetMentionsRequest(&params)
	var expected models.NewsfeedGetMentionsResponse
	fillRandomlyNewsfeedGetMentionsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.getMentions", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGetMentions(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedGetRecommendedRequest(r *requests.NewsfeedGetRecommendedRequest) {
	r.WithStartTime(random.Int())
	r.WithEndTime(random.Int())
	r.WithMaxPhotos(random.Int())
	r.WithStartFrom(random.String())
	r.WithCount(random.Int())
	Fields := new([]models.BaseUserGroupFields)
	lFields := random.IntDiapason(random.MaxArrayLength + 1)
	*Fields = make([]models.BaseUserGroupFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyBaseUserGroupFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
}

func TestVKNewsfeedGetRecommendedSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetRecommendedRequest()
	fillRandomlyNewsfeedGetRecommendedRequest(&params)
	var expected models.NewsfeedGenericResponse
	fillRandomlyNewsfeedGenericResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.getRecommended", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGetRecommended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedGetSuggestedSourcesRequest(r *requests.NewsfeedGetSuggestedSourcesRequest) {
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
	r.WithShuffle(random.Bool())
	Fields := new([]models.BaseUserGroupFields)
	lFields := random.IntDiapason(random.MaxArrayLength + 1)
	*Fields = make([]models.BaseUserGroupFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyBaseUserGroupFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
}

func TestVKNewsfeedGetSuggestedSourcesSuccess(t *testing.T) {
	params := requests.NewNewsfeedGetSuggestedSourcesRequest()
	fillRandomlyNewsfeedGetSuggestedSourcesRequest(&params)
	var expected models.NewsfeedGetSuggestedSourcesResponse
	fillRandomlyNewsfeedGetSuggestedSourcesResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.getSuggestedSources", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedGetSuggestedSources(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedIgnoreItemRequest(r *requests.NewsfeedIgnoreItemRequest) {
	Type := new(models.NewsfeedIgnoreItemType)
	fillRandomlyNewsfeedIgnoreItemType(Type)
	r.WithType(*Type)
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
}

func TestVKNewsfeedIgnoreItemSuccess(t *testing.T) {
	params := requests.NewNewsfeedIgnoreItemRequest()
	fillRandomlyNewsfeedIgnoreItemRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.ignoreItem", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedIgnoreItem(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedSaveListRequest(r *requests.NewsfeedSaveListRequest) {
	r.WithListId(random.Int())
	r.WithTitle(random.String())
	lSourceIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithSourceIds(random.IntArr(lSourceIds))
	r.WithNoReposts(random.Bool())
}

func TestVKNewsfeedSaveListSuccess(t *testing.T) {
	params := requests.NewNewsfeedSaveListRequest()
	fillRandomlyNewsfeedSaveListRequest(&params)
	var expected models.NewsfeedSaveListResponse
	fillRandomlyNewsfeedSaveListResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.saveList", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedSaveList(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedSearchRequest(r *requests.NewsfeedSearchRequest) {
	r.WithQ(random.String())
	r.WithExtended(random.Bool())
	r.WithCount(random.Int())
	r.WithLatitude(random.MustFloat())
	r.WithLongitude(random.MustFloat())
	r.WithStartTime(random.Int())
	r.WithEndTime(random.Int())
	r.WithStartFrom(random.String())
	Fields := new([]models.BaseUserGroupFields)
	lFields := random.IntDiapason(random.MaxArrayLength + 1)
	*Fields = make([]models.BaseUserGroupFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyBaseUserGroupFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
}

func TestVKNewsfeedSearchSuccess(t *testing.T) {
	params := requests.NewNewsfeedSearchRequest()
	fillRandomlyNewsfeedSearchRequest(&params)
	params.WithExtended(false)
	var expected models.NewsfeedSearchResponse
	fillRandomlyNewsfeedSearchResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.search", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedSearch(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKNewsfeedSearchExtendedSuccess(t *testing.T) {
	params := requests.NewNewsfeedSearchRequest()
	fillRandomlyNewsfeedSearchRequest(&params)
	params.WithExtended(true)
	var expected models.NewsfeedSearchExtendedResponse
	fillRandomlyNewsfeedSearchExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.search", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedSearchExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedUnignoreItemRequest(r *requests.NewsfeedUnignoreItemRequest) {
	Type := new(models.NewsfeedIgnoreItemType)
	fillRandomlyNewsfeedIgnoreItemType(Type)
	r.WithType(*Type)
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
	r.WithTrackCode(random.String())
}

func TestVKNewsfeedUnignoreItemSuccess(t *testing.T) {
	params := requests.NewNewsfeedUnignoreItemRequest()
	fillRandomlyNewsfeedUnignoreItemRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.unignoreItem", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedUnignoreItem(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNewsfeedUnsubscribeRequest(r *requests.NewsfeedUnsubscribeRequest) {
	r.WithType(random.String())
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
}

func TestVKNewsfeedUnsubscribeSuccess(t *testing.T) {
	params := requests.NewNewsfeedUnsubscribeRequest()
	fillRandomlyNewsfeedUnsubscribeRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "newsfeed.unsubscribe", params.Params(), expectedJSON))
	resp, err := vk.Api.Newsfeed.NewsfeedUnsubscribe(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
