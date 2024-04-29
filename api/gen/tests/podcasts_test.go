// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/podcasts"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyPodcastsSearchPodcastRequest(r *requests.PodcastsSearchPodcastRequest) {
	r.WithSearchString(random.RandString())
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
}

func TestVKPodcastsSearchPodcastSuccess(t *testing.T) {
	params := requests.NewPodcastsSearchPodcastRequest()
	fillRandomlyPodcastsSearchPodcastRequest(&params)
	var expected models.PodcastsSearchPodcastResponse
	fillRandomlyPodcastsSearchPodcastResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "podcasts.searchPodcast", params.Params(), expectedJSON))
	resp, err := vk.Api.Podcasts.PodcastsSearchPodcast(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
