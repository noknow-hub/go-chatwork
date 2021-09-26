//////////////////////////////////////////////////////////////////////
// rooms.go
//////////////////////////////////////////////////////////////////////
package chatwork

import (
    "net/url"
)


//////////////////////////////////////////////////////////////////////
// POST: messages.
// Send a message to the "room_id".
//////////////////////////////////////////////////////////////////////
func (o *Rooms) PostMessages(roomId string) ([]byte, error) {
    reqUrl := ENDPOINT_ROOMS + "/" + roomId + "/messages"

    formData := &url.Values{}
    formData.Add("body", o.ParamBody)
    if o.ParamSelfUnread {
        formData.Add("self_unread", "1")
    } else {
        formData.Add("self_unread", "0")
    }

    r := NewRequest(o.ApiToken, HTTP_METHOD_POST, reqUrl, formData)
    r.TimeoutSec = o.RequestTimeoutSec

    return r.Do()
}


//////////////////////////////////////////////////////////////////////
// Set "ParamBody".
//////////////////////////////////////////////////////////////////////
func (o *Rooms) SetParamBody(body string) *Rooms {
    o.ParamBody = body
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "ParamSelfUnread".
//////////////////////////////////////////////////////////////////////
func (o *Rooms) SetParamSelfUnread(selfUnread bool) *Rooms {
    o.ParamSelfUnread = selfUnread
    return o
}


//////////////////////////////////////////////////////////////////////
// Set "RequestTimeoutSec".
//////////////////////////////////////////////////////////////////////
func (o *Rooms) SetRequestTimeoutSec(requestTimeoutSec int) *Rooms {
    o.RequestTimeoutSec = requestTimeoutSec
    return o
}