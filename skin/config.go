package skin

import (
	"github.com/IvoryRaptor/postoffice/matrix"
	"github.com/IvoryRaptor/postoffice/mq"
)

type Config struct {
	Matrix matrix.Config            `yaml:"matrix"`
	MQ     mq.Config                `yaml:"mq"`
	Frequency int					`yaml:"frequency"`
}
