// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/downloadedgames"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyDownloadedGamesGetPaidStatusRequest(r *requests.DownloadedGamesGetPaidStatusRequest) {
	r.WithUserId(random.RandInt())
}

func TestVKDownloadedGamesGetPaidStatusSuccess(t *testing.T) {
	params := requests.NewDownloadedGamesGetPaidStatusRequest()
	fillRandomlyDownloadedGamesGetPaidStatusRequest(&params)
	var expected models.DownloadedGamesPaidStatusResponse
	fillRandomlyDownloadedGamesPaidStatusResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "downloadedGames.getPaidStatus", params.Params(), expectedJSON))
	resp, err := vk.Api.DownloadedGames.DownloadedGamesGetPaidStatus(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}