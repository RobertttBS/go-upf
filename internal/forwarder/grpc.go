package forwarder

import (
	"encoding/binary"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/wmnsk/go-pfcp/ie"

	"github.com/free5gc/go-upf/internal/grpcupfc"
	"github.com/free5gc/go-upf/internal/logger"
	"github.com/free5gc/go-upf/internal/report"
	logger_util "github.com/free5gc/util/logger"
)

type Grpc struct {
	grpcClient *grpcupfc.Client
	log        *logrus.Entry
}

func OpenGrpc() (*Grpc, error) {
	g := &Grpc{
		log: logger.FwderLog.WithField(logger_util.FieldCategory, "Grpc"),
	}

	grpcClient, err := grpcupfc.NewClient("localhost:50051") // defaul grpc server address
	if err != nil {
		g.Close()
		return nil, errors.Wrap(err, "new grpc client")
	}
	g.grpcClient = grpcClient

	g.log.Infof("Forwarder started")
	g.grpcClient.SayHello("go-upf")
	return g, nil
}

func (g *Grpc) Close() {
	if g.grpcClient != nil {
		g.grpcClient.Close()
	}
}

func (g *Grpc) newPdi(i *ie.IE) (uint32, uint32, error) {
	var teid uint32
	var ueIp uint32

	ies, err := i.PDI()
	if err != nil {
		return 0, 0, err
	}

	for _, x := range ies {
		switch x.Type {
		case ie.FTEID:
			v, err := x.FTEID()
			if err != nil {
				break
			}
			teid = v.TEID
		case ie.UEIPAddress:
			v, err := x.UEIPAddress()
			if err != nil {
				break
			}
			ueIp = binary.BigEndian.Uint32(v.IPv4Address)
		}
	}

	return teid, ueIp, nil
}

func (g *Grpc) CreatePDR(lSeid uint64, req *ie.IE) error {
	var teid uint32
	var ueIp uint32
	var pdrId uint32
	var farId uint32
	var qerId uint32
	var precedence uint32

	ies, err := req.CreatePDR()
	if err != nil {
		return err
	}
	for _, i := range ies {
		switch i.Type {
		case ie.PDRID:
			v, err := i.PDRID()
			if err != nil {
				break
			}
			pdrId = uint32(v)
		case ie.PDI:
			teid, ueIp, err = g.newPdi(i)
			if err != nil {
				g.log.Error("Failed to get PDI")
			}
		case ie.Precedence:
			precedence, err = i.Precedence()
			if err != nil {
				g.log.Errorf("Failed to get precedence: %+v", err)
			}
		case ie.FARID:
			v, err := i.FARID()
			if err != nil {
				break
			}
			farId = uint32(v)
		case ie.QERID:
			if qerId != 0 {
				continue
			}
			v, err := i.QERID()
			if err != nil {
				break
			}
			qerId = v
		}
	}

	g.grpcClient.AddPdr(teid, ueIp, pdrId, farId, qerId, precedence)
	return nil
}

func (g *Grpc) UpdatePDR(lSeid uint64, req *ie.IE) error {
	var teid uint32
	var ueIp uint32
	var pdrId uint32
	var farId uint32
	var qerId uint32
	var precedence uint32

	ies, err := req.UpdatePDR()
	if err != nil {
		return err
	}
	for _, i := range ies {
		switch i.Type {
		case ie.PDRID:
			v, err := i.PDRID()
			if err != nil {
				break
			}
			pdrId = uint32(v)
		case ie.PDI:
			teid, ueIp, err = g.newPdi(i)
			if err != nil {
				g.log.Error("Failed to get PDI")
			}
		case ie.Precedence:
			precedence, err = i.Precedence()
			if err != nil {
				g.log.Errorf("Failed to get precedence: %+v", err)
			}
		case ie.FARID:
			v, err := i.FARID()
			if err != nil {
				break
			}
			farId = uint32(v)
		case ie.QERID:
			v, err := i.QERID()
			if err != nil {
				break
			}
			qerId = v
		}
	}

	g.grpcClient.AddPdr(teid, ueIp, pdrId, farId, qerId, precedence)
	return nil
}

func (g *Grpc) RemovePDR(lSeid uint64, req *ie.IE) error {
	return nil
}

func (g *Grpc) newForwardingParameter(ies []*ie.IE) (uint32, uint32, error) {
	var teid uint32
	var ipv4 uint32

	for _, x := range ies {
		switch x.Type {
		case ie.DestinationInterface:
		case ie.NetworkInstance:
		case ie.OuterHeaderCreation:
			v, err := x.OuterHeaderCreation()
			if err != nil {
				break
			}

			if x.HasTEID() {
				teid = v.TEID
			}
			if x.HasIPv4() {
				ip := v.IPv4Address
				ipv4 = uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
			}
		}
	}
	return teid, ipv4, nil
}

func (g *Grpc) CreateFAR(lSeid uint64, req *ie.IE) error {
	var teid uint32
	var action uint32
	var farId uint32
	var tunnelSrcIp uint32 // shoulde be UPF ip
	var tunnelDstIp uint32

	ies, err := req.CreateFAR()
	if err != nil {
		return err
	}
	for _, i := range ies {
		switch i.Type {
		case ie.FARID:
			v, err := i.FARID()
			if err != nil {
				return err
			}
			farId = uint32(v)
		case ie.ApplyAction:
			b, err := i.ApplyAction()
			if err != nil {
				return err
			}
			var act report.ApplyAction
			err = act.Unmarshal(b)
			if err != nil {
				return err
			}
			action = uint32(act.Flags)
		case ie.ForwardingParameters:
			xs, err := i.ForwardingParameters()
			if err != nil {
				return err
			}
			teid, tunnelDstIp, err = g.newForwardingParameter(xs)
			if err != nil {
				return err
			}
		}
	}
	g.grpcClient.AddFar(teid, action, farId, tunnelSrcIp, tunnelDstIp)
	return nil
}

func (g *Grpc) UpdateFAR(lSeid uint64, req *ie.IE) error {
	var teid uint32
	var action uint32
	var farId uint32
	var tunnelSrcIp uint32 // shoulde be UPF ip
	var tunnelDstIp uint32

	ies, err := req.UpdateFAR()
	if err != nil {
		return err
	}
	for _, i := range ies {
		switch i.Type {
		case ie.FARID:
			v, err := i.FARID()
			if err != nil {
				return err
			}
			farId = uint32(v)
		case ie.ApplyAction:
			b, err := i.ApplyAction()
			if err != nil {
				return err
			}
			var act report.ApplyAction
			err = act.Unmarshal(b)
			if err != nil {
				return err
			}
			action = uint32(act.Flags)
		case ie.UpdateForwardingParameters:
			xs, err := i.UpdateForwardingParameters()
			if err != nil {
				return err
			}
			teid, tunnelDstIp, err = g.newForwardingParameter(xs)
			if err != nil {
				return err
			}
		}
	}
	g.grpcClient.AddFar(teid, action, farId, tunnelSrcIp, tunnelDstIp)
	return nil
}

func (g *Grpc) RemoveFAR(lSeid uint64, req *ie.IE) error {
	return nil
}

func (g *Grpc) CreateQER(lSeid uint64, req *ie.IE) error {
	var qerId uint32
	var qfi uint32
	var ulgbr uint64
	var ulmbr uint64
	var dlgbr uint64
	var dlmbr uint64

	ies, err := req.CreateQER()
	if err != nil {
		return err
	}
	for _, i := range ies {
		switch i.Type {
		case ie.QERID:
			v, err := i.QERID()
			if err != nil {
				break
			}
			qerId = uint32(v)
		case ie.QFI:
			v, err := i.QFI()
			if err != nil {
				break
			}
			qfi = uint32(v)
		case ie.GBR:
			ulgbr, err = i.GBRUL()
			if err != nil {
				break
			}
			dlgbr, err = i.GBRDL()
			if err != nil {
				break
			}
		case ie.MBR:
			ulmbr, err = i.MBRUL()
			if err != nil {
				break
			}
			dlmbr, err = i.MBRDL()
			if err != nil {
				break
			}
		}
	}

	g.grpcClient.AddQer(qerId, qfi, ulgbr, ulmbr, dlgbr, dlmbr)
	return nil
}

func (g *Grpc) UpdateQER(lSeid uint64, req *ie.IE) error {
	var qerId uint32
	var qfi uint32
	var cir uint64
	var cbs uint64
	var pir uint64
	var pbs uint64

	ies, err := req.UpdateQER()
	if err != nil {
		return err
	}
	for _, i := range ies {
		switch i.Type {
		case ie.QERID:
			v, err := i.QERID()
			if err != nil {
				break
			}
			qerId = uint32(v)
		case ie.QFI:
			v, err := i.QFI()
			if err != nil {
				break
			}
			qfi = uint32(v)
		case ie.GBR:
			gbr, err := i.GBRUL()
			if err != nil {
				break
			}
			cir = gbr
			cbs = cir
		case ie.MBR:
			mbr, err := i.MBRUL()
			if err != nil {
				break
			}
			pir = mbr
			pbs = pir
		}
	}

	g.grpcClient.AddQer(qerId, qfi, cir, cbs, pir, pbs)
	return nil
}

func (g *Grpc) RemoveQER(lSeid uint64, req *ie.IE) error {
	return nil
}

func (g *Grpc) CreateURR(lSeid uint64, req *ie.IE) error {
	return nil
}

func (g *Grpc) UpdateURR(lSeid uint64, req *ie.IE) ([]report.USAReport, error) {
	return nil, nil
}

func (g *Grpc) RemoveURR(lSeid uint64, req *ie.IE) ([]report.USAReport, error) {
	return nil, nil
}

func (g *Grpc) QueryURR(id uint64, query uint32) ([]report.USAReport, error) {
	return nil, nil
}

func (g *Grpc) CreateBAR(lSeid uint64, req *ie.IE) error {
	return nil
}

func (g *Grpc) UpdateBAR(lSeid uint64, req *ie.IE) error {
	return nil
}

func (g *Grpc) RemoveBAR(lSeid uint64, req *ie.IE) error {
	return nil
}

func (g *Grpc) HandleReport(handler report.Handler) {

}
