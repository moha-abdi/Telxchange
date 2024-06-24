package config

const (
	MFS_PROXY_25261     = "https://evcplus.waafi.com"
	MFS_PROXY_25263     = "https://zaad.waafi.com"
	MFS_PROXY_25267     = "https://cashplus.waafi.com"
	MFS_PROXY_25268     = "https://jeeb.waafi.com"
	MFS_PROXY_2529      = "https://sahal.waafi.com"
	MFS_PROXY_253       = "https://sabproxy.waafi.com"
	MFS_PROXY_DEFAULT   = "https://proxy.waafi.com"
	TELECOM_PROXY_25261 = "https://telecom.hormuud.com"
	TELECOM_PROXY_25263 = "https://zaad.waafi.com"
	TELECOM_PROXY_25264 = "https://cashplus.waafi.com"
	TELECOM_PROXY_25268 = "https://telecom.hormuud.com"
	TELECOM_PROXY_2529  = "https://sahal.waafi.com"
)

type (
	ChannelName string
	AuthMode    string
)

const (
	Web       ChannelName = "Web"
	MobileApp ChannelName = "MobileApp"
)

const (
	Pin      AuthMode = "pin"
	Password AuthMode = "password"
)

const (
	DefaultChannel  = MobileApp
	DefaultAuthMode = Pin
)
