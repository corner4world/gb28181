package configdb

import (
	"context"
	"testing"

	"github.com/gowvp/gb28181/internal/core/config"
	"github.com/ixugo/goddd/pkg/orm"
)

func TestConfigGet(t *testing.T) {
	db, mock, err := generateMockDB()
	if err != nil {
		t.Fatal(err)
	}
	userDB := NewConfig(db)

	mock.ExpectQuery(`SELECT \* FROM "configs" WHERE id=\$1 (.+) LIMIT \$2`).WithArgs("jack", 1)
	var out config.Config
	if err := userDB.Get(context.Background(), &out, orm.Where("id=?", "jack")); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("ExpectationsWereMet err:", err)
	}
}
