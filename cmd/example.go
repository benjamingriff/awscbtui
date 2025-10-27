package main

import (
	"context"
	"fmt"
	"log" 
	"time"
	sdkaws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/sts"

)

func main() {
	ctx := context.Background()

	var opts []func(*config.LoadOptions) error

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	stscli := sts.NewFromConfig(cfg)
	idOut, err := stscli.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("get caller identity: %v", err)
	}

	cred, credErr := cfg.Credentials.Retrieve(ctx)
	exp := "n/a"
	if credErr == nil && cred.CanExpire {
		exp = cred.Expires.Local().Format(time.RFC3339)
	}
	
	fmt.Printf("Region: %s\n", cfg.Region)
	fmt.Printf("Account: %s\n", sdkaws.ToString(idOut.Account))
	fmt.Printf("ARN: %s\n", sdkaws.ToString(idOut.Arn))
	fmt.Printf("Creds expire: %s\n\n", exp)

	// CodeBuild projects (first page)
	cb := codebuild.NewFromConfig(cfg)
	out, err := cb.ListProjects(ctx, &codebuild.ListProjectsInput{})
	if err != nil {
		log.Fatalf("list projects: %v", err)
	}
	fmt.Printf("Projects (showing up to 20):\n")
	for i, name := range out.Projects {
		if i >= 20 {
			fmt.Println("(more available...)")
			break
		}
		fmt.Printf("%2d. %s\n", i+1, name)
	}
}
