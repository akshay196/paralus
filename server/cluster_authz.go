package server

import (
	"context"

	sentryrpc "github.com/RafayLabs/rcloud-base/proto/rpc/sentry"

	"github.com/RafayLabs/rcloud-base/pkg/sentry/authz"
	"github.com/RafayLabs/rcloud-base/pkg/service"
)

type clusterAuthzServer struct {
	bs  service.BootstrapService
	aps service.AccountPermissionService
	gps service.GroupPermissionService
	krs service.KubeconfigRevocationService
	kcs service.KubectlClusterSettingsService
	kss service.KubeconfigSettingService
	//apn   models.AccountProjectNamespaceService
}

// GetUserAuthorization return authorization profile of user for a given cluster
func (s *clusterAuthzServer) GetUserAuthorization(ctx context.Context, req *sentryrpc.GetUserAuthorizationRequest) (*sentryrpc.GetUserAuthorizationResponse, error) {
	resp, err := authz.GetAuthorization(ctx, req, s.bs, s.aps, s.gps, s.krs, s.kcs, s.kss)
	if err != nil {
		_log.Errorw("error getting auth profile", "req", req, "error", err.Error())
		return nil, err
	}
	return resp, nil
}

// NewClusterAuthzServer returns New ClusterAuthzServer
func NewClusterAuthzServer(bs service.BootstrapService, aps service.AccountPermissionService, gps service.GroupPermissionService, krs service.KubeconfigRevocationService, kcs service.KubectlClusterSettingsService, kss service.KubeconfigSettingService) sentryrpc.ClusterAuthorizationServer {
	return &clusterAuthzServer{
		bs:  bs,
		aps: aps,
		gps: gps,
		krs: krs,
		kcs: kcs,
		kss: kss,
	}
}
