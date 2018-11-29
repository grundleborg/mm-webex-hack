package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

type StartMeetingRequest struct {
	ChannelId string `json:"channel_id"`
	Personal  bool   `json:"personal"`
	Topic     string `json:"topic"`
	MeetingId int    `json:"meeting_id"`
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	//if err := p.IsConfigurationValid(); err != nil {
	//	http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
	//		return
	//}

	switch path := r.URL.Path; path {
	//case "/webhook":
	//		p.handleWebhook(w, r)
	case "/api/v1/meetings":
		p.handleStartMeeting(w, r)
	default:
		http.NotFound(w, r)
	}
}

func closeBody(r *http.Response) {
	if r.Body != nil {
		ioutil.ReadAll(r.Body)
		r.Body.Close()
	}
}

func (p *Plugin) handleStartMeeting(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("Mattermost-User-Id")

	if userId == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	var req StartMeetingRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var user *model.User
	var appErr *model.AppError
	user, appErr = p.API.GetUser(userId)
	if appErr != nil {
		http.Error(w, appErr.Error(), appErr.StatusCode)
	}

	if _, appErr := p.API.GetChannelMember(req.ChannelId, user.Id); appErr != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	cmr, statusCode, err := CreateMeeting()
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	meetingId := cmr.Body.BodyContent.MeetingKey

	hmur, statusCode, err := GetMeetingHostUrl(meetingId)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	hostUrl := hmur.Body.BodyContent.HostMeetingURL

	jmur, statusCode, err := GetMeetingJoinUrl(meetingId)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	attendeeUrl := jmur.Body.BodyContent.JoinMeetingURL

	post := &model.Post{
		UserId:    user.Id,
		ChannelId: req.ChannelId,
		Message:   fmt.Sprintf("Meeting started at %s.", hostUrl),
		Type:      "custom_zoom",
		Props: map[string]interface{}{
			"meeting_id":   meetingId,
			"host_meeting_link": hostUrl,
			"join_meeting_link": attendeeUrl,
			"meeting_status":    "STARTED",
			"meeting_personal":  false,
			"meeting_topic":     req.Topic,
			"from_webhook":      "true",
			"override_username": "WebEx",
			"override_icon_url": "https://is.wfu.edu/wp-content/uploads/2018/08/WebexMeet_Icon_only-150x150.png",
		},
	}

	if post, err := p.API.CreatePost(post); err != nil {
		http.Error(w, err.Error(), err.StatusCode)
		return
	} else {
		/*
			err = p.API.KVSet(fmt.Sprintf("%v%v", POST_MEETING_KEY, meetingId), []byte(post.Id))
			if err != nil {
				http.Error(w, err.Error(), err.StatusCode)
				return
			}
		*/
		_ = post
	}

	w.Write([]byte(fmt.Sprintf("%v", meetingId)))
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
