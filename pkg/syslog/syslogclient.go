package syslog

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"goutils/pkg/logchain"
)

type Client struct {
	ConnectInfo string
	Filter      Filter

	connection *grpc.ClientConn
}

func NewClient() *Client {
	client := Client{
		ConnectInfo: "127.0.0.1:50600",
	}
	client.Filter.Clear()
	return &client
}

func (cl *Client) Start() {
	if cl.connection == nil {
		cl.connection, _ = grpc.Dial(cl.ConnectInfo, grpc.WithInsecure())
	}
}

func (cl *Client) Stop() {
	if cl.connection != nil {
		_ = cl.connection.Close()
		cl.connection = nil
	}
}

func (cl *Client) IsConnected() bool {
	return cl.connection != nil
}

func (cl *Client) TryConnect() (bool, error) {
	conn, err := grpc.Dial(cl.ConnectInfo, grpc.WithInsecure())
	if err != nil {
		return false, err
	}
	client := NewSyslogServiceClient(conn)
	_, errGet := client.GetLastEvent(context.Background(), &emptypb.Empty{})
	if errGet != nil {
		_ = conn.Close()
		return false, errGet
	}
	_ = conn.Close()
	return true, nil
}

func (cl *Client) GetLastEvent() (*logchain.SyslogMessage, error) {
	cl.Start()

	if cl == nil || cl.connection == nil {
		return nil, errors.New("not connected to the gRPC server")
	}
	client := NewSyslogServiceClient(cl.connection)
	event, errGet := client.GetLastEvent(context.Background(), &emptypb.Empty{})
	if errGet != nil {
		cl.Stop()
		return nil, errGet
	}

	sm := logchain.NewSyslogMessage()
	if event != nil {
		sm.Index = event.Identity
		sm.Severity = logchain.Severity(event.Severity)
		sm.Timestamp = event.Timestamp.AsTime()
		sm.Message = event.Text
	}

	return sm, nil
}

func (cl *Client) GetCount() (uint64, error) {
	cl.Start()

	if cl == nil || cl.connection == nil {
		return 0, errors.New("not connected to the gRPC server")
	}
	client := NewSyslogServiceClient(cl.connection)
	filter := SyslogFilter{
		Level:      nil,
		Facility:   nil,
		TextFilter: "",
		FromTime:   nil,
		ToTime:     nil,
		Data:       nil,
		FromId:     nil,
		Offset:     0,
		Count:      nil,
	}

	count, errGet := client.GetCount(context.Background(), &filter)
	if errGet != nil {
		cl.Stop()
		return 0, errGet
	}
	var nofMessages uint64
	if count != nil {
		nofMessages = count.Count
	}

	return nofMessages, nil
}
