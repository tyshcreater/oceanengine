package file

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// CarouselUpdate 更新图集信息
// 更细图集信息，目前仅支持图集 file_name 和主题修改
func CarouselUpdate(clt *core.SDKClient, accessToken string, req *file.CarouselUpdateRequest) ([]file.CarouselUpdateResult, error) {
	var resp file.CarouselUpdateResponse
	if err := clt.Post("2/file/carousel/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Carousels, nil
}
