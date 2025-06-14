package typec

import (
	"time"

	zzzV1 "github.com/jianbo-zh/jypb/api/zzz/v1"
)

var (
	// 项目开始运营时间
	PROJT_STIME = time.Date(2024, 9, 4, 0, 0, 0, 0, time.Local)
	// 仪表盘2上线时间
	STATS_STIME = time.Date(2024, 4, 2, 0, 0, 0, 0, time.Local)
)

const (
	// 定义车辆是否运营的时间常量
	CarOperationSec = 1800 // 30分钟
)

const (
	Service_WebAdmin      string = "web_admin.GR"
	Service_WxmpUser      string = "wxmp_user.GR"
	Service_WxmpOps       string = "wxmp_ops.GR"
	Service_AppUser       string = "app_user.GR"
	Service_HapiWxNotify  string = "hapi_wxnotify.GR"
	Service_GwCarProxy    string = "gw_carproxy.GR"
	Service_BpTask        string = "bp_task.GR"
	Service_MsOrder       string = "ms_order.GR"
	Service_MsFileStorage string = "ms_filestorage.GR"
	Service_MsMessage     string = "ms_message.GR"
	Service_MsParkMap     string = "ms_parkmap.GR"
	Service_MsCarMeasure  string = "ms_carmeasure.GR"
	Service_MsCarAuth     string = "ms_carauth.GR"
	Service_MsCarConfig   string = "ms_carconfig.GR"
	Service_MsCarDispatch string = "ms_cardispatch.GR"
	Service_MsCarFlight   string = "ms_carflight.GR"
	Service_MsCarControl  string = "ms_carcontrol.GR"
)

type Role zzzV1.Operator_Role

const (
	Role_Admin          Role = Role(zzzV1.Operator_ROLE_ADMIN)
	Role_User           Role = Role(zzzV1.Operator_ROLE_USER)
	Role_Devops         Role = Role(zzzV1.Operator_ROLE_DEVOPS)
	Role_TaskSystem     Role = Role(zzzV1.Operator_ROLE_TASK_SYSTEM)
	Role_DispatchSystem Role = Role(zzzV1.Operator_ROLE_DISPATCH_SYSTEM)
)

type Origin zzzV1.Operator_Origin

const (
	Origin_TaskSystem     Origin = Origin(zzzV1.Operator_ORIGIN_TASK_SYSTEM)
	Origin_DispatchSystem Origin = Origin(zzzV1.Operator_ORIGIN_DISPATCH_SYSTEM)
	Origin_Admin          Origin = Origin(zzzV1.Operator_ORIGIN_ADMIN)
	Origin_UserApplet     Origin = Origin(zzzV1.Operator_ORIGIN_USER_APPLET)
	Origin_UserApp        Origin = Origin(zzzV1.Operator_ORIGIN_USER_APP)
	Origin_DevopsApplet   Origin = Origin(zzzV1.Operator_ORIGIN_DEVOPS_APPLET)
	Origin_CarUI          Origin = Origin(zzzV1.Operator_ORIGIN_CARUI)
	Origin_Vehicle        Origin = Origin(zzzV1.Operator_ORIGIN_VEHICLE)
)

func (o Origin) String() string {
	switch o {
	case Origin_Admin:
		return "管理后台"
	case Origin_Vehicle:
		return "车端"
	case Origin_DevopsApplet:
		return "运维小程序"
	case Origin_DispatchSystem:
		return "调度中心"
	case Origin_UserApp:
		return "用户APP"
	case Origin_UserApplet:
		return "用户小程序"
	case Origin_TaskSystem:
		return "任务系统"
	case Origin_CarUI:
		return "车端UI"
	}
	return "未知"
}
