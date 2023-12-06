package pickup

import (
	"context"
	qrcode "github.com/skip2/go-qrcode"
	"log"
)

func (s *Service) Me(ctx context.Context, userID int) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode("https://amazon.com", qrcode.Medium, 256)
	if err != nil {
		log.Printf("could not create qr code %s", err)
		return nil, err
	}
	log.Println(png)
	return png, nil
}
