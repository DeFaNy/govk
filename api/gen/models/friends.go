// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type FriendsFriendExtendedStatus struct {
	FriendsFriendStatus
	// Is friend request from other user unread
	IsRequestUnread *bool `json:"is_request_unread,omitempty"`
}

type FriendsFriendStatus struct {
	FriendStatus FriendsFriendStatusStatus `json:"friend_status"`
	// MD5 hash for the result validation
	Sign *string `json:"sign,omitempty"`
	// User ID
	//  Format: int64
	//  Minimum: 1
	UserId int `json:"user_id"`
}

// FriendsFriendStatusStatus Friend status with the user
type FriendsFriendStatusStatus int

const (
	FriendsFriendStatusStatusNotAFriend FriendsFriendStatusStatus = 0
	FriendsFriendStatusStatusOutcomingRequest FriendsFriendStatusStatus = 1
	FriendsFriendStatusStatusIncomingRequest FriendsFriendStatusStatus = 2
	FriendsFriendStatusStatusIsFriend FriendsFriendStatusStatus = 3
)

type FriendsFriendsList struct {
	// List ID
	Id int `json:"id"`
	// List title
	Name string `json:"name"`
}

type FriendsMutualFriend struct {
	// Total mutual friends number
	CommonCount *int `json:"common_count,omitempty"`
	//  Minimum: 1
	CommonFriends *[]int `json:"common_friends,omitempty"`
	// User ID
	Id *int `json:"id,omitempty"`
}

type FriendsRequests struct {
	// ID of the user by whom friend has been suggested
	From *string `json:"from,omitempty"`
	Mutual *FriendsRequestsMutual `json:"mutual,omitempty"`
	// User ID
	//  Format: int64
	//  Minimum: 1
	UserId *int `json:"user_id,omitempty"`
}

type FriendsRequestsMutual struct {
	// Total mutual friends number
	//  Minimum: 0
	Count *int `json:"count,omitempty"`
	//  Minimum: 1
	Users *[]int `json:"users,omitempty"`
}

type FriendsRequestsXtrMessage struct {
	// ID of the user by whom friend has been suggested
	From *string `json:"from,omitempty"`
	// Message sent with a request
	Message *string `json:"message,omitempty"`
	Mutual *FriendsRequestsMutual `json:"mutual,omitempty"`
	// User ID
	//  Format: int64
	//  Minimum: 1
	UserId *int `json:"user_id,omitempty"`
}

type FriendsUserXtrPhone struct {
	UsersUserFull
	// User phone
	Phone *string `json:"phone,omitempty"`
}

type FriendsAddListResponse struct {
	Friends struct {
		// List ID
		//  Minimum: 1
		ListId int `json:"list_id"`
	} `json:"friends"`
}

type FriendsAddResponseResponse int

const (
	FriendsAddResponseResponseSEND FriendsAddResponseResponse = 1
	FriendsAddResponseResponseAPPROVED FriendsAddResponseResponse = 2
	FriendsAddResponseResponseRESEND FriendsAddResponseResponse = 4
)

type FriendsAddResponse struct {
	// Friend request status
	Response FriendsAddResponseResponse`json:"response"`
}

type FriendsAreFriendsExtendedResponse struct {
	Response []FriendsFriendExtendedStatus `json:"response"`
}

type FriendsAreFriendsResponse struct {
	Response []FriendsFriendStatus `json:"response"`
}

type FriendsDeleteResponseFriendsFriendDeleted int

const (
	FriendsDeleteResponseFriendsFriendDeletedOk FriendsDeleteResponseFriendsFriendDeleted = 1
)

type FriendsDeleteResponseFriendsInRequestDeleted int

const (
	FriendsDeleteResponseFriendsInRequestDeletedOk FriendsDeleteResponseFriendsInRequestDeleted = 1
)

type FriendsDeleteResponseFriendsOutRequestDeleted int

const (
	FriendsDeleteResponseFriendsOutRequestDeletedOk FriendsDeleteResponseFriendsOutRequestDeleted = 1
)

type FriendsDeleteResponseFriendsSuggestionDeleted int

const (
	FriendsDeleteResponseFriendsSuggestionDeletedOk FriendsDeleteResponseFriendsSuggestionDeleted = 1
)

type FriendsDeleteResponse struct {
	Friends struct {
		// Returns 1 if friend has been deleted
		FriendDeleted *FriendsDeleteResponseFriendsFriendDeleted`json:"friend_deleted,omitempty"`
		// Returns 1 if incoming request has been declined
		InRequestDeleted *FriendsDeleteResponseFriendsInRequestDeleted`json:"in_request_deleted,omitempty"`
		// Returns 1 if out request has been canceled
		OutRequestDeleted *FriendsDeleteResponseFriendsOutRequestDeleted`json:"out_request_deleted,omitempty"`
		//  Default: 1
		Success int `json:"success"`
		// Returns 1 if suggestion has been declined
		SuggestionDeleted *FriendsDeleteResponseFriendsSuggestionDeleted`json:"suggestion_deleted,omitempty"`
	} `json:"friends"`
}

type FriendsGetAppUsersResponse struct {
	//  Format: int64
	//  Minimum: 1
	Response []int `json:"response"`
}

type FriendsGetByPhonesResponse struct {
	Response []FriendsUserXtrPhone `json:"response"`
}

type FriendsGetListsResponse struct {
	Friends struct {
		// Total number of friends lists
		//  Minimum: 0
		Count int `json:"count"`
		Items []FriendsFriendsList `json:"items"`
	} `json:"friends"`
}

type FriendsGetMutualResponse struct {
	//  Format: int64
	//  Minimum: 1
	Response []int `json:"response"`
}

type FriendsGetMutualTargetUidsResponse struct {
	Response []FriendsMutualFriend `json:"response"`
}

type FriendsGetOnlineOnlineMobileResponse struct {
	Friends struct {
		//  Format: int64
		//  Minimum: 1
		Online []int `json:"online"`
		//  Format: int64
		//  Minimum: 1
		OnlineMobile []int `json:"online_mobile"`
	} `json:"friends"`
}

type FriendsGetOnlineResponse struct {
	//  Format: int64
	//  Minimum: 1
	Response []int `json:"response"`
}

type FriendsGetRecentResponse struct {
	//  Minimum: 1
	Response []int `json:"response"`
}

type FriendsGetRequestsExtendedResponse struct {
	Friends struct {
		// Total requests number
		//  Minimum: 0
		Count *int `json:"count,omitempty"`
		Items *[]FriendsRequestsXtrMessage `json:"items,omitempty"`
	} `json:"friends"`
}

type FriendsGetRequestsNeedMutualResponse struct {
	Friends struct {
		// Total requests number
		//  Minimum: 0
		Count *int `json:"count,omitempty"`
		Items *[]FriendsRequests `json:"items,omitempty"`
	} `json:"friends"`
}

type FriendsGetRequestsResponse struct {
	Friends struct {
		// Total requests number
		//  Minimum: 0
		Count *int `json:"count,omitempty"`
		// Total unread requests number
		CountUnread *int `json:"count_unread,omitempty"`
		//  Format: int64
		//  Minimum: 1
		Items *[]int `json:"items,omitempty"`
	} `json:"friends"`
}

type FriendsGetSuggestionsResponse struct {
	Friends struct {
		// Total results number
		//  Minimum: 0
		Count int `json:"count"`
		Items []UsersUserFull `json:"items"`
	} `json:"friends"`
}

type FriendsGetFieldsResponse struct {
	Friends struct {
		// Total friends number
		//  Minimum: 0
		Count int `json:"count"`
		Items []UsersUserFull `json:"items"`
	} `json:"friends"`
}

type FriendsGetResponse struct {
	Friends struct {
		// Total friends number
		//  Minimum: 0
		Count int `json:"count"`
		//  Format: int64
		//  Minimum: 1
		Items []int `json:"items"`
	} `json:"friends"`
}

type FriendsSearchResponse struct {
	Friends struct {
		// Total number
		//  Minimum: 0
		Count int `json:"count"`
		Items []UsersUserFull `json:"items"`
	} `json:"friends"`
}
