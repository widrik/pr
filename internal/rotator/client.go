package rotator

import (
	"context"

	"github.com/widrik/pr/api/spec"
)

type Client struct {
	rotator *Rotator
}

func NewRotatorClient(rotator *Rotator) *Client {
	return &Client{
		rotator: rotator,
	}
}

func (client Client) Add(ctx context.Context, request *spec.AddRequest) (*spec.AddResponse, error) {
	bannerID := uint(request.BannerId)
	slotID := uint(request.SlotId)

	err := client.rotator.Add(bannerID, slotID)

	return &spec.AddResponse{}, err
}

func (client Client) Delete(ctx context.Context, request *spec.DeleteRequest) (*spec.DeleteResponse, error) {
	bannerID := uint(request.BannerId)
	slotID := uint(request.SlotId)

	err := client.rotator.Delete(bannerID, slotID)

	return &spec.DeleteResponse{}, err
}

func (client Client) Hit(ctx context.Context, request *spec.HitRequest) (*spec.HitResponse, error) {
	bannerID := uint(request.BannerId)
	slotID := uint(request.SlotId)
	socialGroupID := uint(request.SdgId)

	err := client.rotator.Hit(bannerID, slotID, socialGroupID)

	return &spec.HitResponse{}, err
}

func (client Client) Get(ctx context.Context, request *spec.GetRequest) (*spec.GetResponse, error) {
	slotID := uint(request.SlotId)
	socialGroupID := uint(request.SdgId)

	selectedBanner, err := client.rotator.Get(slotID, socialGroupID)
	bannerID := uint32(selectedBanner.ID)

	return &spec.GetResponse{BannerId: bannerID}, err
}
