// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Notifications struct {
	api *api.API
}

func NewNotifications(api *api.API) *Notifications {
	return &Notifications{
		api: api,
	}
}

// NotificationsGet Returns a list of notifications about other users' feedback to the current user's wall posts.
type NotificationsGetRequest api.Params

func NewNotificationsGetRequest() NotificationsGetRequest {
	params := make(NotificationsGetRequest, 6)
	return params
}

func (n NotificationsGetRequest) WithCount(n_count int) NotificationsGetRequest {
	n["count"] = n_count
	return n
}

func (n NotificationsGetRequest) WithStartFrom(n_start_from string) NotificationsGetRequest {
	n["start_from"] = n_start_from
	return n
}

func (n NotificationsGetRequest) WithFilters(n_filters []string) NotificationsGetRequest {
	n["filters"] = n_filters
	return n
}

func (n NotificationsGetRequest) WithStartTime(n_start_time int) NotificationsGetRequest {
	n["start_time"] = n_start_time
	return n
}

func (n NotificationsGetRequest) WithEndTime(n_end_time int) NotificationsGetRequest {
	n["end_time"] = n_end_time
	return n
}

func (n NotificationsGetRequest) Params() api.Params {
	return api.Params(n)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/notifications.get
func (n *Notifications) NotificationsGet(params ...api.MethodParams) (resp models.NotificationsGetResponse, err error) {
	req := api.NewRequest[models.NotificationsGetResponse](n.api)

	res, err := req.Execute("notifications.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// NotificationsMarkAsViewed Resets the counter of new notifications about other users' feedback to the current user's wall posts.
// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/notifications.markAsViewed
func (n *Notifications) NotificationsMarkAsViewed(params ...api.MethodParams) (resp models.NotificationsMarkAsViewedResponse, err error) {
	req := api.NewRequest[models.NotificationsMarkAsViewedResponse](n.api)

	res, err := req.Execute("notifications.markAsViewed", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// NotificationsSendMessage ...
type NotificationsSendMessageRequest api.Params

func NewNotificationsSendMessageRequest() NotificationsSendMessageRequest {
	params := make(NotificationsSendMessageRequest, 7)
	return params
}

func (n NotificationsSendMessageRequest) WithUserIds(n_user_ids []int) NotificationsSendMessageRequest {
	n["user_ids"] = n_user_ids
	return n
}

func (n NotificationsSendMessageRequest) WithMessage(n_message string) NotificationsSendMessageRequest {
	n["message"] = n_message
	return n
}

func (n NotificationsSendMessageRequest) WithFragment(n_fragment string) NotificationsSendMessageRequest {
	n["fragment"] = n_fragment
	return n
}

func (n NotificationsSendMessageRequest) WithGroupId(n_group_id int) NotificationsSendMessageRequest {
	n["group_id"] = n_group_id
	return n
}

func (n NotificationsSendMessageRequest) WithRandomId(n_random_id int) NotificationsSendMessageRequest {
	n["random_id"] = n_random_id
	return n
}

func (n NotificationsSendMessageRequest) WithSendingMode(n_sending_mode string) NotificationsSendMessageRequest {
	n["sending_mode"] = n_sending_mode
	return n
}

func (n NotificationsSendMessageRequest) Params() api.Params {
	return api.Params(n)
}

// May execute with listed access token types:
//
//	[ service ]
//
// When executing method, may return one of global or with listed codes API errors:
//
//	[ Error_GroupAppIsNotInstalledInCommunity ]
//
// https://dev.vk.com/method/notifications.sendMessage
func (n *Notifications) NotificationsSendMessage(params ...api.MethodParams) (resp models.NotificationsSendMessageResponse, err error) {
	req := api.NewRequest[models.NotificationsSendMessageResponse](n.api)

	res, err := req.Execute("notifications.sendMessage", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
