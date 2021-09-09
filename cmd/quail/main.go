package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type ttype struct {
	Fields map[string]struct{}
}

type tthing struct {
	Type   string
	Fields map[string]string
}

var (
	types  = map[string]*ttype{}
	things = map[string]*tthing{}
)

func createType(name string) error {
	if _, ok := types[name]; ok {
		return fmt.Errorf("%s exists", name)
	}

	types[name] = &ttype{
		Fields: map[string]struct{}{},
	}
	return nil
}

func createTypeField(name string, field string) error {
	if _, ok := types[name]; !ok {
		return fmt.Errorf("%s does not exist", name)
	}
	if _, ok := types[name].Fields[field]; ok {
		return fmt.Errorf("%s exists", field)
	}

	types[name].Fields[field] = struct{}{}
	return nil
}

func createThing(name string, typename string) error {
	if _, ok := things[name]; ok {
		return fmt.Errorf("%s exists", name)
	}

	things[name] = &tthing{
		Type:   typename,
		Fields: map[string]string{},
	}
	return nil
}

func setThingField(name string, field string, value string) error {
	if _, ok := things[name]; !ok {
		return fmt.Errorf("%s does not exist", name)
	}
	if _, ok := types[things[name].Type].Fields[field]; !ok {
		return fmt.Errorf("%s does not exist", field)
	}

	things[name].Fields[field] = value
	return nil
}

func getThingField(name string, field string) (string, error) {
	if _, ok := things[name]; !ok {
		return "", fmt.Errorf("%s does not exist", name)
	}
	if _, ok := types[things[name].Type].Fields[field]; !ok {
		return "", fmt.Errorf("%s does not exist", field)
	}

	return things[name].Fields[field], nil
}

func processLine() error {
	// Read line
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputSlice := strings.Split(input, " ")

	// Process line
	switch inputSlice[0] {
	case "A":
		fallthrough
	case "An":
		name := inputSlice[1]
		switch inputSlice[2] {
		case "has":
			switch inputSlice[3] {
			case "a", "an":
				field := inputSlice[4]
				createType(name)
				return createTypeField(name, field)
			}
		}
	case "The":
		field := inputSlice[1]
		switch inputSlice[2] {
		case "of":
			name := inputSlice[3]
			switch inputSlice[4] {
			case "is":
				value := inputSlice[5]
				return setThingField(name, field, value)
			}
		}
	case "What":
		switch inputSlice[1] {
		case "is":
			switch inputSlice[2] {
			case "the":
				field := inputSlice[3]
				switch inputSlice[4] {
				case "of":
					name := inputSlice[5]
					value, err := getThingField(name, field)
					if err != nil {
						return err
					}
					if value == "" {
						fmt.Println("none")
					}
					fmt.Println(value)
					return nil
				}
			}
		}
	default:
		name := inputSlice[0]
		switch inputSlice[1] {
		case "is":
			switch inputSlice[2] {
			case "a":
				typename := inputSlice[3]
				return createThing(name, typename)
			}
		}
	}

	return errors.New("bad line")
}

func run() error {
	for {
		if err := processLine(); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
