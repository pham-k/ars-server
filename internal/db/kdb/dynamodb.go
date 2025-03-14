package kdb

import (
	_ "github.com/aws/aws-sdk-go-v2/aws"
)

//type kDB struct {
//	Pool *dynamodb.Client
//	Log  log.Log
//}
//
//func NewKDB(cfg aws.Config, log log.Log) KDB {
//	pool := dynamodb.NewFromConfig(cfg)
//
//	return &kDB{
//		Pool: pool,
//	}
//}
