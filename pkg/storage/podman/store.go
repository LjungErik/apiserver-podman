package podman

import (
	"context"
	"io"
	"log/slog"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/storage"
)

type store struct {
	logger *slog.Logger
}

func NewStore(w io.Writer) storage.Interface {
	return &store{
		logger: slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	}
}

func (s *store) Versioner() storage.Versioner {
	return nil
}

func (s *store) Create(ctx context.Context, key string, obj, out runtime.Object, ttl uint64) error {
	s.logger.Info("Create", "key", key, "obj", obj)

	return nil
}

func (s *store) Delete(
	ctx context.Context, key string, out runtime.Object, preconditions *storage.Preconditions,
	validateDeletion storage.ValidateObjectFunc, cachedExistingObject runtime.Object, opts storage.DeleteOptions) error {
	s.logger.Info("Delete", "key", key)

	return nil
}

func (s *store) Watch(ctx context.Context, key string, opts storage.ListOptions) (watch.Interface, error) {
	s.logger.Info("Watch", "key", key)

	return nil, nil
}

func (s *store) Get(ctx context.Context, key string, opts storage.GetOptions, objPtr runtime.Object) error {
	s.logger.Info("Get", "key", key)

	return nil
}

func (s *store) GetList(ctx context.Context, key string, opts storage.ListOptions, listObj runtime.Object) error {
	s.logger.Info("GetList", "key", key)

	return nil
}

func (s *store) GuaranteedUpdate(
	ctx context.Context, key string, destination runtime.Object, ignoreNotFound bool,
	preconditions *storage.Preconditions, tryUpdate storage.UpdateFunc, cachedExistingObject runtime.Object) error {
	s.logger.Info("GuaranteedUpdate", "key", key, "destination", destination)

	return nil
}

func (s *store) Count(key string) (int64, error) {
	s.logger.Info("Count", "key", key)

	return 0, nil
}

func (s *store) ReadinessCheck() error {
	s.logger.Info("ReadinessCheck")

	return nil
}

func (s *store) RequestWatchProgress(ctx context.Context) error {
	s.logger.Info("RequestWatchProgress")

	return nil
}

func (s *store) GetCurrentResourceVersion(ctx context.Context) (uint64, error) {
	s.logger.Info("GetCurrentResourceVersion")

	return 0, nil
}
