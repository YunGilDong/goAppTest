package main

type SendReport struct {

}

func (f *FinanceReport) SendReport(r *Report, method int) *Report {

}

func (r *FinanceReport) SendReport()

type ReportSender interface {
	SendReport(*Report)
}

type EmailReportSender struct {

}


func (s *EmailReportSender) SendReport(r *ReportSender) {

}

type FileReportSender struct {

}

func (s *FileReportSender) SendReport(r *ReportSender) {

}

type HttpReportSender sruct {

}

func (s *HttpReportSender) SendReport(r *ReportSender)

func main() {

}