package main

import (
	"context"
	"errors"
	"sync"
	"time"

	payment "github.com/ChaoJiCaiNiao3/dymall/app/payment/kitex_gen/payment"
)

// PaymentStatus 支付状态
type PaymentStatus int32

const (
	PaymentStatusPending PaymentStatus = iota
	PaymentStatusSuccess
	PaymentStatusCancelled
)

// PaymentRecord 支付记录
type PaymentRecord struct {
	OrderID     string
	Amount      float64
	Status      PaymentStatus
	CreatedAt   time.Time
	CompletedAt time.Time
	Timer       *time.Timer
}

// PaymentServiceImpl 实现支付服务接口
type PaymentServiceImpl struct {
	payments sync.Map
	timeout  time.Duration
}

// Charge 实现支付功能
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// 参数验证
	if req.OrderId == "" || req.Amount <= 0 {
		return nil, errors.New("invalid parameters")
	}

	// 检查订单是否已存在
	if _, exists := s.payments.Load(req.OrderId); exists {
		return nil, errors.New("order already exists")
	}

	// 创建支付记录
	record := &PaymentRecord{
		OrderID:   req.OrderId,
		Amount:    float64(req.Amount),
		Status:    PaymentStatusPending,
		CreatedAt: time.Now(),
	}

	// 设置自动取消定时器（30分钟后自动取消）
	record.Timer = time.AfterFunc(30*time.Minute, func() {
		s.cancelPayment(req.OrderId)
	})

	// 存储支付记录
	s.payments.Store(req.OrderId, record)

	// 返回响应
	return &payment.ChargeResp{}, nil
}

// 内部方法：取消支付
func (s *PaymentServiceImpl) cancelPayment(orderID string) error {
	value, ok := s.payments.Load(orderID)
	if !ok {
		return errors.New("payment not found")
	}

	record := value.(*PaymentRecord)
	if record.Status != PaymentStatusPending {
		return errors.New("payment cannot be cancelled")
	}

	record.Status = PaymentStatusCancelled
	record.CompletedAt = time.Now()
	if record.Timer != nil {
		record.Timer.Stop()
	}

	s.payments.Store(orderID, record)
	return nil
}

// 内部方法：完成支付
func (s *PaymentServiceImpl) completePayment(orderID string) error {
	value, ok := s.payments.Load(orderID)
	if !ok {
		return errors.New("payment not found")
	}

	record := value.(*PaymentRecord)
	if record.Status != PaymentStatusPending {
		return errors.New("payment cannot be completed")
	}

	record.Status = PaymentStatusSuccess
	record.CompletedAt = time.Now()
	if record.Timer != nil {
		record.Timer.Stop()
	}

	s.payments.Store(orderID, record)
	return nil
}
