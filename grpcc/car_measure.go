package grpcc

import (
	"context"
	"fmt"

	"github.com/jianbo-zh/jylib/errc"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	carmeasureV1 "github.com/jianbo-zh/jypb/api/carmeasure/v1"
)

type ICarMeasure interface {
	CarSpeed(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarSpeedReply, error)
	CarOnline(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarOnlineReply, error)
	CarLonLat(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarLonLatReply, error)
	CarHeading(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarHeadingReply, error)
	CarSupportType(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarSupportTypeReply, error)
	CarDrivingState(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarDrivingStateReply, error)
	CarRemainingPower(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarRemainingPowerReply, error)
	CarInRoad(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarInRoadReply, error)
	CarLoad(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarLoadReply, error)
	CarEstop(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarEstopReply, error)
	CarSupportAuto(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarSupportAutoReply, error)
	CarMeasure(context.Context, *carmeasureV1.CarMeasureRequest, ...filterc.Filter) (*carmeasureV1.CarMeasureReply, error)
	CarMeasures(context.Context, *carmeasureV1.CarMeasuresRequest, ...filterc.Filter) (*carmeasureV1.CarMeasuresReply, error)
	CarPowerEstimate(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.CarPowerEstimateReply, error)
	MetricModuleStatus(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricModuleStatusReply, error)
	MetricBatData(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricBatDataReply, error)
	MetricChassisData(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricChassisDataReply, error)
	MetricGpsData(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricGpsDataReply, error)
	MetricGpsStatus(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricGpsStatusReply, error)
	MetricGuardianData(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricGuardianDataReply, error)
	MetricFenceData(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricFenceDataReply, error)
	MetricFenceId(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricFenceIdReply, error)
	MetricUltrasonicData(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricUltrasonicDataReply, error)
	MetricLocalizationData(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricLocalizationDataReply, error)
	MetricPerceptionObstacles(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricPerceptionObstaclesReply, error)
	MetricSysInfo(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.MetricSysInfoReply, error)
	EventAlarmReport(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventAlarmReportReply, error)
	EventAlarmSupport(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventAlarmSupportReply, error)
	EventChassisState(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventChassisStateReply, error)
	EventChassisError(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventChassisErrorReply, error)
	EventEmergencyBrake(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventEmergencyBrakeReply, error)
	EventCurrentVersion(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventCurrentVersionReply, error)
	EventResourceProgress(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventResourceProgressReply, error)
	EventSolType(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventSolTypeReply, error)
	EventGuardianStatus(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventGuardianStatusReply, error)
	EventVoiceAnnouncements(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventVoiceAnnouncementsReply, error)
	EventPInCarDetect(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventPInCarDetectReply, error)
	EventLocatedType(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventLocatedTypeReply, error)
	EventAutoPhase(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventAutoPhaseReply, error)
	EventVehicleState(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventVehicleStateReply, error)
	EventPowerOn(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventPowerOnReply, error)
	EventNearbyPOI(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventNearbyPOIReply, error)
	EventUOCState(context.Context, *carmeasureV1.Device, ...filterc.Filter) (*carmeasureV1.EventUOCStateReply, error)
}

type CarMeasureGrpc struct {
	client IClient
}

func NewCarMeasureGrpc(cli IClient) ICarMeasure {
	return &CarMeasureGrpc{
		client: cli,
	}
}

func (c *CarMeasureGrpc) CarSpeed(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarSpeedReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarSpeed(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarSpeed error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarOnline(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarOnlineReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarOnline(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarOnline error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarLonLat(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarLonLatReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarLonLat(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarLonLat error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarHeading(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarHeadingReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarHeading(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarHeading error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarSupportType(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarSupportTypeReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarSupportType(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarSupportType error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarDrivingState(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarDrivingStateReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarDrivingState(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarDrivingState error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarRemainingPower(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarRemainingPowerReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarRemainingPower(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarRemainingPower error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarInRoad(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarInRoadReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarInRoad(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarInRoad error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarLoad(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarLoadReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarLoad(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarLoad error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarEstop(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarEstopReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarEstop(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarEstop error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarSupportAuto(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarSupportAutoReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarSupportAuto(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarSupportAuto error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarMeasure(ctx context.Context, req *carmeasureV1.CarMeasureRequest, filters ...filterc.Filter) (*carmeasureV1.CarMeasureReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarMeasure(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarMeasure error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarMeasures(ctx context.Context, req *carmeasureV1.CarMeasuresRequest, filters ...filterc.Filter) (*carmeasureV1.CarMeasuresReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarMeasures(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarMeasures error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) CarPowerEstimate(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.CarPowerEstimateReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.CarPowerEstimate(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CarPowerEstimate error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricModuleStatus(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricModuleStatusReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricModuleStatus(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricModuleStatus error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricBatData(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricBatDataReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricBatData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricBatData error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricChassisData(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricChassisDataReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricChassisData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricChassisData error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricGpsData(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricGpsDataReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricGpsData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricGpsData error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricGpsStatus(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricGpsStatusReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricGpsStatus(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricGpsStatus error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricGuardianData(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricGuardianDataReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricGuardianData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricGuardianData error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricFenceData(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricFenceDataReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricFenceData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricFenceData error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricFenceId(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricFenceIdReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricFenceId(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricFenceId error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricUltrasonicData(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricUltrasonicDataReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricUltrasonicData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricUltrasonicData error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricLocalizationData(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricLocalizationDataReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricLocalizationData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricLocalizationData error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricPerceptionObstacles(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricPerceptionObstaclesReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricPerceptionObstacles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricPerceptionObstacles error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) MetricSysInfo(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.MetricSysInfoReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.MetricSysInfo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.MetricSysInfo error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventAlarmReport(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventAlarmReportReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventAlarmReport(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventAlarmReport error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventAlarmSupport(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventAlarmSupportReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventAlarmSupport(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventAlarmSupport error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventChassisState(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventChassisStateReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventChassisState(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventChassisState error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventChassisError(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventChassisErrorReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventChassisError(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventChassisError error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventEmergencyBrake(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventEmergencyBrakeReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventEmergencyBrake(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventEmergencyBrake error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventCurrentVersion(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventCurrentVersionReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventCurrentVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventCurrentVersion error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventResourceProgress(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventResourceProgressReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventResourceProgress(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventResourceProgress error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventSolType(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventSolTypeReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventSolType(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventSolType error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventGuardianStatus(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventGuardianStatusReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventGuardianStatus(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventGuardianStatus error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventVoiceAnnouncements(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventVoiceAnnouncementsReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventVoiceAnnouncements(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventVoiceAnnouncements error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventPInCarDetect(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventPInCarDetectReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventPInCarDetect(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventPInCarDetect error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventLocatedType(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventLocatedTypeReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventLocatedType(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventLocatedType error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventAutoPhase(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventAutoPhaseReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventAutoPhase(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventAutoPhase error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventVehicleState(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventVehicleStateReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventVehicleState(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventVehicleState error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventPowerOn(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventPowerOnReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventPowerOn(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventPowerOn error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventNearbyPOI(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventNearbyPOIReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventNearbyPOI(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventNearbyPOI error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarMeasureGrpc) EventUOCState(ctx context.Context, req *carmeasureV1.Device, filters ...filterc.Filter) (*carmeasureV1.EventUOCStateReply, error) {
	cli, err := c.client.CarMeasureClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.EventUOCState(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EventUOCState error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}
