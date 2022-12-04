package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-tax-code-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-tax-code-creates-rmq-kube/config"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:  context.Background(),
		conf: conf,
		rmq:  rmq,
	}
}

func (c *DPFMAPICaller) AsyncTaxCodeCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,

	log *logger.Logger,
	// msg rabbitmq.RabbitmqMessage,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	sqlUpdateFin := make(chan error)

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "TaxCode":
			go c.TaxCode(&wg, &mtx, sqlUpdateFin, log, &errs, input)
		case "TaxCodeText":
			go c.TaxCodeText(&wg, &mtx, sqlUpdateFin, log, &errs, input)
		case "TaxRate":
			go c.TaxRate(&wg, &mtx, sqlUpdateFin, log, &errs, input)
		default:
			wg.Done()
		}
	}

	// 後処理
	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-sqlUpdateFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.New("time out"))
		return errs
	}

	return nil
}

func (c *DPFMAPICaller) TaxCode(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs *[]error, sdc *dpfm_api_input_reader.SDC) {
	var err error = nil
	defer wg.Done()
	defer func() {
		errFin <- err
	}()
	sessionID := sdc.RuntimeSessionID
	ctx := context.Background()

	// data_platform_tax_code_tax_code_dataの更新
	taxCodeData := sdc.TaxCode
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": taxCodeData, "function": "TaxCodeTaxCode", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Tax Code Data cannot insert")
		sdc.SQLUpdateResult = getBoolPtr(false)
		sdc.SQLUpdateError = "Tax Code Data cannot insert"
		return
	}

	sdc.SQLUpdateResult = getBoolPtr(true)
	return
}

func (c *DPFMAPICaller) TaxCodeText(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs *[]error, sdc *dpfm_api_input_reader.SDC) {
	var err error = nil
	defer wg.Done()
	defer func() {
		errFin <- err
	}()
	sessionID := sdc.RuntimeSessionID
	ctx := context.Background()

	// data_platform_tax_code_tax_code_text_dataの更新
	taxCodeTextData := sdc.TaxCode.TaxCodeText
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": taxCodeTextData, "function": "TaxCodeTaxCodeText", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Tax Code Text Data cannot insert")
		sdc.SQLUpdateResult = getBoolPtr(false)
		sdc.SQLUpdateError = "Tax Code Text Data cannot insert"
		return
	}

	sdc.SQLUpdateResult = getBoolPtr(true)
	return
}

func (c *DPFMAPICaller) TaxRate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs *[]error, sdc *dpfm_api_input_reader.SDC) {
	var err error = nil
	defer wg.Done()
	defer func() {
		errFin <- err
	}()
	sessionID := sdc.RuntimeSessionID
	ctx := context.Background()

	// data_platform_tax_code_tax_rate_dataの更新
	taxRateData := sdc.TaxCode.TaxRate
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": taxRateData, "function": "TaxCodeTaxRate", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Tax Rate Data cannot insert")
		sdc.SQLUpdateResult = getBoolPtr(false)
		sdc.SQLUpdateError = "Tax Rate Data cannot insert"
		return
	}

	sdc.SQLUpdateResult = getBoolPtr(true)
	return
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	_, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := data["result"].(string)
	if !ok {
		return false
	}
	return result == "success"

}

func getBoolPtr(b bool) *bool {
	return &b
}
