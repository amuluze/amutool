// Package requests
// Date:   2024/12/13 18:27
// Author: Amu
// Description:
package requests

import (
	"testing"
)

type Params struct {
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	EventId   string `json:"event_id"`
}

type AccessLogDetail struct {
	ProtectedObjectName   string `json:"protected_object_name"`
	URLAddress            string `json:"url_address"`
	SrcIp                 string `json:"src_ip"`
	SrcPort               uint32 `json:"src_port"`
	HttpHostPort          uint32 `json:"http_host_port"`
	Method                string `json:"method"`
	Scheme                string `json:"scheme"`
	EventId               string `json:"event_id"`
	ReqDetectorName       string `json:"req_detector_name"`
	RspDetectorName       string `json:"rsp_detector_name"`
	ReqDetectTime         int32  `json:"req_detect_time"`
	RspDetectTime         int32  `json:"rsp_detect_time"`
	ReqHttpBodyIsTruncate int32  `json:"req_http_body_is_truncate"`
	RspHttpBodyIsTruncate int32  `json:"rsp_http_body_is_truncate"`
	UpstreamStatusCode    uint32 `json:"upstream_status_code"`
	StatusCode            uint32 `json:"status_code"`
	InnerVlanId           int32  `json:"inner_vlan_id"`
	OuterVlanId           int32  `json:"outer_vlan_id"`
	Country               string `json:"country"`
	Province              string `json:"province"`
	XForwardedFor         string `json:"x_forwarded_for"`
	HttpBodyIsAbandoned   int32  `json:"http_body_is_abandoned"`
	ReqAttackType         string `json:"req_attack_type"`
	ReqAction             string `json:"req_action"`
	ReqRiskLevel          string `json:"req_risk_level"`
	ReqBlockReason        int32  `json:"req_block_reason"`
	ReqRuleModule         string `json:"req_rule_module"`
	ReqBody               string `json:"req_body"`
}

type AccessLogDetailResponse struct {
	Data []AccessLogDetail `json:"data"`
}

type Response struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    AccessLogDetailResponse `json:"data"`
}

type PostResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

func TestGet(t *testing.T) {
	params := &Params{StartTime: 1733441315000, EndTime: 1733456677927, EventId: "04e774d8406d4383b5b934820e83f207"}
	reply := &Response{}
	err := Get("https://10.9.35.6:9443/api/v3/protected-logger/AccessLogDetail", params, reply)
	if err != nil {
		t.Fatalf("get error: %v", err)
	}
	t.Logf("get success: %#v", reply)
}

type MessageReponse struct {
	Message string `json:"message"`
}

func TestGetTwo(t *testing.T) {
	reply := &MessageReponse{}
	err := Get("http://10.9.35.5:8090/aaa", nil, reply)
	if err != nil {
		t.Fatalf("get error: %v", err)
	}
	t.Logf("get success: %#v", reply)
}

func TestPost(t *testing.T) {
	params := &User{
		Name: "Amu",
		Age:  18,
	}
	reply := &PostResponse{}
	err := Post("http://10.9.35.5:8090/users", params, reply, SetHeader("Referer", "https://example.com/test"), SetCookie("cid", "123456"))
	if err != nil {
		t.Fatalf("post error: %v", err)
	}
	t.Logf("post success: %#v", reply)
}
