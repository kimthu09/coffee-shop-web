package exportnotemodel

type ExportReason string

const (
	Damaged   ExportReason = "Damaged"
	OutOfDate ExportReason = "OutOfDate"
)
