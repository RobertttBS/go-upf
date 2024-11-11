package forwarder

import (
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
	return g, nil
}

func (g *Grpc) Close() {
	if g.grpcClient != nil {
		g.grpcClient.Close()
	}
}

func (g *Grpc) CreatePDR(lSeid uint64, req *ie.IE) error {

	g.grpcClient.AddUplinkPdr(34234, "", 2342, 234, 234)
	return nil
}

func (d *Grpc) UpdatePDR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) RemovePDR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) CreateFAR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) UpdateFAR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) RemoveFAR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) CreateQER(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) UpdateQER(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) RemoveQER(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) CreateURR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) UpdateURR(id uint64, ie *ie.IE) ([]report.USAReport, error) {
	return nil, nil
}

func (d *Grpc) RemoveURR(id uint64, ie *ie.IE) ([]report.USAReport, error) {
	return nil, nil
}

func (d *Grpc) QueryURR(id uint64, query uint32) ([]report.USAReport, error) {
	return nil, nil
}

func (d *Grpc) CreateBAR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) UpdateBAR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) RemoveBAR(id uint64, ie *ie.IE) error {
	return nil
}

func (d *Grpc) HandleReport(handler report.Handler) {}
