package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"example.com/rpc29/common/config"
	"example.com/rpc29/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return model.NewUsersClient(conn)
}

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Batzorig Vaanchig",
		Password: "Chinggiskhaanimagtaal",
		Gender:   model.UserGender_MALE,
	}

	var user2 = &model.User{
		Id:       "u002",
		Name:     "Pikachu",
		Password: "Ichooseyou",
		Gender:   model.UserGender_MALE,
	}

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Mongolia",
		//47.8081° N, 107.5298° E
		Coordinate: &model.GarageCoordinate{
			Latitude:  47.8081,
			Longitude: 107.5298,
		},
	}

	var garage2 = &model.Garage{
		Id:   "g002",
		Name: "Pallet town",
		//47.8081° N, 107.5298° E
		Coordinate: &model.GarageCoordinate{
			Latitude:  47.8081,
			Longitude: 107.5298,
		},
	}

	var garageuser1 = &model.GarageAndUserId{
		UserId: "u001",
		Garage: garage1,
	}

	var garageuser2 = &model.GarageAndUserId{
		UserId: "u002",
		Garage: garage2,
	}

	user := serviceUser()

	fmt.Println("\n", "=======> user test")

	//register user
	user.Register(context.Background(), user1)

	user.Register(context.Background(), user2)

	// returns a list of users
	userlist, err := user.List(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal("error getting list of users")
	}
	fmt.Printf("Printing userlist \n")
	fmt.Printf("%+v", userlist)

	garage := serviceGarage()

	garage.Add(context.Background(), garageuser1)
	garage.Add(context.Background(), garageuser2)

	resgar1, _ := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	resgar2, _ := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	resgar1string, _ := json.Marshal(resgar1)
	resgar2string, _ := json.Marshal(resgar2)
	fmt.Printf("=======> garage list for user 1\n %s", resgar1string)
	fmt.Printf("=======> garage list for user 2\n %s", resgar2string)

	fmt.Println("closing conn")

}
