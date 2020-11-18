package main

import (
	"context"
	"fmt"
	"github.com/newrelic/go-agent/_integrations/nrgrpc"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/logic-building/functional-go/fp"

	"google.golang.org/grpc"

	grpc2 "quote/pkg/grpcquote"
	"quote/pkg/model"
)

const (
	address = "localhost:1923"
)

var operationList = []string{"get", "gets", "put"}
var operation = "get" // get | gets | put
var ID int64 = 41
var IDs = fp.RangeInt64(10, 20)

var updateRecord = `
{
 "id": 41,
 "day": 10,
 "month": 5,
 "year": 2022,
 "title": "Sita Ji appearance day. Jai shree Ram",
 "info": "Sita Navami, Sita Jayanti",
 "links": [
  "https://www.prokerala.com/festivals/sita-navami.html"
 ],
 "type": "different",
 "updatedAt": "2020-10-14 14:03",
 "createdAt": "2020-03-02 14:13"
}
`

func main() {
	if len(os.Args) > 1 {
		if len(os.Args) != 3 {
			help()
			return
		}

		if !fp.ExistsStrIgnoreCase(os.Args[1], operationList) {
			fmt.Printf("1st argument(opration) %s is not allowed. allowed operations are - %v \n", os.Args[1], operationList)
			help()
			return
		}

		operation = strings.ToLower(os.Args[1])

		switch operation {

		case "get":
			var err error
			ID, err = strconv.ParseInt(os.Args[2], 10, 64)
			if err != nil {
				fmt.Println("error converting 3rd argument. 2nd argument must be integer", err)
				help()
				return
			}

		case "put":
			jsonFile, err := os.Open(os.Args[2])
			if err != nil {
				fmt.Printf("\nerr opening file %v, err=%v", jsonFile, err)
				help()
				return
			}
			jsonBytes, err := ioutil.ReadAll(jsonFile)
			if err != nil {
				fmt.Println("error reading json file", err)
				help()
				return
			}

			eventDetail := model.EventDetail{}
			err = eventDetail.UnmarshalJSON(jsonBytes)

			if err != nil {
				fmt.Println("error unmarshalling json. 3rd argument for put is not valid", err)
				return
			}

			updateRecord = string(jsonBytes)

		case "gets":
			var ids []int64
			idsStr := strings.Split(os.Args[2], ",")
			for _, idStr := range idsStr {
				id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
				if err != nil {
					fmt.Println("error converting string to integer for gets - grpc stream")
					return
				}

				ids = append(ids, id)
			}
			IDs = ids
		}

	}
	// steps to update record
	// 1. operation should be get
	// 2. run the client
	// 3. replace updateRecord section with output of client
	// 4. change operation to put
	// 5. Run the client
	// 6. record updated successful.

	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(nrgrpc.StreamClientInterceptor))
	if err != nil {
		fmt.Println("error connecting grpc", err)
	}
	defer conn.Close()

	client := grpc2.NewEventDetailServiceGRPCClient(conn)

	switch operation {
	case "get":

		response, err := client.GetEventDetail(context.Background(), &grpc2.EventDetailRequest{Id: ID})
		if err != nil {
			fmt.Println("error getting event detail's data", err)
		}
		t := time.Unix(response.EventDetail.EventDate.Seconds, 0)

		eventDetail := model.EventDetail{
			ID:           response.EventDetail.Id,
			Day:          t.Day(),
			Month:        int(t.Month()),
			Year:         t.Year(),
			Title:        response.EventDetail.GetTitle(),
			Info:         response.EventDetail.Info,
			Links:        response.EventDetail.Links,
			Type:         response.EventDetail.EventType,
			CreationDate: time.Unix(response.EventDetail.CreatedAt.Seconds, 0).UTC(),
			UpdatedAt:    time.Unix(response.EventDetail.UpdatedAt.Seconds, 0).UTC(),
		}

		fmt.Println("event detail information...")
		json, err := eventDetail.ToJson()
		if err != nil {
			fmt.Println("error in json conversion", err)
		}
		fmt.Println(string(json))

		fmt.Println("Pretty version of json...")
		json, err = eventDetail.ToJsonIndent()
		if err != nil {
			fmt.Println("err json conversion", err)
		}
		fmt.Println(string(json))

	case "put":
		eventDetail := model.EventDetail{}
		err := eventDetail.UnmarshalJSON([]byte(updateRecord))

		if err != nil {
			fmt.Println("error unmarshalling json", err)
		}
		fmt.Println(eventDetail)
		response, err := client.UpdateEventDetail(context.Background(), &grpc2.EventDetailUpdateRequest{
			EventDetail: &grpc2.EventDetail{
				Id:        eventDetail.ID,
				Title:     eventDetail.Title,
				Info:      eventDetail.Info,
				EventType: eventDetail.Type,
				Links:     eventDetail.Links,
			},
			Dd:   int32(eventDetail.Day),
			Mm:   int32(eventDetail.Month),
			Yyyy: int32(eventDetail.Year),
		})
		if err != nil {
			fmt.Println("error updating event detail's data", err)
		}
		fmt.Println(response.Msg)

	case "gets":
		stream, err := client.GetEventDetailStream(context.Background())
		if err != nil {
			println("error in getting event detail stream", err)
		}

		var wg sync.WaitGroup

		var eventDetailList []model.EventDetail

		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				in, err := stream.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					fmt.Println("error receiving stream", err)
					return
				}

				t := time.Unix(in.EventDetail.EventDate.Seconds, 0)
				eventDetail := model.EventDetail{
					ID:           in.EventDetail.Id,
					Day:          t.Day(),
					Month:        int(t.Month()),
					Year:         t.Year(),
					Title:        in.EventDetail.GetTitle(),
					Info:         in.EventDetail.Info,
					Links:        in.EventDetail.Links,
					Type:         in.EventDetail.EventType,
					CreationDate: time.Unix(in.EventDetail.CreatedAt.Seconds, 0).UTC(),
					UpdatedAt:    time.Unix(in.EventDetail.UpdatedAt.Seconds, 0).UTC(),
				}
				eventDetailList = append(eventDetailList, eventDetail)
			}
		}()

		for _, id := range IDs {
			if err := stream.Send(&grpc2.EventDetailRequest{Id: id}); err != nil {
				fmt.Println("error streaming id ", err)
			}
		}
		stream.CloseSend()
		wg.Wait()

		for _, eventDetail := range eventDetailList {
			fmt.Println("event detail information...")
			json, err := eventDetail.ToJson()
			if err != nil {
				fmt.Println("error in json conversion", err)
			}
			fmt.Println(string(json))

			fmt.Println("Pretty version of json...")
			json, err = eventDetail.ToJsonIndent()
			if err != nil {
				fmt.Println("err json conversion", err)
			}
			fmt.Println(string(json))
		}
	}
}

func help() {
	fmt.Println("There must be two arguments with run command as given below")
	fmt.Println("go run cmd/grpc_client/eventDetailClient.go get 49")
	fmt.Println("or")
	fmt.Println(`go run cmd/grpc_client/eventDetailClient.go event.json`)
}
