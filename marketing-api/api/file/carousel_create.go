package file

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// CarouselCreate 上传图集
// 利用图片和音频塑造图集。
// 通过上传广告图片接口按顺序上传图集中的图片，获取图片id；
// 通过上传视频接口上传音频（video_file或video_url），获取音频id（video_id）
// 利用上述两步的图片、视频id塑造图集，获取出参图集mid。不同图片顺序对应不同的图集mid。
func CarouselCreate(clt *core.SDKClient, accessToken string, req *file.CarouselCreateRequest) (*file.Carousel, error) {
	var resp file.CarouselCreateResponse
	if err := clt.Post("2/file/carousel/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Carousel, nil
}
