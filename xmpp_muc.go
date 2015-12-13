// Copyright 2013 Flo Lauber <dev@qatfy.at>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(flo):
//   - support password protected MUC rooms
//   - cleanup signatures of join/leave functions
package xmpp

import (
	"fmt"
)

const (
	nsMUC     = "http://jabber.org/protocol/muc"
	nsMUCUser = "http://jabber.org/protocol/muc#user"
)

// xep-0045 7.2
func (c *Client) JoinMUC(jid, nick string) {
	if nick == "" {
		nick = c.jid
	}
	fmt.Fprintf(c.conn, "<presence to='%s/%s'>\n"+
		"<x xmlns='%s' />\n"+
		"</presence>",
		xmlEscape(jid), xmlEscape(nick), nsMUC)
}

// xep-0045 7.2.6
func (c *Client) JoinProtectedMUC(jid, nick string, password string) {
	if nick == "" {
		nick = c.jid
	}
	fmt.Fprintf(c.conn, "<presence to='%s/%s'>\n"+
		"<x xmlns='%s'>\n"+
		"<password>%s</password>\n"+
		"</x>\n"+
		"</presence>",
		xmlEscape(jid), xmlEscape(nick), nsMUC, xmlEscape(password))
}

// xep-0045 7.14
func (c *Client) LeaveMUC(jid string) {
	fmt.Fprintf(c.conn, "<presence from='%s' to='%s' type='unavailable' />",
		c.jid, xmlEscape(jid))
}

//xep-0249 2.1
func (c *Client) InviteToMUC(srcJid, destJid, roomJid, password, reason string) {
	if password != "" {
		password = "password='" + xmlEscape(password) + "'"
	}
	if reason != "" {
		reason = "reason='" + xmlEscape(reason) + "'"
	}
	fmt.Fprintf(c.conn, "<message from='%s' to='%s'>\n"+
		"<x xmlns='jabber:x:conference'"+
		"jid='%s' %s %s/>"+
		"</message>",
		xmlEscape(srcJid), xmlEscaoe(destJid), xmlEscape(roomJid), password, reason)
}
