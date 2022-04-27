package fingerprint

import (
	"github.com/hanc00l/nemo_go/pkg/task/domainscan"
	"github.com/hanc00l/nemo_go/pkg/task/portscan"
	"testing"
)

func TestHttpx_Run(t *testing.T) {
	domainConfig := domainscan.Config{Target: "800best.com"}
	subdomain := domainscan.NewSubFinder(domainConfig)
	subdomain.Do()
	t.Log(subdomain.Result)

	httpx := NewHttpx()
	httpx.ResultDomainScan = subdomain.Result
	httpx.Do()
	t.Log(httpx.ResultDomainScan)
	for d,da := range httpx.ResultDomainScan.DomainResult{
		t.Log(d,da)
	}
	subdomain.Result.SaveResult(subdomain.Config)
}

func TestHttpx_Run2(t *testing.T) {
	nmapConfig := portscan.Config{
		Target:       "47.98.181.116",
		Port:         "80,443",
		Rate:         1000,
		IsPing:       false,
		Tech:         "-sS",
		CmdBin:       "nmap",
	}
	nmap := portscan.NewNmap(nmapConfig)
	nmap.Do()


	httpx:= NewHttpx()
	httpx.ResultPortScan = nmap.Result
	httpx.Do()
	for ip,r := range httpx.ResultPortScan.IPResult{
		t.Log(ip,r)
		for port,p := range r.Ports{
			t.Log(port,p)
		}
	}
	//httpx.ResultPortScan.SaveResult(nmap.Config)
}