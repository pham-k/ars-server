package kdb

import (
	_ "github.com/aws/aws-sdk-go-v2/aws"
)

//type keyValStore struct {
//	Client *dynamodb.Client
//	logger  logger.logger
//}
//
//func New(cfg aws.config, logger logger.logger) KeyValStore {
//	pool := dynamodb.NewFromConfig(cfg)
//
//	return &keyValStore{
//		Client: pool,
//	}
//}
