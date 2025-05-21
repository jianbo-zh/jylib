package errc

import (
	zzzv1 "github.com/jianbo-zh/jypb/api/zzz/v1"
)

// 错误定义列表（内部抛出的）
var (
	// 公共错误
	CodeUnknowErr        = NewCode(int(zzzv1.ErrCode_UNKNOWN_ERROR))
	CodeParamErr         = NewCode(int(zzzv1.ErrCode_PARAM_ERROR))
	CodeMetadataErr      = NewCode(int(zzzv1.ErrCode_METADATA_ERROR))
	CodeNotFound         = NewCode(int(zzzv1.ErrCode_NOT_FOUND))
	CodeTypeAssertFailed = NewCode(int(zzzv1.ErrCode_TYPE_ASSERT_FAILED))

	// 用户相关错误
	CodeTokenMissing  = NewCode(int(zzzv1.ErrCode_S4_TOKEN_MISSING))
	CodeTokenInvalid  = NewCode(int(zzzv1.ErrCode_S4_TOKEN_INVALID))
	CodeTokenExpired  = NewCode(int(zzzv1.ErrCode_S4_TOKEN_EXPIRED))
	CodeUserNoAuth    = NewCode(int(zzzv1.ErrCode_S4_USER_NO_AUTH))
	CodeUserNotFound  = NewCode(int(zzzv1.ErrCode_S4_USER_NOT_FOUND))
	CodeUserPasswdErr = NewCode(int(zzzv1.ErrCode_S4_USER_PASSWD_WRONG))

	// 车辆相关错误
	CodeCarOffline           = NewCode(int(zzzv1.ErrCode_S4_CAR_OFFLINE))
	CodeCarNotSupportType    = NewCode(int(zzzv1.ErrCode_S4_CAR_NOT_SUPPORT_TYPE))
	CodeCarMeasureNotFound   = NewCode(int(zzzv1.ErrCode_S4_CAR_MEASURE_NOT_FOUND))
	CodeCarLowPower          = NewCode(int(zzzv1.ErrCode_S4_CAR_LOW_POWER))
	CodeCarActivateCodeWrong = NewCode(int(zzzv1.ErrCode_S4_CAR_ACTIVATE_CODE_WRONG))
	CodeCarActivateCodeUsed  = NewCode(int(zzzv1.ErrCode_S4_CAR_ACTIVATE_CODE_USED))
	CodeCarIsDispatching     = NewCode(int(zzzv1.ErrCode_S4_CAR_IS_DISPATCHING))
	CodeCarNotInOperation    = NewCode(int(zzzv1.ErrCode_S4_CAR_NOT_IN_OPERATION))
	CodeCarNotInMaintenance  = NewCode(int(zzzv1.ErrCode_S4_CAR_NOT_IN_MAINTENANCE))
	CodeCarPoiUnreachable    = NewCode(int(zzzv1.ErrCode_S4_CAR_POI_UNREACHABLE))
	CodeCarNotStopped        = NewCode(int(zzzv1.ErrCode_S4_CAR_NOT_STOPPED))
	CodeCarUnlocked          = NewCode(int(zzzv1.ErrCode_S4_CAR_UNLOCKED))
	CodeCarEstopped          = NewCode(int(zzzv1.ErrCode_S4_CAR_ESTOPPED))
	CodeCarCoordUnknown      = NewCode(int(zzzv1.ErrCode_S4_CAR_COORD_UNKNOWN))
	CodeCarNotSupportAuto    = NewCode(int(zzzv1.ErrCode_S4_CAR_NOT_SUPPORT_AUTO))
	CodeCarNotInRoad         = NewCode(int(zzzv1.ErrCode_S4_CAR_NOT_IN_ROAD))

	// 订单相关错误
	CodeOrderNotFound   = NewCode(int(zzzv1.ErrCode_S4_ORDER_NOT_FOUND))
	CodeOrderStateError = NewCode(int(zzzv1.ErrCode_S4_ORDER_STATE_ERROR))
)

var (
	// proto定义
	ErrCode_name  = zzzv1.ErrCode_name
	ErrCode_value = zzzv1.ErrCode_value
)
