// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type AudioAudioGenreId int

const (
	AudioAudioGenreIdRock AudioAudioGenreId = 1
	AudioAudioGenreIdPop AudioAudioGenreId = 2
	AudioAudioGenreIdRapAndHipHop AudioAudioGenreId = 3
	AudioAudioGenreIdEasyListening AudioAudioGenreId = 4
	AudioAudioGenreIdHouseAndDance AudioAudioGenreId = 5
	AudioAudioGenreIdInstrumental AudioAudioGenreId = 6
	AudioAudioGenreIdMetal AudioAudioGenreId = 7
	AudioAudioGenreIdAlternative AudioAudioGenreId = 21
	AudioAudioGenreIdDubstep AudioAudioGenreId = 8
	AudioAudioGenreIdJazzAndBlues AudioAudioGenreId = 1001
	AudioAudioGenreIdDrumAndBass AudioAudioGenreId = 10
	AudioAudioGenreIdTrance AudioAudioGenreId = 11
	AudioAudioGenreIdChanson AudioAudioGenreId = 12
	AudioAudioGenreIdEthnic AudioAudioGenreId = 13
	AudioAudioGenreIdAcousticAndVocal AudioAudioGenreId = 14
	AudioAudioGenreIdReggae AudioAudioGenreId = 15
	AudioAudioGenreIdClassical AudioAudioGenreId = 16
	AudioAudioGenreIdIndiePop AudioAudioGenreId = 17
	AudioAudioGenreIdSpeech AudioAudioGenreId = 19
	AudioAudioGenreIdElectropopAndDisco AudioAudioGenreId = 22
	AudioAudioGenreIdOther AudioAudioGenreId = 18
)

type AudioAudio struct {
	// Access key for the audio
	AccessKey *string `json:"access_key,omitempty"`
	// Album ID
	//  Minimum: 0
	AlbumId *int `json:"album_id,omitempty"`
	// Artist name
	Artist string `json:"artist"`
	// Date when uploaded
	//  Minimum: 0
	Date *int `json:"date,omitempty"`
	// Duration in seconds
	//  Minimum: 0
	Duration int `json:"duration"`
	// Genre ID
	//  Minimum: 0
	GenreId *AudioAudioGenreId`json:"genre_id,omitempty"`
	// Audio ID
	//  Minimum: 0
	Id int `json:"id"`
	// Audio owner's ID
	//  Format: int64
	OwnerId int `json:"owner_id"`
	// Performer name
	Performer *string `json:"performer,omitempty"`
	// Title
	Title string `json:"title"`
	// URL of mp3 file
	//  Format: uri
	Url *string `json:"url,omitempty"`
}

