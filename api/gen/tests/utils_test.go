// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/utils"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyUtilsCheckLinkRequest(r *requests.UtilsCheckLinkRequest) {
	r.WithUrl(random.RandString())
}

func TestVKUtilsCheckLinkSuccess(t *testing.T) {
	params := requests.NewUtilsCheckLinkRequest()
	fillRandomlyUtilsCheckLinkRequest(&params)
	var expected models.UtilsCheckLinkResponse
	fillRandomlyUtilsCheckLinkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "utils.checkLink", params.Params(), expectedJSON))
	resp, err := vk.Api.Utils.UtilsCheckLink(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUtilsDeleteFromLastShortenedRequest(r *requests.UtilsDeleteFromLastShortenedRequest) {
	r.WithKey(random.RandString())
}

func TestVKUtilsDeleteFromLastShortenedSuccess(t *testing.T) {
	params := requests.NewUtilsDeleteFromLastShortenedRequest()
	fillRandomlyUtilsDeleteFromLastShortenedRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "utils.deleteFromLastShortened", params.Params(), expectedJSON))
	resp, err := vk.Api.Utils.UtilsDeleteFromLastShortened(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUtilsGetLastShortenedLinksRequest(r *requests.UtilsGetLastShortenedLinksRequest) {
	r.WithCount(random.RandInt())
	r.WithOffset(random.RandInt())
}

func TestVKUtilsGetLastShortenedLinksSuccess(t *testing.T) {
	params := requests.NewUtilsGetLastShortenedLinksRequest()
	fillRandomlyUtilsGetLastShortenedLinksRequest(&params)
	var expected models.UtilsGetLastShortenedLinksResponse
	fillRandomlyUtilsGetLastShortenedLinksResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "utils.getLastShortenedLinks", params.Params(), expectedJSON))
	resp, err := vk.Api.Utils.UtilsGetLastShortenedLinks(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUtilsGetLinkStatsRequest(r *requests.UtilsGetLinkStatsRequest) {
	r.WithKey(random.RandString())
	r.WithSource(random.RandString())
	r.WithAccessKey(random.RandString())
	r.WithInterval(random.RandString())
	r.WithIntervalsCount(random.RandInt())
	r.WithExtended(random.RandBool())
}

func TestVKUtilsGetLinkStatsSuccess(t *testing.T) {
	params := requests.NewUtilsGetLinkStatsRequest()
	fillRandomlyUtilsGetLinkStatsRequest(&params)
	params.WithExtended(false)
	var expected models.UtilsGetLinkStatsResponse
	fillRandomlyUtilsGetLinkStatsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "utils.getLinkStats", params.Params(), expectedJSON))
	resp, err := vk.Api.Utils.UtilsGetLinkStats(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKUtilsGetLinkStatsExtendedSuccess(t *testing.T) {
	params := requests.NewUtilsGetLinkStatsRequest()
	fillRandomlyUtilsGetLinkStatsRequest(&params)
	params.WithExtended(true)
	var expected models.UtilsGetLinkStatsExtendedResponse
	fillRandomlyUtilsGetLinkStatsExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "utils.getLinkStats", params.Params(), expectedJSON))
	resp, err := vk.Api.Utils.UtilsGetLinkStatsExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKUtilsGetServerTimeSuccess(t *testing.T) {
	var expected models.UtilsGetServerTimeResponse
	fillRandomlyUtilsGetServerTimeResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "utils.getServerTime", nil, expectedJSON))
	resp, err := vk.Api.Utils.UtilsGetServerTime()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUtilsGetShortLinkRequest(r *requests.UtilsGetShortLinkRequest) {
	r.WithUrl(random.RandString())
	r.WithPrivate(random.RandBool())
}

func TestVKUtilsGetShortLinkSuccess(t *testing.T) {
	params := requests.NewUtilsGetShortLinkRequest()
	fillRandomlyUtilsGetShortLinkRequest(&params)
	var expected models.UtilsGetShortLinkResponse
	fillRandomlyUtilsGetShortLinkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "utils.getShortLink", params.Params(), expectedJSON))
	resp, err := vk.Api.Utils.UtilsGetShortLink(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyUtilsResolveScreenNameRequest(r *requests.UtilsResolveScreenNameRequest) {
	r.WithScreenName(random.RandString())
}

func TestVKUtilsResolveScreenNameSuccess(t *testing.T) {
	params := requests.NewUtilsResolveScreenNameRequest()
	fillRandomlyUtilsResolveScreenNameRequest(&params)
	var expected models.UtilsResolveScreenNameResponse
	fillRandomlyUtilsResolveScreenNameResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "utils.resolveScreenName", params.Params(), expectedJSON))
	resp, err := vk.Api.Utils.UtilsResolveScreenName(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}