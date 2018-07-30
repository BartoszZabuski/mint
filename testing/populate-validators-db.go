package main

import "crypto"

func main() {

	crypto.PublicKey

	// session, err := mgo.Dial("localhost")
	// if err != nil {
	// 	panic(err)
	// }
	// db := session.DB("tendermintdb")
	//
	// project := bson.M{
	// 	"$project": bson.M{
	// 		"name":      1,
	// 		"publicKey": 1,
	// 		"votes":     bson.M{"$size": "$votes"},
	// 	},
	// }
	// sort := bson.M{
	// 	"$sort": bson.M{
	// 		"votes": -1,
	// 	},
	// }
	//
	// limit := bson.M{
	// 	"$limit": 4,
	// }
	//
	// iter := db.C("validators").Pipe([]bson.M{project, sort, limit}).Iter()
	// defer iter.Close()
	//
	// var tdValidators []types.Validator
	//
	// var dbResult bson.M
	// for iter.Next(&dbResult) {
	// 	fmt.Println(dbResult)
	// 	validator := types.Validator{
	// 		PubKey: dbResult["publicKey"].([]byte),
	// 		Power:  10,
	// 	}
	// 	tdValidators = append(tdValidators, validator)
	// }
	// if iter.Err() != nil {
	// 	panic(iter.Err())
	// }
	// fmt.Println(tdValidators)

	// err = pipe.Iter()
	// if err != nil {
	// 	//handle error
	// }
	// fmt.Println(validators) // simple print proving it's working

	// validatorId := bson.ObjectIdHex("5b5cde18804584eac813eaee")
	// // err = db.C("validators").Update(bson.M{"_id": validatorId}, bson.M{"$addToSet": bson.M{"votes": bson.NewObjectId()}})
	// err = db.C("validators").Update(bson.M{"_id": validatorId}, bson.M{"$addToSet": bson.M{"votes": bson.ObjectIdHex("5b5ce897804584f2929d4e6b")}})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var validators []jsonstore.Validator
	// err = db.C("validators").Find(nil).All(&validators)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(validators)
	//
	// var tendermintValidators []types.Validator
	// for _, validator := range validators {
	// 	// TODO power should be saved to mongo too
	// 	tdValidator := types.Validator{PubKey: validator.PublicKey, Power: 10}
	// 	tendermintValidators = append(tendermintValidators, tdValidator)
	// }
	// fmt.Println(tendermintValidators)

}
