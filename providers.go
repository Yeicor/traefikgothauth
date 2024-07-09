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

var allProviders = []*ProviderInfo{
	{
		Name: "amazon",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Amazon > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = amazon.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "auth0",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Auth0 > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauth0Domain, ok := custom["auth0Domain"]; ok {
								if _, ok := custom["scopes"]; !ok {
									custom["scopes"] = make([]string, 0)
								}
								if argscopes, ok := custom["scopes"]; ok {
									provider = auth0.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argauth0Domain.(string), argscopes.([]string)...)
								}
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL auth0Domain scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "azuread",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Azuread > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argresources, ok := custom["resources"]; ok {
								if _, ok := custom["scopes"]; !ok {
									custom["scopes"] = make([]string, 0)
								}
								if argscopes, ok := custom["scopes"]; ok {
									provider = azuread.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argresources.([]string), argscopes.([]string)...)
								}
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL resources scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "battlenet",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Battlenet > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = battlenet.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "bitbucket",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Bitbucket > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = bitbucket.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "box",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Box > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = box.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "dailymotion",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Dailymotion > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = dailymotion.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "deezer",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Deezer > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = deezer.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "digitalocean",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Digitalocean > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = digitalocean.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "discord",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Discord > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = discord.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "dropbox",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Dropbox > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = dropbox.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "eveonline",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Eveonline > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = eveonline.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "facebook",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Facebook > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = facebook.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "fitbit",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Fitbit > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = fitbit.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "gitea",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Gitea > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauthURL, ok := custom["authURL"]; ok {
								if argtokenURL, ok := custom["tokenURL"]; ok {
									if argprofileURL, ok := custom["profileURL"]; ok {
										if _, ok := custom["scopes"]; !ok {
											custom["scopes"] = make([]string, 0)
										}
										if argscopes, ok := custom["scopes"]; ok {
											provider = gitea.NewCustomisedURL(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argauthURL.(string), argtokenURL.(string), argprofileURL.(string), argscopes.([]string)...)
										}
									}
								}
							}
						}
					}
				}
			}
			// Gitea > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = gitea.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientKey secret callbackURL authURL tokenURL profileURL scopes \n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "github",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Github > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauthURL, ok := custom["authURL"]; ok {
								if argtokenURL, ok := custom["tokenURL"]; ok {
									if argprofileURL, ok := custom["profileURL"]; ok {
										if argemailURL, ok := custom["emailURL"]; ok {
											if _, ok := custom["scopes"]; !ok {
												custom["scopes"] = make([]string, 0)
											}
											if argscopes, ok := custom["scopes"]; ok {
												provider = github.NewCustomisedURL(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argauthURL.(string), argtokenURL.(string), argprofileURL.(string), argemailURL.(string), argscopes.([]string)...)
											}
										}
									}
								}
							}
						}
					}
				}
			}
			// Github > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = github.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientKey secret callbackURL authURL tokenURL profileURL emailURL scopes \n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "gitlab",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Gitlab > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauthURL, ok := custom["authURL"]; ok {
								if argtokenURL, ok := custom["tokenURL"]; ok {
									if argprofileURL, ok := custom["profileURL"]; ok {
										if _, ok := custom["scopes"]; !ok {
											custom["scopes"] = make([]string, 0)
										}
										if argscopes, ok := custom["scopes"]; ok {
											provider = gitlab.NewCustomisedURL(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argauthURL.(string), argtokenURL.(string), argprofileURL.(string), argscopes.([]string)...)
										}
									}
								}
							}
						}
					}
				}
			}
			// Gitlab > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = gitlab.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientKey secret callbackURL authURL tokenURL profileURL scopes \n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "gplus",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Gplus > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = gplus.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "heroku",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Heroku > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = heroku.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "instagram",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Instagram > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = instagram.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "intercom",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Intercom > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = intercom.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "kakao",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Kakao > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = kakao.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "lastfm",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Lastfm > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							provider = lastfm.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string))
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "line",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Line > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = line.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "linkedin",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Linkedin > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = linkedin.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "mastodon",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Mastodon > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if arginstanceURL, ok := custom["instanceURL"]; ok {
								if _, ok := custom["scopes"]; !ok {
									custom["scopes"] = make([]string, 0)
								}
								if argscopes, ok := custom["scopes"]; ok {
									provider = mastodon.NewCustomisedURL(argclientKey.(string), argsecret.(string), argcallbackURL.(string), arginstanceURL.(string), argscopes.([]string)...)
								}
							}
						}
					}
				}
			}
			// Mastodon > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = mastodon.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientKey secret callbackURL instanceURL scopes \n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "meetup",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Meetup > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = meetup.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "microsoftonline",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Microsoftonline > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = microsoftonline.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "naver",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Naver > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							provider = naver.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string))
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "nextcloud",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Nextcloud > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauthURL, ok := custom["authURL"]; ok {
								if argtokenURL, ok := custom["tokenURL"]; ok {
									if argprofileURL, ok := custom["profileURL"]; ok {
										if _, ok := custom["scopes"]; !ok {
											custom["scopes"] = make([]string, 0)
										}
										if argscopes, ok := custom["scopes"]; ok {
											provider = nextcloud.NewCustomisedURL(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argauthURL.(string), argtokenURL.(string), argprofileURL.(string), argscopes.([]string)...)
										}
									}
								}
							}
						}
					}
				}
			}
			// Nextcloud > NewCustomisedDNS
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argnextcloudURL, ok := custom["nextcloudURL"]; ok {
								if _, ok := custom["scopes"]; !ok {
									custom["scopes"] = make([]string, 0)
								}
								if argscopes, ok := custom["scopes"]; ok {
									provider = nextcloud.NewCustomisedDNS(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argnextcloudURL.(string), argscopes.([]string)...)
								}
							}
						}
					}
				}
			}
			// Nextcloud > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = nextcloud.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientKey secret callbackURL authURL tokenURL profileURL scopes \n - For NewCustomisedDNS: clientKey secret callbackURL nextcloudURL scopes \n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "okta",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Okta > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientID, ok := custom["clientID"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauthURL, ok := custom["authURL"]; ok {
								if argtokenURL, ok := custom["tokenURL"]; ok {
									if argissuerURL, ok := custom["issuerURL"]; ok {
										if argprofileURL, ok := custom["profileURL"]; ok {
											if _, ok := custom["scopes"]; !ok {
												custom["scopes"] = make([]string, 0)
											}
											if argscopes, ok := custom["scopes"]; ok {
												provider = okta.NewCustomisedURL(argclientID.(string), argsecret.(string), argcallbackURL.(string), argauthURL.(string), argtokenURL.(string), argissuerURL.(string), argprofileURL.(string), argscopes.([]string)...)
											}
										}
									}
								}
							}
						}
					}
				}
			}
			// Okta > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientID, ok := custom["clientID"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argorgURL, ok := custom["orgURL"]; ok {
							if argcallbackURL, ok := custom["callbackURL"]; ok {
								if _, ok := custom["scopes"]; !ok {
									custom["scopes"] = make([]string, 0)
								}
								if argscopes, ok := custom["scopes"]; ok {
									provider = okta.New(argclientID.(string), argsecret.(string), argorgURL.(string), argcallbackURL.(string), argscopes.([]string)...)
								}
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientID secret callbackURL authURL tokenURL issuerURL profileURL scopes \n - For New: clientID secret orgURL callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "onedrive",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Onedrive > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = onedrive.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "openidConnect",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// OpenidConnect > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauthURL, ok := custom["authURL"]; ok {
								if argtokenURL, ok := custom["tokenURL"]; ok {
									if argissuerURL, ok := custom["issuerURL"]; ok {
										if arguserInfoURL, ok := custom["userInfoURL"]; ok {
											if argendSessionEndpointURL, ok := custom["endSessionEndpointURL"]; ok {
												if _, ok := custom["scopes"]; !ok {
													custom["scopes"] = make([]string, 0)
												}
												if argscopes, ok := custom["scopes"]; ok {
													var err error
													provider, err = openidConnect.NewCustomisedURL(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argauthURL.(string), argtokenURL.(string), argissuerURL.(string), arguserInfoURL.(string), argendSessionEndpointURL.(string), argscopes.([]string)...)
													if err != nil {
														return nil, err
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			// OpenidConnect > NewNamed
			if provider == nil {
				custom["callbackURL"] = callback
				if argname, ok := custom["name"]; ok {
					if argclientKey, ok := custom["clientKey"]; ok {
						if argsecret, ok := custom["secret"]; ok {
							if argcallbackURL, ok := custom["callbackURL"]; ok {
								if argopenIDAutoDiscoveryURL, ok := custom["openIDAutoDiscoveryURL"]; ok {
									if _, ok := custom["scopes"]; !ok {
										custom["scopes"] = make([]string, 0)
									}
									if argscopes, ok := custom["scopes"]; ok {
										var err error
										provider, err = openidConnect.NewNamed(argname.(string), argclientKey.(string), argsecret.(string), argcallbackURL.(string), argopenIDAutoDiscoveryURL.(string), argscopes.([]string)...)
										if err != nil {
											return nil, err
										}
									}
								}
							}
						}
					}
				}
			}
			// OpenidConnect > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argopenIDAutoDiscoveryURL, ok := custom["openIDAutoDiscoveryURL"]; ok {
								if _, ok := custom["scopes"]; !ok {
									custom["scopes"] = make([]string, 0)
								}
								if argscopes, ok := custom["scopes"]; ok {
									var err error
									provider, err = openidConnect.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argopenIDAutoDiscoveryURL.(string), argscopes.([]string)...)
									if err != nil {
										return nil, err
									}
								}
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientKey secret callbackURL authURL tokenURL issuerURL userInfoURL endSessionEndpointURL scopes \n - For NewNamed: name clientKey secret callbackURL openIDAutoDiscoveryURL scopes \n - For New: clientKey secret callbackURL openIDAutoDiscoveryURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "patreon",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Patreon > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauthURL, ok := custom["authURL"]; ok {
								if argtokenURL, ok := custom["tokenURL"]; ok {
									if argprofileURL, ok := custom["profileURL"]; ok {
										if _, ok := custom["scopes"]; !ok {
											custom["scopes"] = make([]string, 0)
										}
										if argscopes, ok := custom["scopes"]; ok {
											provider = patreon.NewCustomisedURL(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argauthURL.(string), argtokenURL.(string), argprofileURL.(string), argscopes.([]string)...)
										}
									}
								}
							}
						}
					}
				}
			}
			// Patreon > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = patreon.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientKey secret callbackURL authURL tokenURL profileURL scopes \n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "paypal",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Paypal > NewCustomisedURL
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if argauthURL, ok := custom["authURL"]; ok {
								if argtokenURL, ok := custom["tokenURL"]; ok {
									if argprofileURL, ok := custom["profileURL"]; ok {
										if _, ok := custom["scopes"]; !ok {
											custom["scopes"] = make([]string, 0)
										}
										if argscopes, ok := custom["scopes"]; ok {
											provider = paypal.NewCustomisedURL(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argauthURL.(string), argtokenURL.(string), argprofileURL.(string), argscopes.([]string)...)
										}
									}
								}
							}
						}
					}
				}
			}
			// Paypal > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = paypal.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For NewCustomisedURL: clientKey secret callbackURL authURL tokenURL profileURL scopes \n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "salesforce",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Salesforce > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = salesforce.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "seatalk",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Seatalk > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = seatalk.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "shopify",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Shopify > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = shopify.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "slack",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Slack > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = slack.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "soundcloud",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Soundcloud > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = soundcloud.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "spotify",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Spotify > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = spotify.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "steam",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Steam > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argapiKey, ok := custom["apiKey"]; ok {
					if argcallbackURL, ok := custom["callbackURL"]; ok {
						provider = steam.New(argapiKey.(string), argcallbackURL.(string))
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: apiKey callbackURL ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "strava",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Strava > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = strava.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "stripe",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Stripe > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = stripe.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "tiktok",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Tiktok > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = tiktok.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "twitch",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Twitch > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = twitch.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "twitter",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Twitter > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							provider = twitter.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string))
						}
					}
				}
			}
			// Twitter > NewAuthenticate
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							provider = twitter.NewAuthenticate(argclientKey.(string), argsecret.(string), argcallbackURL.(string))
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL \n - For NewAuthenticate: clientKey secret callbackURL ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "twitterv2",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Twitterv2 > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							provider = twitterv2.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string))
						}
					}
				}
			}
			// Twitterv2 > NewAuthenticate
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							provider = twitterv2.NewAuthenticate(argclientKey.(string), argsecret.(string), argcallbackURL.(string))
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL \n - For NewAuthenticate: clientKey secret callbackURL ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "uber",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Uber > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = uber.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "vk",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Vk > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = vk.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "wecom",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Wecom > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argcorpID, ok := custom["corpID"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argagentID, ok := custom["agentID"]; ok {
							if argcallbackURL, ok := custom["callbackURL"]; ok {
								provider = wecom.New(argcorpID.(string), argsecret.(string), argagentID.(string), argcallbackURL.(string))
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: corpID secret agentID callbackURL ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "wepay",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Wepay > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = wepay.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "xero",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Xero > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							provider = xero.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string))
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "yahoo",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Yahoo > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = yahoo.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "yammer",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Yammer > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = yammer.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "yandex",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Yandex > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = yandex.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
	{
		Name: "zoom",
		New: func(callback string, custom map[string]interface{}) (goth.Provider, error) {
			var provider goth.Provider
			// Zoom > New
			if provider == nil {
				custom["callbackURL"] = callback
				if argclientKey, ok := custom["clientKey"]; ok {
					if argsecret, ok := custom["secret"]; ok {
						if argcallbackURL, ok := custom["callbackURL"]; ok {
							if _, ok := custom["scopes"]; !ok {
								custom["scopes"] = make([]string, 0)
							}
							if argscopes, ok := custom["scopes"]; ok {
								provider = zoom.New(argclientKey.(string), argsecret.(string), argcallbackURL.(string), argscopes.([]string)...)
							}
						}
					}
				}
			}
			if provider == nil {
				return nil, fmt.Errorf("failed to create provider with parameters: %v. Required parameters:\n - For New: clientKey secret callbackURL scopes ", custom)
			}
			return provider, nil
		},
	},
}
