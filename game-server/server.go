package gameserver

import (
	"sync"
	"time"

	"github.com/simonvetter/modbus"
)

type Server struct{}

func (s *Server) Start(host string) error {
	var server *modbus.ModbusServer
	var err error

	handler := exampleHandler{}

	server, err = modbus.NewServer(&modbus.ServerConfiguration{
		URL:        host,
		Timeout:    30 * time.Second,
		MaxClients: 5,
	}, &handler)

	if err != nil {
		return err
	}

	err = server.Start()
	if err != nil {
		return err
	}

	return nil
}

type exampleHandler struct {
	lock sync.RWMutex

	uptime uint32

	coils [100]bool
}

func (eh *exampleHandler) HandleCoils(req *modbus.CoilsRequest) ([]bool, error) {
	if req.UnitId != 1 {
		return nil, modbus.ErrIllegalFunction
	}

	if int(req.Addr) > 10_000 {
		return nil, modbus.ErrIllegalDataAddress
	}

	if int(req.Addr)+int(req.Quantity) > len(eh.coils) {
		return nil, modbus.ErrIllegalDataAddress
	}

	eh.lock.Lock()
	defer eh.lock.Unlock()

	var res []bool
	for i := 0; i < int(req.Quantity); i++ {
		res = append(res, eh.coils[int(req.Addr)+i])
	}

	return res, nil
}

func (eh *exampleHandler) HandleDiscreteInputs(req *modbus.DiscreteInputsRequest) ([]bool, error) {
	return nil, modbus.ErrIllegalDataAddress
}

func (eh *exampleHandler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) ([]uint16, error) {
	return nil, modbus.ErrIllegalDataAddress
}

func (eh *exampleHandler) HandleInputRegisters(req *modbus.InputRegistersRequest) ([]uint16, error) {
	return nil, modbus.ErrIllegalDataAddress
}
