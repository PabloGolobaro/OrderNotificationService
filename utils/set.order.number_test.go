package utils

import (
	"WhatsappOrderServer/config"
	"WhatsappOrderServer/db"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
)

func TestSetOrderNumber(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		ctx        context.Context
	}
	config, err := config.LoadConfig("C:\\Users\\Professional\\GolandProjects\\WhatsappOrderServer")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	mongoclient, err := db.NewMongoDBConnection(config, log.Default())
	if err != nil {
		return
	}
	testOrderCollection := mongoclient.Database("golang_mongodb").Collection("test_orders")
	orderCollection := mongoclient.Database("golang_mongodb").Collection("orders")
	testagruments := args{testOrderCollection, context.Background()}
	agruments := args{orderCollection, context.Background()}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "first", args: testagruments, want: 1, wantErr: false},
		{name: "second", args: agruments, want: 5, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetOrderNumber(tt.args.collection, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetOrderNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetOrderNumber() got = %v, want %v", got, tt.want)
			}
			t.Log(got, " : ", tt.want)
		})
	}
}
