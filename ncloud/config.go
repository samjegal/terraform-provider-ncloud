package ncloud

import (
	"time"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/autoscaling"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/cdn"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/clouddb"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/loadbalancer"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/monitoring"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/server"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vpc"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vserver"
)

// DefaultWaitForInterval is Interval for checking status in WaitForXXX method
const DefaultWaitForInterval = 10

// Default timeout
const DefaultTimeout = 5 * time.Minute
const DefaultCreateTimeout = 1 * time.Hour
const DefaultUpdateTimeout = 10 * time.Minute
const DefaultStopTimeout = 5 * time.Minute

type Config struct {
	AccessKey string
	SecretKey string
	Site      string
}

type NcloudAPIClient struct {
	site         string
	server       *server.APIClient
	autoscaling  *autoscaling.APIClient
	loadbalancer *loadbalancer.APIClient
	cdn          *cdn.APIClient
	clouddb      *clouddb.APIClient
	monitoring   *monitoring.APIClient
	vpc          *vpc.APIClient
	vserver      *vserver.APIClient
}

func (c *Config) Client() (*NcloudAPIClient, error) {
	apiKey := &ncloud.APIKey{
		AccessKey: c.AccessKey,
		SecretKey: c.SecretKey,
	}
	return &NcloudAPIClient{
		site:         c.Site,
		server:       server.NewAPIClient(server.NewConfiguration(apiKey)),
		autoscaling:  autoscaling.NewAPIClient(autoscaling.NewConfiguration(apiKey)),
		loadbalancer: loadbalancer.NewAPIClient(loadbalancer.NewConfiguration(apiKey)),
		cdn:          cdn.NewAPIClient(cdn.NewConfiguration(apiKey)),
		clouddb:      clouddb.NewAPIClient(clouddb.NewConfiguration(apiKey)),
		monitoring:   monitoring.NewAPIClient(monitoring.NewConfiguration(apiKey)),
		vpc:          vpc.NewAPIClient(vpc.NewConfiguration(apiKey)),
		vserver:      vserver.NewAPIClient(vserver.NewConfiguration(apiKey)),
	}, nil
}
