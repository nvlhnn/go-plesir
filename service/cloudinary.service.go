package service

import (
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

type CloudinaryService interface {
	UploadImages(ctx *gin.Context, images []multipart.FileHeader) []string
	DeleteImages(ctx *gin.Context, publicIds []string) 
	GetImageUrl(ctx *gin.Context, publicId string) string
}

type cloudinaryService struct {
	cld *cloudinary.Cloudinary
}

//NewAuthService creates a new instance of AuthService
func NewCloudinaryService() CloudinaryService {

	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_APP"), os.Getenv("CLOUDINARY_KEY"), os.Getenv("CLOUDINARY_SECRET"))
	
	return &cloudinaryService{cld:cld}
}


func (s *cloudinaryService) UploadImages(ctx *gin.Context ,images []multipart.FileHeader) []string{

	var urls  []string

	for _, img := range images {
		file, _ := img.Open()
		resp, _ := s.cld.Upload.Upload(ctx, file, uploader.UploadParams{})
		urls = append(urls, resp.SecureURL)
		
	}

	

	return urls

}

func (s *cloudinaryService) GetImageUrl(ctx *gin.Context, id string) (string) {

	url, _ := s.cld.Admin.Asset(ctx, admin.AssetParams{PublicID: id})
	return url.SecureURL
}

func (s *cloudinaryService) DeleteImages(ctx *gin.Context, ids []string) {

	// s.cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{PublicIDs: ids})

}