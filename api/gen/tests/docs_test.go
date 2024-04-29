// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/docs"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyDocsAddRequest(r *requests.DocsAddRequest) {
	r.WithOwnerId(random.RandInt())
	r.WithDocId(random.RandInt())
	r.WithAccessKey(random.RandString())
}

func TestVKDocsAddSuccess(t *testing.T) {
	params := requests.NewDocsAddRequest()
	fillRandomlyDocsAddRequest(&params)
	var expected models.DocsAddResponse
	fillRandomlyDocsAddResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.add", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsAdd(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsDeleteRequest(r *requests.DocsDeleteRequest) {
	r.WithOwnerId(random.RandInt())
	r.WithDocId(random.RandInt())
}

func TestVKDocsDeleteSuccess(t *testing.T) {
	params := requests.NewDocsDeleteRequest()
	fillRandomlyDocsDeleteRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.delete", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsDelete(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsEditRequest(r *requests.DocsEditRequest) {
	r.WithOwnerId(random.RandInt())
	r.WithDocId(random.RandInt())
	r.WithTitle(random.RandString())
	lTags := random.RandIntn(random.MaxArrayLength + 1)
	r.WithTags(random.RandStringArr(lTags))
}

func TestVKDocsEditSuccess(t *testing.T) {
	params := requests.NewDocsEditRequest()
	fillRandomlyDocsEditRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.edit", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsEdit(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsGetRequest(r *requests.DocsGetRequest) {
	r.WithCount(random.RandInt())
	r.WithOffset(random.RandInt())
	r.WithType(random.RandInt())
	r.WithOwnerId(random.RandInt())
	r.WithReturnTags(random.RandBool())
}

func TestVKDocsGetSuccess(t *testing.T) {
	params := requests.NewDocsGetRequest()
	fillRandomlyDocsGetRequest(&params)
	var expected models.DocsGetResponse
	fillRandomlyDocsGetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsGetByIdRequest(r *requests.DocsGetByIdRequest) {
	lDocs := random.RandIntn(random.MaxArrayLength + 1)
	r.WithDocs(random.RandStringArr(lDocs))
	r.WithReturnTags(random.RandBool())
}

func TestVKDocsGetByIdSuccess(t *testing.T) {
	params := requests.NewDocsGetByIdRequest()
	fillRandomlyDocsGetByIdRequest(&params)
	var expected models.DocsGetByIdResponse
	fillRandomlyDocsGetByIdResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.getById", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsGetById(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsGetMessagesUploadServerRequest(r *requests.DocsGetMessagesUploadServerRequest) {
	r.WithType(random.RandString())
	r.WithPeerId(random.RandInt())
}

func TestVKDocsGetMessagesUploadServerSuccess(t *testing.T) {
	params := requests.NewDocsGetMessagesUploadServerRequest()
	fillRandomlyDocsGetMessagesUploadServerRequest(&params)
	var expected models.DocsGetUploadServerResponse
	fillRandomlyDocsGetUploadServerResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.getMessagesUploadServer", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsGetMessagesUploadServer(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsGetTypesRequest(r *requests.DocsGetTypesRequest) {
	r.WithOwnerId(random.RandInt())
}

func TestVKDocsGetTypesSuccess(t *testing.T) {
	params := requests.NewDocsGetTypesRequest()
	fillRandomlyDocsGetTypesRequest(&params)
	var expected models.DocsGetTypesResponse
	fillRandomlyDocsGetTypesResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.getTypes", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsGetTypes(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsGetUploadServerRequest(r *requests.DocsGetUploadServerRequest) {
	r.WithGroupId(random.RandInt())
}

func TestVKDocsGetUploadServerSuccess(t *testing.T) {
	params := requests.NewDocsGetUploadServerRequest()
	fillRandomlyDocsGetUploadServerRequest(&params)
	var expected models.DocsGetUploadServerResponse
	fillRandomlyDocsGetUploadServerResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.getUploadServer", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsGetUploadServer(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsGetWallUploadServerRequest(r *requests.DocsGetWallUploadServerRequest) {
	r.WithGroupId(random.RandInt())
}

func TestVKDocsGetWallUploadServerSuccess(t *testing.T) {
	params := requests.NewDocsGetWallUploadServerRequest()
	fillRandomlyDocsGetWallUploadServerRequest(&params)
	var expected models.BaseGetUploadServerResponse
	fillRandomlyBaseGetUploadServerResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.getWallUploadServer", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsGetWallUploadServer(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsSaveRequest(r *requests.DocsSaveRequest) {
	r.WithFile(random.RandString())
	r.WithTitle(random.RandString())
	r.WithTags(random.RandString())
	r.WithReturnTags(random.RandBool())
}

func TestVKDocsSaveSuccess(t *testing.T) {
	params := requests.NewDocsSaveRequest()
	fillRandomlyDocsSaveRequest(&params)
	var expected models.DocsSaveResponse
	fillRandomlyDocsSaveResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.save", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsSave(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDocsSearchRequest(r *requests.DocsSearchRequest) {
	r.WithQ(random.RandString())
	r.WithSearchOwn(random.RandBool())
	r.WithCount(random.RandInt())
	r.WithOffset(random.RandInt())
	r.WithReturnTags(random.RandBool())
}

func TestVKDocsSearchSuccess(t *testing.T) {
	params := requests.NewDocsSearchRequest()
	fillRandomlyDocsSearchRequest(&params)
	var expected models.DocsSearchResponse
	fillRandomlyDocsSearchResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "docs.search", params.Params(), expectedJSON))
	resp, err := vk.Api.Docs.DocsSearch(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
