package sap_api_caller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-process-order-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-process-order-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostProcessOrder(
	general *requests.General,
	item *requests.Item,
	component *requests.Component,
	operation *requests.Operation,
	status *requests.Status,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(general)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(item)
				wg.Done()
			}()
		case "Component":
			func() {
				c.Component(component)
				wg.Done()
			}()
		case "Operation":
			func() {
				c.Operation(operation)
				wg.Done()
			}()
		case "Status":
			func() {
				c.Status(status)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(general *requests.General) {
	outputDataGeneral, err := c.callProcessOrderSrvAPIRequirementGeneral("A_ProcessOrder_2", general)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataGeneral)
}

func (c *SAPAPICaller) callProcessOrderSrvAPIRequirementGeneral(api string, general *requests.General) (*sap_api_output_formatter.General, error) {
	body, err := json.Marshal(general)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PROCESS_ORDER_2_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(item *requests.Item) {
	outputDataItem, err := c.callProcessOrderSrvAPIRequirementItem("A_ProcessOrderItem_2", item)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataItem)
}

func (c *SAPAPICaller) callProcessOrderSrvAPIRequirementItem(api string, item *requests.Item) (*sap_api_output_formatter.Item, error) {
	body, err := json.Marshal(item)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PROCESS_ORDER_2_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Component(component *requests.Component) {
	outputDataComponent, err := c.callProcessOrderSrvAPIRequirementComponent("A_ProcessOrderComponent_2", component)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataComponent)
}

func (c *SAPAPICaller) callProcessOrderSrvAPIRequirementComponent(api string, component *requests.Component) (*sap_api_output_formatter.Component, error) {
	body, err := json.Marshal(component)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PROCESS_ORDER_2_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToComponent(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Operation(operation *requests.Operation) {
	outputDataOperation, err := c.callProcessOrderSrvAPIRequirementOperation("A_ProcessOrderOperation_2", operation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataOperation)
}

func (c *SAPAPICaller) callProcessOrderSrvAPIRequirementOperation(api string, operation *requests.Operation) (*sap_api_output_formatter.Operation, error) {
	body, err := json.Marshal(operation)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PROCESS_ORDER_2_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Status(status *requests.Status) {
	outputDataStatus, err := c.callProcessOrderSrvAPIRequirementStatus("A_ProcessOrderStatus_2", status)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataStatus)
}

func (c *SAPAPICaller) callProcessOrderSrvAPIRequirementStatus(api string, status *requests.Status) (*sap_api_output_formatter.Status, error) {
	body, err := json.Marshal(status)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PROCESS_ORDER_2_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToStatus(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
