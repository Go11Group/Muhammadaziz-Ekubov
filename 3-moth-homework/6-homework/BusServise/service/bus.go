package service

import (
	pb "Muhammadaziz-Ekubov/3-moth-homework/5-homework/BusServise/genproto/transport_service"
	"Muhammadaziz-Ekubov/3-moth-homework/5-homework/BusServise/storage/postgres"
)

type Bus struct {
	pb.UnimplementedTransportServiceServer
	BusRepo postgres.BusRepo
}

func NewBus(bus *postgres.BusRepo) *Bus {
	return &Bus{
		BusRepo: *bus,
	}
}

//func (b *Bus) GetBusSchedule(ctx context.Context, req *pb.BusScheduleRequest) (*pb.BusScheduleResponse, error) {
//	_, err := b.BusRepo.GetByBusNumber(int(req.GetBusNumber()))
//	if err != nil {
//		return &pb.BusScheduleResponse{}, err
//	}
//
//	return &pb.BusScheduleResponse{Stations: nil}, nil
//}
//func (t *Bus) TrackBusLocation(ctx context.Context, req *pb.BusScheduleResponse) (*pb.Station, error) {
//	bus, err := t.BusRepo.GetByBusNumber(1)
//	if err != nil {
//		return &pb.Station{}, err
//	}
//	ind := rand.Intn(10)
//	curStation := bus.make([]any, ind)
//	return &pb.Station{StationName: curStation}, err
//}
//
//func (t *Bus) ReportTrafficJam(ctx context.Context, req *pb.TrafficJamRequest) (*pb.TrafficJamResponse, error) {
//	return &pb.TrafficJamResponse{IsReport: true}, nil
//}
