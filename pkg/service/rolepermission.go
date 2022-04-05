package service

import (
	"context"

	"github.com/RafayLabs/rcloud-base/internal/dao"
	"github.com/RafayLabs/rcloud-base/internal/models"
	"github.com/RafayLabs/rcloud-base/pkg/query"
	commonv3 "github.com/RafayLabs/rcloud-base/proto/types/commonpb/v3"
	v3 "github.com/RafayLabs/rcloud-base/proto/types/commonpb/v3"
	rolev3 "github.com/RafayLabs/rcloud-base/proto/types/rolepb/v3"
	"github.com/google/uuid"
	bun "github.com/uptrace/bun"
)

const (
	rolepermissionKind     = "RolePermission"
	rolepermissionListKind = "RolePermissionList"
)

// RolepermissionService is the interface for rolepermission operations
type RolepermissionService interface {
	// get rolepermission by name
	GetByName(context.Context, *rolev3.RolePermission) (*rolev3.RolePermission, error)
	// list rolepermissions
	List(context.Context, ...query.Option) (*rolev3.RolePermissionList, error)
}

// rolepermissionService implements RolepermissionService
type rolepermissionService struct {
	db *bun.DB
}

// NewRolepermissionService return new rolepermission service
func NewRolepermissionService(db *bun.DB) RolepermissionService {
	return &rolepermissionService{db: db}
}

func (s *rolepermissionService) toV3Rolepermission(rolepermission *rolev3.RolePermission, rlp *models.ResourcePermission) *rolev3.RolePermission {
	rolepermission.Metadata = &v3.Metadata{
		Name:        rlp.Name,
		Description: rlp.Description,
	}
	rolepermission.Spec = &rolev3.RolePermissionSpec{
		Scope: rlp.Scope,
	}

	return rolepermission
}

func (s *rolepermissionService) getPartnerOrganization(ctx context.Context, rolepermission *rolev3.RolePermission) (uuid.UUID, uuid.UUID, error) {
	partner := rolepermission.GetMetadata().GetPartner()
	org := rolepermission.GetMetadata().GetOrganization()
	partnerId, err := dao.GetPartnerId(ctx, s.db, partner)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}
	organizationId, err := dao.GetOrganizationId(ctx, s.db, org)
	if err != nil {
		return partnerId, uuid.Nil, err
	}
	return partnerId, organizationId, nil

}

func (s *rolepermissionService) GetByName(ctx context.Context, rolepermission *rolev3.RolePermission) (*rolev3.RolePermission, error) {
	name := rolepermission.GetMetadata().GetName()
	entity, err := dao.GetByName(ctx, s.db, name, &models.ResourcePermission{})
	if err != nil {
		return rolepermission, err
	}

	if rle, ok := entity.(*models.ResourcePermission); ok {
		rolepermission = s.toV3Rolepermission(rolepermission, rle)

		return rolepermission, nil
	}
	return rolepermission, nil

}

func (s *rolepermissionService) List(ctx context.Context, opts ...query.Option) (*rolev3.RolePermissionList, error) {
	queryOptions := commonv3.QueryOptions{}
	for _, opt := range opts {
		opt(&queryOptions)
	}
	var rolepermissions []*rolev3.RolePermission
	rolepermissionList := &rolev3.RolePermissionList{
		ApiVersion: apiVersion,
		Kind:       rolepermissionListKind,
		Metadata: &v3.ListMetadata{
			Count: 0,
		},
	}
	var rles []models.ResourcePermission
	if len(queryOptions.Selector) > 0 {
		rles, err := dao.GetRolePermissionsByScope(ctx, s.db, queryOptions.Selector)
		if err != nil {
			return rolepermissionList, err
		}
		for _, rle := range rles {
			entry := &rolev3.RolePermission{}
			entry = s.toV3Rolepermission(entry, &rle)
			rolepermissions = append(rolepermissions, entry)
		}

		//update the list metadata and items response
		rolepermissionList.Metadata = &v3.ListMetadata{
			Count: int64(len(rolepermissions)),
		}
		rolepermissionList.Items = rolepermissions
	} else {
		_, err := dao.List(ctx, s.db, uuid.NullUUID{UUID: uuid.Nil, Valid: false}, uuid.NullUUID{UUID: uuid.Nil, Valid: false}, &rles)
		if err != nil {
			return rolepermissionList, err
		}
		for _, rle := range rles {
			entry := &rolev3.RolePermission{}
			entry = s.toV3Rolepermission(entry, &rle)
			rolepermissions = append(rolepermissions, entry)
		}

		//update the list metadata and items response
		rolepermissionList.Metadata = &v3.ListMetadata{
			Count: int64(len(rolepermissions)),
		}
		rolepermissionList.Items = rolepermissions
	}

	return rolepermissionList, nil
}
