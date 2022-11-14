package types

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/Faaizz/simple_http_chatapp/misc"
)

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
	User      User
}

func (dba *DynamoDBAdapter) SetTableName(tn string) {
	dba.TableName = tn
}

func (dba *DynamoDBAdapter) CheckExists(ctx context.Context) error {
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
func (dba *DynamoDBAdapter) PutConn(ctx context.Context, pcIn User) error {
	err := dba.CheckUsername(ctx, pcIn.Username)
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
		ctx,
		&in,
	)
	if err != nil {
		return err
	}

	// set current user
	dba.SetUser(ctx, pcIn)

	return nil
}

// CheckUsername checks if username already exists on DynamDB table
func (dba *DynamoDBAdapter) CheckUsername(ctx context.Context, username string) error {
	in := dynamodb.ScanInput{
		TableName: aws.String(dba.TableName),
		FilterExpression: aws.String(
			fmt.Sprintf("username = %s", username),
		),
	}

	out, err := ddbSvc.Scan(ctx, &in)
	if err != nil {
		return err
	}

	if out.Count <= 0 {
		return nil
	}

	return fmt.Errorf("username '%s' already exists", username)
}

func (dba *DynamoDBAdapter) SetUser(ctx context.Context, u User) {
	dba.User = u
}

// AvailableUsers lists available users and their connection IDs
// Possible bug: return payload exceeds maximum dataset size limit of 1 MB
func (dba *DynamoDBAdapter) AvailableUsers(ctx context.Context) ([]User, error) {
	in := &dynamodb.ScanInput{
		TableName: &dba.TableName,
	}

	out, err := ddbSvc.Scan(
		ctx,
		in,
	)
	if err != nil {
		return []User{}, err
	}

	au := make([]User, out.Count)
	for idx, item := range out.Items {

		connId := item["connectionId"]
		var connIdStr string
		switch v := connId.(type) {
		case *dynamodbtypes.AttributeValueMemberS:
			connIdStr = v.Value
		default:
			connIdStr = ""
		}

		username := item["username"]
		var usernameStr string
		switch v := username.(type) {
		case *dynamodbtypes.AttributeValueMemberS:
			usernameStr = v.Value
		default:
			usernameStr = ""
		}

		if connIdStr == "" || usernameStr == "" {
			return []User{}, errors.New("could not decode response")
		}

		au[idx] = User{
			ConnectionID: connIdStr,
			Username:     usernameStr,
		}
	}

	return au, nil
}
