// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/orders"
	"github.com/defany/govk/pkg/random"
	"github.com/defany/govk/vk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func fillRandomlyOrdersCancelSubscriptionRequest(r *requests.OrdersCancelSubscriptionRequest) {
	r.WithUserId(random.RandInt())
	r.WithSubscriptionId(random.RandInt())
	r.WithPendingCancel(random.RandBool())
}

func TestVKOrdersCancelSubscriptionSuccess(t *testing.T) {
	params := requests.NewOrdersCancelSubscriptionRequest()
	fillRandomlyOrdersCancelSubscriptionRequest(&params)
	var expected models.OrdersCancelSubscriptionResponse
	fillRandomlyOrdersCancelSubscriptionResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "orders.cancelSubscription", params.Params(), expectedJSON))
	resp, err := vk.Api.Orders.OrdersCancelSubscription(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyOrdersChangeStateRequest(r *requests.OrdersChangeStateRequest) {
	r.WithOrderId(random.RandInt())
	r.WithAction(random.RandString())
	r.WithAppOrderId(random.RandInt())
	r.WithTestMode(random.RandBool())
}

func TestVKOrdersChangeStateSuccess(t *testing.T) {
	params := requests.NewOrdersChangeStateRequest()
	fillRandomlyOrdersChangeStateRequest(&params)
	var expected models.OrdersChangeStateResponse
	fillRandomlyOrdersChangeStateResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "orders.changeState", params.Params(), expectedJSON))
	resp, err := vk.Api.Orders.OrdersChangeState(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyOrdersGetRequest(r *requests.OrdersGetRequest) {
	r.WithOffset(random.RandInt())
	r.WithCount(random.RandInt())
	r.WithTestMode(random.RandBool())
}

func TestVKOrdersGetSuccess(t *testing.T) {
	params := requests.NewOrdersGetRequest()
	fillRandomlyOrdersGetRequest(&params)
	var expected models.OrdersGetResponse
	fillRandomlyOrdersGetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "orders.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Orders.OrdersGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyOrdersGetAmountRequest(r *requests.OrdersGetAmountRequest) {
	r.WithUserId(random.RandInt())
	lVotes := random.RandIntn(random.MaxArrayLength + 1)
	r.WithVotes(random.RandStringArr(lVotes))
}

func TestVKOrdersGetAmountSuccess(t *testing.T) {
	params := requests.NewOrdersGetAmountRequest()
	fillRandomlyOrdersGetAmountRequest(&params)
	var expected models.OrdersGetAmountResponse
	fillRandomlyOrdersGetAmountResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "orders.getAmount", params.Params(), expectedJSON))
	resp, err := vk.Api.Orders.OrdersGetAmount(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyOrdersGetByIdRequest(r *requests.OrdersGetByIdRequest) {
	r.WithOrderId(random.RandInt())
	lOrderIds := random.RandIntn(random.MaxArrayLength + 1)
	r.WithOrderIds(random.RandIntArr(lOrderIds))
	r.WithTestMode(random.RandBool())
}

func TestVKOrdersGetByIdSuccess(t *testing.T) {
	params := requests.NewOrdersGetByIdRequest()
	fillRandomlyOrdersGetByIdRequest(&params)
	var expected models.OrdersGetByIdResponse
	fillRandomlyOrdersGetByIdResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "orders.getById", params.Params(), expectedJSON))
	resp, err := vk.Api.Orders.OrdersGetById(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyOrdersGetUserSubscriptionByIdRequest(r *requests.OrdersGetUserSubscriptionByIdRequest) {
	r.WithUserId(random.RandInt())
	r.WithSubscriptionId(random.RandInt())
}

func TestVKOrdersGetUserSubscriptionByIdSuccess(t *testing.T) {
	params := requests.NewOrdersGetUserSubscriptionByIdRequest()
	fillRandomlyOrdersGetUserSubscriptionByIdRequest(&params)
	var expected models.OrdersGetUserSubscriptionByIdResponse
	fillRandomlyOrdersGetUserSubscriptionByIdResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "orders.getUserSubscriptionById", params.Params(), expectedJSON))
	resp, err := vk.Api.Orders.OrdersGetUserSubscriptionById(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyOrdersGetUserSubscriptionsRequest(r *requests.OrdersGetUserSubscriptionsRequest) {
	r.WithUserId(random.RandInt())
}

func TestVKOrdersGetUserSubscriptionsSuccess(t *testing.T) {
	params := requests.NewOrdersGetUserSubscriptionsRequest()
	fillRandomlyOrdersGetUserSubscriptionsRequest(&params)
	var expected models.OrdersGetUserSubscriptionsResponse
	fillRandomlyOrdersGetUserSubscriptionsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "orders.getUserSubscriptions", params.Params(), expectedJSON))
	resp, err := vk.Api.Orders.OrdersGetUserSubscriptions(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyOrdersUpdateSubscriptionRequest(r *requests.OrdersUpdateSubscriptionRequest) {
	r.WithUserId(random.RandInt())
	r.WithSubscriptionId(random.RandInt())
	r.WithPrice(random.RandInt())
}

func TestVKOrdersUpdateSubscriptionSuccess(t *testing.T) {
	params := requests.NewOrdersUpdateSubscriptionRequest()
	fillRandomlyOrdersUpdateSubscriptionRequest(&params)
	var expected models.OrdersUpdateSubscriptionResponse
	fillRandomlyOrdersUpdateSubscriptionResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := random.RandString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "orders.updateSubscription", params.Params(), expectedJSON))
	resp, err := vk.Api.Orders.OrdersUpdateSubscription(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}