package executor

import (
	"fmt"
	"os"
	"strings"

	"secretKeeper/internal/client/app"
	"secretKeeper/internal/client/model"
)

type Executor struct {
	app *app.App
}

func NewExecutor() *Executor {
	appL, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	return &Executor{app: appL}
}

func (e *Executor) Execute(s string) {
	var isForce bool

	setCommand, options := getCommandArgsAndOptions(s)
	if options["force"] || options["f"] {
		isForce = true
	}

	switch setCommand[0] {
	case "login":
		if err := e.login(setCommand); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("successfully authorized")
		return
	case "register":
		if err := e.register(setCommand); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("User successfully created. You are logged in.")

		return
	case "delete-user":
		if err := e.deleteUser(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("you successfully deleted account and logged out")

		return
	case "logout":
		if err := e.logout(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("you successfully logged out")
		return
	case "types":
		types, err := e.types()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, t := range types {
			fmt.Printf("%+v\n", t)
		}

		return
	case "create-auth":
		if err := e.createAuth(setCommand); err != nil {
			fmt.Println(err)
			return
		}
		return
	case "create-text":
		if err := e.createText(setCommand); err != nil {
			fmt.Println(err)
			return
		}

		return
	case "create-binary":
		if err := e.createBinary(setCommand); err != nil {
			fmt.Println(err)
			return
		}

		return
	case "create-card":
		if err := e.createCard(setCommand); err != nil {
			fmt.Println(err)
			return
		}

		return
	case "delete-secret":
		if err := e.deleteSecret(setCommand); err != nil {
			fmt.Println(err)
			return
		}

		return
	case "get-secrets-by-type":
		list, err := e.getSecretsByTypeId(setCommand)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, secret := range list {
			fmt.Printf("ID:%v Title: %v\n", secret.Id, secret.Title)
		}

		return
	case "get-secret":
		secret, err := e.getSecret(setCommand)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Content:%+v\n", secret)

		return
	case "get-secret-binary":
		if err := e.getSecretBinary(setCommand); err != nil {
			fmt.Println(err)
			return
		}
		return
	case "edit-secret":
		if err := e.editSecret(setCommand, isForce); err != nil {
			fmt.Println(err)
			return
		}

		return
	case "exit":
		fmt.Println("bye bye...application is closing")

		e.app.Cancel()
		e.app.Cron.Stop()

		os.Exit(0)
	}
}

// types - is executor for "types" case in Execute method.
func (e *Executor) types() ([]model.SecretType, error) {
	secrets, err := e.app.SecretTypeService.List()
	if err != nil {
		return nil, err
	}

	var models []model.SecretType
	for _, secret := range secrets.Secrets {

		models = append(models, model.SecretType{
			Id:    int(secret.Id),
			Title: secret.Title,
		})
	}

	return models, nil
}

// getCommandArgsAndOptions - splits args to command args and options
func getCommandArgsAndOptions(s string) ([]string, map[string]bool) {
	s = strings.TrimSpace(s)

	setCommand := strings.Split(s, " ")

	l := len(setCommand)

	filtered := make([]string, 0, l)
	options := make(map[string]bool)

	for i := 0; i < len(setCommand); i++ {
		if strings.HasPrefix(setCommand[i], "-") {
			opt := strings.TrimPrefix(setCommand[i], "--")
			opt = strings.TrimPrefix(opt, "-")

			optSplited := strings.Split(opt, "=")
			options[optSplited[0]] = true

			continue
		}
		filtered = append(filtered, setCommand[i])
	}

	return filtered, options
}
