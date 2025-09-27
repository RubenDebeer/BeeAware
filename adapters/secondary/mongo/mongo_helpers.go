package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

//
// Small helpers
//

func oidToHex(oid primitive.ObjectID) string {
	if oid == primitive.NilObjectID {
		return ""
	}
	return oid.Hex()
}

func hexToOIDOrNew(hex string) (primitive.ObjectID, error) {
	if hex == "" {
		return primitive.NewObjectID(), nil
	}
	return primitive.ObjectIDFromHex(hex)
}
