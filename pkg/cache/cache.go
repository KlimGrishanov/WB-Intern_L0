package cache

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/usecase"
	"go.uber.org/zap"
)

type CacheService struct {
	services *usecase.UseCase
	Cache    map[string]entity.Order
}

func NewCache(services *usecase.UseCase) *CacheService {
	return &CacheService{services: services, Cache: make(map[string]entity.Order)}
}

func (s *CacheService) NewOrder(order entity.Order) {
	s.Cache[order.OrderUid] = order
	zap.L().Info("Cache: Add new Order")
}

func (s *CacheService) GetOrder(orderUID string) (entity.Order, bool) {
	val, isExist := s.Cache[orderUID]
	zap.L().Info("Cache: Send order from cache")
	return val, isExist
}

func (s *CacheService) LoadDataFromDatabase() error {
	orderArr, err := s.services.GetOrders()
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	for i := 0; i < len(orderArr); i++ {
		Delivery, err := s.services.GetDeliveryByOrderUID(orderArr[i].OrderUid)
		if err != nil {
			return err
		}
		Payment, err := s.services.GetPaymentByOrderUID(orderArr[i].OrderUid)
		if err != nil {
			return err
		}
		Items, err := s.services.GetItemsByOrderUID(orderArr[i].OrderUid)
		if err != nil {
			return err
		}
		orderArr[i].Delivery = Delivery
		orderArr[i].Payment = Payment
		orderArr[i].Items = Items
		s.Cache[orderArr[i].OrderUid] = orderArr[i]
	}
	zap.L().Info("Cache: Load Cache from Database -> Success")
	return nil
}
