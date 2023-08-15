package chaosmaker

import (
	"github.com/yaklang/yaklang/common/chaosmaker/rule"
	"github.com/yaklang/yaklang/common/log"
	"github.com/yaklang/yaklang/common/suricata"
	"github.com/yaklang/yaklang/common/utils/lowhttp"
)

const GENCOUNT = 5

func init() {
	chaosMap.Store("suricata-http", &httpHandler{})
}

type httpHandler struct {
}

var _ chaosHandler = (*httpHandler)(nil)

func (h *httpHandler) MatchBytes(i any) bool {
	//TODO implement me
	panic("implement me")
}

func (h *httpHandler) Generator(maker *ChaosMaker, chaosRule *rule.Storage, originRule *suricata.Rule) chan *ChaosTraffic {
	if originRule == nil {
		return nil
	}
	if originRule.Protocol != "http" {
		return nil
	}

	ch := make(chan *ChaosTraffic)
	go (&httpGenerator{
		chaosRule:  chaosRule,
		originRule: originRule,
		maker:      maker,
		out:        ch,
	}).generator(GENCOUNT)

	return ch
}

type httpGenerator struct {
	chaosRule  *rule.Storage
	originRule *suricata.Rule
	maker      *ChaosMaker
	out        chan *ChaosTraffic
}

func (h *httpGenerator) generator(count int) {
	surigen, err := suricata.NewSurigen(h.originRule.ContentRuleConfig.ContentRules)
	if err != nil {
		log.Warnf("suricata.NewSurigen failed: %v", err)
	}

	for i := 0; i < count; i++ {
		raw, err := surigen.Gen()
		if err != nil {
			log.Warnf("surigen.Gen failed: %v", err)
			continue
		}
		h.toChaosTraffic(raw)
	}

	close(h.out)
}

func (h *httpGenerator) toChaosTraffic(raw []byte) {
	if lowhttp.IsResp(raw) {
		h.out <- &ChaosTraffic{
			ChaosRule:    h.chaosRule,
			SuricataRule: h.originRule,
			HttpResponse: raw,
		}
	} else {
		h.out <- &ChaosTraffic{
			ChaosRule:    h.chaosRule,
			SuricataRule: h.originRule,
			HttpRequest:  raw,
		}
	}
}
