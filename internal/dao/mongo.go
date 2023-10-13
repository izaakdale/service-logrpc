package dao

import (
	"context"
	"log"
	"time"

	logrpc "github.com/izaakdale/service-logrpc/api/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Service = "service"
	Message = "message"
	Level   = "log_level"
	TraceId = "trace_id"
	Time    = "time"
)

type mongoConn struct {
	client *mongo.Client
}

func NewMongoConn(client *mongo.Client) *mongoConn {
	return &mongoConn{client}
}

func (m *mongoConn) Insert(ctx context.Context, service, message, trace string, level logrpc.Level) (string, error) {
	res, err := m.client.Database("service-logrpc").Collection("logs").InsertOne(context.TODO(),
		bson.D{
			{Key: Service, Value: service},
			{Key: Message, Value: message},
			{Key: Level, Value: level.String()},
			{Key: TraceId, Value: trace},
			{Key: Time, Value: time.Now()},
		},
	)
	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Printf("not object id\n")
	}
	return id.Hex(), nil
}

func (m *mongoConn) Fetch(ctx context.Context) ([]*logrpc.LogRecord, error) {
	var ret []*logrpc.LogRecord

	filter := bson.D{{}}                                                // Empty filter to select all documents
	options := options.Find().SetSort(bson.D{{Key: "time", Value: -1}}) // Sort by 'time' field in descending order

	// Create a cursor for querying the database
	cur, err := m.client.Database("service-logrpc").Collection("logs").Find(ctx, filter, options)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var record bson.M
		if err := cur.Decode(&record); err != nil {
			return nil, err
		}

		s := record[Service].(string)
		m := record[Message].(string)
		ll := record[Level].(string)
		t, ok := record[TraceId].(string)
		if !ok {
			log.Printf("no trace in log\n")
		}

		ret = append(ret, &logrpc.LogRecord{
			Service: s, Message: m, TraceId: t, LogLevel: logrpc.Level(logrpc.Level_value[ll]),
		})

	}
	return ret, nil
}
