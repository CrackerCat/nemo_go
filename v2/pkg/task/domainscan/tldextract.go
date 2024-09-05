package domainscan

import (
	"fmt"
	"github.com/hanc00l/nemo_go/v2/pkg/conf"
	"github.com/joeguo/tldextract"
	"path"
)

type TldExtract struct {
	extract *tldextract.TLDExtract
}

// NewTldExtract 创建TldExtract对象
func NewTldExtract() TldExtract {
	tldCacheFile := path.Join(conf.GetRootPath(), "thirdparty/tldextract/tld.cache")
	t := TldExtract{}
	t.extract, _ = tldextract.New(tldCacheFile, false)

	return t
}

// ExtractFLD 从url或domain中提取FLD
func (t *TldExtract) ExtractFLD(url string) (fldDomain string) {
	result := t.extract.Extract(url)
	if result == nil {
		return
	}
	if result.Flag == tldextract.Domain {
		fldDomain = fmt.Sprintf("%s.%s", result.Root, result.Tld)
	}
	return
}
