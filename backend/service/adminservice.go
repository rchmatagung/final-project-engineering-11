package service

import "context"

func (r *AuthServiceimpl) DeleteById(id int) error {
	ctx := context.Background()
	err := r.userRepo.DeleteUserById(ctx, id)

	if err != nil {
		return err
	}
	return nil
}
