package address

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/christhianjesus/bia-challenge/internal/domain/address"
	"github.com/christhianjesus/bia-challenge/internal/infrastructure"
)

type msAddressRepository struct {
	client infrastructure.HTTPClient
}

func NewMSAddressRepository(client infrastructure.HTTPClient) address.AddressRepository {
	return &msAddressRepository{client}
}

func (ms *msAddressRepository) GetByMetersIDs(ctx context.Context, metersIDs []int) (map[int]string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://nginx:8080/api/address", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	var resp *http.Response
	if resp, err = ms.client.Do(req); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var response map[string]map[string]string
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	addresses := make(map[int]string, len(response["addresses"]))
	for k, v := range response["addresses"] {
		meterID, _ := strconv.Atoi(k)
		addresses[meterID] = v
	}

	return addresses, nil
}
