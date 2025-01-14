// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Board struct {
	api *api.API
}

func NewBoard(api *api.API) *Board {
	return &Board{
		api: api,
	}
}

// BoardAddTopic Creates a new topic on a community's discussion board.
type BoardAddTopicRequest api.Params

func NewBoardAddTopicRequest() BoardAddTopicRequest {
	params := make(BoardAddTopicRequest, 6)
	return params
}

func (b BoardAddTopicRequest) WithGroupId(b_group_id int) BoardAddTopicRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardAddTopicRequest) WithTitle(b_title string) BoardAddTopicRequest {
	b["title"] = b_title
	return b
}

func (b BoardAddTopicRequest) WithText(b_text string) BoardAddTopicRequest {
	b["text"] = b_text
	return b
}

func (b BoardAddTopicRequest) WithFromGroup(b_from_group bool) BoardAddTopicRequest {
	b["from_group"] = b_from_group
	return b
}

func (b BoardAddTopicRequest) WithAttachments(b_attachments []string) BoardAddTopicRequest {
	b["attachments"] = b_attachments
	return b
}

func (b BoardAddTopicRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.addTopic
func (b *Board) BoardAddTopic(params ...api.MethodParams) (resp models.BoardAddTopicResponse, err error) {
	req := api.NewRequest[models.BoardAddTopicResponse](b.api)

	res, err := req.Execute("board.addTopic", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardCloseTopic Closes a topic on a community's discussion board so that comments cannot be posted.
type BoardCloseTopicRequest api.Params

func NewBoardCloseTopicRequest() BoardCloseTopicRequest {
	params := make(BoardCloseTopicRequest, 3)
	return params
}

func (b BoardCloseTopicRequest) WithGroupId(b_group_id int) BoardCloseTopicRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardCloseTopicRequest) WithTopicId(b_topic_id int) BoardCloseTopicRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardCloseTopicRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.closeTopic
func (b *Board) BoardCloseTopic(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.closeTopic", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardCreateComment Adds a comment on a topic on a community's discussion board.
type BoardCreateCommentRequest api.Params

func NewBoardCreateCommentRequest() BoardCreateCommentRequest {
	params := make(BoardCreateCommentRequest, 8)
	return params
}

func (b BoardCreateCommentRequest) WithGroupId(b_group_id int) BoardCreateCommentRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardCreateCommentRequest) WithTopicId(b_topic_id int) BoardCreateCommentRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardCreateCommentRequest) WithMessage(b_message string) BoardCreateCommentRequest {
	b["message"] = b_message
	return b
}

func (b BoardCreateCommentRequest) WithAttachments(b_attachments []string) BoardCreateCommentRequest {
	b["attachments"] = b_attachments
	return b
}

func (b BoardCreateCommentRequest) WithFromGroup(b_from_group bool) BoardCreateCommentRequest {
	b["from_group"] = b_from_group
	return b
}

func (b BoardCreateCommentRequest) WithStickerId(b_sticker_id int) BoardCreateCommentRequest {
	b["sticker_id"] = b_sticker_id
	return b
}

func (b BoardCreateCommentRequest) WithGuid(b_guid string) BoardCreateCommentRequest {
	b["guid"] = b_guid
	return b
}

func (b BoardCreateCommentRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.createComment
func (b *Board) BoardCreateComment(params ...api.MethodParams) (resp models.BoardCreateCommentResponse, err error) {
	req := api.NewRequest[models.BoardCreateCommentResponse](b.api)

	res, err := req.Execute("board.createComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardDeleteComment Deletes a comment on a topic on a community's discussion board.
type BoardDeleteCommentRequest api.Params

func NewBoardDeleteCommentRequest() BoardDeleteCommentRequest {
	params := make(BoardDeleteCommentRequest, 4)
	return params
}

func (b BoardDeleteCommentRequest) WithGroupId(b_group_id int) BoardDeleteCommentRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardDeleteCommentRequest) WithTopicId(b_topic_id int) BoardDeleteCommentRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardDeleteCommentRequest) WithCommentId(b_comment_id int) BoardDeleteCommentRequest {
	b["comment_id"] = b_comment_id
	return b
}

func (b BoardDeleteCommentRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user, group ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.deleteComment
func (b *Board) BoardDeleteComment(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.deleteComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardDeleteTopic Deletes a topic from a community's discussion board.
type BoardDeleteTopicRequest api.Params

func NewBoardDeleteTopicRequest() BoardDeleteTopicRequest {
	params := make(BoardDeleteTopicRequest, 3)
	return params
}

func (b BoardDeleteTopicRequest) WithGroupId(b_group_id int) BoardDeleteTopicRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardDeleteTopicRequest) WithTopicId(b_topic_id int) BoardDeleteTopicRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardDeleteTopicRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.deleteTopic
func (b *Board) BoardDeleteTopic(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.deleteTopic", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardEditComment Edits a comment on a topic on a community's discussion board.
type BoardEditCommentRequest api.Params

func NewBoardEditCommentRequest() BoardEditCommentRequest {
	params := make(BoardEditCommentRequest, 6)
	return params
}

func (b BoardEditCommentRequest) WithGroupId(b_group_id int) BoardEditCommentRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardEditCommentRequest) WithTopicId(b_topic_id int) BoardEditCommentRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardEditCommentRequest) WithCommentId(b_comment_id int) BoardEditCommentRequest {
	b["comment_id"] = b_comment_id
	return b
}

func (b BoardEditCommentRequest) WithMessage(b_message string) BoardEditCommentRequest {
	b["message"] = b_message
	return b
}

func (b BoardEditCommentRequest) WithAttachments(b_attachments []string) BoardEditCommentRequest {
	b["attachments"] = b_attachments
	return b
}

func (b BoardEditCommentRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.editComment
func (b *Board) BoardEditComment(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.editComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardEditTopic Edits the title of a topic on a community's discussion board.
type BoardEditTopicRequest api.Params

func NewBoardEditTopicRequest() BoardEditTopicRequest {
	params := make(BoardEditTopicRequest, 4)
	return params
}

func (b BoardEditTopicRequest) WithGroupId(b_group_id int) BoardEditTopicRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardEditTopicRequest) WithTopicId(b_topic_id int) BoardEditTopicRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardEditTopicRequest) WithTitle(b_title string) BoardEditTopicRequest {
	b["title"] = b_title
	return b
}

func (b BoardEditTopicRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.editTopic
func (b *Board) BoardEditTopic(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.editTopic", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardFixTopic Pins a topic (fixes its place) to the top of a community's discussion board.
type BoardFixTopicRequest api.Params

func NewBoardFixTopicRequest() BoardFixTopicRequest {
	params := make(BoardFixTopicRequest, 3)
	return params
}

func (b BoardFixTopicRequest) WithGroupId(b_group_id int) BoardFixTopicRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardFixTopicRequest) WithTopicId(b_topic_id int) BoardFixTopicRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardFixTopicRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.fixTopic
func (b *Board) BoardFixTopic(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.fixTopic", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardGetComments Returns a list of comments on a topic on a community's discussion board.
type BoardGetCommentsRequest api.Params

func NewBoardGetCommentsRequest() BoardGetCommentsRequest {
	params := make(BoardGetCommentsRequest, 10)
	return params
}

func (b BoardGetCommentsRequest) WithGroupId(b_group_id int) BoardGetCommentsRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardGetCommentsRequest) WithTopicId(b_topic_id int) BoardGetCommentsRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardGetCommentsRequest) WithNeedLikes(b_need_likes bool) BoardGetCommentsRequest {
	b["need_likes"] = b_need_likes
	return b
}

func (b BoardGetCommentsRequest) WithStartCommentId(b_start_comment_id int) BoardGetCommentsRequest {
	b["start_comment_id"] = b_start_comment_id
	return b
}

func (b BoardGetCommentsRequest) WithOffset(b_offset int) BoardGetCommentsRequest {
	b["offset"] = b_offset
	return b
}

func (b BoardGetCommentsRequest) WithCount(b_count int) BoardGetCommentsRequest {
	b["count"] = b_count
	return b
}

func (b BoardGetCommentsRequest) WithExtended(b_extended bool) BoardGetCommentsRequest {
	b["extended"] = b_extended
	return b
}

func (b BoardGetCommentsRequest) WithSort(b_sort string) BoardGetCommentsRequest {
	b["sort"] = b_sort
	return b
}

func (b BoardGetCommentsRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.getComments
func (b *Board) BoardGetComments(params ...api.MethodParams) (resp models.BoardGetCommentsResponse, err error) {
	req := api.NewRequest[models.BoardGetCommentsResponse](b.api)

	res, err := req.Execute("board.getComments", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardGetCommentsExtended Returns a list of comments on a topic on a community's discussion board.
// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.getComments
func (b *Board) BoardGetCommentsExtended(params ...api.MethodParams) (resp models.BoardGetCommentsExtendedResponse, err error) {
	req := api.NewRequest[models.BoardGetCommentsExtendedResponse](b.api)

	res, err := req.Execute("board.getComments", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardGetTopics Returns a list of topics on a community's discussion board.
type BoardGetTopicsRequest api.Params

func NewBoardGetTopicsRequest() BoardGetTopicsRequest {
	params := make(BoardGetTopicsRequest, 10)
	return params
}

func (b BoardGetTopicsRequest) WithGroupId(b_group_id int) BoardGetTopicsRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardGetTopicsRequest) WithTopicIds(b_topic_ids []int) BoardGetTopicsRequest {
	b["topic_ids"] = b_topic_ids
	return b
}

func (b BoardGetTopicsRequest) WithOrder(b_order int) BoardGetTopicsRequest {
	b["order"] = b_order
	return b
}

func (b BoardGetTopicsRequest) WithOffset(b_offset int) BoardGetTopicsRequest {
	b["offset"] = b_offset
	return b
}

func (b BoardGetTopicsRequest) WithCount(b_count int) BoardGetTopicsRequest {
	b["count"] = b_count
	return b
}

func (b BoardGetTopicsRequest) WithExtended(b_extended bool) BoardGetTopicsRequest {
	b["extended"] = b_extended
	return b
}

func (b BoardGetTopicsRequest) WithPreview(b_preview int) BoardGetTopicsRequest {
	b["preview"] = b_preview
	return b
}

func (b BoardGetTopicsRequest) WithPreviewLength(b_preview_length int) BoardGetTopicsRequest {
	b["preview_length"] = b_preview_length
	return b
}

func (b BoardGetTopicsRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.getTopics
func (b *Board) BoardGetTopics(params ...api.MethodParams) (resp models.BoardGetTopicsResponse, err error) {
	req := api.NewRequest[models.BoardGetTopicsResponse](b.api)

	res, err := req.Execute("board.getTopics", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardGetTopicsExtended Returns a list of topics on a community's discussion board.
// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.getTopics
func (b *Board) BoardGetTopicsExtended(params ...api.MethodParams) (resp models.BoardGetTopicsExtendedResponse, err error) {
	req := api.NewRequest[models.BoardGetTopicsExtendedResponse](b.api)

	res, err := req.Execute("board.getTopics", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardOpenTopic Re-opens a previously closed topic on a community's discussion board.
type BoardOpenTopicRequest api.Params

func NewBoardOpenTopicRequest() BoardOpenTopicRequest {
	params := make(BoardOpenTopicRequest, 3)
	return params
}

func (b BoardOpenTopicRequest) WithGroupId(b_group_id int) BoardOpenTopicRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardOpenTopicRequest) WithTopicId(b_topic_id int) BoardOpenTopicRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardOpenTopicRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.openTopic
func (b *Board) BoardOpenTopic(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.openTopic", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardRestoreComment Restores a comment deleted from a topic on a community's discussion board.
type BoardRestoreCommentRequest api.Params

func NewBoardRestoreCommentRequest() BoardRestoreCommentRequest {
	params := make(BoardRestoreCommentRequest, 4)
	return params
}

func (b BoardRestoreCommentRequest) WithGroupId(b_group_id int) BoardRestoreCommentRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardRestoreCommentRequest) WithTopicId(b_topic_id int) BoardRestoreCommentRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardRestoreCommentRequest) WithCommentId(b_comment_id int) BoardRestoreCommentRequest {
	b["comment_id"] = b_comment_id
	return b
}

func (b BoardRestoreCommentRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user, group ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.restoreComment
func (b *Board) BoardRestoreComment(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.restoreComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardUnfixTopic Unpins a pinned topic from the top of a community's discussion board.
type BoardUnfixTopicRequest api.Params

func NewBoardUnfixTopicRequest() BoardUnfixTopicRequest {
	params := make(BoardUnfixTopicRequest, 3)
	return params
}

func (b BoardUnfixTopicRequest) WithGroupId(b_group_id int) BoardUnfixTopicRequest {
	b["group_id"] = b_group_id
	return b
}

func (b BoardUnfixTopicRequest) WithTopicId(b_topic_id int) BoardUnfixTopicRequest {
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardUnfixTopicRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.unfixTopic
func (b *Board) BoardUnfixTopic(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.unfixTopic", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
