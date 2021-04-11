package coModels

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	CustomerOrder struct {
		ID               primitive.ObjectID `bson:"_id,omitempty"`
		Email            string             `bson:"Email"`
		OwnReferralCode  int                `bson:"OwnReferralCode"`
		UsedReferralCode int                `bson:"UsedReferralCode"`
		WaiterName       string             `bson:"WaiterName"`
		Orders           []string           `bson:"Orders"`
	}
)

func (m CustomerOrder) IsHighPriority() bool {

	// have 3 order and selectedServicePackage minimum 4 stars, 1 order if 5 stars

	return len(m.Orders) > 3
}
