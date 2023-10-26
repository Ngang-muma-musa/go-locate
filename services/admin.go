package services

import "go-locate/repository"

type Admin struct {
	businessRepo *repository.Business
}

func NewAdmin(businessRepo *repository.Business) *Admin {
	return &Admin{businessRepo: businessRepo}
}

func (a *Admin) VerifyBusiness(ID uint) error {
	business, err := a.businessRepo.GetByID(ID)
	if err != nil {
		return err
	}
	business.Verified = true
	err = a.businessRepo.Update(business)
	return err
}
