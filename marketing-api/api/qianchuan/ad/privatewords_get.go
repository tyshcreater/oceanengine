package ad

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// PrivatewordsGet 获取否定词列表
func PrivatewordsGet(clt *core.SDKClient, accessToken string, req *ad.PrivatewordsGetRequest) (*ad.PrivateWords, error) {
	var resp ad.PrivatewordsGetResponse
	if err := clt.Get("v1.0/qianchuan/ad/privatewords/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
