package types

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"go.uber.org/zap"

	"github.com/Faaizz/simple_http_chatapp/misc"
)

var logger *zap.SugaredLogger
var ddbSvc *dynamodb.Client

func init() {
	logger = misc.Logger()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Fatalf("could not initialize AWS client %v", err)
	}

	ddbSvc = dynamodb.NewFromConfig(cfg)
}

// A DynamoDBAdapter provides a layer of abstraction for interaction an underlying AWS DynamoDB database
type DynamoDBAdapter struct {
	TableName string
}

func (dba *DynamoDBAdapter) SetTableName(tn string) {
	dba.TableName = tn
}

func (dba *DynamoDBAdapter) CheckExists() error {
	in := dynamodb.DescribeTableInput{
		TableName: aws.String(dba.TableName),
	}
	_, err := ddbSvc.DescribeTable(context.TODO(), &in)
	if err != nil {
		return err
	}

	return nil
}

// PutConn inserts a username and connectionId in the underlying DynamoDB table
func (dba *DynamoDBAdapter) PutConn(pcIn PutConnInput) error {
	err := dba.CheckUsername(pcIn.Username)
	if err != nil {
		return err
	}

	in := dynamodb.PutItemInput{
		TableName: aws.String(dba.TableName),
		Item: map[string]dynamodbtypes.AttributeValue{
			"connectionId": &dynamodbtypes.AttributeValueMemberS{
				Value: pcIn.ConnectionID,
			},
			"username": &dynamodbtypes.AttributeValueMemberS{
				Value: pcIn.Username,
			},
		},
	}

	_, err = ddbSvc.PutItem(
		context.TODO(),
		&in,
	)
	if err != nil {
		return err
	}

	return nil
}

// CheckUsername checks if username already exists on DynamDB table
func (dba *DynamoDBAdapter) CheckUsername(username string) error {
	in := dynamodb.GetItemInput{
		TableName: aws.String(dba.TableName),
		Key: map[string]dynamodbtypes.AttributeValue{
			"username": &dynamodbtypes.AttributeValueMemberS{
				Value: username,
			},
		},
	}
	out, err := ddbSvc.GetItem(context.TODO(), &in)
	if err != nil {
		return err
	}

	if out.Item == nil {
		return nil
	}

	return fmt.Errorf("username '%s' already exists", username)
}
