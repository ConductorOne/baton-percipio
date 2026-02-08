package config

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	ApiTokenField = field.StringField(
		"api-token",
		field.WithDescription("The Percipio Bearer Token"),
		field.WithIsSecret(true),
		field.WithRequired(true),
	)
	OrganizationIdField = field.StringField(
		"organization-id",
		field.WithDescription("The Percipio Organization ID"),
		field.WithRequired(true),
	)

	LimitCoursesField = field.StringSliceField(
		"limited-courses",
		field.WithDescription("Limit imported courses to a specific list by Course ID"),
		field.WithRequired(false),
	)

	BaseURLField = field.StringField(
		"base-url",
		field.WithDescription("Override the Percipio API URL (for testing)"),
	)

	// ConfigurationFields defines the external configuration required for the
	// connector to run. Note: these fields can be marked as optional or
	// required.
	ConfigurationFields = []field.SchemaField{
		ApiTokenField,
		OrganizationIdField,
		LimitCoursesField,
		BaseURLField,
	}

	// FieldRelationships defines relationships between the fields listed in
	// ConfigurationFields that can be automatically validated. For example, a
	// username and password can be required together, or an access token can be
	// marked as mutually exclusive from the username password pair.
	FieldRelationships = []field.SchemaFieldRelationship{}
)

//go:generate go run ./gen
var Config = field.NewConfiguration(
	ConfigurationFields,
	field.WithConstraints(FieldRelationships...),
)
