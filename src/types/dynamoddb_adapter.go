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
// It expects a DynamoDB table with a string-valued partition key "username".
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
	in := dynamodb.GetItemInput{
		TableName: aws.String(dba.TableName),
		Key: map[string]dynamodbtypes.AttributeValue{
			"username": &dynamodbtypes.AttributeValueMemberS{
				Value: username,
			},
		},
		ConsistentRead: aws.Bool(true),
	}

	out, err := ddbSvc.GetItem(ctx, &in)
	if err != nil {
		return err
	}

	if len(out.Item) <= 0 {
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

	if dba.User.Username != "" {
		in.FilterExpression = aws.String(
			"username <> :val",
		)
		in.ExpressionAttributeValues = map[string]dynamodbtypes.AttributeValue{
			":val": &dynamodbtypes.AttributeValueMemberS{
				Value: dba.User.Username,
			},
		}
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

// Disconnect disconnects current User by deleting the user from DB
func (dba *DynamoDBAdapter) Disconnect(ctx context.Context) error {
	if dba.User.Username == "" {
		return errors.New("no connected")
	}
	in := &dynamodb.DeleteItemInput{
		TableName: &dba.TableName,
		Key: map[string]dynamodbtypes.AttributeValue{
			"username": &dynamodbtypes.AttributeValueMemberS{
				Value: dba.User.Username,
			},
		},
	}

	_, err := ddbSvc.DeleteItem(ctx, in)
	if err != nil {
		return err
	}

	dba.User = User{}

	return nil
}
