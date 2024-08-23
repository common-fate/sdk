package integration

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/notification/v1alpha1/notificationv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) notificationv1alpha1connect.UserNotificationSettingsServiceClient {
	return notificationv1alpha1connect.NewUserNotificationSettingsServiceClient(cfg.HTTPClient, cfg.APIURL)
}
