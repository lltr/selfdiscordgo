// Discordgo - Discord bindings for Go
// Available at https://github.com/bwmarrin/discordgo

// Copyright 2015-2016 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains high level helper functions and easy entry points for the
// entire discordgo package.  These functions are being developed and are very
// experimental at this point.  They will most likely change so please use the
// low level functions if that's a problem.

// package selfdiscordgo provides Discord binding for Go
package selfdiscordgo

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// VERSION of DiscordGo, follows Semantic Versioning. (http://semver.org/)
const VERSION = "0.27.7"

// New creates a new Discord session with provided token.
// If the token is for a bot, it must be prefixed with "Bot "
//
//	e.g. "Bot ..."
//
// Or if it is an OAuth2 token, it must be prefixed with "Bearer "
//
//	e.g. "Bearer ..."
func New(token string) (s *Session, err error) {

	// Create an empty Session interface.
	s = &Session{
		State:                              NewState(),
		Ratelimiter:                        NewRatelimiter(),
		StateEnabled:                       true,
		ShouldReconnectOnError:             true,
		ShouldReconnectVoiceOnSessionError: true,
		ShouldRetryOnRateLimit:             true,
		MaxRestRetries:                     3,
		Client:                             &http.Client{Timeout: (20 * time.Second)},
		Dialer:                             websocket.DefaultDialer,
		UserAgent:                          "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		sequence:                           new(int64),
		LastHeartbeatAck:                   time.Now().UTC(),
	}

	// Initialize the Identify Package with defaults
	// These can be modified prior to calling Open()
	s.Identify.Compress = true
	s.Identify.Token = token
	s.Identify.Properties = IdentifyProperties{
		// Windows, Linux etc.
		OS: "Mac OS X",
		// Brave, Safari etc.
		Browser: "Chrome",
		// Perhaps filled if its a phone like iPhone or Samsung ? idk
		Device: "",
		// en-US, de-DE etc.
		SystemLocale: "en-AU",
		// Any user agent that matches OS and Browser
		BrowserUserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.61 Safari/537.36",
		BrowserVersion:   "109.0.5414.61",
		// This is actually NOT my OS version, idk how they get this value,
		// perhaps its the OS version you signed up to your account with,
		// but then how does the webssite know to send that in the packet ? idk
		OSVersion: "10.15.7",
		// idk
		Referrer:               "",
		ReferringDomain:        "",
		ReferrerCurrent:        "",
		ReferringDomainCurrent: "",
		ReleaseChannel:         "stable",
		ClientBuildNumber:      94294,
		ClientEventSource:      nil,
	}
	s.Identify.Presence = IdentifyPresence{
		// Online etcc.
		Status: "invisible",
		// Leave at 0
		Since: 0,
		// idk how you would fill this in
		Activities: nil,
		// idk what this does
		AFK: false,
	}
	s.Identify.Compress = false
	// No clue, I suggest you do not touch this.
	s.Identify.ClientState = IdentifyClientState{
		GuildHashes:              struct{}{},
		HighestLastMessageID:     "0",
		ReadStateVersion:         0,
		UserGuildSettingsVersion: -1,
	}
	// 61 for Discord client and 125 for browser?
	// Reddit people report seeing 61, I saw 125 on browser.
	s.Identify.Capabilities = 125
	s.Identify.Token = token

	return
}
