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
	bannerId := uint(request.BannerId)
	slotId := uint(request.SlotId)

	err := client.rotator.Add(bannerId, slotId)

	return &spec.AddResponse{}, err
}

func (client Client) Delete(ctx context.Context, request *spec.DeleteRequest) (*spec.DeleteResponse, error) {
	bannerId := uint(request.BannerId)
	slotId := uint(request.SlotId)

	err := client.rotator.Delete(bannerId, slotId)

	return &spec.DeleteResponse{}, err
}

func (client Client) Hit(ctx context.Context, request *spec.HitRequest) (*spec.HitResponse, error) {
	bannerId := uint(request.BannerId)
	slotId := uint(request.SlotId)
	socialGroupId := uint(request.SdgId)

	err := client.rotator.Hit(bannerId, slotId, socialGroupId)

	return &spec.HitResponse{}, err
}

func (client Client) Get(ctx context.Context, request *spec.GetRequest) (*spec.GetResponse, error) {
	slotId := uint(request.SlotId)
	socialGroupId := uint(request.SdgId)

	selectedBanner, err := client.rotator.Get(slotId, socialGroupId)
	bannerId := uint32(selectedBanner.ID)

	return &spec.GetResponse{BannerId: bannerId}, err
}