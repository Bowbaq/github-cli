package main

import (
	"fmt"
	"log"
)

type flag struct {
	Name  string
	Typ   string
	Usage string
}

func (f flag) String() string {
	return f.Name + " " + f.Typ + " - " + f.Usage
}

func (f flag) Declaration() string {
	switch f.Typ {
	case "int":
		fallthrough
	case "*int":
		return fmt.Sprintf("cli.IntFlag{Name: `%s`, Usage: `%s`}", dasherize(f.Name), f.Usage)
	case "bool":
		fallthrough
	case "*bool":
		return fmt.Sprintf("cli.BoolFlag{Name: `%s`, Usage: `%s`}", dasherize(f.Name), f.Usage)
	case "string":
		fallthrough
	case "*string":
		fallthrough
	case "time.Time":
		fallthrough
	case "*time.Time":
		fallthrough
	case "*github.Timestamp":
		return fmt.Sprintf("cli.StringFlag{Name: `%s`, Usage: `%s`}", dasherize(f.Name), f.Usage)
	case "[]string":
		fallthrough
	case "*[]string":
		return fmt.Sprintf("cli.StringSliceFlag{Name: `%s`, Usage: `%s`}", dasherize(f.Name), f.Usage)
	default:
		log.Println("no declaration for flag type " + f.Typ)
		return ""
	}
}

func (f flag) Accessor() string {
	switch f.Typ {
	case "int":
		return fmt.Sprintf(`c.Int("%s")`, dasherize(f.Name))
	case "*int":
		return fmt.Sprintf(`github.Int(c.Int("%s"))`, dasherize(f.Name))
	case "bool":
		return fmt.Sprintf(`c.Bool("%s")`, dasherize(f.Name))
	case "*bool":
		return fmt.Sprintf(`github.Bool(c.Bool("%s"))`, dasherize(f.Name))
	case "string":
		return fmt.Sprintf(`c.String("%s")`, dasherize(f.Name))
	case "*string":
		return fmt.Sprintf(`github.String(c.String("%s"))`, dasherize(f.Name))
	case "[]string":
		return fmt.Sprintf(`c.StringSlice("%s")`, dasherize(f.Name))
	case "*[]string":
		return fmt.Sprintf(`stringSlicePointer(c.StringSlice("%s"))`, dasherize(f.Name))
	case "time.Time":
		return fmt.Sprintf(`now.MustParse(c.String("%s"))`, dasherize(f.Name))
	case "*time.Time":
		return fmt.Sprintf(`timePointer(now.MustParse(c.String("%s")))`, dasherize(f.Name))
	case "*github.Timestamp":
		return fmt.Sprintf(`&github.Timestamp{now.MustParse(c.String("%s"))}`, dasherize(f.Name))
	default:
		log.Println("no accessor for flag type " + f.Typ)
		return ""
	}
}
