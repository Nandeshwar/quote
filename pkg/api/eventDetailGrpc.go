package api

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"quote/pkg/model"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	grpc "quote/pkg/grpcquote"
)

func (s *Server) GetEventDetail(ctx context.Context, in *grpc.EventDetailRequest) (*grpc.EventDetailReply, error) {
	eventDetailDB, err := s.eventDetailService.GetEventDetailByID(in.Id)
	if err != nil {
		return nil, err
	}
	eventDate := time.Date(eventDetailDB.Year, time.Month(eventDetailDB.Month), eventDetailDB.Day, 0, 0, 0, 0, time.Local)

	eventDetail := &grpc.EventDetailReply{
		EventDetail: &grpc.EventDetail{
			Id:        eventDetailDB.ID,
			Title:     eventDetailDB.Title,
			Info:      eventDetailDB.Info,
			EventType: eventDetailDB.Type,
			EventDate: &timestamp.Timestamp{
				Seconds: eventDate.Unix(),
			},

			UpdatedAt: &timestamp.Timestamp{
				Seconds: eventDetailDB.UpdatedAt.UTC().Unix(),
			},

			CreatedAt: &timestamp.Timestamp{
				Seconds: eventDetailDB.CreationDate.UTC().Unix(),
			},
			Links: eventDetailDB.Links,
		},
	}
	return eventDetail, nil
}

func (s *Server) UpdateEventDetail(ctx context.Context, in *grpc.EventDetailUpdateRequest) (*grpc.EventDetailUpdateReply, error) {
	eventDetail := model.EventDetail{
		ID:    in.EventDetail.Id,
		Day:   int(in.Dd),
		Month: int(in.Mm),
		Year:  int(in.Yyyy),
		Title: in.EventDetail.Title,
		Info:  in.EventDetail.Info,
		Type:  in.EventDetail.EventType,
		Links: in.EventDetail.Links,
	}

	IDs, err := s.eventDetailService.GetEventDetailLinkIDs(strings.Join(eventDetail.Links, ","))
	if err != nil {
		logrus.Errorf("error checking existence of links=%v", err)
		return nil, err
	}

	err = s.eventDetailService.UpdateEventDetailByID(eventDetail)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("id=%v created successfully.", eventDetail.ID)
	if len(IDs) > 0 {
		msg += fmt.Sprintf("Link already exists for IDs=%v", IDs)
	}

	eventDetailReply := &grpc.EventDetailUpdateReply{Id: eventDetail.ID, Msg: msg}
	return eventDetailReply, nil
}

func (s *Server) GetEventDetailStream(stream grpc.EventDetailServiceGRPC_GetEventDetailStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		id := in.Id
		eventDetailDB, err := s.eventDetailService.GetEventDetailByID(id)
		if err != nil {
			return err
		}
		eventDate := time.Date(eventDetailDB.Year, time.Month(eventDetailDB.Month), eventDetailDB.Day, 0, 0, 0, 0, time.Local)

		eventDetail := &grpc.EventDetailReply{
			EventDetail: &grpc.EventDetail{
				Id:        eventDetailDB.ID,
				Title:     eventDetailDB.Title,
				Info:      eventDetailDB.Info,
				EventType: eventDetailDB.Type,
				EventDate: &timestamp.Timestamp{
					Seconds: eventDate.Unix(),
				},

				UpdatedAt: &timestamp.Timestamp{
					Seconds: eventDetailDB.UpdatedAt.UTC().Unix(),
				},

				CreatedAt: &timestamp.Timestamp{
					Seconds: eventDetailDB.CreationDate.UTC().Unix(),
				},
				Links: eventDetailDB.Links,
			},
		}

		if err := stream.Send(eventDetail); err != nil {
			return err
		}
	}

	return nil
}
