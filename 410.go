// MIT License
//
// Copyright (c) 2024 chaunsin
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
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package api

import (
	"context"
	"fmt"
)

type ECIV4GetBasicDetailsByNameReq struct {
	Keyword string
}

type ECIV4GetBasicDetailsByNameResp struct {
	Response[ECIV4GetBasicDetailsByNameRespResult]
}

type ECIV4GetBasicDetailsByNameRespResult struct {
	KeyNo                 string `json:"KeyNo"`                 // 主键
	Name                  string `json:"Name"`                  // 企业名称
	No                    string `json:"No"`                    // 根据企业性质返回不同值（对中国境内企业（EntType=0/1/4/7/9/10/11）返回工商注册号，对中国香港企业返回企业编号，对中国台湾企业返回企业编号）
	BelongOrg             string `json:"BelongOrg"`             // 登记机关
	OperId                string `json:"OperId"`                // 法定代表人ID
	OperName              string `json:"OperName"`              // 法定代表人名称
	StartDate             string `json:"StartDate"`             // 成立日期
	EndDate               string `json:"EndDate"`               // 吊销日期（保留字段）
	Status                string `json:"Status"`                // 登记状态
	Province              string `json:"Province"`              // 省份
	UpdatedDate           string `json:"UpdatedDate"`           // 更新日期
	CreditCode            string `json:"CreditCode"`            // 根据企业性质返回不同值（对中国境内企业（EntType=0/1/4/7/9/10/11）返回统一社会信用代码，对中国香港企业返回商业登记号码）
	RegistCapi            string `json:"RegistCapi"`            // 注册资本
	RegisteredCapital     string `json:"RegisteredCapital"`     // 注册资本数额
	RegisteredCapitalUnit string `json:"RegisteredCapitalUnit"` // 注册资本单位
	RegisteredCapitalCCY  string `json:"RegisteredCapitalCCY"`  // 注册资本币种
	EconKind              string `json:"EconKind"`              // 企业类型
	Address               string `json:"Address"`               // 注册地址
	Scope                 string `json:"Scope"`                 // 经营范围
	TermStart             string `json:"TermStart"`             // 营业期限始
	TermEnd               string `json:"TermEnd"`               // 营业期限至
	CheckDate             string `json:"CheckDate"`             // 核准日期
	OrgNo                 string `json:"OrgNo"`                 // 组织机构代码
	IsOnStock             string `json:"IsOnStock"`             // 是否上市（0-未上市，1-上市）
	StockNumber           string `json:"StockNumber"`           // 股票代码（若A股和港股同时存在，优先返回A股代码）
	StockType             string `json:"StockType"`             // 上市类型（A股、港股、美股、新三板、新四板）
	OriginalName          []struct {
		Name       string `json:"Name"`       // 	曾用名
		ChangeDate string `json:"ChangeDate"` // 	变更日期
	} `json:"OriginalName"` // 曾用名
	ImageUrl          string `json:"ImageUrl"`          // 企业Logo地址
	EntType           string `json:"EntType"`           // 企业性质（0-大陆企业，1-社会组织，4-事业单位，7-医院，9-律师事务所，10-学校，11-机关单位，-1-其他）
	RecCap            string `json:"RecCap"`            // 实缴资本
	PaidUpCapital     string `json:"PaidUpCapital"`     // 实缴出资额数额
	PaidUpCapitalUnit string `json:"PaidUpCapitalUnit"` // 实缴出资额单位
	PaidUpCapitalCCY  string `json:"PaidUpCapitalCCY"`  // 实缴出资额币种
	RevokeInfo        struct {
		CancelDate   string `json:"CancelDate"`   // 注销日期（如“2022-01-01”）
		CancelReason string `json:"CancelReason"` // 注销原因
		RevokeDate   string `json:"RevokeDate"`   // 吊销日期
		RevokeReason string `json:"RevokeReason"` // 吊销原因
	} `json:"RevokeInfo"` // 注销吊销信息
	Area struct {
		Province string `json:"Province"` // 省份
		City     string `json:"City"`     // 城市
		County   string `json:"County"`   // 区域
	} `json:"Area"` // 行政区域
	AreaCode string `json:"AreaCode"` // 行政区划代码
}

// ECIV4GetBasicDetailsByName 企业工商信息（含工商照面信息） https://openapi.qcc.com/dataApi/410
func (a *Api) ECIV4GetBasicDetailsByName(ctx context.Context, req *ECIV4GetBasicDetailsByNameReq) (*ECIV4GetBasicDetailsByNameResp, error) {
	var resp ECIV4GetBasicDetailsByNameResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("keyword", req.Keyword).
		SetResult(&resp).
		Get("/ECIV4/GetBasicDetailsByName")
	if err != nil {
		return nil, err
	}
	if reply.StatusCode() != 200 {
		return nil, fmt.Errorf("request status code [%v] body: %s", reply.StatusCode(), string(reply.Body()))
	}
	if resp.Status != "200" {
		return nil, fmt.Errorf("err: %+v", resp)
	}
	return &resp, nil
}
