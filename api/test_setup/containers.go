// Package testing_setup provides ways of creating containers for integration tests
package tests_setup

import (
	"context"
	"fmt"
	"strconv"

	"github.com/looped-dev/cms/api/emails"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	mail "github.com/xhit/go-simple-mail/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestContainers struct {
	Pool *dockertest.Pool
}

// NewContainer creates a new container for running integration tests against.
// This container can be any docker container that can be pulled and run so
// tests can run tests against it.
func (t TestContainers) NewContainer(runOptions *dockertest.RunOptions, retryFunction func(resource *dockertest.Resource) func() error) (*dockertest.Resource, error) {
	pool := t.Pool
	hcOpts := func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	}
	resource, err := pool.RunWithOptions(runOptions, hcOpts)
	if err != nil {
		return nil, fmt.Errorf("Could not start resource: %s", err)
	}
	// set timeout to 5 minutes
	if err := resource.Expire(300); err != nil {
		return nil, fmt.Errorf("Couldn't setup resource expiration to 5 minutes: %v", err)
	}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(retryFunction(resource)) // use a closure to pass the resource to the retry function
	if err != nil {
		return nil, fmt.Errorf("Could not connect to docker: %s", err)
	}
	return resource, nil
}

// NewMongoContainer create a new mongodb container for running integration
// tests against.
func (t TestContainers) NewMongoContainer(ctx context.Context) (*mongo.Client, *dockertest.Resource, error) {
	var db *mongo.Client
	// resource will be passed in by the caller
	retryFunction := func(resource *dockertest.Resource) func() error {
		return func() error {
			var err error
			connectionURL := options.Client().ApplyURI(
				fmt.Sprintf("mongodb://looped:root@localhost:%s", resource.GetPort("27017/tcp")),
			)
			db, err = mongo.Connect(ctx, connectionURL)
			if err != nil {
				return err
			}
			return db.Ping(ctx, nil)
		}
	}
	runOptions := &dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "5.0",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=looped",
			"MONGO_INITDB_ROOT_PASSWORD=root",
		},
	}
	resource, err := t.NewContainer(runOptions, retryFunction)
	return db, resource, err
}

func (t TestContainers) NewMailTestServer(ctx context.Context) (*mail.SMTPClient, *dockertest.Resource, error) {
	var smtpClient *mail.SMTPClient
	// resource will be passed in by the caller
	retryFunction := func(resource *dockertest.Resource) func() error {
		return func() error {
			port, err := strconv.Atoi(resource.GetPort("1025/tcp"))
			if err != nil {
				return err
			}
			smtpClient, err = emails.NewMockSMTPClient("localhost", port)
			if err != nil {
				return err
			}
			return smtpClient.Noop()
		}
	}
	runOptions := &dockertest.RunOptions{
		Repository: "mailhog/mailhog",
		Tag:        "latest",
		Env:        []string{},
	}
	resource, err := t.NewContainer(runOptions, retryFunction)
	return smtpClient, resource, err
}
