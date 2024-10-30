// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/market"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyMarketAddRequest(r *requests.MarketAddRequest) {
	r.WithOwnerId(random.Int())
	r.WithName(random.String())
	r.WithDescription(random.String())
	r.WithCategoryId(random.Int())
	r.WithPrice(random.MustFloat())
	r.WithOldPrice(random.MustFloat())
	r.WithDeleted(random.Bool())
	r.WithMainPhotoId(random.Int())
	lPhotoIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithPhotoIds(random.IntArr(lPhotoIds))
	r.WithUrl(random.String())
	r.WithDimensionWidth(random.Int())
	r.WithDimensionHeight(random.Int())
	r.WithDimensionLength(random.Int())
	r.WithWeight(random.Int())
	r.WithSku(random.String())
}

func TestVKMarketAddSuccess(t *testing.T) {
	params := requests.NewMarketAddRequest()
	fillRandomlyMarketAddRequest(&params)
	var expected models.MarketAddResponse
	fillRandomlyMarketAddResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.add", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketAdd(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketAddAlbumRequest(r *requests.MarketAddAlbumRequest) {
	r.WithOwnerId(random.Int())
	r.WithTitle(random.String())
	r.WithPhotoId(random.Int())
	r.WithMainAlbum(random.Bool())
	r.WithIsHidden(random.Bool())
}

func TestVKMarketAddAlbumSuccess(t *testing.T) {
	params := requests.NewMarketAddAlbumRequest()
	fillRandomlyMarketAddAlbumRequest(&params)
	var expected models.MarketAddAlbumResponse
	fillRandomlyMarketAddAlbumResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.addAlbum", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketAddAlbum(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketAddToAlbumRequest(r *requests.MarketAddToAlbumRequest) {
	r.WithOwnerId(random.Int())
	lItemIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithItemIds(random.IntArr(lItemIds))
	lAlbumIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithAlbumIds(random.IntArr(lAlbumIds))
}

func TestVKMarketAddToAlbumSuccess(t *testing.T) {
	params := requests.NewMarketAddToAlbumRequest()
	fillRandomlyMarketAddToAlbumRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.addToAlbum", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketAddToAlbum(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketCreateCommentRequest(r *requests.MarketCreateCommentRequest) {
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
	r.WithMessage(random.String())
	lAttachments := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithAttachments(random.StringArr(lAttachments))
	r.WithFromGroup(random.Bool())
	r.WithReplyToComment(random.Int())
	r.WithStickerId(random.Int())
	r.WithGuid(random.String())
}

func TestVKMarketCreateCommentSuccess(t *testing.T) {
	params := requests.NewMarketCreateCommentRequest()
	fillRandomlyMarketCreateCommentRequest(&params)
	var expected models.MarketCreateCommentResponse
	fillRandomlyMarketCreateCommentResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.createComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketCreateComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketDeleteRequest(r *requests.MarketDeleteRequest) {
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
}

func TestVKMarketDeleteSuccess(t *testing.T) {
	params := requests.NewMarketDeleteRequest()
	fillRandomlyMarketDeleteRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.delete", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketDelete(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketDeleteAlbumRequest(r *requests.MarketDeleteAlbumRequest) {
	r.WithOwnerId(random.Int())
	r.WithAlbumId(random.Int())
}

func TestVKMarketDeleteAlbumSuccess(t *testing.T) {
	params := requests.NewMarketDeleteAlbumRequest()
	fillRandomlyMarketDeleteAlbumRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.deleteAlbum", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketDeleteAlbum(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketDeleteCommentRequest(r *requests.MarketDeleteCommentRequest) {
	r.WithOwnerId(random.Int())
	r.WithCommentId(random.Int())
}

func TestVKMarketDeleteCommentSuccess(t *testing.T) {
	params := requests.NewMarketDeleteCommentRequest()
	fillRandomlyMarketDeleteCommentRequest(&params)
	var expected models.MarketDeleteCommentResponse
	fillRandomlyMarketDeleteCommentResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.deleteComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketDeleteComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketEditRequest(r *requests.MarketEditRequest) {
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
	r.WithName(random.String())
	r.WithDescription(random.String())
	r.WithCategoryId(random.Int())
	r.WithPrice(random.MustFloat())
	r.WithOldPrice(random.MustFloat())
	r.WithDeleted(random.Bool())
	r.WithMainPhotoId(random.Int())
	lPhotoIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithPhotoIds(random.IntArr(lPhotoIds))
	r.WithUrl(random.String())
	r.WithDimensionWidth(random.Int())
	r.WithDimensionHeight(random.Int())
	r.WithDimensionLength(random.Int())
	r.WithWeight(random.Int())
	r.WithSku(random.String())
}

func TestVKMarketEditSuccess(t *testing.T) {
	params := requests.NewMarketEditRequest()
	fillRandomlyMarketEditRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.edit", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketEdit(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketEditAlbumRequest(r *requests.MarketEditAlbumRequest) {
	r.WithOwnerId(random.Int())
	r.WithAlbumId(random.Int())
	r.WithTitle(random.String())
	r.WithPhotoId(random.Int())
	r.WithMainAlbum(random.Bool())
	r.WithIsHidden(random.Bool())
}

func TestVKMarketEditAlbumSuccess(t *testing.T) {
	params := requests.NewMarketEditAlbumRequest()
	fillRandomlyMarketEditAlbumRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.editAlbum", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketEditAlbum(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketEditCommentRequest(r *requests.MarketEditCommentRequest) {
	r.WithOwnerId(random.Int())
	r.WithCommentId(random.Int())
	r.WithMessage(random.String())
	lAttachments := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithAttachments(random.StringArr(lAttachments))
}

func TestVKMarketEditCommentSuccess(t *testing.T) {
	params := requests.NewMarketEditCommentRequest()
	fillRandomlyMarketEditCommentRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.editComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketEditComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketEditOrderRequest(r *requests.MarketEditOrderRequest) {
	r.WithUserId(random.Int())
	r.WithOrderId(random.Int())
	r.WithMerchantComment(random.String())
	r.WithStatus(random.Int())
	r.WithTrackNumber(random.String())
	r.WithPaymentStatus(random.String())
	r.WithDeliveryPrice(random.Int())
	r.WithWidth(random.Int())
	r.WithLength(random.Int())
	r.WithHeight(random.Int())
	r.WithWeight(random.Int())
}

func TestVKMarketEditOrderSuccess(t *testing.T) {
	params := requests.NewMarketEditOrderRequest()
	fillRandomlyMarketEditOrderRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.editOrder", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketEditOrder(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetRequest(r *requests.MarketGetRequest) {
	r.WithOwnerId(random.Int())
	r.WithAlbumId(random.Int())
	r.WithCount(random.Int())
	r.WithOffset(random.Int())
	r.WithExtended(random.Bool())
	r.WithDateFrom(random.String())
	r.WithDateTo(random.String())
	r.WithNeedVariants(random.Bool())
	r.WithWithDisabled(random.Bool())
}

func TestVKMarketGetSuccess(t *testing.T) {
	params := requests.NewMarketGetRequest()
	fillRandomlyMarketGetRequest(&params)
	params.WithExtended(false)
	var expected models.MarketGetResponse
	fillRandomlyMarketGetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKMarketGetExtendedSuccess(t *testing.T) {
	params := requests.NewMarketGetRequest()
	fillRandomlyMarketGetRequest(&params)
	params.WithExtended(true)
	var expected models.MarketGetExtendedResponse
	fillRandomlyMarketGetExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetAlbumByIdRequest(r *requests.MarketGetAlbumByIdRequest) {
	r.WithOwnerId(random.Int())
	lAlbumIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithAlbumIds(random.IntArr(lAlbumIds))
}

func TestVKMarketGetAlbumByIdSuccess(t *testing.T) {
	params := requests.NewMarketGetAlbumByIdRequest()
	fillRandomlyMarketGetAlbumByIdRequest(&params)
	var expected models.MarketGetAlbumByIdResponse
	fillRandomlyMarketGetAlbumByIdResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getAlbumById", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetAlbumById(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetAlbumsRequest(r *requests.MarketGetAlbumsRequest) {
	r.WithOwnerId(random.Int())
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
}

func TestVKMarketGetAlbumsSuccess(t *testing.T) {
	params := requests.NewMarketGetAlbumsRequest()
	fillRandomlyMarketGetAlbumsRequest(&params)
	var expected models.MarketGetAlbumsResponse
	fillRandomlyMarketGetAlbumsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getAlbums", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetAlbums(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetByIdRequest(r *requests.MarketGetByIdRequest) {
	lItemIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithItemIds(random.StringArr(lItemIds))
	r.WithExtended(random.Bool())
}

func TestVKMarketGetByIdSuccess(t *testing.T) {
	params := requests.NewMarketGetByIdRequest()
	fillRandomlyMarketGetByIdRequest(&params)
	params.WithExtended(false)
	var expected models.MarketGetByIdResponse
	fillRandomlyMarketGetByIdResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getById", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetById(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKMarketGetByIdExtendedSuccess(t *testing.T) {
	params := requests.NewMarketGetByIdRequest()
	fillRandomlyMarketGetByIdRequest(&params)
	params.WithExtended(true)
	var expected models.MarketGetByIdExtendedResponse
	fillRandomlyMarketGetByIdExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getById", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetByIdExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetCategoriesRequest(r *requests.MarketGetCategoriesRequest) {
	r.WithCount(random.Int())
	r.WithOffset(random.Int())
}

func TestVKMarketGetCategoriesSuccess(t *testing.T) {
	params := requests.NewMarketGetCategoriesRequest()
	fillRandomlyMarketGetCategoriesRequest(&params)
	var expected models.MarketGetCategoriesResponse
	fillRandomlyMarketGetCategoriesResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getCategories", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetCategories(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetCommentsRequest(r *requests.MarketGetCommentsRequest) {
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
	r.WithNeedLikes(random.Bool())
	r.WithStartCommentId(random.Int())
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
	r.WithSort(random.String())
	r.WithExtended(random.Bool())
	Fields := new([]models.UsersFields)
	lFields := random.IntDiapason(random.MaxArrayLength + 1)
	*Fields = make([]models.UsersFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyUsersFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
}

func TestVKMarketGetCommentsSuccess(t *testing.T) {
	params := requests.NewMarketGetCommentsRequest()
	fillRandomlyMarketGetCommentsRequest(&params)
	var expected models.MarketGetCommentsResponse
	fillRandomlyMarketGetCommentsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getComments", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetComments(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetGroupOrdersRequest(r *requests.MarketGetGroupOrdersRequest) {
	r.WithGroupId(random.Int())
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
}

func TestVKMarketGetGroupOrdersSuccess(t *testing.T) {
	params := requests.NewMarketGetGroupOrdersRequest()
	fillRandomlyMarketGetGroupOrdersRequest(&params)
	var expected models.MarketGetGroupOrdersResponse
	fillRandomlyMarketGetGroupOrdersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getGroupOrders", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetGroupOrders(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetOrderByIdRequest(r *requests.MarketGetOrderByIdRequest) {
	r.WithUserId(random.Int())
	r.WithOrderId(random.Int())
	r.WithExtended(random.Bool())
}

func TestVKMarketGetOrderByIdSuccess(t *testing.T) {
	params := requests.NewMarketGetOrderByIdRequest()
	fillRandomlyMarketGetOrderByIdRequest(&params)
	var expected models.MarketGetOrderByIdResponse
	fillRandomlyMarketGetOrderByIdResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getOrderById", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetOrderById(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetOrderItemsRequest(r *requests.MarketGetOrderItemsRequest) {
	r.WithUserId(random.Int())
	r.WithOrderId(random.Int())
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
}

func TestVKMarketGetOrderItemsSuccess(t *testing.T) {
	params := requests.NewMarketGetOrderItemsRequest()
	fillRandomlyMarketGetOrderItemsRequest(&params)
	var expected models.MarketGetOrderItemsResponse
	fillRandomlyMarketGetOrderItemsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getOrderItems", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetOrderItems(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketGetOrdersRequest(r *requests.MarketGetOrdersRequest) {
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
	r.WithExtended(random.Bool())
	r.WithDateFrom(random.String())
	r.WithDateTo(random.String())
}

func TestVKMarketGetOrdersSuccess(t *testing.T) {
	params := requests.NewMarketGetOrdersRequest()
	fillRandomlyMarketGetOrdersRequest(&params)
	params.WithExtended(false)
	var expected models.MarketGetOrdersResponse
	fillRandomlyMarketGetOrdersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getOrders", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetOrders(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKMarketGetOrdersExtendedSuccess(t *testing.T) {
	params := requests.NewMarketGetOrdersRequest()
	fillRandomlyMarketGetOrdersRequest(&params)
	params.WithExtended(true)
	var expected models.MarketGetOrdersExtendedResponse
	fillRandomlyMarketGetOrdersExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.getOrders", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketGetOrdersExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketRemoveFromAlbumRequest(r *requests.MarketRemoveFromAlbumRequest) {
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
	lAlbumIds := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithAlbumIds(random.IntArr(lAlbumIds))
}

func TestVKMarketRemoveFromAlbumSuccess(t *testing.T) {
	params := requests.NewMarketRemoveFromAlbumRequest()
	fillRandomlyMarketRemoveFromAlbumRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.removeFromAlbum", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketRemoveFromAlbum(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketReorderAlbumsRequest(r *requests.MarketReorderAlbumsRequest) {
	r.WithOwnerId(random.Int())
	r.WithAlbumId(random.Int())
	r.WithBefore(random.Int())
	r.WithAfter(random.Int())
}

func TestVKMarketReorderAlbumsSuccess(t *testing.T) {
	params := requests.NewMarketReorderAlbumsRequest()
	fillRandomlyMarketReorderAlbumsRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.reorderAlbums", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketReorderAlbums(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketReorderItemsRequest(r *requests.MarketReorderItemsRequest) {
	r.WithOwnerId(random.Int())
	r.WithAlbumId(random.Int())
	r.WithItemId(random.Int())
	r.WithBefore(random.Int())
	r.WithAfter(random.Int())
}

func TestVKMarketReorderItemsSuccess(t *testing.T) {
	params := requests.NewMarketReorderItemsRequest()
	fillRandomlyMarketReorderItemsRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.reorderItems", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketReorderItems(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketReportRequest(r *requests.MarketReportRequest) {
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
	r.WithReason(random.Int())
}

func TestVKMarketReportSuccess(t *testing.T) {
	params := requests.NewMarketReportRequest()
	fillRandomlyMarketReportRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.report", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketReport(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketReportCommentRequest(r *requests.MarketReportCommentRequest) {
	r.WithOwnerId(random.Int())
	r.WithCommentId(random.Int())
	r.WithReason(random.Int())
}

func TestVKMarketReportCommentSuccess(t *testing.T) {
	params := requests.NewMarketReportCommentRequest()
	fillRandomlyMarketReportCommentRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.reportComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketReportComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketRestoreRequest(r *requests.MarketRestoreRequest) {
	r.WithOwnerId(random.Int())
	r.WithItemId(random.Int())
}

func TestVKMarketRestoreSuccess(t *testing.T) {
	params := requests.NewMarketRestoreRequest()
	fillRandomlyMarketRestoreRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.restore", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketRestore(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketRestoreCommentRequest(r *requests.MarketRestoreCommentRequest) {
	r.WithOwnerId(random.Int())
	r.WithCommentId(random.Int())
}

func TestVKMarketRestoreCommentSuccess(t *testing.T) {
	params := requests.NewMarketRestoreCommentRequest()
	fillRandomlyMarketRestoreCommentRequest(&params)
	var expected models.MarketRestoreCommentResponse
	fillRandomlyMarketRestoreCommentResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.restoreComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketRestoreComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketSearchRequest(r *requests.MarketSearchRequest) {
	r.WithOwnerId(random.Int())
	r.WithAlbumId(random.Int())
	r.WithQ(random.String())
	r.WithPriceFrom(random.Int())
	r.WithPriceTo(random.Int())
	r.WithSort(random.Int())
	r.WithRev(random.Int())
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
	r.WithExtended(random.Bool())
	lStatus := random.IntDiapason(random.MaxArrayLength + 1)
	r.WithStatus(random.IntArr(lStatus))
	r.WithNeedVariants(random.Bool())
}

func TestVKMarketSearchSuccess(t *testing.T) {
	params := requests.NewMarketSearchRequest()
	fillRandomlyMarketSearchRequest(&params)
	params.WithExtended(false)
	var expected models.MarketSearchResponse
	fillRandomlyMarketSearchResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.search", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketSearch(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKMarketSearchExtendedSuccess(t *testing.T) {
	params := requests.NewMarketSearchRequest()
	fillRandomlyMarketSearchRequest(&params)
	params.WithExtended(true)
	var expected models.MarketSearchExtendedResponse
	fillRandomlyMarketSearchExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.search", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketSearchExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyMarketSearchItemsRequest(r *requests.MarketSearchItemsRequest) {
	r.WithQ(random.String())
	r.WithOffset(random.Int())
	r.WithCount(random.Int())
	r.WithCategoryId(random.Int())
	r.WithPriceFrom(random.Int())
	r.WithPriceTo(random.Int())
	r.WithSortBy(random.Int())
	r.WithSortDirection(random.Int())
	r.WithCountry(random.Int())
	r.WithCity(random.Int())
}

func TestVKMarketSearchItemsSuccess(t *testing.T) {
	params := requests.NewMarketSearchItemsRequest()
	fillRandomlyMarketSearchItemsRequest(&params)
	var expected models.MarketSearchResponse
	fillRandomlyMarketSearchResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.String()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "market.searchItems", params.Params(), expectedJSON))
	resp, err := vk.Api.Market.MarketSearchItems(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
