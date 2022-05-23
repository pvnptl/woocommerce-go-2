package woocommerce

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReportService_All(t *testing.T) {
	_, err := wooClient.Services.Report.All()
	assert.Equal(t, nil, err)
}

func TestReportService_SalesReports(t *testing.T) {
	req := SalesReportsQueryParams{}
	_, err := wooClient.Services.Report.SalesReports(req)
	assert.Equal(t, nil, err)
}
