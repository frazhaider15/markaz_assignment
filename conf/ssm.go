package conf

import (
	"context"
	"fmt"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var client *ssm.Client

// SSMGetParameterAPI defines the interface for the GetParameter function.
// We use this interface to test the function using a mocked service.
type SSMGetParameterAPI interface {
	GetParameter(ctx context.Context,
		params *ssm.GetParameterInput,
		optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

// FindParameter retrieves an AWS Systems Manager string parameter
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region
//	api is the interface that defines the method call
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a GetParameterOutput object containing the result of the service call and nil
//	Otherwise, nil and an error from the call to GetParameter
func FindParameter(c context.Context, api SSMGetParameterAPI, input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	return api.GetParameter(c, input)

}

// InItAWSssm init aws ssm sdk for getting values from AWS vault
func InItAWSssm() {
	cfg, err := awsConfig.LoadDefaultConfig(
		context.Background(),
		awsConfig.WithRegion(Config.AWS.AWSRegion),
	)
	if err != nil {
		fmt.Println("configuration error, " + err.Error())
	}

	client = ssm.NewFromConfig(cfg)
}

func GetValueFromSSMByKey(key string, withDecryptn bool) string {
	input := &ssm.GetParameterInput{
		Name:           &key,
		WithDecryption: &withDecryptn,
	}
	results, err := FindParameter(context.TODO(), client, input)
	if err != nil {
		fmt.Printf("error getting key: %v from ssm: %v", key, err.Error())
		return ""
	}
	return *results.Parameter.Value
}
