// +build oci

package impt

import (
	_ "github.com/micro-plat/hydra/engines"
	_ "github.com/micro-plat/hydra/registry/local"
	_ "github.com/micro-plat/hydra/registry/zookeeper"
	_ "github.com/micro-plat/hydra/rpc"
	_ "github.com/micro-plat/hydra/servers/cron"
	_ "github.com/micro-plat/hydra/servers/http"
	_ "github.com/micro-plat/hydra/servers/mqc"
	_ "github.com/micro-plat/hydra/servers/rpc"
	_ "github.com/qxnw/lib4go/cache/memcache"
	_ "github.com/qxnw/lib4go/cache/redis"
	_ "github.com/qxnw/lib4go/mq/mqtt"
	_ "github.com/qxnw/lib4go/mq/redis"
	_ "github.com/qxnw/lib4go/mq/stomp"
	_ "github.com/qxnw/lib4go/mq/xmq"
	_ "github.com/qxnw/lib4go/queue"
	_ "github.com/qxnw/lib4go/queue/redis"
	_ "github.com/zkfy/go-oci8"
)