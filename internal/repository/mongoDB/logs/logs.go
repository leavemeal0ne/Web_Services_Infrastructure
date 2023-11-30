package repositorylogs

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"lab4/internal/models"
	"time"
)

type LogsMongo struct {
	Collection *mongo.Collection
}

func NewLogsMongo(collection *mongo.Collection) *LogsMongo {
	return &LogsMongo{Collection: collection}
}

func (l *LogsMongo) InsertLog(data models.LogData) error {
	_, err := l.Collection.InsertOne(
		context.TODO(),
		bson.D{
			{"Time", time.Now().Format("2006-01-02 15:04:05 Monday")},
			{"IP", data.IP},
			{"Method", data.Method},
			{"URI", data.URI},
		},
	)
	return err
}
