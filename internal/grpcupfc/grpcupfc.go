package grpcupfc

import (
	"context"
	"time"

	pb "github.com/free5gc/go-upf/internal/grpcupfc/upfc"
	"github.com/free5gc/go-upf/internal/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Client pb.UPFAppClient
	Ctx    context.Context
	conn   *grpc.ClientConn
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		<-ctx.Done()
		cancel()
	}()
	c := pb.NewUPFAppClient(conn)

	logger.Log.Infof("Connected to UPF gRPC server at %s", addr)

	return &Client{
		Client: c,
		Ctx:    ctx,
		conn:   conn,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

// 在這裡添加其他方法，例如 TestSetBuffer, TestAddUplinkPdr 等
func (c *Client) TestSetBuffer() error {
	// 實現 TestSetBuffer 邏輯
	return nil
}

func (c *Client) TestAddUplinkPdr(pdrId uint32, ueIpAddr string, f_teid uint32, qerId uint32, farId uint32) error {
	// 實現 TestAddUplinkPdr 邏輯
	return nil
}
