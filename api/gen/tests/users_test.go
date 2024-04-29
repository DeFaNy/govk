// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/users"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyUsersGetRequest(r *requests.UsersGetRequest) {
	lUserIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithUserIds(random.RandStringArr(lUserIds))
	Fields := new([]models.UsersFields)
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
	r.WithNameCase(random.RandString())
}

func TestVKUsersGetSuccess(t *testing.T) {
	params := requests.NewUsersGetRequest()
	fillRandomlyUsersGetRequest(&params)
	var expected models.UsersGetResponse
	fillRandomlyUsersGetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "users.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Users.UsersGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUsersGetFollowersRequest(r *requests.UsersGetFollowersRequest) {
	r.WithUserId(random.RandInt())
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	Fields := new([]models.UsersFields)
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
	r.WithNameCase(random.RandString())
}

func TestVKUsersGetFollowersSuccess(t *testing.T) {
	params := requests.NewUsersGetFollowersRequest()
	fillRandomlyUsersGetFollowersRequest(&params)
	var expected models.UsersGetFollowersResponse
	fillRandomlyUsersGetFollowersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "users.getFollowers", params.Params(), expectedJSON))
	resp, err := vk.Api.Users.UsersGetFollowers(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUsersGetSubscriptionsRequest(r *requests.UsersGetSubscriptionsRequest) {
	r.WithUserId(random.RandInt())
	r.WithExtended(random.RandBool())
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	Fields := new([]models.UsersFields)
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
}

func TestVKUsersGetSubscriptionsSuccess(t *testing.T) {
	params := requests.NewUsersGetSubscriptionsRequest()
	fillRandomlyUsersGetSubscriptionsRequest(&params)
	params.WithExtended(false)
	var expected models.UsersGetSubscriptionsResponse
	fillRandomlyUsersGetSubscriptionsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "users.getSubscriptions", params.Params(), expectedJSON))
	resp, err := vk.Api.Users.UsersGetSubscriptions(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKUsersGetSubscriptionsExtendedSuccess(t *testing.T) {
	params := requests.NewUsersGetSubscriptionsRequest()
	fillRandomlyUsersGetSubscriptionsRequest(&params)
	params.WithExtended(true)
	var expected models.UsersGetSubscriptionsExtendedResponse
	fillRandomlyUsersGetSubscriptionsExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "users.getSubscriptions", params.Params(), expectedJSON))
	resp, err := vk.Api.Users.UsersGetSubscriptionsExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUsersReportRequest(r *requests.UsersReportRequest) {
	r.WithUserId(random.RandInt())
	r.WithType(random.RandString())
	r.WithComment(random.RandString())
}

func TestVKUsersReportSuccess(t *testing.T) {
	params := requests.NewUsersReportRequest()
	fillRandomlyUsersReportRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "users.report", params.Params(), expectedJSON))
	resp, err := vk.Api.Users.UsersReport(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUsersSearchRequest(r *requests.UsersSearchRequest) {
	r.WithQ(random.RandString())
	r.WithSort(random.RandInt())
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	Fields := new([]models.UsersFields)
	lFields := random.RandIntn(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
	r.WithCity(random.RandInt())
	r.WithCountry(random.RandInt())
	r.WithHometown(random.RandString())
	r.WithUniversityCountry(random.RandInt())
	r.WithUniversity(random.RandInt())
	r.WithUniversityYear(random.RandInt())
	r.WithUniversityFaculty(random.RandInt())
	r.WithUniversityChair(random.RandInt())
	r.WithSex(random.RandInt())
	r.WithStatus(random.RandInt())
	r.WithAgeFrom(random.RandInt())
	r.WithAgeTo(random.RandInt())
	r.WithBirthDay(random.RandInt())
	r.WithBirthMonth(random.RandInt())
	r.WithBirthYear(random.RandInt())
	r.WithOnline(random.RandBool())
	r.WithHasPhoto(random.RandBool())
	r.WithSchoolCountry(random.RandInt())
	r.WithSchoolCity(random.RandInt())
	r.WithSchoolClass(random.RandInt())
	r.WithSchool(random.RandInt())
	r.WithSchoolYear(random.RandInt())
	r.WithReligion(random.RandString())
	r.WithCompany(random.RandString())
	r.WithPosition(random.RandString())
	r.WithGroupId(random.RandInt())
	lFromList := random.RandIntn(random.MaxArrayLength + 1)
	r.WithFromList(random.RandStringArr(lFromList))
}

func TestVKUsersSearchSuccess(t *testing.T) {
	params := requests.NewUsersSearchRequest()
	fillRandomlyUsersSearchRequest(&params)
	var expected models.UsersSearchResponse
	fillRandomlyUsersSearchResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "users.search", params.Params(), expectedJSON))
	resp, err := vk.Api.Users.UsersSearch(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}