package cli

import (
	"errors"
	"strings"
)

func (container *Registry) AddCommand(callable Command) {
	if _, ok := container.Commands[callable.Command]; ok {
		panic("Command " + callable.Command + " already exists")
	}
	container.Commands[callable.Command] = callable
}

func (container *Registry) GetCommand(callable string) (Command, error) {
	if _, ok := container.Commands[callable]; ok {
		return container.Commands[callable], nil
	}
	return Command{}, errors.New("Command " + callable + " not found")
}

func (container *Registry) GetCommands() map[string]Command {
	return container.Commands
}

func (c Command) HasCorrectAmountOfArgs(args []string) bool {
	requiredArgs := make(map[string]CommandArgument)
	specifiedRequiredArgs := 0

	// Get all required args
	for _, arg := range c.Args {
		if arg.Required {
			requiredArgs[arg.Name] = arg
		}
	}

	// Check if all required args are specified
	for _, arg := range args {
		parts := strings.Split(arg, ":")
		argsKey := parts[0]

		// Check if the arg is required
		if _, ok := requiredArgs[argsKey]; ok {
			specifiedRequiredArgs++
		}
	}

	// Check if the amount of required args is correct
	return len(requiredArgs) == specifiedRequiredArgs
}

func (c Command) FindMissingArgs(args []string) []CommandArgument {
	missingArgs := make([]CommandArgument, 0)
	requiredArgs := make(map[string]CommandArgument)

	// Get all required args
	for _, arg := range c.Args {
		if arg.Required {
			requiredArgs[arg.Name] = arg
		}
	}

	// Check if all required args are specified
	for _, arg := range args {
		parts := strings.Split(arg, ":")
		argsKey := parts[0]

		// Check if the arg is required
		if _, ok := requiredArgs[argsKey]; ok {
			delete(requiredArgs, argsKey)
		}
	}

	// Check if the amount of required args is correct
	for _, arg := range requiredArgs {
		missingArgs = append(missingArgs, arg)
	}

	return missingArgs
}

func UnpackArguments(args []string) map[string]string {
	store := make(map[string]string)

	for _, arg := range args {
		parts := strings.Split(arg, ":")
		store[parts[0]] = parts[1]
	}

	return store
}

func (c CommandArgumentValue) GetArgument(key string, defaultValue string) string {

	if _, ok := c.Store[key]; ok {
		return c.Store[key]
	}

	return defaultValue
}
