package localService

import (
	"context"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/countSyncMap"
	"go2ban/pkg/osUtil"
	"go2ban/pkg/validator"
	"strings"
)

func checkLogAndBlock(ctx context.Context, service config.Service, countFailsMap *countSyncMap.Counters, maxFails int) {
	//TODO ReadAt //find byte //sync map name+date = last byte
	sts, _ := osUtil.ReadStsFile(service.LogFile)

	for _, st := range sts {
		if strings.Contains(st, service.Regxp) {

			ip, err := validator.CheckIp(st)

			if err == nil {
				countFailsMap.Inc(ip)
				if countFailsMap.Load(ip) > maxFails {

					go firewall.BlockIP(ctx, ip)

				}
			}
		}
	}
}
