package sap_api_input_reader

import (
	"sap-api-integrations-process-order-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.ProcessOrder
	return &requests.General{
		ManufacturingOrder:            data.ManufacturingOrder,
		ManufacturingOrderCategory:    data.ManufacturingOrderCategory,
		ManufacturingOrderType:        data.ManufacturingOrderType,
		ManufacturingOrderImportance:  data.ManufacturingOrderImportance,
		OrderIsCreated:                data.OrderIsCreated,
		OrderIsReleased:               data.OrderIsReleased,
		OrderIsPrinted:                data.OrderIsPrinted,
		OrderIsConfirmed:              data.OrderIsConfirmed,
		OrderIsPartiallyConfirmed:     data.OrderIsPartiallyConfirmed,
		OrderIsDelivered:              data.OrderIsDelivered,
		OrderIsDeleted:                data.OrderIsDeleted,
		OrderIsPreCosted:              data.OrderIsPreCosted,
		SettlementRuleIsCreated:       data.SettlementRuleIsCreated,
		OrderIsPartiallyReleased:      data.OrderIsPartiallyReleased,
		OrderIsLocked:                 data.OrderIsLocked,
		OrderIsTechnicallyCompleted:   data.OrderIsTechnicallyCompleted,
		OrderIsClosed:                 data.OrderIsClosed,
		OrderIsPartiallyDelivered:     data.OrderIsPartiallyDelivered,
		OrderIsMarkedForDeletion:      data.OrderIsMarkedForDeletion,
		SettlementRuleIsCrtedManually: data.SettlementRuleIsCrtedManually,
		OrderIsScheduled:              data.OrderIsScheduled,
		OrderHasGeneratedOperations:   data.OrderHasGeneratedOperations,
		OrderIsToBeHandledInBatches:   data.OrderIsToBeHandledInBatches,
		MaterialAvailyIsNotChecked:    data.MaterialAvailyIsNotChecked,
		MfgOrderCreationDate:          data.MfgOrderCreationDate,
		MfgOrderCreationTime:          data.MfgOrderCreationTime,
		LastChangeDateTime:            data.LastChangeDateTime,
		Material:                      data.Material,
		StorageLocation:               data.StorageLocation,
		GoodsRecipientName:            data.GoodsRecipientName,
		UnloadingPointName:            data.UnloadingPointName,
		InventoryUsabilityCode:        data.InventoryUsabilityCode,
		MaterialGoodsReceiptDuration:  data.MaterialGoodsReceiptDuration,
		QuantityDistributionKey:       data.QuantityDistributionKey,
		StockSegment:                  data.StockSegment,
		OrderInternalBillOfOperations: data.OrderInternalBillOfOperations,
		ProductionPlant:               data.ProductionPlant,
		Plant:                         data.Plant,
		MRPArea:                       data.MRPArea,
		MRPController:                 data.MRPController,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductionVersion:             data.ProductionVersion,
		PlannedOrder:                  data.PlannedOrder,
		SalesOrder:                    data.SalesOrder,
		SalesOrderItem:                data.SalesOrderItem,
		BasicSchedulingType:           data.BasicSchedulingType,
		ManufacturingObject:           data.ManufacturingObject,
		ProductConfiguration:          data.ProductConfiguration,
		OrderSequenceNumber:           data.OrderSequenceNumber,
		BusinessArea:                  data.BusinessArea,
		CompanyCode:                   data.CompanyCode,
		ProfitCenter:                  data.ProfitCenter,
		ActualCostsCostingVariant:     data.ActualCostsCostingVariant,
		PlannedCostsCostingVariant:    data.PlannedCostsCostingVariant,
		FunctionalArea:                data.FunctionalArea,
		MfgOrderPlannedStartDate:      data.MfgOrderPlannedStartDate,
		MfgOrderPlannedStartTime:      data.MfgOrderPlannedStartTime,
		MfgOrderPlannedEndDate:        data.MfgOrderPlannedEndDate,
		MfgOrderPlannedEndTime:        data.MfgOrderPlannedEndTime,
		MfgOrderScheduledStartDate:    data.MfgOrderScheduledStartDate,
		MfgOrderScheduledStartTime:    data.MfgOrderScheduledStartTime,
		MfgOrderScheduledEndDate:      data.MfgOrderScheduledEndDate,
		MfgOrderScheduledEndTime:      data.MfgOrderScheduledEndTime,
		MfgOrderActualReleaseDate:     data.MfgOrderActualReleaseDate,
		ProductionUnit:                data.ProductionUnit,
		TotalQuantity:                 data.TotalQuantity,
		MfgOrderPlannedScrapQty:       data.MfgOrderPlannedScrapQty,
		MfgOrderConfirmedYieldQty:     data.MfgOrderConfirmedYieldQty,
		CustomerName:                  data.CustomerName,
		WBSElementExternalID:          data.WBSElementExternalID,
		OrderLongText:                 data.OrderLongText,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataProcessOrder := sdc.ProcessOrder
	data := sdc.ProcessOrder.ProcessOrderItem
	return &requests.Item{
		ManufacturingOrder:             dataProcessOrder.ManufacturingOrder,
		ManufacturingOrderItem:         data.ManufacturingOrderItem,
		IsCompletelyDelivered:          data.IsCompletelyDelivered,
		Material:                       data.Material,
		ProductionPlant:                data.ProductionPlant,
		Plant:                          data.Plant,
		MRPArea:                        data.MRPArea,
		QuantityDistributionKey:        data.QuantityDistributionKey,
		MaterialGoodsReceiptDuration:   data.MaterialGoodsReceiptDuration,
		StorageLocation:                data.StorageLocation,
		Batch:                          data.Batch,
		InventoryUsabilityCode:         data.InventoryUsabilityCode,
		GoodsRecipientName:             data.GoodsRecipientName,
		UnloadingPointName:             data.UnloadingPointName,
		StockSegment:                   data.StockSegment,
		MfgOrderItemPlndDeliveryDate:   data.MfgOrderItemPlndDeliveryDate,
		MfgOrderItemActualDeliveryDate: data.MfgOrderItemActualDeliveryDate,
		ProductionUnit:                 data.ProductionUnit,
		MfgOrderItemPlannedTotalQty:    data.MfgOrderItemPlannedTotalQty,
		MfgOrderItemPlannedScrapQty:    data.MfgOrderItemPlannedScrapQty,
		MfgOrderItemGoodsReceiptQty:    data.MfgOrderItemGoodsReceiptQty,
		MfgOrderItemActualDeviationQty: data.MfgOrderItemActualDeviationQty,
	}
}

func (sdc *SDC) ConvertToOperation() *requests.Operation {
	dataProcessOrder := sdc.ProcessOrder
	data := sdc.ProcessOrder.Operation
	return &requests.Operation{
		OrderInternalBillOfOperations:  data.OrderInternalBillOfOperations,
		OrderIntBillOfOperationsItem:   data.OrderIntBillOfOperationsItem,
		ManufacturingOrder:             dataProcessOrder.ManufacturingOrder,
		ManufacturingOrderSequence:     data.ManufacturingOrderSequence,
		MfgOrderSequenceText:           data.MfgOrderSequenceText,
		ManufacturingOrderOperation:    data.ManufacturingOrderOperation,
		ManufacturingOrderSubOperation: data.ManufacturingOrderSubOperation,
		MfgOrderOperationText:          data.MfgOrderOperationText,
		OperationIsCreated:             data.OperationIsCreated,
		OperationIsReleased:            data.OperationIsReleased,
		OperationIsPrinted:             data.OperationIsPrinted,
		OperationIsConfirmed:           data.OperationIsConfirmed,
		OperationIsPartiallyConfirmed:  data.OperationIsPartiallyConfirmed,
		OperationIsDeleted:             data.OperationIsDeleted,
		OperationIsTechlyCompleted:     data.OperationIsTechlyCompleted,
		OperationIsClosed:              data.OperationIsClosed,
		OperationIsScheduled:           data.OperationIsScheduled,
		OperationIsPartiallyDelivered:  data.OperationIsPartiallyDelivered,
		OperationIsDelivered:           data.OperationIsDelivered,
		ProductionPlant:                data.ProductionPlant,
		WorkCenterInternalID:           data.WorkCenterInternalID,
		WorkCenterTypeCode:             data.WorkCenterTypeCode,
		WorkCenter:                     data.WorkCenter,
		OperationControlProfile:        data.OperationControlProfile,
		OpErlstSchedldExecStrtDte:      data.OpErlstSchedldExecStrtDte,
		OpErlstSchedldExecStrtTme:      data.OpErlstSchedldExecStrtTme,
		OpErlstSchedldExecEndDte:       data.OpErlstSchedldExecEndDte,
		OpErlstSchedldExecEndTme:       data.OpErlstSchedldExecEndTme,
		OpActualExecutionStartDate:     data.OpActualExecutionStartDate,
		OpActualExecutionStartTime:     data.OpActualExecutionStartTime,
		OpActualExecutionEndDate:       data.OpActualExecutionEndDate,
		OpActualExecutionEndTime:       data.OpActualExecutionEndTime,
		ErlstSchedldExecDurnInWorkdays: data.ErlstSchedldExecDurnInWorkdays,
		OpActualExecutionDays:          data.OpActualExecutionDays,
		OperationUnit:                  data.OperationUnit,
		OpPlannedTotalQuantity:         data.OpPlannedTotalQuantity,
		OpTotalConfirmedYieldQty:       data.OpTotalConfirmedYieldQty,
		LastChangeDateTime:             data.LastChangeDateTime,
	}
}

func (sdc *SDC) ConvertToComponent() *requests.Component {
	dataProcessOrder := sdc.ProcessOrder
	data := sdc.ProcessOrder.Component
	return &requests.Component{
		Reservation:                    data.Reservation,
		ReservationItem:                data.ReservationItem,
		MaterialGroup:                  data.MaterialGroup,
		Material:                       data.Material,
		Plant:                          data.Plant,
		ManufacturingOrder:             dataProcessOrder.ManufacturingOrder,
		ManufacturingOrderSequence:     data.ManufacturingOrderSequence,
		ManufacturingOrderOperation:    data.ManufacturingOrderOperation,
		ProductionPlant:                data.ProductionPlant,
		OrderInternalBillOfOperations:  data.OrderInternalBillOfOperations,
		MatlCompRequirementDate:        data.MatlCompRequirementDate,
		MatlCompRequirementTime:        data.MatlCompRequirementTime,
		ReservationIsFinallyIssued:     data.ReservationIsFinallyIssued,
		IsBulkMaterialComponent:        data.IsBulkMaterialComponent,
		MatlCompIsMarkedForBackflush:   data.MatlCompIsMarkedForBackflush,
		MaterialCompIsCostRelevant:     data.MaterialCompIsCostRelevant,
		SalesOrder:                     data.SalesOrder,
		SalesOrderItem:                 data.SalesOrderItem,
		SortField:                      data.SortField,
		BillOfMaterialCategory:         data.BillOfMaterialCategory,
		BOMItem:                        data.BOMItem,
		BOMItemCategory:                data.BOMItemCategory,
		BillOfMaterialItemNumber:       data.BillOfMaterialItemNumber,
		BOMItemDescription:             data.BOMItemDescription,
		StorageLocation:                data.StorageLocation,
		Batch:                          data.Batch,
		BatchSplitType:                 data.BatchSplitType,
		GoodsMovementType:              data.GoodsMovementType,
		SupplyArea:                     data.SupplyArea,
		GoodsRecipientName:             data.GoodsRecipientName,
		UnloadingPointName:             data.UnloadingPointName,
		MaterialCompIsAlternativeItem:  data.MaterialCompIsAlternativeItem,
		AlternativeItemGroup:           data.AlternativeItemGroup,
		AlternativeItemStrategy:        data.AlternativeItemStrategy,
		AlternativeItemPriority:        data.AlternativeItemPriority,
		UsageProbabilityPercent:        data.UsageProbabilityPercent,
		MaterialComponentIsPhantomItem: data.MaterialComponentIsPhantomItem,
		LeadTimeOffset:                 data.LeadTimeOffset,
		QuantityIsFixed:                data.QuantityIsFixed,
		IsNetScrap:                     data.IsNetScrap,
		ComponentScrapInPercent:        data.ComponentScrapInPercent,
		OperationScrapInPercent:        data.OperationScrapInPercent,
		BaseUnit:                       data.BaseUnit,
		RequiredQuantity:               data.RequiredQuantity,
		WithdrawnQuantity:              data.WithdrawnQuantity,
		ConfirmedAvailableQuantity:     data.ConfirmedAvailableQuantity,
		MaterialCompOriginalQuantity:   data.MaterialCompOriginalQuantity,
		Currency:                       data.Currency,
		WithdrawnQuantityAmount:        data.WithdrawnQuantityAmount,
	}
}

func (sdc *SDC) ConvertToStatus() *requests.Status {
	dataProcessOrder := sdc.ProcessOrder
	data := sdc.ProcessOrder.Status
	return &requests.Status{
		ManufacturingOrder: dataProcessOrder.ManufacturingOrder,
		StatusCode:         data.StatusCode,
		IsUserStatus:       data.IsUserStatus,
		StatusShortName:    data.StatusShortName,
		StatusName:         data.StatusName,
	}
}