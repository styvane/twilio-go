package twilio

import (
	"context"
	"net/url"
)

const activityPathPart = "Activities"

type TaskRouterActivityService struct {
	client *Client
}

type Activity struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	FriendlyName string     `json:"friendly_name"`
	Available    bool       `json:"available"`
	DateCreated  TwilioTime `json:"date_created"`
	DateUpdated  TwilioTime `json:"date_updated"`
	URL          string     `json:"url"`
	WorkspaceSid string     `json:"workspace_sid"`
}

func (r *TaskRouterActivityService) Get(ctx context.Context, sid string) (*Activity, error) {
	activity := new(Activity)
	err := r.client.GetResource(ctx, activityPathPart, sid, activity)
	return activity, err
}

func (r *TaskRouterActivityService) Create(ctx context.Context, data url.Values) (*Activity, error) {
	activity := new(Activity)
	err := r.client.CreateResource(ctx, activityPathPart, data, activity)
	return activity, err
}

func (r *TaskRouterActivityService) Delete(ctx context.Context, sid string) error {
	return r.client.DeleteResource(ctx, activityPathPart, sid)
}

func (ipn *TaskRouterActivityService) Update(ctx context.Context, sid string, data url.Values) (*Activity, error) {
	activity := new(Activity)
	err := ipn.client.UpdateResource(ctx, activityPathPart, sid, data, activity)
	return activity, err
}
