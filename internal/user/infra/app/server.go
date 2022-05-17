package userapp

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ pbv1.UserServiceServer = (*server)(nil)

func NewServer() pbv1.UserServiceServer {
	return new(server)
}

type server struct {
}

func (s server) Register(ctx context.Context, request *pbv1.RegisterRequest) (*pbv1.TokenResponse, error) {
	return &pbv1.TokenResponse{
		Token: randomdata.Alphanumeric(128),
	}, nil
}

func (s server) GetByUUID(ctx context.Context, request *pbv1.GetByUserUUIDRequest) (*pbv1.User, error) {
	return &pbv1.User{
		Uuid:        uuid.New().String(),
		Group:       pbv1.UserGroupCustomer,
		Status:      pbv1.UserStatusActive,
		PhoneNumber: randomdata.PhoneNumber(),
		Lastname:    randomdata.LastName(),
		Firstname:   randomdata.FirstName(1),
		CreatedAt:   timestamppb.New(time.Now()),
	}, nil
}

func (s server) GetUserProducts(ctx context.Context, request *pbv1.GetByUserUUIDRequest) (*pbv1.UserProductsResponse, error) {
	return &pbv1.UserProductsResponse{
		Products: []*pbv1.UserProduct{
			{
				Uuid:        uuid.NewString(),
				ProductUuid: uuid.NewString(),
				Count:       uint32(randomdata.Number(1, 10)),
			},
			{
				Uuid:        uuid.NewString(),
				ProductUuid: uuid.NewString(),
				Count:       uint32(randomdata.Number(1, 10)),
			},
			{
				Uuid:        uuid.NewString(),
				ProductUuid: uuid.NewString(),
				Count:       uint32(randomdata.Number(1, 10)),
			},
		},
	}, nil
}

func (s server) PutProduct(ctx context.Context, request *pbv1.PutProductRequest) (*pbv1.UserProductsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) ChangeProduct(ctx context.Context, request *pbv1.ChangeProductRequest) (*pbv1.UserProductsResponse, error) {
	//TODO implement me
	panic("implement me")
}
