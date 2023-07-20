package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"example.com/rpc29/model"
	"github.com/golang/protobuf/jsonpb"
)

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Batzorig Vaanchig",
		Password: "Chinggis khaani magtaal",
		Gender:   model.UserGender_FEMALE,
	}

	var userlist = &model.UserList{
		List: []*model.User{
			user1,
		},
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

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	var garageListByUser = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	_ = userlist
	_ = garageListByUser

	fmt.Printf("original \n 		%#v\n", user1)
	fmt.Printf("string \n		%#v\n", user1.String())

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonStr := buf.String()
	fmt.Printf("Json \n		%v \n", jsonStr)

	buf2 := strings.NewReader(jsonStr)
	protoObj := new(model.GarageList)

	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObj)
	if err1 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("string \n		%v \n", protoObj.String())
	err2 = jsonpb.UnmarshalString(jsonStr, protoObj)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== As String\n       %v \n", protoObj.String())
}
