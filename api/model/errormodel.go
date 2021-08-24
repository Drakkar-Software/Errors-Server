package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// ErrorData stores usage info about a specific error
type ErrorData struct {
	Title          string    `json:"title,omitempty" bson:"title"`
	Type           string    `json:"type,omitempty" bson:"type"`
	StackTrace     []string  `json:"stacktrace,omitempty" bson:"stacktrace"`
	FirstTimestamp float64   `json:"firsttimestamp,omitempty" bson:"firstTimestamp"`
	LastTimestamp  float64   `json:"lasttimestamp,omitempty" bson:"lastTimestamp"`
	Count  		   int16     `json:"count,omitempty" bson:"count"`
	MetricsId      string    `json:"metricsid,omitempty" bson:"metricsId"`
}

// Error stores an ErrorData and its id
type Error struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Error ErrorData          `json:"_data,omitempty" bson:"_data"`
}
