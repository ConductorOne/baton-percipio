package main

import (
	"context"
	"fmt"
	"os"

	config2 "github.com/conductorone/baton-percipio/pkg/config"
	"github.com/conductorone/baton-percipio/pkg/connector"
	"github.com/conductorone/baton-sdk/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/types"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	connectorName = "baton-percipio"
	version       = "dev"
)

func main() {
	ctx := context.Background()

	_, cmd, err := config.DefineConfiguration(
		ctx,
		connectorName,
		getConnector,
		config2.ConfigurationSchema,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	cmd.Version = version

	err = cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getConnector(ctx context.Context, v *viper.Viper) (types.ConnectorServer, error) {
	l := ctxzap.Extract(ctx)
	cb, err := connector.New(
		ctx,
		v.GetString(config2.OrganizationIdField.FieldName),
		v.GetString(config2.ApiTokenField.FieldName),
	)
	if err != nil {
		l.Error("error creating connector", zap.Error(err))
		return nil, err
	}
	connector, err := connectorbuilder.NewConnector(ctx, cb)
	if err != nil {
		l.Error("error creating connector", zap.Error(err))
		return nil, err
	}
	return connector, nil
}
