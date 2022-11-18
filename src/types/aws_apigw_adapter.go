package types

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awsagma "github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

type AWSApiGwAdapter struct {
}

func (aaga *AWSApiGwAdapter) Spawn(url string) (*awsagma.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Errorf("could not initialize AWS client %v", err)
		return nil, err
	}

	erFunc := awsagma.EndpointResolverFromURL(url)
	agmaSvc := awsagma.NewFromConfig(
		cfg,
		awsagma.WithEndpointResolver(erFunc),
	)

	return agmaSvc, nil
}

func (aaga *AWSApiGwAdapter) Message(ctx context.Context, cID, msg, fromUsername, url string) error {
	dataStr := fmt.Sprintf("%s: %s", fromUsername, msg)

	client, err := aaga.Spawn(url)
	if err != nil {
		logger.Errorln(err)
		return err
	}

	in := &awsagma.PostToConnectionInput{
		ConnectionId: aws.String(cID),
		Data:         []byte(dataStr),
	}

	_, err = client.PostToConnection(ctx, in)

	return err
}
