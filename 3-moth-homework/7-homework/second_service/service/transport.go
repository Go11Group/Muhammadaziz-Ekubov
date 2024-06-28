package service

import (
	"context"
	"fmt"
	p "github.com/Go11Group/at_lesson/lesson47/second_service/genproto"
)

type TransportService struct {
	p.UnimplementedTransportServiceServer
	//Storage storage.Tr
	// Clinets grpcclinet.ServiceMeneger
}

func (u *TransportService) GetBusSchedule(ctx context.Context, in *p.BusNumber) (*p.BusSchudle, error) {
	fmt.Println("WORKED")
	return nil, nil
}
