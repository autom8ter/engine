package lib

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine/util"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

func MongoClient(addr string, ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(addr))
	if err != nil {
		grpclog.Warningln(err.Error())
		return nil
	}
	return client
}

func NewUnaryPingMongoInterceptor(c *mongo.Client, ctx context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = c.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Warnln(err.Error())
			return nil, status.Errorf(codes.Internal, "database unavailable: %s", err.Error())
		}
		return handler(ctx, req)
	}
}

func NewUnarySaveMongoInterceptor(c *mongo.Client, database, collection string, keyFromCtx interface{}) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		coll := c.Database(database).Collection(collection)

		res, err := coll.InsertOne(ctx, bson.M{
			util.FromContext(ctx, keyFromCtx): req,
		})
		if err != nil {
			log.Warnln(err.Error())
			return nil, status.Error(codes.Internal, fmt.Sprintf("database insertion failure: %s", err.Error()))
		}
		id := res.InsertedID
		grpclog.Infof("mongo insertion id: %v\n", id)
		return handler(ctx, req)
	}
}
