package service

import (
	"context"
	"fmt"

	p "Muhammadaziz-Ekubov/3-moth-homework/7-homework/second_service/genproto"
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
