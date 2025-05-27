package zelos

import (
	"fmt"
	"time"
)

type HttpMethod string

const (
	HttpMethod_GET    HttpMethod = "GET"
	HttpMethod_POST   HttpMethod = "POST"
	HttpMethod_PUT    HttpMethod = "PUT"
	HttpMethod_DELETE HttpMethod = "DELETE"
)

type Error struct {
	Code    string
	Message string
}

func (err *Error) Error() string {
	return fmt.Sprintf("zelo %s %s", err.Code, err.Message)
}

func NewError(code string, msg string) error {
	return &Error{Code: code, Message: msg}
}

type Status struct {
	Success   bool   `json:"success,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	Message   string `json:"message,omitempty"`
}

type AppIdKey struct {
	AppId  string `json:"appId,omitempty"`
	AppKey string `json:"appKey,omitempty"`
}

type AccessTokenResponse struct {
	Status
	Data struct {
		Id           int    `json:"id,omitempty"`
		Token        string `json:"token,omitempty"`
		ExpiresAfter int    `json:"expiresAfter,omitempty"` // 多少分钟后过期
	} `json:"data,omitempty"`
}

type Auth struct {
	Id        int
	Token     string
	ExpiredAt time.Time
}

type Org struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Station struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	LonWgs84 float64 `json:"lon,omitempty"`
	LatWgs84 float64 `json:"lat,omitempty"`
	LonGcj02 float64 `json:"gcj02Lon,omitempty"`
	LatGcj02 float64 `json:"gcj02Lat,omitempty"`
}

type Stop struct {
	Id          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	StationId   int     `json:"stationId,omitempty"`
	StationName string  `json:"stationName,omitempty"`
	LonWgs84    float64 `json:"lon,omitempty"`
	LatWgs84    float64 `json:"lat,omitempty"`
	LonGcj02    float64 `json:"gcj02Lon,omitempty"`
	LatGcj02    float64 `json:"gcj02Lat,omitempty"`
}

type Goal struct {
	StopId            int    `json:"stopId,omitempty"`
	StopName          string `json:"stopName,omitempty"`
	GoalIndex         int    `json:"goalIndex,omitempty"`
	ArrivalDatetime   string `json:"arrivalDatetime,omitempty"`
	DepartureDatetime string `json:"departureDatetime,omitempty"`
}

type Vehicle struct {
	Id                 int     `json:"id,omitempty"`
	Name               string  `json:"name,omitempty"`
	Number             string  `json:"number,omitempty"`
	Vin                string  `json:"vin,omitempty"`
	StationId          int     `json:"stationId,omitempty"`
	StationName        string  `json:"stationName,omitempty"`
	BusinessStatus     string  `json:"businessStatus,omitempty"`
	BusinessStatusName string  `json:"businessStatusName,omitempty"`
	Model              string  `json:"model,omitempty"`
	BoxVolume          float64 `json:"boxVolume,omitempty"`
	BoxSize            string  `json:"boxSize,omitempty"`
	Weight             float64 `json:"weight,omitempty"`
	VehicleSize        string  `json:"vehicleSize,omitempty"`
	NoLoadEndurance    float64 `json:"noLoadEndurance,omitempty"`
	FullLoadEndurance  float64 `json:"fullLoadEndurance,omitempty"`
	MaxSpeed           float64 `json:"maxSpeed,omitempty"`
}

type Dispatch struct {
	VehicleId                 int    `json:"vehicleId,omitempty"`
	VehicleName               string `json:"vehicleName,omitempty"`
	VehicleNumber             string `json:"vehicleNumber,omitempty"`
	VehicleBusinessStatus     string `json:"vehicleBusinessStatus,omitempty"`
	VehicleBusinessStatusName string `json:"vehicleBusinessStatusName,omitempty"`
	DispatchId                int    `json:"dispatchId,omitempty"`
}

type OrgListRequest struct {
	PageNo   int
	PageSize int
}
type OrgListResponse struct {
	Status
	Data struct {
		Total           int    `json:"total,omitempty"`
		List            []*Org `json:"list,omitempty"`
		PageNum         int    `json:"pageNum,omitempty"`
		PageSize        int    `json:"pageSize,omitempty"`
		Size            int    `json:"size,omitempty"`
		StartRow        int    `json:"startRow,omitempty"`
		EndRow          int    `json:"endRow,omitempty"`
		Pages           int    `json:"pages,omitempty"`
		PrePage         int    `json:"prePage,omitempty"`
		NextPage        int    `json:"nextPage,omitempty"`
		IsFirstPage     bool   `json:"isFirstPage,omitempty"`
		IsLastPage      bool   `json:"isLastPage,omitempty"`
		HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
		HasNextPage     bool   `json:"hasNextPage,omitempty"`
	} `json:"data,omitempty"`
}

type StationListRequest struct {
	OrgId    int
	PageNo   int
	PageSize int
}
type StationListResponse struct {
	Status
	Data struct {
		Total           int        `json:"total,omitempty"`
		List            []*Station `json:"list,omitempty"`
		PageNum         int        `json:"pageNum,omitempty"`
		PageSize        int        `json:"pageSize,omitempty"`
		Size            int        `json:"size,omitempty"`
		StartRow        int        `json:"startRow,omitempty"`
		EndRow          int        `json:"endRow,omitempty"`
		Pages           int        `json:"pages,omitempty"`
		PrePage         int        `json:"prePage,omitempty"`
		NextPage        int        `json:"nextPage,omitempty"`
		IsFirstPage     bool       `json:"isFirstPage,omitempty"`
		IsLastPage      bool       `json:"isLastPage,omitempty"`
		HasPreviousPage bool       `json:"hasPreviousPage,omitempty"`
		HasNextPage     bool       `json:"hasNextPage,omitempty"`
	} `json:"data,omitempty"`
}

type VehicleListRequest struct {
	StationId int
	PageNo    int
	PageSize  int
}
type VehicleListResponse struct {
	Status
	Data struct {
		Total           int        `json:"total,omitempty"`
		List            []*Vehicle `json:"list,omitempty"`
		PageNum         int        `json:"pageNum,omitempty"`
		PageSize        int        `json:"pageSize,omitempty"`
		Size            int        `json:"size,omitempty"`
		StartRow        int        `json:"startRow,omitempty"`
		EndRow          int        `json:"endRow,omitempty"`
		Pages           int        `json:"pages,omitempty"`
		PrePage         int        `json:"prePage,omitempty"`
		NextPage        int        `json:"nextPage,omitempty"`
		IsFirstPage     bool       `json:"isFirstPage,omitempty"`
		IsLastPage      bool       `json:"isLastPage,omitempty"`
		HasPreviousPage bool       `json:"hasPreviousPage,omitempty"`
		HasNextPage     bool       `json:"hasNextPage,omitempty"`
	} `json:"data,omitempty"`
}

type StopListRequest struct {
	StationId int
	PageNo    int
	PageSize  int
}
type StopListResponse struct {
	Status
	Data struct {
		Total           int     `json:"total,omitempty"`
		List            []*Stop `json:"list,omitempty"`
		PageNum         int     `json:"pageNum,omitempty"`
		PageSize        int     `json:"pageSize,omitempty"`
		Size            int     `json:"size,omitempty"`
		StartRow        int     `json:"startRow,omitempty"`
		EndRow          int     `json:"endRow,omitempty"`
		Pages           int     `json:"pages,omitempty"`
		PrePage         int     `json:"prePage,omitempty"`
		NextPage        int     `json:"nextPage,omitempty"`
		IsFirstPage     bool    `json:"isFirstPage,omitempty"`
		IsLastPage      bool    `json:"isLastPage,omitempty"`
		HasPreviousPage bool    `json:"hasPreviousPage,omitempty"`
		HasNextPage     bool    `json:"hasNextPage,omitempty"`
	} `json:"data,omitempty"`
}

type VehicleDetail struct {
	Vehicle
	Stops            []*Stop `json:"stops,omitempty"`
	LonWgs84         float64 `json:"lon,omitempty"`
	LatWgs84         float64 `json:"lat,omitempty"`
	LonGcj02         float64 `json:"gcj02Lon,omitempty"`
	LatGcj02         float64 `json:"gcj02Lat,omitempty"`
	Online           bool    `json:"online,omitempty"`
	BatteryCharging  bool    `json:"batteryCharging,omitempty"`
	BatteryPower     float64 `json:"batteryPower,omitempty"`
	DispatchId       int64   `json:"dispatchId,omitempty"`
	CurrentGoalIndex int     `json:"currentGoalIndex,omitempty"`
	Goals            []*Goal `json:"goals,omitempty"`
}

type VehicleDetailResp struct {
	Status
	Data *VehicleDetail `json:"data,omitempty"`
}

type AddDispatchRequest struct {
	VehicleName string  `json:"vehicleName,omitempty"`
	FromStopId  int     `json:"fromStopId,omitempty"`
	ToStopIds   []int   `json:"toStopIds,omitempty"`
	SpeedLimit  float32 `json:"speedLimit,omitempty"`
}

type StartDispatchResponse struct {
	Status
	Data *Dispatch `json:"data,omitempty"`
}

type AddStopsRequest struct {
	Name      string  `json:"name,omitempty"`
	LonWgs84  float64 `json:"lon,omitempty"`
	LatWgs84  float64 `json:"lat,omitempty"`
	Heading   float64 `json:"heading,omitempty"`
	StationId int     `json:"stationId,omitempty"`
	Comment   string  `json:"comment,omitempty"`
	UserId    int     `json:"userId,omitempty"`
	UserName  string  `json:"userName,omitempty"`
}

type UpdateStopsRequest struct {
	Id        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	LonWgs84  float64 `json:"lon,omitempty"`
	LatWgs84  float64 `json:"lat,omitempty"`
	Heading   float64 `json:"heading,omitempty"`
	StationId int     `json:"stationId,omitempty"`
	Comment   string  `json:"comment,omitempty"`
	UserId    int     `json:"userId,omitempty"`
	UserName  string  `json:"userName,omitempty"`
}

type DeleteStopsRequest struct {
	Id       int    `json:"id,omitempty"`
	UserId   int    `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
}
