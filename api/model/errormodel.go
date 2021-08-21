package model


// Error stores usage info about a specific error
type Error struct {
	Title       string   `json:"title,omitempty" bson:"title"`
	Type        string   `json:"type,omitempty" bson:"type"`
	StackTrace  string   `json:"stacktrace,omitempty" bson:"stacktrace"`
	Timestamp   float64  `json:"timestamp,omitempty" bson:"timestamp"`
	MetricsId   string   `json:"metricsid,omitempty" bson:"metricsId"`
}
