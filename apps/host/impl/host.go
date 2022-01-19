package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/HAOlowkey/restfulapi-demo/apps/host"
)

func (i *impl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {

	if err := ins.Validate(); err != nil {
		return nil, err
	}

	ins.Resource.Id = xid.New().String()
	if ins.Resource.CreateAt == 0 {
		ins.Resource.CreateAt = ftime.Now().Timestamp()
	}

	var (
		resStmt  *sql.Stmt
		descStmt *sql.Stmt
		err      error
	)

	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				i.log.Debugf("tx rollback error,%v", err)
			}
		} else {
			err = tx.Commit()
			if err != nil {
				i.log.Debugf("tx commit error,%v", err)
			}
		}
	}()

	resStmt, err = tx.Prepare(insertResourceSQL)
	if err != nil {
		return nil, fmt.Errorf("prepare resource sql error, %v", err)
	}

	defer resStmt.Close()

	_, err = resStmt.Exec(ins.Resource.Id, ins.Resource.Vendor, ins.Resource.Region, ins.Resource.Zone, ins.Resource.CreateAt, ins.Resource.ExpireAt, ins.Resource.Category, ins.Resource.Type, ins.Resource.InstanceId,
		ins.Resource.Name, ins.Resource.Description, ins.Resource.Status, ins.Resource.UpdateAt, ins.Resource.SyncAt, ins.Resource.SyncAccount, ins.Resource.PublicIp,
		ins.Resource.PrivateIp, ins.Resource.PayType, ins.ResourceHash, ins.DescribeHash)
	if err != nil {
		return nil, fmt.Errorf("execute resource sql error, %v", err)
	}

	descStmt, err = tx.Prepare(insertDescribeSQL)
	if err != nil {
		return nil, fmt.Errorf("prepare describe sql error, %v", err)
	}

	defer descStmt.Close()

	_, err = descStmt.Exec(ins.Resource.Id, ins.Describe.Cpu, ins.Describe.Memory, ins.Describe.GpuAmount, ins.Describe.GpuSpec, ins.Describe.OsType, ins.Describe.OsName,
		ins.Describe.SerialNumber, ins.Describe.ImageId, ins.Describe.InternetMaxBandwidthOut, ins.Describe.InternetMaxBandwidthIn,
		ins.Describe.KeyPairName, ins.Describe.SecurityGroups)

	if err != nil {
		return nil, fmt.Errorf("execute describe sql error, %v", err)
	}

	return ins, nil
}

func (i *impl) QueryHost(ctx context.Context, q *host.QueryHostRequest) (*host.Set, error) {
	var (
		err  error
		stmt *sql.Stmt
		rows *sql.Rows
	)
	if q.Keywords == "" {
		i.log.Debugf("sql: %s args: %s %s", queryHostSQL, strconv.Itoa(int(q.Offset())), strconv.Itoa(int(q.PageSize)))
		stmt, err = i.db.Prepare(queryHostSQL)
	} else {
		i.log.Debugf("sql: %s args: %s %s %s", queryHostSQLKeyWord, q.Keywords, strconv.Itoa(int(q.Offset())), strconv.Itoa(int(q.PageSize)))
		stmt, err = i.db.Prepare(queryHostSQLKeyWord)
	}

	if err != nil {
		return nil, fmt.Errorf("prepare query host sql error, %v", err)
	}
	defer stmt.Close()

	if q.Keywords == "" {
		rows, err = stmt.Query(strconv.Itoa(int(q.Offset())), strconv.Itoa(int(q.PageSize)))
	} else {
		rows, err = stmt.Query("%"+q.Keywords+"%", strconv.Itoa(int(q.Offset())), strconv.Itoa(int(q.PageSize)))
	}
	if err != nil {
		return nil, err
	}

	set := host.NewDefaultSet()

	for rows.Next() {
		ins := host.NewDefaultHost()
		err := rows.Scan(&ins.Resource.Id, &ins.Resource.Vendor, &ins.Resource.Region, &ins.Resource.Zone, &ins.Resource.CreateAt, &ins.Resource.ExpireAt,
			&ins.Resource.Category, &ins.Resource.Type, &ins.Resource.InstanceId, &ins.Resource.Name,
			&ins.Resource.Description, &ins.Resource.Status, &ins.Resource.UpdateAt, &ins.Resource.SyncAt, &ins.Resource.SyncAccount,
			&ins.Resource.PublicIp, &ins.Resource.PrivateIp, &ins.Resource.PayType, &ins.ResourceHash, &ins.DescribeHash,
			&ins.Resource.Id, &ins.Describe.Cpu,
			&ins.Describe.Memory, &ins.Describe.GpuAmount, &ins.Describe.GpuSpec, &ins.Describe.OsType, &ins.Describe.OsName,
			&ins.Describe.SerialNumber, &ins.Describe.ImageId, &ins.Describe.InternetMaxBandwidthOut, &ins.Describe.InternetMaxBandwidthIn,
			&ins.Describe.KeyPairName, &ins.Describe.SecurityGroups)
		if err != nil {
			return nil, err
		}
		set.Add(ins)
	}
	return set, nil
}

func (i *impl) DescribeHost(ctx context.Context, q *host.DescribeHostRequest) (*host.Host, error) {

	i.log.Debugf("sql: %s args: %s", describeHostSQL, q.Id)
	stmt, err := i.db.Prepare(describeHostSQL)

	if err != nil {
		return nil, fmt.Errorf("prepare query host sql error, %v", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(q.Id)

	ins := host.NewDefaultHost()

	err = row.Scan(&ins.Resource.Id, &ins.Resource.Vendor, &ins.Resource.Region, &ins.Resource.Zone, &ins.Resource.CreateAt, &ins.Resource.ExpireAt,
		&ins.Resource.Category, &ins.Resource.Type, &ins.Resource.InstanceId, &ins.Resource.Name,
		&ins.Resource.Description, &ins.Resource.Status, &ins.Resource.UpdateAt, &ins.Resource.SyncAt, &ins.Resource.SyncAccount,
		&ins.Resource.PublicIp, &ins.Resource.PrivateIp, &ins.Resource.PayType, &ins.ResourceHash, &ins.DescribeHash,
		&ins.Resource.Id, &ins.Describe.Cpu,
		&ins.Describe.Memory, &ins.Describe.GpuAmount, &ins.Describe.GpuSpec, &ins.Describe.OsType, &ins.Describe.OsName,
		&ins.Describe.SerialNumber, &ins.Describe.ImageId, &ins.Describe.InternetMaxBandwidthOut, &ins.Describe.InternetMaxBandwidthIn,
		&ins.Describe.KeyPairName, &ins.Describe.SecurityGroups)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("host %s not found", q.Id)
		}
		return nil, err
	}

	return ins, nil
}

func (i *impl) UpdateHost(ctx context.Context, q *host.UpdateHostRequest) (*host.Host, error) {

	ins, err := i.DescribeHost(ctx, &host.DescribeHostRequest{Id: q.Host.Resource.Id})
	if err != nil {
		return nil, err
	}

	switch q.UpdateMode {
	case host.UpdateMode_PUT:
		ins.PutUpdate(q.Host.Resource, q.Host.Describe)
	case host.UpdateMode_PATCH:
		err := ins.PatchUpdate(q.Host.Resource, q.Host.Describe)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("not support %v method", q.UpdateMode)
	}

	if err := ins.Validate(); err != nil {
		return nil, err
	}

	stmt, err := i.db.Prepare(updateResourceSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// DML
	// vendor=?,region=?,zone=?,expire_at=?,name=?,description=? WHERE id = ?
	_, err = stmt.Exec(ins.Resource.Vendor, ins.Resource.Region, ins.Resource.Zone, ins.Resource.ExpireAt, ins.Resource.Name, ins.Resource.Description, ins.Resource.Id)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) DeleteHost(ctx context.Context, q *host.DeleteHostRequest) (*host.Host, error) {

	var (
		resstmt  *sql.Stmt
		descstmt *sql.Stmt
		err      error
	)

	ins, err := i.DescribeHost(ctx, &host.DescribeHostRequest{Id: q.Id})
	if err != nil {
		return nil, err
	}
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				i.log.Debugf("tx rollback error,%v", err)
			}
		} else {
			err = tx.Commit()
			if err != nil {
				i.log.Debugf("tx commit error,%v", err)
			}
		}
	}()

	resstmt, err = i.db.Prepare(deleteResourceSQL)
	if err != nil {
		return nil, err
	}
	defer resstmt.Close()
	_, err = resstmt.Exec(q.Id)
	if err != nil {
		return nil, err
	}

	descstmt, err = i.db.Prepare(deleteDescribeSQL)
	if err != nil {
		return nil, err
	}
	defer descstmt.Close()
	_, err = descstmt.Exec(q.Id)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
