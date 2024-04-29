// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Orders struct {
	api *api.API
}

func NewOrders(api *api.API) *Orders {
	return &Orders{
		api: api,
	}
}

// OrdersCancelSubscription ...
type OrdersCancelSubscriptionRequest api.Params

func NewOrdersCancelSubscriptionRequest() OrdersCancelSubscriptionRequest {
	params := make(OrdersCancelSubscriptionRequest, 4)
	return params
}

func (o OrdersCancelSubscriptionRequest) WithUserId(o_user_id int) OrdersCancelSubscriptionRequest {
	o["user_id"] = o_user_id
	return o
}

func (o OrdersCancelSubscriptionRequest) WithSubscriptionId(o_subscription_id int) OrdersCancelSubscriptionRequest {
	o["subscription_id"] = o_subscription_id
	return o
}

func (o OrdersCancelSubscriptionRequest) WithPendingCancel(o_pending_cancel bool) OrdersCancelSubscriptionRequest {
	o["pending_cancel"] = o_pending_cancel
	return o
}

func (o OrdersCancelSubscriptionRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AppsSubscriptionNotFound, Error_AppsSubscriptionInvalidStatus ]
//
// https://dev.vk.com/method/orders.cancelSubscription
func (o *Orders) OrdersCancelSubscription(params ...api.MethodParams) (resp models.OrdersCancelSubscriptionResponse, err error) {
	req := api.NewRequest[models.OrdersCancelSubscriptionResponse](o.api)

	res, err := req.Execute("orders.cancelSubscription", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// OrdersChangeState Changes order status.
type OrdersChangeStateRequest api.Params

func NewOrdersChangeStateRequest() OrdersChangeStateRequest {
	params := make(OrdersChangeStateRequest, 5)
	return params
}

func (o OrdersChangeStateRequest) WithOrderId(o_order_id int) OrdersChangeStateRequest {
	o["order_id"] = o_order_id
	return o
}

func (o OrdersChangeStateRequest) WithAction(o_action string) OrdersChangeStateRequest {
	o["action"] = o_action
	return o
}

func (o OrdersChangeStateRequest) WithAppOrderId(o_app_order_id int) OrdersChangeStateRequest {
	o["app_order_id"] = o_app_order_id
	return o
}

func (o OrdersChangeStateRequest) WithTestMode(o_test_mode bool) OrdersChangeStateRequest {
	o["test_mode"] = o_test_mode
	return o
}

func (o OrdersChangeStateRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_Limits, Error_ActionFailed ]
//
// https://dev.vk.com/method/orders.changeState
func (o *Orders) OrdersChangeState(params ...api.MethodParams) (resp models.OrdersChangeStateResponse, err error) {
	req := api.NewRequest[models.OrdersChangeStateResponse](o.api)

	res, err := req.Execute("orders.changeState", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// OrdersGet Returns a list of orders.
type OrdersGetRequest api.Params

func NewOrdersGetRequest() OrdersGetRequest {
	params := make(OrdersGetRequest, 4)
	return params
}

func (o OrdersGetRequest) WithOffset(o_offset int) OrdersGetRequest {
	o["offset"] = o_offset
	return o
}

func (o OrdersGetRequest) WithCount(o_count int) OrdersGetRequest {
	o["count"] = o_count
	return o
}

func (o OrdersGetRequest) WithTestMode(o_test_mode bool) OrdersGetRequest {
	o["test_mode"] = o_test_mode
	return o
}

func (o OrdersGetRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/orders.get
func (o *Orders) OrdersGet(params ...api.MethodParams) (resp models.OrdersGetResponse, err error) {
	req := api.NewRequest[models.OrdersGetResponse](o.api)

	res, err := req.Execute("orders.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// OrdersGetAmount ...
type OrdersGetAmountRequest api.Params

func NewOrdersGetAmountRequest() OrdersGetAmountRequest {
	params := make(OrdersGetAmountRequest, 3)
	return params
}

func (o OrdersGetAmountRequest) WithUserId(o_user_id int) OrdersGetAmountRequest {
	o["user_id"] = o_user_id
	return o
}

func (o OrdersGetAmountRequest) WithVotes(o_votes []string) OrdersGetAmountRequest {
	o["votes"] = o_votes
	return o
}

func (o OrdersGetAmountRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/orders.getAmount
func (o *Orders) OrdersGetAmount(params ...api.MethodParams) (resp models.OrdersGetAmountResponse, err error) {
	req := api.NewRequest[models.OrdersGetAmountResponse](o.api)

	res, err := req.Execute("orders.getAmount", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// OrdersGetById Returns information about orders by their IDs.
type OrdersGetByIdRequest api.Params

func NewOrdersGetByIdRequest() OrdersGetByIdRequest {
	params := make(OrdersGetByIdRequest, 4)
	return params
}

func (o OrdersGetByIdRequest) WithOrderId(o_order_id int) OrdersGetByIdRequest {
	o["order_id"] = o_order_id
	return o
}

func (o OrdersGetByIdRequest) WithOrderIds(o_order_ids []int) OrdersGetByIdRequest {
	o["order_ids"] = o_order_ids
	return o
}

func (o OrdersGetByIdRequest) WithTestMode(o_test_mode bool) OrdersGetByIdRequest {
	o["test_mode"] = o_test_mode
	return o
}

func (o OrdersGetByIdRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/orders.getById
func (o *Orders) OrdersGetById(params ...api.MethodParams) (resp models.OrdersGetByIdResponse, err error) {
	req := api.NewRequest[models.OrdersGetByIdResponse](o.api)

	res, err := req.Execute("orders.getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// OrdersGetUserSubscriptionById ...
type OrdersGetUserSubscriptionByIdRequest api.Params

func NewOrdersGetUserSubscriptionByIdRequest() OrdersGetUserSubscriptionByIdRequest {
	params := make(OrdersGetUserSubscriptionByIdRequest, 3)
	return params
}

func (o OrdersGetUserSubscriptionByIdRequest) WithUserId(o_user_id int) OrdersGetUserSubscriptionByIdRequest {
	o["user_id"] = o_user_id
	return o
}

func (o OrdersGetUserSubscriptionByIdRequest) WithSubscriptionId(o_subscription_id int) OrdersGetUserSubscriptionByIdRequest {
	o["subscription_id"] = o_subscription_id
	return o
}

func (o OrdersGetUserSubscriptionByIdRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AppsSubscriptionNotFound ]
//
// https://dev.vk.com/method/orders.getUserSubscriptionById
func (o *Orders) OrdersGetUserSubscriptionById(params ...api.MethodParams) (resp models.OrdersGetUserSubscriptionByIdResponse, err error) {
	req := api.NewRequest[models.OrdersGetUserSubscriptionByIdResponse](o.api)

	res, err := req.Execute("orders.getUserSubscriptionById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// OrdersGetUserSubscriptions ...
type OrdersGetUserSubscriptionsRequest api.Params

func NewOrdersGetUserSubscriptionsRequest() OrdersGetUserSubscriptionsRequest {
	params := make(OrdersGetUserSubscriptionsRequest, 2)
	return params
}

func (o OrdersGetUserSubscriptionsRequest) WithUserId(o_user_id int) OrdersGetUserSubscriptionsRequest {
	o["user_id"] = o_user_id
	return o
}

func (o OrdersGetUserSubscriptionsRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/orders.getUserSubscriptions
func (o *Orders) OrdersGetUserSubscriptions(params ...api.MethodParams) (resp models.OrdersGetUserSubscriptionsResponse, err error) {
	req := api.NewRequest[models.OrdersGetUserSubscriptionsResponse](o.api)

	res, err := req.Execute("orders.getUserSubscriptions", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// OrdersUpdateSubscription ...
type OrdersUpdateSubscriptionRequest api.Params

func NewOrdersUpdateSubscriptionRequest() OrdersUpdateSubscriptionRequest {
	params := make(OrdersUpdateSubscriptionRequest, 4)
	return params
}

func (o OrdersUpdateSubscriptionRequest) WithUserId(o_user_id int) OrdersUpdateSubscriptionRequest {
	o["user_id"] = o_user_id
	return o
}

func (o OrdersUpdateSubscriptionRequest) WithSubscriptionId(o_subscription_id int) OrdersUpdateSubscriptionRequest {
	o["subscription_id"] = o_subscription_id
	return o
}

func (o OrdersUpdateSubscriptionRequest) WithPrice(o_price int) OrdersUpdateSubscriptionRequest {
	o["price"] = o_price
	return o
}

func (o OrdersUpdateSubscriptionRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_AppsSubscriptionNotFound, Error_AppsSubscriptionInvalidStatus ]
//
// https://dev.vk.com/method/orders.updateSubscription
func (o *Orders) OrdersUpdateSubscription(params ...api.MethodParams) (resp models.OrdersUpdateSubscriptionResponse, err error) {
	req := api.NewRequest[models.OrdersUpdateSubscriptionResponse](o.api)

	res, err := req.Execute("orders.updateSubscription", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
