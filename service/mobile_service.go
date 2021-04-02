package service

import (
	"context"
	"errors"
	"log"

	"github.com/MobileStore-Grpc/product/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//MobileService provide mobile CURD operations
type MobileService struct {
	mobileStore MobileStore
	pb.UnimplementedMobileServiceServer
}

//NewMobileService returns a new mobile service
func NewMobileService(mobileStore MobileStore) *MobileService {
	return &MobileService{
		mobileStore: mobileStore,
	}
}

//CreateMobile is a unary RPC to create a new mobile
func (server *MobileService) CreateMobile(ctx context.Context, req *pb.CreateMobileRequest) (*pb.CreateMobileResponse, error) {
	mobile := req.GetMobile()
	log.Printf("receive a create-mobile request with id: %s", mobile.Id)

	if len(mobile.Id) > 0 {
		// check if it's a valid UUID
		_, err := uuid.Parse(mobile.Id)
		if err != nil {
			return nil, logError(status.Errorf(codes.InvalidArgument, "mobile ID is not a valid UUID: %v", err))
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, logError(status.Errorf(codes.Internal, "cannot generate a new mobile id: %v", err))
		}
		mobile.Id = id.String()
	}

	//some heavy processing
	// time.Sleep(6 * time.Second)

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	//save the laptop to in-memory store
	err := server.mobileStore.Save(mobile)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExits) {
			code = codes.AlreadyExists
		}
		return nil, status.Error(code, "cannot save mobile to the in-memory store")
	}
	log.Printf("saved mobile with id: %s", mobile.Id)
	// will change laptop object value here so that value in Inmemeory database laptop object also get changed {without having deep copy save}
	// and after change value from here , verify that value has changed by geeting latop value from database.
	res := &pb.CreateMobileResponse{
		Id: mobile.Id,
	}
	return res, nil
}
func (server *MobileService) SearchMobile(ctx context.Context, req *pb.SearchMobileRequest) (*pb.SearchMobileResponse, error) {
	mobile_id := req.GetMobileId()
	log.Printf("receive a search-laptop request with mobile_id: %v", mobile_id)
	mobile, err := server.mobileStore.Search(mobile_id)
	res := &pb.SearchMobileResponse{
		Mobile: mobile,
	}
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}

	if mobile == nil {
		return nil, logError(status.Errorf(codes.InvalidArgument, "laptop %s doesn't exist", err))
	}
	return res, nil

}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
