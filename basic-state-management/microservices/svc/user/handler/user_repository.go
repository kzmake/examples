package handler

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type UserRepository interface {
	List(ctx context.Context) ([]*AggregateRootUser, error)
	Get(ctx context.Context, id UserID) (*AggregateRootUser, error)
	Append(ctx context.Context, user *AggregateRootUser) error
	Remove(ctx context.Context, user *AggregateRootUser) error
	Update(ctx context.Context, user *AggregateRootUser) error
}

type userRepository struct{}

func NewUserRepository() UserRepository { return &userRepository{} }

func (r *userRepository) List(ctx context.Context) ([]*AggregateRootUser, error) {
	keys, err := r.keys(ctx)
	if err != nil {
		return nil, err
	}

	users, err := r.find(ctx, keys)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Get(ctx context.Context, id UserID) (*AggregateRootUser, error) {
	keys := []string{string(id)}

	users, err := r.find(ctx, keys)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.Errorf("the user(id: %s) not found", id)
	}

	if len(users) > 1 {
		return nil, errors.Errorf("multiple users found")
	}

	return users[0], nil
}

func (r *userRepository) Append(ctx context.Context, user *AggregateRootUser) error {
	client, err := dapr.NewClient()
	if err != nil {
		return err
	}

	key := string(user.ID)
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}

	keys, err := r.keys(ctx)
	if err != nil {
		return err
	}

	appendedKeys := append(keys, key)
	keysValue, err := json.Marshal(appendedKeys)
	if err != nil {
		return err
	}

	if err := client.SaveBulkState(ctx, "state-user", []*dapr.SetStateItem{
		{Key: "__keys__", Value: keysValue},
		{Key: key, Value: value},
	}...); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Remove(ctx context.Context, user *AggregateRootUser) error {
	client, err := dapr.NewClient()
	if err != nil {
		return err
	}

	key := string(user.ID)
	keys, err := r.keys(ctx)
	if err != nil {
		return err
	}

	removedKeys := []string{}
	for _, v := range keys {
		if v == key {
			continue
		}
		removedKeys = append(removedKeys, v)
	}

	keysValue, err := json.Marshal(removedKeys)
	if err != nil {
		return err
	}

	if err := client.SaveState(ctx, "state-user", "__keys__", keysValue); err != nil {
		return err
	}
	if err := client.DeleteState(ctx, "state-user", key); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, user *AggregateRootUser) error {
	client, err := dapr.NewClient()
	if err != nil {
		return err
	}

	key := string(user.ID)
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := client.SaveState(ctx, "state-user", key, value); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) keys(ctx context.Context) ([]string, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	keys := []string{}
	keysItem, err := client.GetState(ctx, "state-user", "__keys__")
	if err != nil {
		return nil, err
	}

	if keysItem != nil && len(keysItem.Value) != 0 {
		if err := json.Unmarshal(keysItem.Value, &keys); err != nil {
			return nil, err
		}
	}

	return keys, nil
}

func (r *userRepository) find(ctx context.Context, keys []string) ([]*AggregateRootUser, error) {
	if len(keys) == 0 {
		return []*AggregateRootUser{}, nil
	}

	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	items, err := client.GetBulkState(ctx, "state-user", keys, nil, 100)
	if err != nil {
		return nil, err
	}

	users := []*AggregateRootUser{}
	for _, item := range items {
		if item != nil && len(item.Value) != 0 {
			user := &AggregateRootUser{}
			if err := json.Unmarshal(item.Value, user); err != nil {
				log.Warn().Msgf("failed to unmarshal: %v", err)
				continue
			}

			users = append(users, user)
		}
	}

	return users, nil
}
