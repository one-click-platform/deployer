package handlers

import (
	"context"
	"github.com/one-click-platform/deployer/internal/data"
	"net/http"

	"github.com/one-click-platform/deployer/resources"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	githubKeyCtxKey
	storageCtxKey
	tasksCtxKey
	accountsQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxGithubKey(entry string) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, githubKeyCtxKey, entry)
	}
}

func GithubKey(r *http.Request) string {
	return r.Context().Value(githubKeyCtxKey).(string)
}

func CtxStorage(entry map[string]resources.EnvConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, storageCtxKey, entry)
	}
}

func Storage(r *http.Request) map[string]resources.EnvConfig {
	return r.Context().Value(storageCtxKey).(map[string]resources.EnvConfig)
}

func CtxTasks(entry chan string) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, tasksCtxKey, entry)
	}
}

func Tasks(r *http.Request) chan string {
	return r.Context().Value(tasksCtxKey).(chan string)
}

func CtxAccountsQ(entry data.AccountsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, accountsQCtxKey, entry)
	}
}

func AccountsQ(r *http.Request) data.AccountsQ {
	return r.Context().Value(accountsQCtxKey).(data.AccountsQ).New()
}
