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

var agmaSvc *awsagma.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Fatalf("could not initialize AWS client %v", err)
	}

	agmaSvc = awsagma.NewFromConfig(cfg)
}

func (aaga *AWSApiGwAdapter) Message(ctx context.Context, cID, msg, fromUsername string) error {
	dataStr := fmt.Sprintf("%s: %s", fromUsername, msg)

	in := &awsagma.PostToConnectionInput{
		ConnectionId: aws.String(cID),
		Data:         []byte(dataStr),
	}

	_, err := agmaSvc.PostToConnection(ctx, in)

	return err
}
