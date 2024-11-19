package grpcupfc

import (
	"context"
	"encoding/binary"
	"log"
	"net"
	"time"

	pb "github.com/free5gc/go-upf/internal/grpcupfc/upfc"
	"github.com/free5gc/go-upf/internal/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Client pb.UPFAppClient
	conn   *grpc.ClientConn
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := pb.NewUPFAppClient(conn)

	logger.Log.Infof("Connected to UPF gRPC server at %s", addr)

	return &Client{
		Client: c,
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

func ipv4ToUint32(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func (c *Client) SayHello(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}

func (c *Client) AddPdr(teid, ueIp, pdrId, farId, qerId, precedence uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.AddUplinkPdr(ctx, &pb.UplinkPdrRequest{
		Teid:       teid,
		UeIp:       ueIp,
		PdrId:      pdrId,
		FarId:      farId,
		QerId:      qerId,
		Precedence: precedence})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}

func (c *Client) SetBuffer(action bool, farId uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.SetBuffer(ctx, &pb.BufferRequest{
		Action: action,
		FarId:  farId,
	})
	if err != nil {
		log.Fatalf("could not set buffer: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}

func (c *Client) AddFar(teid uint32, action uint32, farId uint32, tunnelSrcAddr uint32, tunnelDstAddr uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.AddUplinkFar(ctx, &pb.FarRequest{
		Teid:          teid,
		Action:        action,
		FarId:         farId,
		TunnelSrcAddr: tunnelSrcAddr,
		TunnelDstAddr: tunnelDstAddr,
	})
	if err != nil {
		log.Fatalf("could not add uplink FAR: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}

func (c *Client) AddDownlinkPdr(seid uint64, ueIp uint32, pdrId uint32, farId uint32, qerId uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.AddDownlinkPdr(ctx, &pb.DownlinkPdrRequest{
		Seid:  seid,
		UeIp:  ueIp,
		PdrId: pdrId,
		FarId: farId,
		QerId: qerId,
	})
	if err != nil {
		log.Fatalf("could not add downlink PDR: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}

func (c *Client) AddDownlinkFar(teid uint32, action uint32, farId uint32, tunnelSrcAddr uint32, tunnelDstAddr uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.AddDownlinkFar(ctx, &pb.FarRequest{
		Teid:          teid,
		Action:        action,
		FarId:         farId,
		TunnelSrcAddr: tunnelSrcAddr,
		TunnelDstAddr: tunnelDstAddr,
	})
	if err != nil {
		log.Fatalf("could not add downlink FAR: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}

func (c *Client) AddRoute(isUplink bool, srcMac []byte, dstMac []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.AddRoute(ctx, &pb.RouteRequest{
		IsUplink: isUplink,
		SrcMac:   srcMac,
		DstMac:   dstMac,
	})
	if err != nil {
		log.Fatalf("could not add route: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}

func (c *Client) AddQer(qerId uint32, qfi uint32, ulgbr, ulmbr, dlgbr, dlmbr uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.AddQer(ctx, &pb.QerRequest{
		QerId: qerId,
		Qfi:   qfi,
		Ulgbr: ulgbr,
		Ulmbr: ulmbr,
		Dlgbr: dlgbr,
		Dlmbr: dlmbr,
	})
	if err != nil {
		log.Fatalf("could not add QER: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}

func (c *Client) AddSDFTemplate(srcIp uint32, dstIp uint32, srcPort uint32, dstPort uint32, proto uint32, qerId uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.Client.AddSDFTemplate(ctx, &pb.SDFTemplateRequest{
		SrcIp:   srcIp,
		DstIp:   dstIp,
		SrcPort: srcPort,
		DstPort: dstPort,
		Proto:   proto,
		QerId:   qerId,
	})
	if err != nil {
		log.Fatalf("could not add SDF template: %v", err)
	}
	log.Printf("Received: %s", r.GetMessage())
}
