package codeupdate

type ScriptOptions struct {
	DecafBase64 string     `json:"strDecaf,omitempty"`
	DeXMLBase64 string     `json:"strDexml,omitempty"`
	Label       string     `json:"strLabel,omitempty"`
	ScriptType  ScriptType `json:"intScriptType"`
	TriggerORM  TriggerORM `json:"strTriggerOrm,omitempty"`
}

type ReportOptions struct {
	Name             string `json:"strName"`
	IsDev            uint   `json:"blnInDev"`
	DateFilter       uint   `json:"intDateFilter"`
	CompanyFilter    uint   `json:"blnCompanyFilter"`
	Category         uint   `json:"intCategory"`
	HTMLBase64       string `json:"strBody"`
	JavascriptBase64 string `json:"strJsCode"`
	DecafBase64      string `json:"strDecaf,omitempty"`
	DeXMLBase64      string `json:"strDexml,omitempty"`
	CSSBase64        string `json:"strCss"`
	WidgetDexml      string `json:"strWidgetDexml,omitempty"`
}

type CustomAppOptions struct {
	Name             string `json:"strName"`
	HTMLBase64       string `json:"strBody"`
	JavascriptBase64 string `json:"strJsCode"`
	DecafBase64      string `json:"strDecaf,omitempty"`
	DeXMLBase64      string `json:"strDexml,omitempty"`
	CSSBase64        string `json:"strCss"`
}

type ScriptType uint
type TriggerORM string

const (
	HelperSupportScript  ScriptType = 1  // 1 - Helper Support Script
	ExecuteOnLogin                  = 2  // 2 - Execute on Login
	ExecuteAtCron                   = 3  // 3 - Execute at Cron
	RecordValidation                = 4  // 4 - Record Validation*
	PeriodPayCalculation            = 6  // 6 - Period Pay Calculation
	RecordAfterSave                 = 7  // 7 - Record After Save*
	ExecuteViaApi                   = 8  // 8 - Executable via SOAP/REST api
	PayrollExport                   = 9  // 9 - Payroll Export Script
	ApplicationLaunch               = 10 // 10 - Application Launch*
	ShiftPayCalculation             = 11 // 11 - Shift Pay Calculation
	KioskLaunch                     = 12 // 12 - Kiosk Launch
	Notification                    = 13 // 13 - Notification
	Widget                          = 14 // 14 - Widget*
	RecordAfterDelete               = 15 // 15 - Record After Delete*
)

const (
	TriggerCategory                     TriggerORM = "DeputecCategory"
	TriggerAddress                                 = "DeputecAddress"
	TriggerCompany                                 = "DeputecCompany"
	TriggerCompanyPeriod                           = "DeputecCompanyPeriod"
	TriggerContact                                 = "DeputecContact"
	TriggerCountry                                 = "DeputecCountry"
	TriggerCustomAppData                           = "DeputecCustomAppData"
	TriggerCustomField                             = "DeputecCustomField"
	TriggerCustomFieldDataComment                  = "DeputecCustomFieldDataComment"
	TriggerEmployee                                = "DeputecEmployee"
	TriggerEmployeeAgreement                       = "DeputecEmployeeAgreement"
	TriggerEmployeeAgreementHistory                = "DeputecEmployeeAgreementHistory"
	TriggerEmployeeAppraisal                       = "DeputecEmployeeAppraisal"
	TriggerEmployeeAvailability                    = "DeputecEmployeeAvailability"
	TriggerEmployeeHistory                         = "DeputecEmployeeHistory"
	TriggerEmployeePaycycle                        = "DeputecEmployeePaycycle"
	TriggerEmployeePaycycleReturn                  = "DeputecEmployeePaycycleReturn"
	TriggerEmployeeRole                            = "DeputecEmployeeRole"
	TriggerEmployeeSalaryOpunitCosting             = "DeputecEmployeeSalaryOpunitCosting"
	TriggerEmployeeWorkplace                       = "DeputecEmployeeWorkplace"
	TriggerEmploymentCondition                     = "DeputecEmploymentCondition"
	TriggerEmploymentContract                      = "DeputecEmploymentContract"
	TriggerEmploymentContractLeaveRules            = "DeputecEmploymentContractLeaveRules"
	TriggerEvent                                   = "DeputecEvent"
	TriggerGeo                                     = "DeputecGeo"
	TriggerJournal                                 = "DeputecJournal"
	TriggerKiosk                                   = "DeputecKiosk"
	TriggerLeave                                   = "DeputecLeave"
	TriggerLeavePayLine                            = "DeputecLeavePayLine"
	TriggerLeaveAccrual                            = "DeputecLeaveAccrual"
	TriggerLeaveRules                              = "DeputecLeaveRules"
	TriggerMemo                                    = "DeputecMemo"
	TriggerOperationalUnit                         = "DeputecOperationalUnit"
	TriggerPayRules                                = "DeputecPayRules"
	TriggerPayPeriod                               = "DeputecPayPeriod"
	TriggerPublicHoliday                           = "DeputecPublicHoliday"
	TriggerRoster                                  = "DeputecRoster"
	TriggerRosterOpen                              = "DeputecRosterOpen"
	TriggerRosterSwap                              = "DeputecRosterSwap"
	TriggerRosterTemplate                          = "DeputecRosterTemplate"
	TriggerSalesData                               = "DeputecSalesData"
	TriggerSchedule                                = "DeputecSchedule"
	TriggerSmsLog                                  = "DeputecSmsLog"
	TriggerState                                   = "DeputecState"
	TriggerStressProfile                           = "DeputecStressProfile"
	TriggerTask                                    = "DeputecTask"
	TriggerTaskGroup                               = "DeputecTaskGroup"
	TriggerTaskGroupSetup                          = "DeputecTaskGroupSetup"
	TriggerTaskOpunitConfig                        = "DeputecTaskOpunitConfig"
	TriggerTaskSetup                               = "DeputecTaskSetup"
	TriggerTimesheet                               = "DeputecTimesheet"
	TriggerTimesheetPayReturn                      = "DeputecTimesheetPayReturn"
	TriggerTeam                                    = "DeputecTeam"
	TriggerTrainingModule                          = "DeputecTrainingModule"
	TriggerTrainingRecord                          = "DeputecTrainingRecord"
	TriggerWebhook                                 = "DeputecWebhook"
	TriggerShiftTemplate                           = "DeputecShiftTemplate"
)
