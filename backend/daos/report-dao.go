package daos

import (
	"context"
	"errors"
	"fmt"
	"play-on-eleven/backend/models"
	"play-on-eleven/backend/server"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "reports"
)

type ReportDao struct {
	*mongo.Collection
}

func NewReportDao(db *server.Database) *ReportDao {
	collection := db.Collection(collectionName)
	return &ReportDao{collection}
}

func (r *ReportDao) FetchAllReports() ([]models.Report, error) {
	var reports []models.Report
	cur, err := r.Find(context.Background(), bson.D{})

	if err != nil {
		fmt.Println("error fetching reports from collection", err)
		return nil, err
	}

	// defer close cursor
	defer func() {
		err := cur.Close(context.Background())
		if err != nil {
			// logger.WithField("error", err).WithField("query", query).Warn("error closing cursor")
			// return nil, err
		}
	}()

	err = cur.All(context.Background(), &reports)

	if err != nil {
		fmt.Println("error fetching reports")
		return nil, err
	}

	if len(reports) == 0 {
		return nil, errors.New("no result found")
	}
	fmt.Println("sending reports")
	fmt.Print(reports)
	return reports, nil
}

func (r *ReportDao) CreateReport(report *models.Report) error {
	fmt.Println("creating report")
	_, err := r.InsertOne(context.TODO(), report)
	return err
}
