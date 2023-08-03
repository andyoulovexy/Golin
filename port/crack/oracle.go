package crack

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	"time"
)

func oraclecon(ctx context.Context, cancel context.CancelFunc, ip, user, passwd string, port, timeout int) {
	defer func() {
		wg.Done()
		<-ch
	}()
	select {
	case <-ctx.Done():
		return
	default:
	}

	dataSourceName := fmt.Sprintf("oracle://%s:%s@%s:%d/orcl", user, passwd, ip, port)
	db, err := sql.Open("oracle", dataSourceName)
	if err == nil {
		db.SetConnMaxLifetime(time.Duration(timeout) * time.Second)
		db.SetConnMaxIdleTime(time.Duration(timeout) * time.Second)
		db.SetMaxIdleConns(0)
		defer db.Close()
		err = db.Ping()
		if err == nil {
			end(ip, user, passwd, port, "oracle")
			cancel()
		}
	}

}