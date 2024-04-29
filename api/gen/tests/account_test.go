// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/account"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyAccountBanRequest(r *requests.AccountBanRequest) {
	r.WithOwnerId(random.RandInt())
}

func TestVKAccountBanSuccess(t *testing.T) {
	params := requests.NewAccountBanRequest()
	fillRandomlyAccountBanRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.ban", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountBan(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountChangePasswordRequest(r *requests.AccountChangePasswordRequest) {
	r.WithRestoreSid(random.RandString())
	r.WithChangePasswordHash(random.RandString())
	r.WithOldPassword(random.RandString())
	r.WithNewPassword(random.RandString())
}

func TestVKAccountChangePasswordSuccess(t *testing.T) {
	params := requests.NewAccountChangePasswordRequest()
	fillRandomlyAccountChangePasswordRequest(&params)
	var expected models.AccountChangePasswordResponse
	fillRandomlyAccountChangePasswordResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.changePassword", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountChangePassword(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetActiveOffersRequest(r *requests.AccountGetActiveOffersRequest) {
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
}

func TestVKAccountGetActiveOffersSuccess(t *testing.T) {
	params := requests.NewAccountGetActiveOffersRequest()
	fillRandomlyAccountGetActiveOffersRequest(&params)
	var expected models.AccountGetActiveOffersResponse
	fillRandomlyAccountGetActiveOffersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getActiveOffers", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetActiveOffers(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetAppPermissionsRequest(r *requests.AccountGetAppPermissionsRequest) {
	r.WithUserId(random.RandInt())
}

func TestVKAccountGetAppPermissionsSuccess(t *testing.T) {
	params := requests.NewAccountGetAppPermissionsRequest()
	fillRandomlyAccountGetAppPermissionsRequest(&params)
	var expected models.AccountGetAppPermissionsResponse
	fillRandomlyAccountGetAppPermissionsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getAppPermissions", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetAppPermissions(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetBannedRequest(r *requests.AccountGetBannedRequest) {
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
}

func TestVKAccountGetBannedSuccess(t *testing.T) {
	params := requests.NewAccountGetBannedRequest()
	fillRandomlyAccountGetBannedRequest(&params)
	var expected models.AccountGetBannedResponse
	fillRandomlyAccountGetBannedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getBanned", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetBanned(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetCountersRequest(r *requests.AccountGetCountersRequest) {
	lFilter := random.RandIntn(random.MaxArrayLength + 1)
	r.WithFilter(random.RandStringArr(lFilter))
	r.WithUserId(random.RandInt())
}

func TestVKAccountGetCountersSuccess(t *testing.T) {
	params := requests.NewAccountGetCountersRequest()
	fillRandomlyAccountGetCountersRequest(&params)
	var expected models.AccountGetCountersResponse
	fillRandomlyAccountGetCountersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getCounters", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetCounters(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetInfoRequest(r *requests.AccountGetInfoRequest) {
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	r.WithFields(random.RandStringArr(lFields))
}

func TestVKAccountGetInfoSuccess(t *testing.T) {
	params := requests.NewAccountGetInfoRequest()
	fillRandomlyAccountGetInfoRequest(&params)
	var expected models.AccountGetInfoResponse
	fillRandomlyAccountGetInfoResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getInfo", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetInfo(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKAccountGetProfileInfoSuccess(t *testing.T) {
	var expected models.AccountGetProfileInfoResponse
	fillRandomlyAccountGetProfileInfoResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getProfileInfo", nil, expectedJSON))
	resp, err := vk.Api.Account.AccountGetProfileInfo()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetPushSettingsRequest(r *requests.AccountGetPushSettingsRequest) {
	r.WithDeviceId(random.RandString())
}

func TestVKAccountGetPushSettingsSuccess(t *testing.T) {
	params := requests.NewAccountGetPushSettingsRequest()
	fillRandomlyAccountGetPushSettingsRequest(&params)
	var expected models.AccountGetPushSettingsResponse
	fillRandomlyAccountGetPushSettingsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getPushSettings", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetPushSettings(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountRegisterDeviceRequest(r *requests.AccountRegisterDeviceRequest) {
	r.WithToken(random.RandString())
	r.WithDeviceModel(random.RandString())
	r.WithDeviceYear(random.RandInt())
	r.WithDeviceId(random.RandString())
	r.WithSystemVersion(random.RandString())
	r.WithSettings(random.RandString())
	r.WithSandbox(random.RandBool())
}

func TestVKAccountRegisterDeviceSuccess(t *testing.T) {
	params := requests.NewAccountRegisterDeviceRequest()
	fillRandomlyAccountRegisterDeviceRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.registerDevice", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountRegisterDevice(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSaveProfileInfoRequest(r *requests.AccountSaveProfileInfoRequest) {
	r.WithFirstName(random.RandString())
	r.WithLastName(random.RandString())
	r.WithMaidenName(random.RandString())
	r.WithScreenName(random.RandString())
	r.WithCancelRequestId(random.RandInt())
	r.WithSex(random.RandInt())
	r.WithRelation(random.RandInt())
	r.WithRelationPartnerId(random.RandInt())
	r.WithBdate(random.RandString())
	r.WithBdateVisibility(random.RandInt())
	r.WithHomeTown(random.RandString())
	r.WithCountryId(random.RandInt())
	r.WithCityId(random.RandInt())
	r.WithStatus(random.RandString())
}

func TestVKAccountSaveProfileInfoSuccess(t *testing.T) {
	params := requests.NewAccountSaveProfileInfoRequest()
	fillRandomlyAccountSaveProfileInfoRequest(&params)
	var expected models.AccountSaveProfileInfoResponse
	fillRandomlyAccountSaveProfileInfoResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.saveProfileInfo", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSaveProfileInfo(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSetInfoRequest(r *requests.AccountSetInfoRequest) {
	r.WithName(random.RandString())
	r.WithValue(random.RandString())
}

func TestVKAccountSetInfoSuccess(t *testing.T) {
	params := requests.NewAccountSetInfoRequest()
	fillRandomlyAccountSetInfoRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setInfo", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSetInfo(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKAccountSetOfflineSuccess(t *testing.T) {
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setOffline", nil, expectedJSON))
	resp, err := vk.Api.Account.AccountSetOffline()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSetOnlineRequest(r *requests.AccountSetOnlineRequest) {
	r.WithVoip(random.RandBool())
}

func TestVKAccountSetOnlineSuccess(t *testing.T) {
	params := requests.NewAccountSetOnlineRequest()
	fillRandomlyAccountSetOnlineRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setOnline", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSetOnline(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSetPushSettingsRequest(r *requests.AccountSetPushSettingsRequest) {
	r.WithDeviceId(random.RandString())
	r.WithSettings(random.RandString())
	r.WithKey(random.RandString())
	lValue := random.RandIntn(random.MaxArrayLength + 1)
	r.WithValue(random.RandStringArr(lValue))
}

func TestVKAccountSetPushSettingsSuccess(t *testing.T) {
	params := requests.NewAccountSetPushSettingsRequest()
	fillRandomlyAccountSetPushSettingsRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setPushSettings", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSetPushSettings(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSetSilenceModeRequest(r *requests.AccountSetSilenceModeRequest) {
	r.WithDeviceId(random.RandString())
	r.WithTime(random.RandInt())
	r.WithPeerId(random.RandInt())
	r.WithSound(random.RandInt())
}

func TestVKAccountSetSilenceModeSuccess(t *testing.T) {
	params := requests.NewAccountSetSilenceModeRequest()
	fillRandomlyAccountSetSilenceModeRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setSilenceMode", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSetSilenceMode(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountUnbanRequest(r *requests.AccountUnbanRequest) {
	r.WithOwnerId(random.RandInt())
}

func TestVKAccountUnbanSuccess(t *testing.T) {
	params := requests.NewAccountUnbanRequest()
	fillRandomlyAccountUnbanRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.unban", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountUnban(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountUnregisterDeviceRequest(r *requests.AccountUnregisterDeviceRequest) {
	r.WithDeviceId(random.RandString())
	r.WithSandbox(random.RandBool())
}

func TestVKAccountUnregisterDeviceSuccess(t *testing.T) {
	params := requests.NewAccountUnregisterDeviceRequest()
	fillRandomlyAccountUnregisterDeviceRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.unregisterDevice", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountUnregisterDevice(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}