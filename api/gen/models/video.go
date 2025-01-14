// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type VideoLiveInfo struct {
	Enabled BaseBoolInt `json:"enabled"`
	IsNotificationsBlocked *BaseBoolInt `json:"is_notifications_blocked,omitempty"`
}

// VideoLiveSettings Video live settings
type VideoLiveSettings struct {
	// If user car rewind live or not
	CanRewind *BaseBoolInt `json:"can_rewind,omitempty"`
	// If live is endless or not
	IsEndless *BaseBoolInt `json:"is_endless,omitempty"`
	// Max possible time for rewind
	MaxDuration *int `json:"max_duration,omitempty"`
}

type VideoSaveResult struct {
	// Video access key
	AccessKey *string `json:"access_key,omitempty"`
	// Video description
	Description *string `json:"description,omitempty"`
	// Video owner ID
	//  Format: int64
	OwnerId *int `json:"owner_id,omitempty"`
	// Video title
	Title *string `json:"title,omitempty"`
	// URL for the video uploading
	//  Format: uri
	UploadUrl *string `json:"upload_url,omitempty"`
	// Video ID
	VideoId *int `json:"video_id,omitempty"`
}

type VideoVideoLiveStatus string

const (
	VideoVideoLiveStatusWaiting VideoVideoLiveStatus = "waiting"
	VideoVideoLiveStatusStarted VideoVideoLiveStatus = "started"
	VideoVideoLiveStatusFinished VideoVideoLiveStatus = "finished"
	VideoVideoLiveStatusFailed VideoVideoLiveStatus = "failed"
	VideoVideoLiveStatusUpcoming VideoVideoLiveStatus = "upcoming"
)

type VideoVideoType string

const (
	VideoVideoTypeVideo VideoVideoType = "video"
	VideoVideoTypeMusicVideo VideoVideoType = "music_video"
	VideoVideoTypeMovie VideoVideoType = "movie"
)

type VideoVideo struct {
	// Video access key
	AccessKey *string `json:"access_key,omitempty"`
	// 1 if video is added to user's albums
	Added *BaseBoolInt `json:"added,omitempty"`
	// Date when the video has been added in Unixtime
	//  Minimum: 0
	AddingDate *int `json:"adding_date,omitempty"`
	// Live donations balance
	//  Minimum: 0
	Balance *int `json:"balance,omitempty"`
	// Information whether current user can add the video
	CanAdd *BaseBoolInt `json:"can_add,omitempty"`
	// Information whether current user can add the video to favourites
	CanAddToFaves *BaseBoolInt `json:"can_add_to_faves,omitempty"`
	// Information whether current user can attach action button to the video
	CanAttachLink *BaseBoolInt `json:"can_attach_link,omitempty"`
	// Information whether current user can comment the video
	CanComment *BaseBoolInt `json:"can_comment,omitempty"`
	// Information whether current user can edit the video
	CanEdit *BaseBoolInt `json:"can_edit,omitempty"`
	// Information whether current user can like the video
	CanLike *BaseBoolInt `json:"can_like,omitempty"`
	// Information whether current user can repost the video
	CanRepost *BaseBoolInt `json:"can_repost,omitempty"`
	// Information whether current user can subscribe to author of the video
	CanSubscribe *BaseBoolInt `json:"can_subscribe,omitempty"`
	// Number of comments
	//  Minimum: 0
	Comments *int `json:"comments,omitempty"`
	// Restriction code
	//  Minimum: 0
	ContentRestricted *int `json:"content_restricted,omitempty"`
	// Restriction text
	ContentRestrictedMessage *string `json:"content_restricted_message,omitempty"`
	// 1 if  video is being converted
	Converting *BaseBoolInt `json:"converting,omitempty"`
	// Date when video has been uploaded in Unixtime
	//  Minimum: 0
	Date *int `json:"date,omitempty"`
	// Video description
	Description *string `json:"description,omitempty"`
	// Video duration in seconds
	//  Minimum: 0
	Duration *int `json:"duration,omitempty"`
	FirstFrame *[]VideoVideoImage `json:"first_frame,omitempty"`
	// Video height
	//  Minimum: 0
	Height *int `json:"height,omitempty"`
	// Video ID
	//  Minimum: 0
	Id *int `json:"id,omitempty"`
	Image *[]VideoVideoImage `json:"image,omitempty"`
	// Whether video is added to bookmarks
	IsFavorite *bool `json:"is_favorite,omitempty"`
	// 1 if video is private
	IsPrivate *BaseBoolInt `json:"is_private,omitempty"`
	// 1 if user is subscribed to author of the video
	IsSubscribed *BaseBoolInt `json:"is_subscribed,omitempty"`
	Likes *BaseLikes `json:"likes,omitempty"`
	// 1 if the video is a live stream
	Live *BasePropertyExists `json:"live,omitempty"`
	// Whether current user is subscribed to the upcoming live stream notification (if not subscribed to the author)
	LiveNotify *BaseBoolInt `json:"live_notify,omitempty"`
	// Date in Unixtime when the live stream is scheduled to start by the author
	//  Minimum: 0
	LiveStartTime *int `json:"live_start_time,omitempty"`
	// Live stream status
	LiveStatus *VideoVideoLiveStatus`json:"live_status,omitempty"`
	// If video is external, number of views on vk
	//  Minimum: 0
	LocalViews *int `json:"local_views,omitempty"`
	// Video owner ID
	//  Format: int64
	OwnerId *int `json:"owner_id,omitempty"`
	// External platform
	Platform *string `json:"platform,omitempty"`
	// Video embed URL
	//  Format: uri
	Player *string `json:"player,omitempty"`
	// Returns if the video is processing
	Processing *BasePropertyExists `json:"processing,omitempty"`
	// Information whether the video is repeated
	Repeat *BasePropertyExists `json:"repeat,omitempty"`
	Reposts *BaseRepostsInfo `json:"reposts,omitempty"`
	// Number of spectators of the stream
	//  Minimum: 0
	Spectators *int `json:"spectators,omitempty"`
	// Video title
	Title *string `json:"title,omitempty"`
	TrackCode *string `json:"track_code,omitempty"`
	Type *VideoVideoType`json:"type,omitempty"`
	// 1 if the video is an upcoming stream
	Upcoming *BasePropertyExists `json:"upcoming,omitempty"`
	// Id of the user who uploaded the video if it was uploaded to a group by member
	//  Format: int64
	//  Minimum: 0
	UserId *int `json:"user_id,omitempty"`
	// Number of views
	//  Minimum: 0
	Views *int `json:"views,omitempty"`
	// Video width
	//  Minimum: 0
	Width *int `json:"width,omitempty"`
}

type VideoVideoAlbum struct {
	// Album ID
	Id int `json:"id"`
	// Album owner's ID
	//  Format: int64
	OwnerId int `json:"owner_id"`
	// Album title
	Title string `json:"title"`
}

type VideoVideoAlbumFull struct {
	VideoVideoAlbum
	// Total number of videos in album
	//  Minimum: 0
	Count int `json:"count"`
	// Album cover image in different sizes
	Image *[]VideoVideoImage `json:"image,omitempty"`
	// Need blur album thumb or not
	ImageBlur *BasePropertyExists `json:"image_blur,omitempty"`
	// Information whether album is system
	IsSystem *BasePropertyExists `json:"is_system,omitempty"`
	// Date when the album has been updated last time in Unixtime
	//  Minimum: 0
	UpdatedTime int `json:"updated_time"`
}

type VideoVideoFiles struct {
	// URL of the external player
	//  Format: uri
	External *string `json:"external,omitempty"`
	// URL of the flv file with 320p quality
	//  Format: uri
	Flv320 *string `json:"flv_320,omitempty"`
	// URL of the mpeg4 file with 1080p quality
	//  Format: uri
	Mp41080 *string `json:"mp4_1080,omitempty"`
	// URL of the mpeg4 file with 144p quality
	//  Format: uri
	Mp4144 *string `json:"mp4_144,omitempty"`
	// URL of the mpeg4 file with 2K quality
	//  Format: uri
	Mp41440 *string `json:"mp4_1440,omitempty"`
	// URL of the mpeg4 file with 4K quality
	//  Format: uri
	Mp42160 *string `json:"mp4_2160,omitempty"`
	// URL of the mpeg4 file with 240p quality
	//  Format: uri
	Mp4240 *string `json:"mp4_240,omitempty"`
	// URL of the mpeg4 file with 360p quality
	//  Format: uri
	Mp4360 *string `json:"mp4_360,omitempty"`
	// URL of the mpeg4 file with 480p quality
	//  Format: uri
	Mp4480 *string `json:"mp4_480,omitempty"`
	// URL of the mpeg4 file with 720p quality
	//  Format: uri
	Mp4720 *string `json:"mp4_720,omitempty"`
}

type VideoVideoFull struct {
	VideoVideo
	Files *VideoVideoFiles `json:"files,omitempty"`
	// Settings for live stream
	LiveSettings *VideoLiveSettings `json:"live_settings,omitempty"`
	Trailer *VideoVideoFiles `json:"trailer,omitempty"`
}

type VideoVideoImage struct {
	BaseImage
	WithPadding *BasePropertyExists `json:"with_padding,omitempty"`
}

type VideoAddAlbumResponse struct {
	Video struct {
		// Created album ID
		AlbumId int `json:"album_id"`
	} `json:"video"`
}

type VideoChangeVideoAlbumsResponse struct {
	Response []int `json:"response"`
}

type VideoCreateCommentResponse struct {
	// Created comment ID
	Response int `json:"response"`
}

type VideoGetAlbumByIdResponse struct {
	Response VideoVideoAlbumFull `json:"response"`
}

type VideoGetAlbumsByVideoExtendedResponse struct {
	Video struct {
		// Total number
		Count *int `json:"count,omitempty"`
		Items *[]VideoVideoAlbumFull `json:"items,omitempty"`
	} `json:"video"`
}

type VideoGetAlbumsByVideoResponse struct {
	Response []int `json:"response"`
}

type VideoGetAlbumsExtendedResponse struct {
	Video struct {
		// Total number
		Count int `json:"count"`
		Items []VideoVideoAlbumFull `json:"items"`
	} `json:"video"`
}

type VideoGetAlbumsResponse struct {
	Video struct {
		// Total number
		Count int `json:"count"`
		Items []VideoVideoAlbum `json:"items"`
	} `json:"video"`
}

type VideoGetCommentsExtendedResponse struct {
	Video struct {
		// Information whether current user can comment the post
		CanPost *bool `json:"can_post,omitempty"`
		// Total number
		Count int `json:"count"`
		// Count of replies of current level
		CurrentLevelCount *int `json:"current_level_count,omitempty"`
		Groups []GroupsGroup `json:"groups"`
		// Information whether groups can comment the post
		GroupsCanPost *bool `json:"groups_can_post,omitempty"`
		Items []WallWallComment `json:"items"`
		Profiles []UsersUser `json:"profiles"`
		RealOffset *int `json:"real_offset,omitempty"`
		ShowReplyButton *bool `json:"show_reply_button,omitempty"`
	} `json:"video"`
}

type VideoGetCommentsResponse struct {
	Video struct {
		// Information whether current user can comment the post
		CanPost *bool `json:"can_post,omitempty"`
		// Total number
		Count int `json:"count"`
		// Count of replies of current level
		CurrentLevelCount *int `json:"current_level_count,omitempty"`
		// Information whether groups can comment the post
		GroupsCanPost *bool `json:"groups_can_post,omitempty"`
		Items []WallWallComment `json:"items"`
		RealOffset *int `json:"real_offset,omitempty"`
		ShowReplyButton *bool `json:"show_reply_button,omitempty"`
	} `json:"video"`
}

type VideoGetResponse struct {
	Video struct {
		// Total number
		Count int `json:"count"`
		Groups *[]GroupsGroupFull `json:"groups,omitempty"`
		Items []VideoVideoFull `json:"items"`
		Profiles *[]UsersUserFull `json:"profiles,omitempty"`
	} `json:"video"`
}

type VideoRestoreCommentResponse struct {
	// Returns 1 if request has been processed successfully, 0 if the comment is not found
	Response BaseBoolInt `json:"response"`
}

type VideoSaveResponse struct {
	Response VideoSaveResult `json:"response"`
}

type VideoSearchExtendedResponse struct {
	Video struct {
		// Total number
		Count int `json:"count"`
		Groups []GroupsGroupFull `json:"groups"`
		Items []VideoVideoFull `json:"items"`
		Profiles []UsersUser `json:"profiles"`
	} `json:"video"`
}

type VideoSearchResponse struct {
	Video struct {
		// Total number
		Count int `json:"count"`
		Items []VideoVideoFull `json:"items"`
	} `json:"video"`
}

type VideoUploadResponse struct {
	Video struct {
		// Video size
		Size *int `json:"size,omitempty"`
		// Video ID
		VideoId *int `json:"video_id,omitempty"`
	} `json:"video"`
}

