package traefikgothauth

import (
	"fmt"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/amazon"
	"github.com/markbates/goth/providers/auth0"
	"github.com/markbates/goth/providers/azuread"
	"github.com/markbates/goth/providers/battlenet"
	"github.com/markbates/goth/providers/bitbucket"
	"github.com/markbates/goth/providers/box"
	"github.com/markbates/goth/providers/dailymotion"
	"github.com/markbates/goth/providers/deezer"
	"github.com/markbates/goth/providers/digitalocean"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/dropbox"
	"github.com/markbates/goth/providers/eveonline"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/fitbit"
	"github.com/markbates/goth/providers/gitea"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/heroku"
	"github.com/markbates/goth/providers/instagram"
	"github.com/markbates/goth/providers/intercom"
	"github.com/markbates/goth/providers/kakao"
	"github.com/markbates/goth/providers/lastfm"
	"github.com/markbates/goth/providers/line"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/mastodon"
	"github.com/markbates/goth/providers/meetup"
	"github.com/markbates/goth/providers/microsoftonline"
	"github.com/markbates/goth/providers/naver"
	"github.com/markbates/goth/providers/nextcloud"
	"github.com/markbates/goth/providers/okta"
	"github.com/markbates/goth/providers/onedrive"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/markbates/goth/providers/patreon"
	"github.com/markbates/goth/providers/paypal"
	"github.com/markbates/goth/providers/salesforce"
	"github.com/markbates/goth/providers/seatalk"
	"github.com/markbates/goth/providers/shopify"
	"github.com/markbates/goth/providers/slack"
	"github.com/markbates/goth/providers/soundcloud"
	"github.com/markbates/goth/providers/spotify"
	"github.com/markbates/goth/providers/steam"
	"github.com/markbates/goth/providers/strava"
	"github.com/markbates/goth/providers/stripe"
	"github.com/markbates/goth/providers/tiktok"
	"github.com/markbates/goth/providers/twitch"
	"github.com/markbates/goth/providers/twitter"
	"github.com/markbates/goth/providers/twitterv2"
	"github.com/markbates/goth/providers/uber"
	"github.com/markbates/goth/providers/vk"
	"github.com/markbates/goth/providers/wecom"
	"github.com/markbates/goth/providers/wepay"
	"github.com/markbates/goth/providers/xero"
	"github.com/markbates/goth/providers/yahoo"
	"github.com/markbates/goth/providers/yammer"
	"github.com/markbates/goth/providers/yandex"
	"github.com/markbates/goth/providers/zoom"
)

// NewProvider creates a New provider based on the given Name and parameters.
func NewProvider(name, clientclientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
	p, ok := getProviderInfo(name)
	if !ok {
		return nil, fmt.Errorf("provider %s not found", name)
	}
	return p.New(clientclientKey, secret, callback, custom, scopes...)
}

func getProviderInfo(name string) (*ProviderInfo, bool) {
	for _, p := range allProviders {
		if p.Name == name {
			return p, true
		}
	}
	return nil, false
}

// ProviderInfo contains static metadata for a provider.
type ProviderInfo struct {
	Name, DisplayName, Icon string
	New                     func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error)
}

// allProviders is a list of static metadata for all supported providers.
//
// Based on https://github.com/markbates/goth/blob/master/examples/main.go
var allProviders = []*ProviderInfo{
	{
		Name:        "amazon",
		DisplayName: "Amazon",
		Icon:        "https://icons.duckduckgo.com/ip3/amazon.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return amazon.New(clientKey, secret, callback, scopes...), nil
		},
	},
	//{ // Incompatible with yaegi (for now...)
	//	Name:        "apple",
	//	DisplayName: "Apple",
	//	Icon:        "https://icons.duckduckgo.com/ip3/apple.com.ico",
	//	New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
	//		client, ok := custom["client"].(*http.Client)
	//		if !ok {
	//			client = http.DefaultClient
	//		}
	//		return apple.New(clientKey, secret, callback, client, scopes...), nil
	//	},
	//},
	{
		Name:        "auth0",
		DisplayName: "Auth0",
		Icon:        "https://icons.duckduckgo.com/ip3/auth0.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return auth0.New(clientKey, secret, callback, custom["domain"].(string), scopes...), nil
		},
	},
	{
		Name:        "azuread",
		DisplayName: "Azure AD",
		Icon:        "https://icons.duckduckgo.com/ip3/azure.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			resources, ok := custom["resources"].([]string)
			if !ok {
				resources = []string{}
			}
			return azuread.New(clientKey, secret, callback, resources, scopes...), nil
		},
	},
	{
		Name:        "battlenet",
		DisplayName: "Battle.net",
		Icon:        "https://icons.duckduckgo.com/ip3/battle.net.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return battlenet.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "bitbucket",
		DisplayName: "Bitbucket",
		Icon:        "https://icons.duckduckgo.com/ip3/bitbucket.org.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return bitbucket.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "box",
		DisplayName: "Box",
		Icon:        "https://icons.duckduckgo.com/ip3/box.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return box.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "dailymotion",
		DisplayName: "Dailymotion",
		Icon:        "https://icons.duckduckgo.com/ip3/dailymotion.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return dailymotion.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "deezer",
		DisplayName: "Deezer",
		Icon:        "https://icons.duckduckgo.com/ip3/deezer.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return deezer.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "digitalocean",
		DisplayName: "DigitalOcean",
		Icon:        "https://icons.duckduckgo.com/ip3/digitalocean.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return digitalocean.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "discord",
		DisplayName: "Discord",
		Icon:        "https://icons.duckduckgo.com/ip3/discord.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return discord.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "dropbox",
		DisplayName: "Dropbox",
		Icon:        "https://icons.duckduckgo.com/ip3/dropbox.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return dropbox.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "eveonline",
		DisplayName: "EVE Online",
		Icon:        "https://icons.duckduckgo.com/ip3/eveonline.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return eveonline.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "facebook",
		DisplayName: "Facebook",
		Icon:        "https://icons.duckduckgo.com/ip3/facebook.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return facebook.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "fitbit",
		DisplayName: "Fitbit",
		Icon:        "https://icons.duckduckgo.com/ip3/fitbit.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return fitbit.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "gitea",
		DisplayName: "Gitea",
		Icon:        "https://icons.duckduckgo.com/ip3/gitea.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return gitea.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "github",
		DisplayName: "GitHub",
		Icon:        "https://icons.duckduckgo.com/ip3/github.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return github.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "gitlab",
		DisplayName: "GitLab",
		Icon:        "https://icons.duckduckgo.com/ip3/gitlab.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return gitlab.New(clientKey, secret, callback, scopes...), nil
		},
	},
	//{ // Incompatible with yaegi (for now...)
	//	Name:        "google",
	//	DisplayName: "Google",
	//	Icon:        "https://icons.duckduckgo.com/ip3/google.com.ico",
	//	New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
	//		scopes, ok := custom["scopes"].([]string)
	//		if !ok {
	//			scopes = []string{}
	//		}
	//		return google.New(clientKey, secret, callback, scopes...), nil
	//	},
	//},
	{
		Name:        "gplus",
		DisplayName: "Google Plus",
		Icon:        "https://icons.duckduckgo.com/ip3/google.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return gplus.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "heroku",
		DisplayName: "Heroku",
		Icon:        "https://icons.duckduckgo.com/ip3/heroku.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return heroku.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "instagram",
		DisplayName: "Instagram",
		Icon:        "https://icons.duckduckgo.com/ip3/instagram.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return instagram.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "intercom",
		DisplayName: "Intercom",
		Icon:        "https://icons.duckduckgo.com/ip3/intercom.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return intercom.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "kakao",
		DisplayName: "Kakao",
		Icon:        "https://icons.duckduckgo.com/ip3/kakao.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return kakao.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "lastfm",
		DisplayName: "Last.fm",
		Icon:        "https://icons.duckduckgo.com/ip3/last.fm.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return lastfm.New(clientKey, secret, callback), nil
		},
	},
	{
		Name:        "line",
		DisplayName: "Line",
		Icon:        "https://icons.duckduckgo.com/ip3/line.me.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return line.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "linkedin",
		DisplayName: "LinkedIn",
		Icon:        "https://icons.duckduckgo.com/ip3/linkedin.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return linkedin.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "mastodon",
		DisplayName: "Mastodon",
		Icon:        "https://icons.duckduckgo.com/ip3/mastodon.social.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return mastodon.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "meetup",
		DisplayName: "Meetup",
		Icon:        "https://icons.duckduckgo.com/ip3/meetup.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return meetup.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "microsoftonline",
		DisplayName: "Microsoft Online",
		Icon:        "https://icons.duckduckgo.com/ip3/www.microsoft.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return microsoftonline.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "naver",
		DisplayName: "Naver",
		Icon:        "https://icons.duckduckgo.com/ip3/naver.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return naver.New(clientKey, secret, callback), nil
		},
	},
	{
		Name:        "nextcloud",
		DisplayName: "Nextcloud",
		Icon:        "https://icons.duckduckgo.com/ip3/nextcloud.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			if nextcloudURL, ok := custom["nextcloudURL"].(string); ok {
				return nextcloud.NewCustomisedDNS(clientKey, secret, callback, nextcloudURL, scopes...), nil
			} else {
				return nextcloud.NewCustomisedURL(clientKey, secret, callback, custom["authURL"].(string), custom["tokenURL"].(string), custom["profileURL"].(string), scopes...), nil
			}
		},
	},
	{
		Name:        "okta",
		DisplayName: "Okta",
		Icon:        "https://icons.duckduckgo.com/ip3/okta.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return okta.New(clientKey, secret, custom["orgURL"].(string), callback), nil
		},
	},
	{
		Name:        "onedrive",
		DisplayName: "OneDrive",
		Icon:        "https://icons.duckduckgo.com/ip3/onedrive.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return onedrive.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "openid-connect",
		DisplayName: "OpenID Connect",
		Icon:        "https://icons.duckduckgo.com/ip3/openid.net.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			var (
				c   *openidConnect.Provider
				err error
			)
			if autoDiscoveryURL, ok := custom["openIDAutoDiscoveryURL"].(string); ok {
				c, err = openidConnect.New(clientKey, secret, callback, autoDiscoveryURL, scopes...)
			} else {
				c, err = openidConnect.NewCustomisedURL(clientKey, secret, callback, custom["authURL"].(string), custom["tokenURL"].(string), custom["issuerURL"].(string), custom["userInfoURL"].(string), custom["endSessionEndpointURL"].(string), scopes...)
			}
			if err != nil {
				return nil, err
			}
			if skipUserInfoRequest, ok := custom["skipUserInfoRequest"].(bool); ok {
				c.SkipUserInfoRequest = skipUserInfoRequest
			}
			return c, err
		},
	},
	{
		Name:        "patreon",
		DisplayName: "Patreon",
		Icon:        "https://icons.duckduckgo.com/ip3/patreon.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return patreon.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "paypal",
		DisplayName: "PayPal",
		Icon:        "https://icons.duckduckgo.com/ip3/paypal.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return paypal.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "salesforce",
		DisplayName: "Salesforce",
		Icon:        "https://icons.duckduckgo.com/ip3/salesforce.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return salesforce.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "seatalk",
		DisplayName: "SeaTalk",
		Icon:        "https://icons.duckduckgo.com/ip3/seatalk.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return seatalk.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "shopify",
		DisplayName: "Shopify",
		Icon:        "https://icons.duckduckgo.com/ip3/shopify.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return shopify.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "slack",
		DisplayName: "Slack",
		Icon:        "https://icons.duckduckgo.com/ip3/slack.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return slack.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "soundcloud",
		DisplayName: "SoundCloud",
		Icon:        "https://icons.duckduckgo.com/ip3/soundcloud.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return soundcloud.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "spotify",
		DisplayName: "Spotify",
		Icon:        "https://icons.duckduckgo.com/ip3/spotify.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return spotify.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "steam",
		DisplayName: "Steam",
		Icon:        "https://icons.duckduckgo.com/ip3/steamcommunity.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return steam.New(clientKey, callback), nil
		},
	},
	{
		Name:        "strava",
		DisplayName: "Strava",
		Icon:        "https://icons.duckduckgo.com/ip3/strava.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return strava.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "stripe",
		DisplayName: "Stripe",
		Icon:        "https://icons.duckduckgo.com/ip3/stripe.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return stripe.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "tiktok",
		DisplayName: "TikTok",
		Icon:        "https://icons.duckduckgo.com/ip3/tiktok.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return tiktok.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "twitch",
		DisplayName: "Twitch",
		Icon:        "https://icons.duckduckgo.com/ip3/twitch.tv.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return twitch.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "twitter",
		DisplayName: "Twitter (v1)",
		Icon:        "https://icons.duckduckgo.com/ip3/twitter.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return twitter.New(clientKey, secret, callback), nil
		},
	},
	{
		Name:        "twitterv2",
		DisplayName: "Twitter (v2)",
		Icon:        "https://icons.duckduckgo.com/ip3/twitter.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return twitterv2.New(clientKey, secret, callback), nil
		},
	},
	//{ // https://nulab.com/blog/company-news/typetalk-sunsetting/
	//	Name:        "typetalk",
	//	DisplayName: "TypeTalk",
	//	Icon:   "https://icons.duckduckgo.com/ip3/typetalk.com.ico",
	//	New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
	//		return typetalk.New(clientKey, secret, callback, scopes...), nil
	//	},
	//},
	{
		Name:        "uber",
		DisplayName: "Uber",
		Icon:        "https://icons.duckduckgo.com/ip3/uber.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return uber.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "vk",
		DisplayName: "VK",
		Icon:        "https://icons.duckduckgo.com/ip3/vk.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return vk.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "wecom",
		DisplayName: "WeCom",
		Icon:        "https://icons.duckduckgo.com/ip3/wework.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return wecom.New(clientKey, secret, custom["agentID"].(string), callback), nil
		},
	},
	{
		Name:        "wepay",
		DisplayName: "WePay",
		Icon:        "https://icons.duckduckgo.com/ip3/wepay.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return wepay.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "xero",
		DisplayName: "Xero",
		Icon:        "https://icons.duckduckgo.com/ip3/xero.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return xero.New(clientKey, secret, callback), nil
		},
	},
	{
		Name:        "yahoo",
		DisplayName: "Yahoo",
		Icon:        "https://icons.duckduckgo.com/ip3/yahoo.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return yahoo.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "yammer",
		DisplayName: "Yammer",
		Icon:        "https://icons.duckduckgo.com/ip3/yammer.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return yammer.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "yandex",
		DisplayName: "Yandex",
		Icon:        "https://icons.duckduckgo.com/ip3/yandex.com.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return yandex.New(clientKey, secret, callback, scopes...), nil
		},
	},
	{
		Name:        "zoom",
		DisplayName: "Zoom",
		Icon:        "https://icons.duckduckgo.com/ip3/zoom.us.ico",
		New: func(clientKey, secret, callback string, custom map[string]interface{}, scopes ...string) (goth.Provider, error) {
			return zoom.New(clientKey, secret, callback, scopes...), nil
		},
	},
}
