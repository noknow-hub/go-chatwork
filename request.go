//////////////////////////////////////////////////////////////////////
// request.go
//////////////////////////////////////////////////////////////////////
package chatwork

import (
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)


//////////////////////////////////////////////////////////////////////
// Post.
//////////////////////////////////////////////////////////////////////
func (o *Request) Do() ([]byte, error) {
    var req *http.Request
    var err error

    if o.FormData != nil {
        formData := *o.FormData
        req, err = http.NewRequest(o.Method, o.Url, strings.NewReader(formData.Encode()))
    } else {
        req, err = http.NewRequest(o.Method, o.Url, nil)
    }
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", HTTP_HEADER_CONTENT_TYPE_FORM)
    req.Header.Set("X-ChatWorkToken", o.ApiToken)

    client := &http.Client{}

    if o.TimeoutSec > 0 {
        client.Timeout = time.Second * time.Duration(o.TimeoutSec)
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
