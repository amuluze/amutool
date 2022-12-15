// Package requests
// Date: 2022/9/1 23:06
// Author: Amu
// Description:
package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Responses struct {
	Ok          bool
	Error       error
	RawResponse *http.Response
	StatusCode  int
	Header      http.Header

	internalByteBuffer *bytes.Buffer
}

func buildResponses(resp *http.Response, err error) (*Responses, error) {
	if err != nil {
		return &Responses{Error: err}, err
	}

	goodResp := &Responses{
		Ok:                 resp.StatusCode >= 200 && resp.StatusCode < 300,
		Error:              nil,
		RawResponse:        resp,
		StatusCode:         resp.StatusCode,
		Header:             resp.Header,
		internalByteBuffer: bytes.NewBuffer([]byte{}),
	}

	return goodResp, nil
}

func (r *Responses) Read(p []byte) (n int, err error) {
	if r.Error != nil {
		return -1, r.Error
	}
	return r.RawResponse.Body.Read(p)
}

func (r *Responses) Close() error {
	if r.Error != nil {
		return r.Error
	}
	io.Copy(ioutil.Discard, r)
	return r.RawResponse.Body.Close()
}

func (r *Responses) getInternalReader() io.Reader {
	if r.internalByteBuffer.Len() != 0 {
		return r.internalByteBuffer
	}
	return r
}

func (r *Responses) JSON(userStruct interface{}) error {
	if r.Error != nil {
		return r.Error
	}
	jsonDecoder := json.NewDecoder(r.getInternalReader())
	defer r.Close()
	return jsonDecoder.Decode(&userStruct)
}

func (r *Responses) populateResponseByteBuffer() {
	if r.internalByteBuffer.Len() != 0 {
		return
	}

	defer r.Close()

	// if there have any content
	if r.RawResponse.ContentLength == 0 {
		return
	}

	if r.RawResponse.ContentLength > 0 {
		r.internalByteBuffer.Grow(int(r.RawResponse.ContentLength))
	}

	if _, err := io.Copy(r.internalByteBuffer, r); err != nil && err != io.EOF {
		r.Error = err
		r.RawResponse.Body.Close()
	}
}

func (r *Responses) String() string {
	if r.Error != nil {
		return ""
	}
	r.populateResponseByteBuffer()
	return r.internalByteBuffer.String()
}

func (r *Responses) Bytes() []byte {
	if r.Error != nil {
		return nil
	}
	r.populateResponseByteBuffer()

	// Are we will emtpy
	if r.internalByteBuffer.Len() == 0 {
		return nil
	}
	return r.internalByteBuffer.Bytes()
}
