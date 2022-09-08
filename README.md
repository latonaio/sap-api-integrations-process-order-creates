# sap-api-integrations-process-order-creates  
sap-api-integrations-process-order-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で プロセス指図データ を登録するマイクロサービスです。  
sap-api-integrations-process-order-creates には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-process-order-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PROCESS_ORDER_2_SRV_0001/overview  

## 動作環境  
sap-api-integrations-process-order-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-process-order-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-process-order-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PROCESS_ORDER_2_SRV_0001/overview  
* APIサービス名(=baseURL): API_PROCESS_ORDER_2_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-process-order-creates には、次の API をコールするためのリソースが含まれています。  

* A_General（プロセス指図 - ヘッダ）
* Component（プロセス指図 - 構成品目）
* Item（プロセス指図 - 明細）
* Operation（プロセス指図 - 作業手順）
* Status（プロセス指図 - ステータス）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
	"api_schema": "SAPProcessOrderCreates",
	"accepter": ["General"],
	"process_order": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPProcessOrderCreates",
	"accepter": ["All"],
	"process_order": "",
	"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP プロセス指図 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"XXXXX" ～ "XXXXX" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-process-order-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-process-order-reads/SAP_API_Caller.(*SAPAPICaller).General",
	"level": "INFO",
	"message": [
		{
			"ManufacturingOrder": "1003463",
			"ManufacturingOrderCategory": "40",
			"ManufacturingOrderType": "YBM2",
			"ManufacturingOrderImportance": "",
			"OrderIsCreated": "",
			"OrderIsReleased": "",
			"OrderIsPrinted": "",
			"OrderIsConfirmed": "",
			"OrderIsPartiallyConfirmed": "",
			"OrderIsDelivered": "",
			"OrderIsDeleted": "",
			"OrderIsPreCosted": "X",
			"SettlementRuleIsCreated": "X",
			"OrderIsPartiallyReleased": "",
			"OrderIsLocked": "",
			"OrderIsTechnicallyCompleted": "X",
			"OrderIsClosed": "",
			"OrderIsPartiallyDelivered": "",
			"OrderIsMarkedForDeletion": "",
			"SettlementRuleIsCrtedManually": "",
			"OrderIsScheduled": "",
			"OrderHasGeneratedOperations": "",
			"OrderIsToBeHandledInBatches": "X",
			"MaterialAvailyIsNotChecked": "X",
			"MfgOrderCreationDate": "2017-08-30T09:00:00+09:00",
			"MfgOrderCreationTime": "PT09H05M40S",
			"LastChangeDateTime": "20200723125734",
			"Material": "FG29",
			"StorageLocation": "171A",
			"GoodsRecipientName": "",
			"UnloadingPointName": "",
			"InventoryUsabilityCode": "",
			"MaterialGoodsReceiptDuration": "0",
			"QuantityDistributionKey": "",
			"StockSegment": "",
			"OrderInternalBillOfOperations": "3508",
			"ProductionPlant": "1710",
			"Plant": "1710",
			"MRPArea": "1710",
			"MRPController": "001",
			"ProductionSupervisor": "YB2",
			"ProductionVersion": "0001",
			"PlannedOrder": "491",
			"SalesOrder": "",
			"SalesOrderItem": "0",
			"BasicSchedulingType": "1",
			"ManufacturingObject": "OR000001003463",
			"ProductConfiguration": "0",
			"OrderSequenceNumber": "0",
			"BusinessArea": "",
			"CompanyCode": "1710",
			"ProfitCenter": "YB110",
			"ActualCostsCostingVariant": "PYG2",
			"PlannedCostsCostingVariant": "PYG1",
			"FunctionalArea": "",
			"MfgOrderPlannedStartDate": "2017-08-30T09:00:00+09:00",
			"MfgOrderPlannedStartTime": "PT00H00M00S",
			"MfgOrderPlannedEndDate": "2017-09-06T09:00:00+09:00",
			"MfgOrderPlannedEndTime": "PT00H00M00S",
			"MfgOrderScheduledStartDate": "2017-09-01T09:00:00+09:00",
			"MfgOrderScheduledStartTime": "PT06H00M00S",
			"MfgOrderScheduledEndDate": "2017-09-01T09:00:00+09:00",
			"MfgOrderScheduledEndTime": "PT06H11M38S",
			"MfgOrderActualReleaseDate": "2017-08-30T09:00:00+09:00",
			"ProductionUnit": "BT",
			"TotalQuantity": "969",
			"MfgOrderPlannedScrapQty": "0",
			"MfgOrderConfirmedYieldQty": "0",
			"CustomerName": "",
			"WBSElementExternalID": "",
			"OrderLongText": "",
			"to_ProcessOrderComponent": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PROCESS_ORDER_2_SRV/A_ProcessOrder_2('1003463')/to_ProcessOrderComponent",
			"to_ProcessOrderItem": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PROCESS_ORDER_2_SRV/A_ProcessOrder_2('1003463')/to_ProcessOrderItem",
			"to_ProcessOrderOperation": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PROCESS_ORDER_2_SRV/A_ProcessOrder_2('1003463')/to_ProcessOrderOperation",
			"to_ProcessOrderStatus": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PROCESS_ORDER_2_SRV/A_ProcessOrder_2('1003463')/to_ProcessOrderStatus"
		}
	],
	"time": "2022-01-28T16:17:24+09:00"
```