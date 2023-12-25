package shopgeneralbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

type SeeShopGeneralStore interface {
	FindShopGeneral(
		ctx context.Context) (*shopgeneralmodel.ShopGeneral, error)
}

type seeShopGeneralBiz struct {
	store     SeeShopGeneralStore
	requester middleware.Requester
}

func NewSeeShopGeneralBiz(
	store SeeShopGeneralStore,
	requester middleware.Requester) *seeShopGeneralBiz {
	return &seeShopGeneralBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *seeShopGeneralBiz) SeeShopGeneral(
	ctx context.Context) (*shopgeneralmodel.ShopGeneral, error) {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return nil, shopgeneralmodel.ErrGeneralShopViewNoPermission
	}

	general, errGetGeneral := biz.store.FindShopGeneral(ctx)
	if errGetGeneral != nil {
		return nil, errGetGeneral
	}

	return general, nil
}
