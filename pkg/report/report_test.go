package report_test

import (
	"testing"

	"github.com/ONSdigital/who-goes-there/pkg/report"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	r := report.New()
	assert.NotNil(t, r)
	assert.NotNil(t, r.Generated, "Generated time is populated")
}
