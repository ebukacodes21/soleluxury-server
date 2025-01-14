package db

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DatabaseContract interface {
	CreateUser(ctx context.Context, args *pb.CreateUserRequest) (*User, error)
	FindUser(ctx context.Context, arg *pb.LoginUserRequest) (*User, error)
	CreateSession(ctx context.Context, args SessionReq) (*Session, error)
	LogOut(ctx context.Context, id bson.ObjectID) error

	CreateStore(ctx context.Context, arg *pb.CreateStoreRequest) (*Store, error)
	GetFirstStore(ctx context.Context) (*Store, error)
	GetStoreByID(ctx context.Context, req *pb.GetStoreRequest) (*Store, error)
	GetAllStores(ctx context.Context) ([]Store, error)
	UpdateStore(ctx context.Context, args *pb.UpdateStoreRequest) (string, error)
	DeleteStore(ctx context.Context, arg *pb.DeleteStoreRequest) (string, error)

	CreateBillboard(ctx context.Context, args *pb.CreateBillboardRequest) (*Billboard, error)
	GetBillboard(ctx context.Context, args *pb.GetBillboardRequest) (*Billboard, error)
	GetAllBillboards(ctx context.Context, args *pb.GetBillboardsRequest) ([]Billboard, error)
	UpdateBillboard(ctx context.Context, args *pb.UpdateBillboardRequest) (string, error)
	DeleteBillboard(ctx context.Context, args *pb.DeleteBillboardRequest) (string, error)

	CreateCategory(ctx context.Context, args *pb.CreateCategoryRequest) (*Category, error)
	GetCatgeoryByID(ctx context.Context, req *pb.GetCategoryRequest) (*Category, error)
	GetAllCategories(ctx context.Context, args *pb.GetCategoriesRequest) ([]Category, error)
	UpdateCategory(ctx context.Context, args *pb.UpdateCategoryRequest) (string, error)
	DeleteCategory(ctx context.Context, arg *pb.DeleteCategoryRequest) (string, error)

	CreateSize(ctx context.Context, args *pb.CreateSizeRequest) (*Size, error)
	GetSize(ctx context.Context, args *pb.GetSizeRequest) (*Size, error)
	GetAllSizes(ctx context.Context, args *pb.GetSizesRequest) ([]Size, error)
	UpdateSize(ctx context.Context, args *pb.UpdateSizeRequest) (string, error)
	DeleteSize(ctx context.Context, arg *pb.DeleteSizeRequest) (string, error)

	CreateColor(ctx context.Context, args *pb.CreateColorRequest) (*Color, error)
	GetColor(ctx context.Context, args *pb.GetColorRequest) (*Color, error)
	GetAllColors(ctx context.Context, args *pb.GetColorsRequest) ([]Color, error)
	UpdateColor(ctx context.Context, args *pb.UpdateColorRequest) (string, error)
	DeleteColor(ctx context.Context, arg *pb.DeleteColorRequest) (string, error)

	CreateProduct(ctx context.Context, args *pb.CreateProductRequest) (*Product, error)
	GetProducts(ctx context.Context, args *pb.GetProductsRequest) ([]Product, error)
	GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*Product, error)
	UpdateProduct(ctx context.Context, args *pb.UpdateProductRequest) (string, error)
	DeleteProduct(ctx context.Context, arg *pb.DeleteProductRequest) (string, error)
}

type Repository struct {
	client        *mongo.Client
	database      *mongo.Database
	userColl      *mongo.Collection
	sessionColl   *mongo.Collection
	productColl   *mongo.Collection
	storeColl     *mongo.Collection
	colorColl     *mongo.Collection
	sizeColl      *mongo.Collection
	categoryColl  *mongo.Collection
	billboardColl *mongo.Collection
	orderColl     *mongo.Collection
}

func NewRepository(client *mongo.Client, databaseName string) DatabaseContract {
	db := client.Database(databaseName)
	return &Repository{
		client:        client,
		database:      db,
		userColl:      db.Collection("users"),
		sessionColl:   db.Collection("sessions"),
		productColl:   db.Collection("products"),
		storeColl:     db.Collection("stores"),
		colorColl:     db.Collection("colors"),
		sizeColl:      db.Collection("sizes"),
		categoryColl:  db.Collection("categories"),
		billboardColl: db.Collection("billboards"),
		orderColl:     db.Collection("orders"),
	}
}
