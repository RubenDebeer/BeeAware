package mongo

import (
	"github.com/RubenDeBeer/BeeAware/internal/domain"
)

//
// Mappers: Mongo  domain
//
// From Domain to Mongo Mappings
/*
FromDomain converts domain.User -> MongoUser
If u.ID is empty, a new ObjectID is generated.
*/
func FromDomainUser(u domain.User) (MongoUser, error) {
	oid, err := hexToOIDOrNew(u.ID)
	if err != nil {
		return MongoUser{}, err
	}

	hives := make([]MongoHive, len(u.Hives))
	for i := range u.Hives {
		h, err := FromDomainHive(u.Hives[i])
		if err != nil {
			return MongoUser{}, err
		}
		hives[i] = h
	}

	return MongoUser{
		ID:    oid,
		Name:  u.Name,
		Hives: hives,
	}, nil
}

/*
FromDomainHive converts domain.Hive -> MongoHive
If h.ID is empty, a new ObjectID is generated.
*/
func FromDomainHive(h domain.Hive) (MongoHive, error) {
	oid, err := hexToOIDOrNew(h.ID)
	if err != nil {
		return MongoHive{}, err
	}

	return MongoHive{
		ID:              oid,
		GeoLocation:     FromDomainGeo(h.GeoLocation),
		Config:          h.Config,
		Description:     h.Description,
		LastServiceDate: h.LastServiceDate,
		NextServiceDate: h.NextServiceDate,
	}, nil
}

/*
FromDomainGoe converts domain.GeoLOcation -> MongoGeoLocation
If h.ID is empty, a new ObjectID is generated.
*/
func FromDomainGeo(g domain.GeoLocation) MongoGeoLocation {
	return MongoGeoLocation{Lat: g.Lat, Lon: g.Lon}
}

// From Mongo to Domain mappings
/*
ToDomain converts MongoUser -> domain.User

This Mapper function retuns a domain user
*/
func (mongoUser MongoUser) ToDomain() domain.User {
	hives := make([]domain.Hive, len(mongoUser.Hives))

	for i := range mongoUser.Hives {
		hives[i] = mongoUser.Hives[i].ToDomain()
	}

	return domain.User{
		ID:    oidToHex(mongoUser.ID),
		Name:  mongoUser.Name,
		Hives: hives,
	}
}

/*
ToDomain converts MongoHive -> domain.Hive

This Mapper function retuns a domain hive
*/
func (mongoHive MongoHive) ToDomain() domain.Hive {
	return domain.Hive{
		ID:              oidToHex(mongoHive.ID),
		GeoLocation:     mongoHive.GeoLocation.ToDomain(),
		Config:          mongoHive.Config,
		Description:     mongoHive.Description,
		LastServiceDate: mongoHive.LastServiceDate,
		NextServiceDate: mongoHive.NextServiceDate,
	}
}

/*
ToDomain converts MongoGeoLocation -> domain.GeoLocation

This Mapper function retuns a domain GeoLocation
*/
func (mongoLocation MongoGeoLocation) ToDomain() domain.GeoLocation {
	return domain.GeoLocation{Lat: mongoLocation.Lat, Lon: mongoLocation.Lon}
}
