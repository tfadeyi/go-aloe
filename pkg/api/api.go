// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package api

import "encoding/json"
import "fmt"

type ErrorMetaLoc struct {
	// The line number of error location
	Line int `json:"line" yaml:"line" mapstructure:"line"`

	// Full path of the error location
	Path string `json:"path" yaml:"path" mapstructure:"path"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ErrorMetaLoc) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["line"]; !ok || v == nil {
		return fmt.Errorf("field line in ErrorMetaLoc: required")
	}
	if v, ok := raw["path"]; !ok || v == nil {
		return fmt.Errorf("field path in ErrorMetaLoc: required")
	}
	type Plain ErrorMetaLoc
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ErrorMetaLoc(plain)
	return nil
}

type Error struct {
	// Unique code of the error. (max: 40 characters)
	Code string `json:"code" yaml:"code" mapstructure:"code"`

	// Detailed description of the error.
	Details *string `json:"details,omitempty" yaml:"details,omitempty" mapstructure:"details,omitempty"`

	// Metadata information about the error.
	Meta *ErrorMeta `json:"meta,omitempty" yaml:"meta,omitempty" mapstructure:"meta,omitempty"`

	// Short summary of the error. (max: 70 characters)
	Summary string `json:"summary" yaml:"summary" mapstructure:"summary"`

	// Name of the error, to be displayed. (max: 40 characters)
	Title string `json:"title" yaml:"title" mapstructure:"title"`
}

// Metadata information about the error.
type ErrorMeta struct {
	// Loc corresponds to the JSON schema field "loc".
	Loc *ErrorMetaLoc `json:"loc,omitempty" yaml:"loc,omitempty" mapstructure:"loc,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Error) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["code"]; !ok || v == nil {
		return fmt.Errorf("field code in Error: required")
	}
	if v, ok := raw["summary"]; !ok || v == nil {
		return fmt.Errorf("field summary in Error: required")
	}
	if v, ok := raw["title"]; !ok || v == nil {
		return fmt.Errorf("field title in Error: required")
	}
	type Plain Error
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Error(plain)
	return nil
}

type Application struct {
	// BaseUrl corresponds to the JSON schema field "base_url".
	BaseUrl string `json:"base_url" yaml:"base_url" mapstructure:"base_url"`

	// Description corresponds to the JSON schema field "description".
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// List of application error definitions.
	ErrorsDefinitions ErrorDefinitions `json:"errors_definitions,omitempty" yaml:"errors_definitions,omitempty" mapstructure:"errors_definitions,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// Display name of the application
	Title *string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`

	// Version corresponds to the JSON schema field "version".
	Version string `json:"version" yaml:"version" mapstructure:"version"`
}

type ErrorDefinitions map[string]Error

// UnmarshalJSON implements json.Unmarshaler.
func (j *Application) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["base_url"]; !ok || v == nil {
		return fmt.Errorf("field base_url in Application: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in Application: required")
	}
	if v, ok := raw["version"]; !ok || v == nil {
		return fmt.Errorf("field version in Application: required")
	}
	type Plain Application
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Application(plain)
	return nil
}