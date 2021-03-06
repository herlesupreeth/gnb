package util

import (
	"free5gc/lib/openapi/models"
	"free5gc/src/gnb/context"
	"free5gc/src/gnb/factory"
	"free5gc/src/gnb/logger"

	"github.com/google/uuid"
)

func InitRanContext(context *context.RANContext) {
	config := factory.RanConfig
	logger.UtilLog.Infof("gNB config Info: Version[%s] Description[%s]", config.Info.Version, config.Info.Description)
	configuration := config.Configuration
	context.NfId = uuid.New().String()
	if configuration.RanName != "" {
		context.Name = configuration.RanName
	}
	sbi := configuration.Sbi
	context.UriScheme = models.UriScheme(sbi.Scheme)
	context.HttpIPv4Address = "127.0.0.1" // default localhost
	context.HttpIpv4Port = 32000          // default port
	if sbi != nil {
		if sbi.IPv4Addr != "" {
			context.HttpIPv4Address = sbi.IPv4Addr
		}
		if sbi.Port != 0 {
			context.HttpIpv4Port = sbi.Port
		}
	}

	// for i := range context.SupportTaiLists {
	// 	context.SupportTaiLists[i].Tac = TACConfigToModels(context.SupportTaiLists[i].Tac)
	// }
	context.NetworkName = configuration.NetworkName
	context.AmfInterface = configuration.AmfInterface
	context.UpfInterface = configuration.UpfInterface
	context.UEList = configuration.UEList
	context.NGRANInterface = configuration.NGRANInterface
	context.GTPInterface = configuration.GTPInterface
	context.Security = configuration.Security
	context.Snssai = configuration.Snssai
	context.PLMN = configuration.PLMN
}
