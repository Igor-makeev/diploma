package executor

import (
	"fmt"
	"secretKeeper/internal/client/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// login - is executor for "login" case in Execute method.
func (e *Executor) login(args []string) error {
	switch len(args) - 1 {
	case 1:
		return fmt.Errorf("validation error: Password is missing")
	case 0:
		return fmt.Errorf("validation error: Login and Password is missing")
	}

	user := model.User{Login: args[1], Password: args[2]}

	if err := e.app.UserService.Login(user); err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.NotFound:
			return fmt.Errorf("error: User not found")
		default:
			return fmt.Errorf("error:" + st.Message())
		}
	}

	// firstly we sync all on start up
	e.app.Syncer.SyncAll()

	// then we spawn goroutin with cron job to sync data every minute
	go e.app.Cron.Run()

	return nil
}

// register - is executor for "register" case in Execute method.
func (e *Executor) register(args []string) error {
	switch len(args) - 1 {
	case 1:
		return fmt.Errorf("validation error: Password is missing")
	case 0:
		return fmt.Errorf("validation error: Login and Password is missing")
	}

	user := model.User{Login: args[1], Password: args[2]}
	if err := e.app.UserService.Register(user); err != nil {
		switch status.Code(err) {
		case codes.InvalidArgument:
			return fmt.Errorf("error: data is invalid or user already exists")
		default:
			return err
		}
	}

	return nil
}

// deleteUser - is executor for "delete-user" case in Execute method.
func (e *Executor) deleteUser() error {
	return e.app.UserService.Delete()
}

// logout - is executor for "types" case in Execute method.
func (e *Executor) logout() error {
	e.app.UserService.Logout()

	e.app.Cron.Stop()

	e.app.Storage.ResetStorage()

	return nil
}
