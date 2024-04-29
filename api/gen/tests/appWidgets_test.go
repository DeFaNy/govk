// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/appwidgets"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyAppWidgetsGetAppImageUploadServerRequest(r *requests.AppWidgetsGetAppImageUploadServerRequest) {
	r.WithImageType(random.RandString())
}

func TestVKAppWidgetsGetAppImageUploadServerSuccess(t *testing.T) {
	params := requests.NewAppWidgetsGetAppImageUploadServerRequest()
	fillRandomlyAppWidgetsGetAppImageUploadServerRequest(&params)
	var expected models.AppWidgetsGetAppImageUploadServerResponse
	fillRandomlyAppWidgetsGetAppImageUploadServerResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "appWidgets.getAppImageUploadServer", params.Params(), expectedJSON))
	resp, err := vk.Api.AppWidgets.AppWidgetsGetAppImageUploadServer(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppWidgetsGetAppImagesRequest(r *requests.AppWidgetsGetAppImagesRequest) {
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	r.WithImageType(random.RandString())
}

func TestVKAppWidgetsGetAppImagesSuccess(t *testing.T) {
	params := requests.NewAppWidgetsGetAppImagesRequest()
	fillRandomlyAppWidgetsGetAppImagesRequest(&params)
	var expected models.AppWidgetsGetAppImagesResponse
	fillRandomlyAppWidgetsGetAppImagesResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "appWidgets.getAppImages", params.Params(), expectedJSON))
	resp, err := vk.Api.AppWidgets.AppWidgetsGetAppImages(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppWidgetsGetGroupImageUploadServerRequest(r *requests.AppWidgetsGetGroupImageUploadServerRequest) {
	r.WithImageType(random.RandString())
}

func TestVKAppWidgetsGetGroupImageUploadServerSuccess(t *testing.T) {
	params := requests.NewAppWidgetsGetGroupImageUploadServerRequest()
	fillRandomlyAppWidgetsGetGroupImageUploadServerRequest(&params)
	var expected models.AppWidgetsGetGroupImageUploadServerResponse
	fillRandomlyAppWidgetsGetGroupImageUploadServerResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "appWidgets.getGroupImageUploadServer", params.Params(), expectedJSON))
	resp, err := vk.Api.AppWidgets.AppWidgetsGetGroupImageUploadServer(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppWidgetsGetGroupImagesRequest(r *requests.AppWidgetsGetGroupImagesRequest) {
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	r.WithImageType(random.RandString())
}

func TestVKAppWidgetsGetGroupImagesSuccess(t *testing.T) {
	params := requests.NewAppWidgetsGetGroupImagesRequest()
	fillRandomlyAppWidgetsGetGroupImagesRequest(&params)
	var expected models.AppWidgetsGetGroupImagesResponse
	fillRandomlyAppWidgetsGetGroupImagesResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "appWidgets.getGroupImages", params.Params(), expectedJSON))
	resp, err := vk.Api.AppWidgets.AppWidgetsGetGroupImages(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppWidgetsGetImagesByIdRequest(r *requests.AppWidgetsGetImagesByIdRequest) {
	lImages := random.RandIntn(random.MaxArrayLength + 1)
	r.WithImages(random.RandStringArr(lImages))
}

func TestVKAppWidgetsGetImagesByIdSuccess(t *testing.T) {
	params := requests.NewAppWidgetsGetImagesByIdRequest()
	fillRandomlyAppWidgetsGetImagesByIdRequest(&params)
	var expected models.AppWidgetsGetImagesByIdResponse
	fillRandomlyAppWidgetsGetImagesByIdResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "appWidgets.getImagesById", params.Params(), expectedJSON))
	resp, err := vk.Api.AppWidgets.AppWidgetsGetImagesById(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppWidgetsSaveAppImageRequest(r *requests.AppWidgetsSaveAppImageRequest) {
	r.WithHash(random.RandString())
	r.WithImage(random.RandString())
}

func TestVKAppWidgetsSaveAppImageSuccess(t *testing.T) {
	params := requests.NewAppWidgetsSaveAppImageRequest()
	fillRandomlyAppWidgetsSaveAppImageRequest(&params)
	var expected models.AppWidgetsSaveAppImageResponse
	fillRandomlyAppWidgetsSaveAppImageResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "appWidgets.saveAppImage", params.Params(), expectedJSON))
	resp, err := vk.Api.AppWidgets.AppWidgetsSaveAppImage(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppWidgetsSaveGroupImageRequest(r *requests.AppWidgetsSaveGroupImageRequest) {
	r.WithHash(random.RandString())
	r.WithImage(random.RandString())
}

func TestVKAppWidgetsSaveGroupImageSuccess(t *testing.T) {
	params := requests.NewAppWidgetsSaveGroupImageRequest()
	fillRandomlyAppWidgetsSaveGroupImageRequest(&params)
	var expected models.AppWidgetsSaveGroupImageResponse
	fillRandomlyAppWidgetsSaveGroupImageResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "appWidgets.saveGroupImage", params.Params(), expectedJSON))
	resp, err := vk.Api.AppWidgets.AppWidgetsSaveGroupImage(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAppWidgetsUpdateRequest(r *requests.AppWidgetsUpdateRequest) {
	r.WithCode(random.RandString())
	r.WithType(random.RandString())
}

func TestVKAppWidgetsUpdateSuccess(t *testing.T) {
	params := requests.NewAppWidgetsUpdateRequest()
	fillRandomlyAppWidgetsUpdateRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "appWidgets.update", params.Params(), expectedJSON))
	resp, err := vk.Api.AppWidgets.AppWidgetsUpdate(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
