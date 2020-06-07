package colorlog_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/donbattery/colorlog"
)

type ColorLogSuite struct {
	suite.Suite
	logger *colorlog.ColorLogger
}

func (cls *ColorLogSuite) TestComponent() {
	component := "Server"
	cls.logger = colorlog.NewColorLogger(component, false)
	cls.NotNil(cls.logger, "A new Color Logger should have been created")

	buf := bytes.NewBuffer([]byte{})

	cls.logger.Output = buf

	cls.logger.Info("a")

	cls.Containsf(buf.String(), component+":", "The output should contain the component name '%s:'", component)
}

func TestColorLogSuite(t *testing.T) {
	suite.Run(t, new(ColorLogSuite))
}
