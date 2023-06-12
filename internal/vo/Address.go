package vo

import (
	"mall-demo/internal/model"
)

type AddressResponse struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"userID"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func BuildAddressResponse(address *model.Address) AddressResponse {
	return AddressResponse{
		ID:      address.ID,
		UserID:  address.UserID,
		Name:    address.Name,
		Phone:   address.Phone,
		Address: address.Address,
	}
}

func BuildAddressResponseList(addressList []model.Address) (addressResponseList []AddressResponse) {
	for _, address := range addressList {
		addressResponse := BuildAddressResponse(&address)
		addressResponseList = append(addressResponseList, addressResponse)
	}
	return addressResponseList
}
