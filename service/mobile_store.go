package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/MobileStore-Grpc/product/pb"
	"github.com/jinzhu/copier"
)

//ErrAlreadyExits is returned when a record with the same ID already exist in the data store
var ErrAlreadyExits = errors.New("mobile with same specification already exist")

// NewInMemoryMobileStore returns a new NewInMemoryMobileStore object
func NewInMemoryMobileStore() *InMemoryMobileStore {
	return &InMemoryMobileStore{
		data: make(map[string]*pb.Mobile),
	}
}

// MobileStore is an interface to store and search mobile
type MobileStore interface {
	// Save the mobile to the store
	Save(mobile *pb.Mobile) error
	//Find finds a mobile by ID
	Search(id string) (*pb.Mobile, error)
	// Search searches for laptops with filter, return one by one via the found function
	// Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

// InMemoryMobileStore stores mobile in memory
type InMemoryMobileStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Mobile
}

// Save saves the mobile to the store
func (store *InMemoryMobileStore) Save(mobile *pb.Mobile) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.data[mobile.Id] != nil {
		return ErrAlreadyExits
	}

	// deep copy of the laptop object in the in-memory store
	other, err := deepCopy(mobile)
	if err != nil {
		return err
	}
	store.data[other.Id] = other
	return nil
}

//Find finds a laptop by ID
func (store *InMemoryMobileStore) Search(id string) (*pb.Mobile, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	mobile := store.data[id]

	if mobile == nil {
		return nil, nil
	}

	return deepCopy(mobile)
}

// Search searches for laptops with filter, return one by one via the found function
// func (store *InMemoryMobileStore) Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error {
// 	store.mutex.RLock()
// 	defer store.mutex.RUnlock()

// 	for _, laptop := range store.data {

// 		// heavy processing
// 		// time.Sleep(time.Second)
// 		// log.Print("checking laptop id: ", laptop.GetId())

// 		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
// 			log.Print("context is cancelled")
// 			return errors.New("context is cancelled")
// 		}

// 		if isQualified(filter, laptop) {
// 			other, err := deepCopy(laptop)
// 			if err != nil {
// 				return err
// 			}
// 			err = found(other)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
// 	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
// 		return false
// 	}

// 	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
// 		return false
// 	}

// 	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
// 		return false
// 	}

// 	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
// 		return false
// 	}
// 	return true
// }

// func toBit(memory *pb.Memory) uint64 {
// 	value := memory.GetValue()

// 	switch memory.GetUnit() {
// 	case pb.Memory_BIT:
// 		return value
// 	case pb.Memory_BYTE:
// 		return value << 3 // 8 = 2^3
// 	case pb.Memory_KILOBYTE:
// 		return value << 13 // 1024 * 8 = 2^10 * 2^3 = 2^13
// 	case pb.Memory_MEGABYTE:
// 		return value << 23
// 	case pb.Memory_GIGABYTE:
// 		return value << 33
// 	case pb.Memory_TERABYTE:
// 		return value << 43
// 	default:
// 		return 0
// 	}
// }

func deepCopy(mobile *pb.Mobile) (*pb.Mobile, error) {
	other := &pb.Mobile{}
	err := copier.Copy(other, mobile)
	if err != nil {
		return nil, fmt.Errorf("cannot deep-copy mobile data: %w", err)
	}
	return other, nil
}
