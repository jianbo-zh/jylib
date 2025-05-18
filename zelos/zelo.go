// 九识Api库
package zelos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type IZelos interface {
	// 公司列表
	OrgList(ctx context.Context, req *OrgListRequest) (int, []*Org, error)
	// 站点列表（对应景区）
	StationList(ctx context.Context, req *StationListRequest) (int, []*Station, error)
	// 车辆列表（景区车辆）
	VehicleList(ctx context.Context, req *VehicleListRequest) (int, []*Vehicle, error)
	// 停车点列表（景区POI点）
	StopList(ctx context.Context, req *StopListRequest) (int, []*Stop, error)
	// 车辆详情
	VehicleDetail(ctx context.Context, vehicleName string) (*VehicleDetail, error)
	// 添加任务并出发（发起调度）
	AddDispatch(ctx context.Context, req *AddDispatchRequest) (*Dispatch, error)
	// 车辆出发（重启调度）
	GoDispatch(ctx context.Context, vehicleName string) error
	// 取消任务（取消调度）
	CancelDispatch(ctx context.Context, vehicleName string) error
	// 添加停车点
	AddStops(ctx context.Context, req *AddStopsRequest) error
	// 更新停车点
	UpdateStops(ctx context.Context, req *UpdateStopsRequest) error
	// 删除停车点
	DeleteStops(ctx context.Context, req *DeleteStopsRequest) error
}

type Zelos struct {
	mutex sync.Mutex

	DomainAddr string
	AppId      string
	AppKey     string

	Auth *Auth

	httpCli *http.Client
}

func NewZelos(domainAddr, appId, appKey string) IZelos {
	return &Zelos{
		DomainAddr: domainAddr,
		AppId:      appId,
		AppKey:     appKey,

		httpCli: &http.Client{Timeout: 5 * time.Second},
	}
}

// OrgList 公司列表
func (z *Zelos) OrgList(ctx context.Context, req *OrgListRequest) (int, []*Org, error) {
	query := url.Values{}
	query.Add("pageNumber", strconv.Itoa(req.PageNo))
	query.Add("pageSize", strconv.Itoa(req.PageSize))

	respBs, err := z.get(ctx, "/business-server/open-apis/organizations", query)
	if err != nil {
		return 0, nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data OrgListResponse
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return 0, nil, fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return 0, nil, NewError(data.ErrorCode, data.Message)
	}

	return data.Data.Total, data.Data.List, nil
}

// StationList 车辆列表
func (z *Zelos) StationList(ctx context.Context, req *StationListRequest) (int, []*Station, error) {
	query := url.Values{}
	query.Add("pageNumber", strconv.Itoa(req.PageNo))
	query.Add("pageSize", strconv.Itoa(req.PageSize))
	query.Add("organizationId", strconv.Itoa(req.OrgId))

	respBs, err := z.get(ctx, "/business-server/open-apis/stations", query)
	if err != nil {
		return 0, nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data StationListResponse
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return 0, nil, fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return 0, nil, NewError(data.ErrorCode, data.Message)
	}

	return data.Data.Total, data.Data.List, nil
}

// VehicleList 车辆列表
func (z *Zelos) VehicleList(ctx context.Context, req *VehicleListRequest) (int, []*Vehicle, error) {
	query := url.Values{}
	query.Add("pageNumber", strconv.Itoa(req.PageNo))
	query.Add("pageSize", strconv.Itoa(req.PageSize))
	query.Add("stationId", strconv.Itoa(req.StationId))

	respBs, err := z.get(ctx, "/business-server/open-apis/vehicles", query)
	if err != nil {
		return 0, nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data VehicleListResponse
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return 0, nil, fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return 0, nil, NewError(data.ErrorCode, data.Message)
	}

	return data.Data.Total, data.Data.List, nil
}

// StopList 停车点列表
func (z *Zelos) StopList(ctx context.Context, req *StopListRequest) (int, []*Stop, error) {
	query := url.Values{}
	query.Add("pageNumber", strconv.Itoa(req.PageNo))
	query.Add("pageSize", strconv.Itoa(req.PageSize))
	query.Add("stationId", strconv.Itoa(req.StationId))

	respBs, err := z.get(ctx, "/business-server/open-apis/stops", query)
	if err != nil {
		return 0, nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data StopListResponse
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return 0, nil, fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return 0, nil, NewError(data.ErrorCode, data.Message)
	}

	return data.Data.Total, data.Data.List, nil
}

// VehicleData 车辆详情
func (z *Zelos) VehicleDetail(ctx context.Context, vehicleName string) (*VehicleDetail, error) {
	query := url.Values{}
	query.Add("vehicleName", vehicleName)

	respBs, err := z.get(ctx, "/business-server/open-apis/vehicle", query)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data VehicleDetailResp
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return nil, NewError(data.ErrorCode, data.Message)
	}

	return data.Data, nil
}

// AddDispatch 开始任务并出发
func (z *Zelos) AddDispatch(ctx context.Context, req *AddDispatchRequest) (*Dispatch, error) {
	bs, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal error: %w", err)
	}

	respBs, err := z.post(ctx, "/business-server/open-apis/vehicle/add_dispatch", nil, bs)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data StartDispatchResponse
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return nil, NewError(data.ErrorCode, data.Message)
	}

	return data.Data, nil
}

// GoDispatch 车辆出发
func (z *Zelos) GoDispatch(ctx context.Context, vehicleName string) error {
	bs, err := json.Marshal(&struct {
		VehicleName string `json:"vehicleName,omitempty"`
	}{
		VehicleName: vehicleName,
	})
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}

	respBs, err := z.post(ctx, "/business-server/open-apis/vehicle/go", nil, bs)
	if err != nil {
		return fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data Status
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return NewError(data.ErrorCode, data.Message)
	}

	return nil
}

// CancelDispatch 取消任务
func (z *Zelos) CancelDispatch(ctx context.Context, vehicleName string) error {
	bs, err := json.Marshal(&struct {
		VehicleName string `json:"vehicleName,omitempty"`
	}{
		VehicleName: vehicleName,
	})
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}

	respBs, err := z.post(ctx, "/business-server/open-apis/vehicle/cancel_dispatch", nil, bs)
	if err != nil {
		return fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data Status
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return NewError(data.ErrorCode, data.Message)
	}

	return nil
}

// CreateStops 新增停靠点
func (z *Zelos) AddStops(ctx context.Context, req *AddStopsRequest) error {
	bs, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}

	respBs, err := z.post(ctx, "/business-server/open-apis/stop/add", nil, bs)
	if err != nil {
		return fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data Status
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return NewError(data.ErrorCode, data.Message)
	}

	return nil
}

// UpdateStops 修改停靠点
func (z *Zelos) UpdateStops(ctx context.Context, req *UpdateStopsRequest) error {
	bs, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}

	respBs, err := z.put(ctx, "/business-server/open-apis/stop/update", nil, bs)
	if err != nil {
		return fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data Status
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return NewError(data.ErrorCode, data.Message)
	}

	return nil
}

// DeleteStops 删除停靠点
func (z *Zelos) DeleteStops(ctx context.Context, req *DeleteStopsRequest) error {
	bs, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}

	respBs, err := z.delete(ctx, "/business-server/open-apis/stop/delete", nil, bs)
	if err != nil {
		return fmt.Errorf("io.ReadAll error: %w", err)
	}

	var data Status
	err = json.Unmarshal(respBs, &data)
	if err != nil {
		return fmt.Errorf("json.Unmarshal error: %w", err)

	} else if !data.Success {
		return NewError(data.ErrorCode, data.Message)
	}

	return nil
}

func (z *Zelos) get(ctx context.Context, path string, query url.Values) ([]byte, error) {
	return z.nobody(ctx, HttpMethod_GET, path, query)
}

func (z *Zelos) post(ctx context.Context, path string, query url.Values, body []byte) ([]byte, error) {
	return z.body(ctx, HttpMethod_PUT, path, query, body)
}

func (z *Zelos) put(ctx context.Context, path string, query url.Values, body []byte) ([]byte, error) {
	return z.body(ctx, HttpMethod_PUT, path, query, body)
}

func (z *Zelos) delete(ctx context.Context, path string, query url.Values, body []byte) ([]byte, error) {
	return z.body(ctx, HttpMethod_DELETE, path, query, body)
}

func (z *Zelos) body(ctx context.Context, method HttpMethod, path string, query url.Values, data []byte) ([]byte, error) {
	u := z.url(path)
	if query != nil {
		u.RawQuery = query.Encode()
	}
	request, err := http.NewRequest(string(method), u.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %w", err)
	}
	token, err := z.getAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("getAccessToken error: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("token", token)

	resp, err := z.httpCli.Do(request)
	if err != nil {
		return nil, fmt.Errorf("client.Do error: %w", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (z *Zelos) nobody(ctx context.Context, method HttpMethod, path string, query url.Values) ([]byte, error) {
	u := z.url(path)
	if query != nil {
		u.RawQuery = query.Encode()
	}
	request, err := http.NewRequest(string(method), u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %w", err)
	}
	token, err := z.getAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("getAccessToken error: %w", err)
	}
	request.Header.Set("token", token)

	resp, err := z.httpCli.Do(request)
	if err != nil {
		return nil, fmt.Errorf("client.Do error: %w", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (z *Zelos) url(path string) *url.URL {
	u, _ := url.Parse(fmt.Sprintf("%s%s", strings.TrimRight(z.DomainAddr, "/"), path))
	return u
}

func (z *Zelos) getAccessToken(ctx context.Context) (string, error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.Auth == nil || z.Auth.ExpiredAt.Before(time.Now()) {
		bs, _ := json.Marshal(&AppIdKey{
			AppId:  z.AppId,
			AppKey: z.AppKey,
		})

		body, err := z.post(ctx, "/app/accessToken", nil, bs)
		if err != nil {
			return "", fmt.Errorf("http.post error: %w", err)
		}

		var result AccessTokenResponse
		err = json.Unmarshal(body, &result)
		if err != nil {
			return "", fmt.Errorf("json.Unmarshal error: %w", err)

		} else if !result.Success {
			return "", NewError(result.ErrorCode, result.Message)
		}

		z.Auth = &Auth{
			Id:        result.Data.Id,
			Token:     result.Data.Token,
			ExpiredAt: time.Now().Add(time.Duration(result.Data.ExpiresAfter-1) * time.Minute),
		}
	}

	return z.Auth.Token, nil
}
