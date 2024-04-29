// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/board"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyBoardAddTopicRequest(r *requests.BoardAddTopicRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTitle(random.RandString())
	r.WithText(random.RandString())
	r.WithFromGroup(random.RandBool())
	lAttachments := random.RandIntn(random.MaxArrayLength + 1)
	r.WithAttachments(random.RandStringArr(lAttachments))
}

func TestVKBoardAddTopicSuccess(t *testing.T) {
	params := requests.NewBoardAddTopicRequest()
	fillRandomlyBoardAddTopicRequest(&params)
	var expected models.BoardAddTopicResponse
	fillRandomlyBoardAddTopicResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.addTopic", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardAddTopic(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardCloseTopicRequest(r *requests.BoardCloseTopicRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
}

func TestVKBoardCloseTopicSuccess(t *testing.T) {
	params := requests.NewBoardCloseTopicRequest()
	fillRandomlyBoardCloseTopicRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.closeTopic", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardCloseTopic(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardCreateCommentRequest(r *requests.BoardCreateCommentRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
	r.WithMessage(random.RandString())
	lAttachments := random.RandIntn(random.MaxArrayLength + 1)
	r.WithAttachments(random.RandStringArr(lAttachments))
	r.WithFromGroup(random.RandBool())
	r.WithStickerId(random.RandInt())
	r.WithGuid(random.RandString())
}

func TestVKBoardCreateCommentSuccess(t *testing.T) {
	params := requests.NewBoardCreateCommentRequest()
	fillRandomlyBoardCreateCommentRequest(&params)
	var expected models.BoardCreateCommentResponse
	fillRandomlyBoardCreateCommentResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.createComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardCreateComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardDeleteCommentRequest(r *requests.BoardDeleteCommentRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
	r.WithCommentId(random.RandInt())
}

func TestVKBoardDeleteCommentSuccess(t *testing.T) {
	params := requests.NewBoardDeleteCommentRequest()
	fillRandomlyBoardDeleteCommentRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.deleteComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardDeleteComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardDeleteTopicRequest(r *requests.BoardDeleteTopicRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
}

func TestVKBoardDeleteTopicSuccess(t *testing.T) {
	params := requests.NewBoardDeleteTopicRequest()
	fillRandomlyBoardDeleteTopicRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.deleteTopic", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardDeleteTopic(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardEditCommentRequest(r *requests.BoardEditCommentRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
	r.WithCommentId(random.RandInt())
	r.WithMessage(random.RandString())
	lAttachments := random.RandIntn(random.MaxArrayLength + 1)
	r.WithAttachments(random.RandStringArr(lAttachments))
}

func TestVKBoardEditCommentSuccess(t *testing.T) {
	params := requests.NewBoardEditCommentRequest()
	fillRandomlyBoardEditCommentRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.editComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardEditComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardEditTopicRequest(r *requests.BoardEditTopicRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
	r.WithTitle(random.RandString())
}

func TestVKBoardEditTopicSuccess(t *testing.T) {
	params := requests.NewBoardEditTopicRequest()
	fillRandomlyBoardEditTopicRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.editTopic", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardEditTopic(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardFixTopicRequest(r *requests.BoardFixTopicRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
}

func TestVKBoardFixTopicSuccess(t *testing.T) {
	params := requests.NewBoardFixTopicRequest()
	fillRandomlyBoardFixTopicRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.fixTopic", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardFixTopic(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardGetCommentsRequest(r *requests.BoardGetCommentsRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
	r.WithNeedLikes(random.RandBool())
	r.WithStartCommentId(random.RandInt())
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	r.WithExtended(random.RandBool())
	r.WithSort(random.RandString())
}

func TestVKBoardGetCommentsSuccess(t *testing.T) {
	params := requests.NewBoardGetCommentsRequest()
	fillRandomlyBoardGetCommentsRequest(&params)
	params.WithExtended(false)
	var expected models.BoardGetCommentsResponse
	fillRandomlyBoardGetCommentsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.getComments", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardGetComments(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKBoardGetCommentsExtendedSuccess(t *testing.T) {
	params := requests.NewBoardGetCommentsRequest()
	fillRandomlyBoardGetCommentsRequest(&params)
	params.WithExtended(true)
	var expected models.BoardGetCommentsExtendedResponse
	fillRandomlyBoardGetCommentsExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.getComments", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardGetCommentsExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardGetTopicsRequest(r *requests.BoardGetTopicsRequest) {
	r.WithGroupId(random.RandInt())
	lTopicIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithTopicIds(random.RandIntArr(lTopicIds))
	r.WithOrder(random.RandInt())
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	r.WithExtended(random.RandBool())
	r.WithPreview(random.RandInt())
	r.WithPreviewLength(random.RandInt())
}

func TestVKBoardGetTopicsSuccess(t *testing.T) {
	params := requests.NewBoardGetTopicsRequest()
	fillRandomlyBoardGetTopicsRequest(&params)
	params.WithExtended(false)
	var expected models.BoardGetTopicsResponse
	fillRandomlyBoardGetTopicsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.getTopics", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardGetTopics(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKBoardGetTopicsExtendedSuccess(t *testing.T) {
	params := requests.NewBoardGetTopicsRequest()
	fillRandomlyBoardGetTopicsRequest(&params)
	params.WithExtended(true)
	var expected models.BoardGetTopicsExtendedResponse
	fillRandomlyBoardGetTopicsExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.getTopics", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardGetTopicsExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardOpenTopicRequest(r *requests.BoardOpenTopicRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
}

func TestVKBoardOpenTopicSuccess(t *testing.T) {
	params := requests.NewBoardOpenTopicRequest()
	fillRandomlyBoardOpenTopicRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.openTopic", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardOpenTopic(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardRestoreCommentRequest(r *requests.BoardRestoreCommentRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
	r.WithCommentId(random.RandInt())
}

func TestVKBoardRestoreCommentSuccess(t *testing.T) {
	params := requests.NewBoardRestoreCommentRequest()
	fillRandomlyBoardRestoreCommentRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.restoreComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardRestoreComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyBoardUnfixTopicRequest(r *requests.BoardUnfixTopicRequest) {
	r.WithGroupId(random.RandInt())
	r.WithTopicId(random.RandInt())
}

func TestVKBoardUnfixTopicSuccess(t *testing.T) {
	params := requests.NewBoardUnfixTopicRequest()
	fillRandomlyBoardUnfixTopicRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "board.unfixTopic", params.Params(), expectedJSON))
	resp, err := vk.Api.Board.BoardUnfixTopic(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
