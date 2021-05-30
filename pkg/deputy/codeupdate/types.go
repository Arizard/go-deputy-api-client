package codeupdate

type ScriptOptions struct {
	DecafBase64 string     `json:"strDecaf,omitempty"`
	DeXMLBase64 string     `json:"strDexml,omitempty"`
	Label       string     `json:"strLabel,omitempty"`
	ScriptType  ScriptType `json:"intScriptType"`
	TriggerORM  TriggerORM `json:"strTriggerOrm,omitempty"`
}

type ReportOptions struct {
	Name             string         `json:"strName"`
	IsDev            DeputyBoolean  `json:"blnInDev"`
	DateFilter       DatePickerMode `json:"intDateFilter"`
	CompanyFilter    DeputyBoolean  `json:"blnCompanyFilter"`
	Category         ReportCategory `json:"intCategory"`
	HTMLBase64       string         `json:"strBody"`
	JavascriptBase64 string         `json:"strJsCode"`
	DecafBase64      string         `json:"strDecaf,omitempty"`
	DeXMLBase64      string         `json:"strDexml,omitempty"`
	CSSBase64        string         `json:"strCss"`
	WidgetDexml      string         `json:"strWidgetDexml,omitempty"`
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
	ExecuteOnLogin       ScriptType = 2  // 2 - Execute on Login
	ExecuteAtCron        ScriptType = 3  // 3 - Execute at Cron
	RecordValidation     ScriptType = 4  // 4 - Record Validation*
	PeriodPayCalculation ScriptType = 6  // 6 - Period Pay Calculation
	RecordAfterSave      ScriptType = 7  // 7 - Record After Save*
	ExecuteViaApi        ScriptType = 8  // 8 - Executable via SOAP/REST api
	PayrollExport        ScriptType = 9  // 9 - Payroll Export Script
	ApplicationLaunch    ScriptType = 10 // 10 - Application Launch*
	ShiftPayCalculation  ScriptType = 11 // 11 - Shift Pay Calculation
	KioskLaunch          ScriptType = 12 // 12 - Kiosk Launch
	Notification         ScriptType = 13 // 13 - Notification
	Widget               ScriptType = 14 // 14 - Widget*
	RecordAfterDelete    ScriptType = 15 // 15 - Record After Delete*
)

const (
	TriggerCategory                     TriggerORM = "DeputecCategory"
	TriggerAddress                      TriggerORM = "DeputecAddress"
	TriggerCompany                      TriggerORM = "DeputecCompany"
	TriggerCompanyPeriod                TriggerORM = "DeputecCompanyPeriod"
	TriggerContact                      TriggerORM = "DeputecContact"
	TriggerCountry                      TriggerORM = "DeputecCountry"
	TriggerCustomAppData                TriggerORM = "DeputecCustomAppData"
	TriggerCustomField                  TriggerORM = "DeputecCustomField"
	TriggerCustomFieldDataComment       TriggerORM = "DeputecCustomFieldDataComment"
	TriggerEmployee                     TriggerORM = "DeputecEmployee"
	TriggerEmployeeAgreement            TriggerORM = "DeputecEmployeeAgreement"
	TriggerEmployeeAgreementHistory     TriggerORM = "DeputecEmployeeAgreementHistory"
	TriggerEmployeeAppraisal            TriggerORM = "DeputecEmployeeAppraisal"
	TriggerEmployeeAvailability         TriggerORM = "DeputecEmployeeAvailability"
	TriggerEmployeeHistory              TriggerORM = "DeputecEmployeeHistory"
	TriggerEmployeePaycycle             TriggerORM = "DeputecEmployeePaycycle"
	TriggerEmployeePaycycleReturn       TriggerORM = "DeputecEmployeePaycycleReturn"
	TriggerEmployeeRole                 TriggerORM = "DeputecEmployeeRole"
	TriggerEmployeeSalaryOpunitCosting  TriggerORM = "DeputecEmployeeSalaryOpunitCosting"
	TriggerEmployeeWorkplace            TriggerORM = "DeputecEmployeeWorkplace"
	TriggerEmploymentCondition          TriggerORM = "DeputecEmploymentCondition"
	TriggerEmploymentContract           TriggerORM = "DeputecEmploymentContract"
	TriggerEmploymentContractLeaveRules TriggerORM = "DeputecEmploymentContractLeaveRules"
	TriggerEvent                        TriggerORM = "DeputecEvent"
	TriggerGeo                          TriggerORM = "DeputecGeo"
	TriggerJournal                      TriggerORM = "DeputecJournal"
	TriggerKiosk                        TriggerORM = "DeputecKiosk"
	TriggerLeave                        TriggerORM = "DeputecLeave"
	TriggerLeavePayLine                 TriggerORM = "DeputecLeavePayLine"
	TriggerLeaveAccrual                 TriggerORM = "DeputecLeaveAccrual"
	TriggerLeaveRules                   TriggerORM = "DeputecLeaveRules"
	TriggerMemo                         TriggerORM = "DeputecMemo"
	TriggerOperationalUnit              TriggerORM = "DeputecOperationalUnit"
	TriggerPayRules                     TriggerORM = "DeputecPayRules"
	TriggerPayPeriod                    TriggerORM = "DeputecPayPeriod"
	TriggerPublicHoliday                TriggerORM = "DeputecPublicHoliday"
	TriggerRoster                       TriggerORM = "DeputecRoster"
	TriggerRosterOpen                   TriggerORM = "DeputecRosterOpen"
	TriggerRosterSwap                   TriggerORM = "DeputecRosterSwap"
	TriggerRosterTemplate               TriggerORM = "DeputecRosterTemplate"
	TriggerSalesData                    TriggerORM = "DeputecSalesData"
	TriggerSchedule                     TriggerORM = "DeputecSchedule"
	TriggerSmsLog                       TriggerORM = "DeputecSmsLog"
	TriggerState                        TriggerORM = "DeputecState"
	TriggerStressProfile                TriggerORM = "DeputecStressProfile"
	TriggerTask                         TriggerORM = "DeputecTask"
	TriggerTaskGroup                    TriggerORM = "DeputecTaskGroup"
	TriggerTaskGroupSetup               TriggerORM = "DeputecTaskGroupSetup"
	TriggerTaskOpunitConfig             TriggerORM = "DeputecTaskOpunitConfig"
	TriggerTaskSetup                    TriggerORM = "DeputecTaskSetup"
	TriggerTimesheet                    TriggerORM = "DeputecTimesheet"
	TriggerTimesheetPayReturn           TriggerORM = "DeputecTimesheetPayReturn"
	TriggerTeam                         TriggerORM = "DeputecTeam"
	TriggerTrainingModule               TriggerORM = "DeputecTrainingModule"
	TriggerTrainingRecord               TriggerORM = "DeputecTrainingRecord"
	TriggerWebhook                      TriggerORM = "DeputecWebhook"
	TriggerShiftTemplate                TriggerORM = "DeputecShiftTemplate"
)

type DatePickerMode uint

const (
	NoDatePicker     DatePickerMode = 0
	DateRangePicker  DatePickerMode = 1
	SingleDatePicker DatePickerMode = 2
)

type DeputyBoolean uint

const (
	False DeputyBoolean = 0
	True  DeputyBoolean = 1
)

type ReportCategory uint

const (
	ReportCategoryMiscellaneous      ReportCategory = 1
	ReportCategoryCustom             ReportCategory = 2
	ReportCategoryManage             ReportCategory = 3
	ReportCategoryEmployeeDetails    ReportCategory = 4
	ReportCategoryAttendanceAndLeave ReportCategory = 5
	ReportCategoryPayroll            ReportCategory = 6
	ReportCategoryBudgets            ReportCategory = 7
	ReportCategoryPerformance        ReportCategory = 8
)
