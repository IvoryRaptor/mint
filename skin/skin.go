package skin

import (
	"github.com/IvoryRaptor/postoffice/source"
	"github.com/IvoryRaptor/postoffice/matrix"
	"github.com/IvoryRaptor/postoffice/mq"
	"sync"
)

type Skin struct {
	host         string
	ConfigFile   string
	run          bool
	source       []source.ISource
	matrixManger matrix.Manager
	from         mq.IMQ
	to           mq.IMQ
	clients      sync.Map
	redisMutex   sync.Mutex
}

