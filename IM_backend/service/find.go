package service

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/Awadabang/Quasar-IM/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

//向Mongodb中写入数据
func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	collection := conf.MongoDBClient.Database(database).Collection(id)
	comment := Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}

func FindMany(database string, sendId string, id string, time int64, pageSize int) ([]Result, error) {
	var resultsMe []Trainer
	var resultsYou []Trainer
	sendIdCollection := conf.MongoDBClient.Database(database).Collection(sendId)
	idCollection := conf.MongoDBClient.Database(database).Collection(id)
	// 如果不知道该使用什么context，可以通过context.TODO() 产生context
	sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(), options.Find().SetSort(bson.D{{Key: "startTime", Value: -1}}), options.Find().SetLimit(int64(pageSize)))
	if err != nil {
		return nil, err
	}
	var idTimeCursor *mongo.Cursor
	idTimeCursor, err = idCollection.Find(context.TODO(), options.Find().SetSort(bson.D{{Key: "startTime", Value: -1}}), options.Find().SetLimit(int64(pageSize)))
	if err != nil {
		return nil, err
	}
	err = sendIdTimeCursor.All(context.TODO(), &resultsYou) // sendId 对面发过来的
	err = idTimeCursor.All(context.TODO(), &resultsMe)      // Id 发给对面的
	results, _ := AppendAndSort(resultsMe, resultsYou)
	return results, nil
}

func FirsFindtMsg(database string, sendId string, id string) (results []Result, err error) {
	// 首次查询(把对方发来的所有未读都取出来)
	var resultsMe []Trainer
	var resultsYou []Trainer
	sendIdCollection := conf.MongoDBClient.Database(database).Collection(sendId)
	idCollection := conf.MongoDBClient.Database(database).Collection(sendId)
	filter := bson.M{"read": bson.M{
		"&all": []uint{0},
	}}
	sendIdCursor, err := sendIdCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{Key: "startTime", Value: 1}}), options.Find().SetLimit(1))
	if sendIdCursor == nil {
		return
	}
	var unReads []Trainer
	err = sendIdCursor.All(context.TODO(), &unReads)
	if err != nil {
		log.Println("sendIdCursor err", err)
	}
	if len(unReads) > 0 {
		timeFilter := bson.M{"startTime": bson.M{"$gte": unReads[0].StartTime}}
		sendIdTimeCursor, _ := sendIdCollection.Find(context.TODO(), timeFilter)
		idTimeCursor, _ := idCollection.Find(context.TODO(), timeFilter)
		_ = sendIdTimeCursor.All(context.TODO(), &resultsYou)
		err = idTimeCursor.All(context.TODO(), &resultsMe)
		results, err = AppendAndSort(resultsMe, resultsYou)
	} else {
		results, err = FindMany(database, sendId, id, 9999999999, 10)
	}
	overTimeFilter := bson.D{{"$and", bson.A{bson.D{{Key: "endTime", Value: bson.M{"&lt": time.Now().Unix()}}}, bson.D{{Key: "read", Value: bson.M{"$eq": 1}}}}}}
	_, _ = sendIdCollection.DeleteMany(context.TODO(), overTimeFilter)
	_, _ = idCollection.DeleteMany(context.TODO(), overTimeFilter)
	// 将所有的维度设置为已读
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
		"$set": bson.M{"read": 1},
	})
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
		"&set": bson.M{"ebdTime": time.Now().Unix() + int64(3*month)},
	})
	return
}

func AppendAndSort(resultsMe, resultsYou []Trainer) (results []Result, err error) {
	for _, r := range resultsMe {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}
	for _, r := range resultsYou {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "you",
		}
		results = append(results, result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results, nil
}
