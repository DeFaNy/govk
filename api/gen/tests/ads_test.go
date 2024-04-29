// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/ads"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyAdsAddOfficeUsersRequest(r *requests.AdsAddOfficeUsersRequest) {
	r.WithAccountId(random.RandInt())
	Data := new([]models.AdsUserSpecificationCutted)
	lData := random.RandIntn(random.MaxArrayLength + 1)
	*Data = make([]models.AdsUserSpecificationCutted, lData)
	for i0 := 0; i0 < lData; i0++ {
		fillRandomlyAdsUserSpecificationCutted(&(*Data)[i0])
	}
	r.WithData(*Data)
}

func TestVKAdsAddOfficeUsersSuccess(t *testing.T) {
	params := requests.NewAdsAddOfficeUsersRequest()
	fillRandomlyAdsAddOfficeUsersRequest(&params)
	var expected models.AdsAddOfficeUsersResponse
	fillRandomlyAdsAddOfficeUsersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.addOfficeUsers", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsAddOfficeUsers(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsCheckLinkRequest(r *requests.AdsCheckLinkRequest) {
	r.WithAccountId(random.RandInt())
	r.WithLinkType(random.RandString())
	r.WithLinkUrl(random.RandString())
	r.WithCampaignId(random.RandInt())
}

func TestVKAdsCheckLinkSuccess(t *testing.T) {
	params := requests.NewAdsCheckLinkRequest()
	fillRandomlyAdsCheckLinkRequest(&params)
	var expected models.AdsCheckLinkResponse
	fillRandomlyAdsCheckLinkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.checkLink", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsCheckLink(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsCreateAdsRequest(r *requests.AdsCreateAdsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithData(random.RandString())
}

func TestVKAdsCreateAdsSuccess(t *testing.T) {
	params := requests.NewAdsCreateAdsRequest()
	fillRandomlyAdsCreateAdsRequest(&params)
	var expected models.AdsCreateAdsResponse
	fillRandomlyAdsCreateAdsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.createAds", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsCreateAds(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsCreateCampaignsRequest(r *requests.AdsCreateCampaignsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithData(random.RandString())
}

func TestVKAdsCreateCampaignsSuccess(t *testing.T) {
	params := requests.NewAdsCreateCampaignsRequest()
	fillRandomlyAdsCreateCampaignsRequest(&params)
	var expected models.AdsCreateCampaignsResponse
	fillRandomlyAdsCreateCampaignsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.createCampaigns", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsCreateCampaigns(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsCreateClientsRequest(r *requests.AdsCreateClientsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithData(random.RandString())
}

func TestVKAdsCreateClientsSuccess(t *testing.T) {
	params := requests.NewAdsCreateClientsRequest()
	fillRandomlyAdsCreateClientsRequest(&params)
	var expected models.AdsCreateClientsResponse
	fillRandomlyAdsCreateClientsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.createClients", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsCreateClients(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsCreateTargetGroupRequest(r *requests.AdsCreateTargetGroupRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithName(random.RandString())
	r.WithLifetime(random.RandInt())
	r.WithTargetPixelId(random.RandInt())
	r.WithTargetPixelRules(random.RandString())
}

func TestVKAdsCreateTargetGroupSuccess(t *testing.T) {
	params := requests.NewAdsCreateTargetGroupRequest()
	fillRandomlyAdsCreateTargetGroupRequest(&params)
	var expected models.AdsCreateTargetGroupResponse
	fillRandomlyAdsCreateTargetGroupResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.createTargetGroup", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsCreateTargetGroup(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsDeleteAdsRequest(r *requests.AdsDeleteAdsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithIds(random.RandString())
}

func TestVKAdsDeleteAdsSuccess(t *testing.T) {
	params := requests.NewAdsDeleteAdsRequest()
	fillRandomlyAdsDeleteAdsRequest(&params)
	var expected models.AdsDeleteAdsResponse
	fillRandomlyAdsDeleteAdsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.deleteAds", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsDeleteAds(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsDeleteCampaignsRequest(r *requests.AdsDeleteCampaignsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithIds(random.RandString())
}

func TestVKAdsDeleteCampaignsSuccess(t *testing.T) {
	params := requests.NewAdsDeleteCampaignsRequest()
	fillRandomlyAdsDeleteCampaignsRequest(&params)
	var expected models.AdsDeleteCampaignsResponse
	fillRandomlyAdsDeleteCampaignsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.deleteCampaigns", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsDeleteCampaigns(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsDeleteClientsRequest(r *requests.AdsDeleteClientsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithIds(random.RandString())
}

func TestVKAdsDeleteClientsSuccess(t *testing.T) {
	params := requests.NewAdsDeleteClientsRequest()
	fillRandomlyAdsDeleteClientsRequest(&params)
	var expected models.AdsDeleteClientsResponse
	fillRandomlyAdsDeleteClientsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.deleteClients", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsDeleteClients(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsDeleteTargetGroupRequest(r *requests.AdsDeleteTargetGroupRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithTargetGroupId(random.RandInt())
}

func TestVKAdsDeleteTargetGroupSuccess(t *testing.T) {
	params := requests.NewAdsDeleteTargetGroupRequest()
	fillRandomlyAdsDeleteTargetGroupRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.deleteTargetGroup", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsDeleteTargetGroup(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKAdsGetAccountsSuccess(t *testing.T) {
	var expected models.AdsGetAccountsResponse
	fillRandomlyAdsGetAccountsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getAccounts", nil, expectedJSON))
	resp, err := vk.Api.Ads.AdsGetAccounts()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetAdsRequest(r *requests.AdsGetAdsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithAdIds(random.RandString())
	r.WithCampaignIds(random.RandString())
	r.WithClientId(random.RandInt())
	r.WithIncludeDeleted(random.RandBool())
	r.WithOnlyDeleted(random.RandBool())
	r.WithLimit(random.RandInt())
	r.WithOffset(random.RandInt())
}

func TestVKAdsGetAdsSuccess(t *testing.T) {
	params := requests.NewAdsGetAdsRequest()
	fillRandomlyAdsGetAdsRequest(&params)
	var expected models.AdsGetAdsResponse
	fillRandomlyAdsGetAdsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getAds", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetAds(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetAdsLayoutRequest(r *requests.AdsGetAdsLayoutRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithIncludeDeleted(random.RandBool())
	r.WithOnlyDeleted(random.RandBool())
	r.WithCampaignIds(random.RandString())
	r.WithAdIds(random.RandString())
	r.WithLimit(random.RandInt())
	r.WithOffset(random.RandInt())
}

func TestVKAdsGetAdsLayoutSuccess(t *testing.T) {
	params := requests.NewAdsGetAdsLayoutRequest()
	fillRandomlyAdsGetAdsLayoutRequest(&params)
	var expected models.AdsGetAdsLayoutResponse
	fillRandomlyAdsGetAdsLayoutResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getAdsLayout", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetAdsLayout(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetAdsTargetingRequest(r *requests.AdsGetAdsTargetingRequest) {
	r.WithAccountId(random.RandInt())
	r.WithAdIds(random.RandString())
	r.WithCampaignIds(random.RandString())
	r.WithClientId(random.RandInt())
	r.WithIncludeDeleted(random.RandBool())
	r.WithLimit(random.RandInt())
	r.WithOffset(random.RandInt())
}

func TestVKAdsGetAdsTargetingSuccess(t *testing.T) {
	params := requests.NewAdsGetAdsTargetingRequest()
	fillRandomlyAdsGetAdsTargetingRequest(&params)
	var expected models.AdsGetAdsTargetingResponse
	fillRandomlyAdsGetAdsTargetingResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getAdsTargeting", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetAdsTargeting(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetBudgetRequest(r *requests.AdsGetBudgetRequest) {
	r.WithAccountId(random.RandInt())
}

func TestVKAdsGetBudgetSuccess(t *testing.T) {
	params := requests.NewAdsGetBudgetRequest()
	fillRandomlyAdsGetBudgetRequest(&params)
	var expected models.AdsGetBudgetResponse
	fillRandomlyAdsGetBudgetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getBudget", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetBudget(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetCampaignsRequest(r *requests.AdsGetCampaignsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithIncludeDeleted(random.RandBool())
	r.WithCampaignIds(random.RandString())
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	r.WithFields(random.RandStringArr(lFields))
}

func TestVKAdsGetCampaignsSuccess(t *testing.T) {
	params := requests.NewAdsGetCampaignsRequest()
	fillRandomlyAdsGetCampaignsRequest(&params)
	var expected models.AdsGetCampaignsResponse
	fillRandomlyAdsGetCampaignsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getCampaigns", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetCampaigns(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetCategoriesRequest(r *requests.AdsGetCategoriesRequest) {
	r.WithLang(random.RandString())
}

func TestVKAdsGetCategoriesSuccess(t *testing.T) {
	params := requests.NewAdsGetCategoriesRequest()
	fillRandomlyAdsGetCategoriesRequest(&params)
	var expected models.AdsGetCategoriesResponse
	fillRandomlyAdsGetCategoriesResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getCategories", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetCategories(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetClientsRequest(r *requests.AdsGetClientsRequest) {
	r.WithAccountId(random.RandInt())
}

func TestVKAdsGetClientsSuccess(t *testing.T) {
	params := requests.NewAdsGetClientsRequest()
	fillRandomlyAdsGetClientsRequest(&params)
	var expected models.AdsGetClientsResponse
	fillRandomlyAdsGetClientsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getClients", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetClients(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetDemographicsRequest(r *requests.AdsGetDemographicsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithIdsType(random.RandString())
	r.WithIds(random.RandString())
	r.WithPeriod(random.RandString())
	r.WithDateFrom(random.RandString())
	r.WithDateTo(random.RandString())
}

func TestVKAdsGetDemographicsSuccess(t *testing.T) {
	params := requests.NewAdsGetDemographicsRequest()
	fillRandomlyAdsGetDemographicsRequest(&params)
	var expected models.AdsGetDemographicsResponse
	fillRandomlyAdsGetDemographicsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getDemographics", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetDemographics(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetFloodStatsRequest(r *requests.AdsGetFloodStatsRequest) {
	r.WithAccountId(random.RandInt())
}

func TestVKAdsGetFloodStatsSuccess(t *testing.T) {
	params := requests.NewAdsGetFloodStatsRequest()
	fillRandomlyAdsGetFloodStatsRequest(&params)
	var expected models.AdsGetFloodStatsResponse
	fillRandomlyAdsGetFloodStatsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getFloodStats", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetFloodStats(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetLookalikeRequestsRequest(r *requests.AdsGetLookalikeRequestsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithRequestsIds(random.RandString())
	r.WithOffset(random.RandInt())
	r.WithLimit(random.RandInt())
	r.WithSortBy(random.RandString())
}

func TestVKAdsGetLookalikeRequestsSuccess(t *testing.T) {
	params := requests.NewAdsGetLookalikeRequestsRequest()
	fillRandomlyAdsGetLookalikeRequestsRequest(&params)
	var expected models.AdsGetLookalikeRequestsResponse
	fillRandomlyAdsGetLookalikeRequestsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getLookalikeRequests", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetLookalikeRequests(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetMusiciansRequest(r *requests.AdsGetMusiciansRequest) {
	r.WithArtistName(random.RandString())
}

func TestVKAdsGetMusiciansSuccess(t *testing.T) {
	params := requests.NewAdsGetMusiciansRequest()
	fillRandomlyAdsGetMusiciansRequest(&params)
	var expected models.AdsGetMusiciansResponse
	fillRandomlyAdsGetMusiciansResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getMusicians", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetMusicians(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetMusiciansByIdsRequest(r *requests.AdsGetMusiciansByIdsRequest) {
	lIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithIds(random.RandIntArr(lIds))
}

func TestVKAdsGetMusiciansByIdsSuccess(t *testing.T) {
	params := requests.NewAdsGetMusiciansByIdsRequest()
	fillRandomlyAdsGetMusiciansByIdsRequest(&params)
	var expected models.AdsGetMusiciansResponse
	fillRandomlyAdsGetMusiciansResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getMusiciansByIds", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetMusiciansByIds(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetOfficeUsersRequest(r *requests.AdsGetOfficeUsersRequest) {
	r.WithAccountId(random.RandInt())
}

func TestVKAdsGetOfficeUsersSuccess(t *testing.T) {
	params := requests.NewAdsGetOfficeUsersRequest()
	fillRandomlyAdsGetOfficeUsersRequest(&params)
	var expected models.AdsGetOfficeUsersResponse
	fillRandomlyAdsGetOfficeUsersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getOfficeUsers", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetOfficeUsers(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetPostsReachRequest(r *requests.AdsGetPostsReachRequest) {
	r.WithAccountId(random.RandInt())
	r.WithIdsType(random.RandString())
	r.WithIds(random.RandString())
}

func TestVKAdsGetPostsReachSuccess(t *testing.T) {
	params := requests.NewAdsGetPostsReachRequest()
	fillRandomlyAdsGetPostsReachRequest(&params)
	var expected models.AdsGetPostsReachResponse
	fillRandomlyAdsGetPostsReachResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getPostsReach", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetPostsReach(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetRejectionReasonRequest(r *requests.AdsGetRejectionReasonRequest) {
	r.WithAccountId(random.RandInt())
	r.WithAdId(random.RandInt())
}

func TestVKAdsGetRejectionReasonSuccess(t *testing.T) {
	params := requests.NewAdsGetRejectionReasonRequest()
	fillRandomlyAdsGetRejectionReasonRequest(&params)
	var expected models.AdsGetRejectionReasonResponse
	fillRandomlyAdsGetRejectionReasonResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getRejectionReason", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetRejectionReason(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetStatisticsRequest(r *requests.AdsGetStatisticsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithIdsType(random.RandString())
	r.WithIds(random.RandString())
	r.WithPeriod(random.RandString())
	r.WithDateFrom(random.RandString())
	r.WithDateTo(random.RandString())
	lStatsFields := random.RandIntn(random.MaxArrayLength + 1)
	r.WithStatsFields(random.RandStringArr(lStatsFields))
}

func TestVKAdsGetStatisticsSuccess(t *testing.T) {
	params := requests.NewAdsGetStatisticsRequest()
	fillRandomlyAdsGetStatisticsRequest(&params)
	var expected models.AdsGetStatisticsResponse
	fillRandomlyAdsGetStatisticsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getStatistics", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetStatistics(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetSuggestionsRequest(r *requests.AdsGetSuggestionsRequest) {
	r.WithSection(random.RandString())
	r.WithIds(random.RandString())
	r.WithQ(random.RandString())
	r.WithCountry(random.RandInt())
	r.WithCities(random.RandString())
	r.WithLang(random.RandString())
}

func TestVKAdsGetSuggestionsSuccess(t *testing.T) {
	params := requests.NewAdsGetSuggestionsRequest()
	fillRandomlyAdsGetSuggestionsRequest(&params)
	var expected models.AdsGetSuggestionsResponse
	fillRandomlyAdsGetSuggestionsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getSuggestions", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetSuggestions(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetTargetGroupsRequest(r *requests.AdsGetTargetGroupsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithExtended(random.RandBool())
}

func TestVKAdsGetTargetGroupsSuccess(t *testing.T) {
	params := requests.NewAdsGetTargetGroupsRequest()
	fillRandomlyAdsGetTargetGroupsRequest(&params)
	var expected models.AdsGetTargetGroupsResponse
	fillRandomlyAdsGetTargetGroupsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getTargetGroups", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetTargetGroups(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetTargetingStatsRequest(r *requests.AdsGetTargetingStatsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithCriteria(random.RandString())
	r.WithAdId(random.RandInt())
	r.WithAdFormat(random.RandInt())
	r.WithAdPlatform(random.RandString())
	r.WithAdPlatformNoWall(random.RandString())
	r.WithAdPlatformNoAdNetwork(random.RandString())
	r.WithPublisherPlatforms(random.RandString())
	r.WithLinkUrl(random.RandString())
	r.WithLinkDomain(random.RandString())
	r.WithNeedPrecise(random.RandBool())
	r.WithImpressionsLimitPeriod(random.RandInt())
}

func TestVKAdsGetTargetingStatsSuccess(t *testing.T) {
	params := requests.NewAdsGetTargetingStatsRequest()
	fillRandomlyAdsGetTargetingStatsRequest(&params)
	var expected models.AdsGetTargetingStatsResponse
	fillRandomlyAdsGetTargetingStatsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getTargetingStats", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetTargetingStats(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsGetUploadURLRequest(r *requests.AdsGetUploadURLRequest) {
	r.WithAdFormat(random.RandInt())
	r.WithIcon(random.RandInt())
}

func TestVKAdsGetUploadURLSuccess(t *testing.T) {
	params := requests.NewAdsGetUploadURLRequest()
	fillRandomlyAdsGetUploadURLRequest(&params)
	var expected models.AdsGetUploadURLResponse
	fillRandomlyAdsGetUploadURLResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getUploadURL", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsGetUploadURL(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKAdsGetVideoUploadURLSuccess(t *testing.T) {
	var expected models.AdsGetVideoUploadURLResponse
	fillRandomlyAdsGetVideoUploadURLResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.getVideoUploadURL", nil, expectedJSON))
	resp, err := vk.Api.Ads.AdsGetVideoUploadURL()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsImportTargetContactsRequest(r *requests.AdsImportTargetContactsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithTargetGroupId(random.RandInt())
	r.WithContacts(random.RandString())
}

func TestVKAdsImportTargetContactsSuccess(t *testing.T) {
	params := requests.NewAdsImportTargetContactsRequest()
	fillRandomlyAdsImportTargetContactsRequest(&params)
	var expected models.AdsImportTargetContactsResponse
	fillRandomlyAdsImportTargetContactsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.importTargetContacts", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsImportTargetContacts(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsRemoveOfficeUsersRequest(r *requests.AdsRemoveOfficeUsersRequest) {
	r.WithAccountId(random.RandInt())
	r.WithIds(random.RandString())
}

func TestVKAdsRemoveOfficeUsersSuccess(t *testing.T) {
	params := requests.NewAdsRemoveOfficeUsersRequest()
	fillRandomlyAdsRemoveOfficeUsersRequest(&params)
	var expected models.AdsRemoveOfficeUsersResponse
	fillRandomlyAdsRemoveOfficeUsersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.removeOfficeUsers", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsRemoveOfficeUsers(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsUpdateAdsRequest(r *requests.AdsUpdateAdsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithData(random.RandString())
}

func TestVKAdsUpdateAdsSuccess(t *testing.T) {
	params := requests.NewAdsUpdateAdsRequest()
	fillRandomlyAdsUpdateAdsRequest(&params)
	var expected models.AdsUpdateAdsResponse
	fillRandomlyAdsUpdateAdsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.updateAds", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsUpdateAds(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsUpdateCampaignsRequest(r *requests.AdsUpdateCampaignsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithData(random.RandString())
}

func TestVKAdsUpdateCampaignsSuccess(t *testing.T) {
	params := requests.NewAdsUpdateCampaignsRequest()
	fillRandomlyAdsUpdateCampaignsRequest(&params)
	var expected models.AdsUpdateCampaignsResponse
	fillRandomlyAdsUpdateCampaignsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.updateCampaigns", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsUpdateCampaigns(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsUpdateClientsRequest(r *requests.AdsUpdateClientsRequest) {
	r.WithAccountId(random.RandInt())
	r.WithData(random.RandString())
}

func TestVKAdsUpdateClientsSuccess(t *testing.T) {
	params := requests.NewAdsUpdateClientsRequest()
	fillRandomlyAdsUpdateClientsRequest(&params)
	var expected models.AdsUpdateClientsResponse
	fillRandomlyAdsUpdateClientsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.updateClients", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsUpdateClients(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsUpdateOfficeUsersRequest(r *requests.AdsUpdateOfficeUsersRequest) {
	r.WithAccountId(random.RandInt())
	Data := new([]models.AdsUserSpecification)
	lData := random.RandIntn(random.MaxArrayLength + 1)
	*Data = make([]models.AdsUserSpecification, lData)
	for i0 := 0; i0 < lData; i0++ {
		fillRandomlyAdsUserSpecification(&(*Data)[i0])
	}
	r.WithData(*Data)
}

func TestVKAdsUpdateOfficeUsersSuccess(t *testing.T) {
	params := requests.NewAdsUpdateOfficeUsersRequest()
	fillRandomlyAdsUpdateOfficeUsersRequest(&params)
	var expected models.AdsUpdateOfficeUsersResponse
	fillRandomlyAdsUpdateOfficeUsersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.updateOfficeUsers", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsUpdateOfficeUsers(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAdsUpdateTargetGroupRequest(r *requests.AdsUpdateTargetGroupRequest) {
	r.WithAccountId(random.RandInt())
	r.WithClientId(random.RandInt())
	r.WithTargetGroupId(random.RandInt())
	r.WithName(random.RandString())
	r.WithDomain(random.RandString())
	r.WithLifetime(random.RandInt())
	r.WithTargetPixelId(random.RandInt())
	r.WithTargetPixelRules(random.RandString())
}

func TestVKAdsUpdateTargetGroupSuccess(t *testing.T) {
	params := requests.NewAdsUpdateTargetGroupRequest()
	fillRandomlyAdsUpdateTargetGroupRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "ads.updateTargetGroup", params.Params(), expectedJSON))
	resp, err := vk.Api.Ads.AdsUpdateTargetGroup(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}