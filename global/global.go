package global

import (
	"SamWaf/model"
	"SamWaf/model/spec"
	"github.com/bytedance/godlp/dlpheader"
	"gorm.io/gorm"
	"time"
)

const (
	Version_name = "v1.0.1"
	Version_num  = 1
)

var (
	GWAF_LOCAL_DB          *gorm.DB                    //通用本地数据库，尊重用户隐私
	GWAF_REMOTE_DB         *gorm.DB                    //仅当用户使用云数据库
	GWAF_LOCAL_SERVER_PORT int                 = 26666 // 本地local端口
	GWAF_USER_CODE         string                      // 当前识别号
	GWAF_TENANT_ID         string                      // 当前租户ID
	GWAF_RELEASE           bool                = true  // 当前是否为发行版
	GWAF_LAST_UPDATE_TIME  time.Time                   // 上次时间
	GWAF_DLP               dlpheader.EngineAPI         // 脱敏引擎
	/**链聚合**/
	GWAF_CHAN_HOST   = make(chan model.Hosts, 10) //主机链
	GWAF_CHAN_ENGINE = make(chan int, 10)         //引擎链

	GWAF_CHAN_MSG = make(chan spec.ChanCommonHost, 10) //全局通讯包

)
