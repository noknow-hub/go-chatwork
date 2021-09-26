//////////////////////////////////////////////////////////////////////
// chatwork.go
// 
// 
// MIT License
//
// Copyright (c) 2020 noknow.info
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTW//ARE.
//////////////////////////////////////////////////////////////////////
package chatwork

import (
    "net/url"
)


const (
    ENDPOINT_ROOMS = "https://api.chatwork.com/v2/rooms"
    HTTP_METHOD_GET = "GET"
    HTTP_METHOD_POST = "POST"
    HTTP_HEADER_CONTENT_TYPE_FORM = "application/x-www-form-urlencoded"
)


type Request struct {
    ApiToken string
    FormData *url.Values
    Method string
    TimeoutSec int
    Url string
}
type Rooms struct {
    ApiToken string
    ParamBody string
    ParamSelfUnread bool
    RequestTimeoutSec int
    RoomId string
}


//////////////////////////////////////////////////////////////////////
// New request.
//////////////////////////////////////////////////////////////////////
func NewRequest(apiToken, method, url string, formData *url.Values) *Request {
    return &Request{
        ApiToken: apiToken,
        FormData: formData,
        Method: method,
        Url: url,
    }
}


//////////////////////////////////////////////////////////////////////
// New rooms client.
//////////////////////////////////////////////////////////////////////
func NewRoomsClient(apiToken string) *Rooms {
    return &Rooms{
        ApiToken: apiToken,
    }
}
