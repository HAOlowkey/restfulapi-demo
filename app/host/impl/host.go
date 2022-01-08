package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/HAOlowkey/restfulapi-demo/app/host"
)

func (i *impl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	if err := ins.Validate(); err != nil {
		return nil, err
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
			err := tx.Rollback()
			i.log.Debugf("tx rollback error,%s", err)
		} else {
			err := tx.Commit()
			i.log.Debugf("tx commit error,%s", err)
		}
	}()

	resStmt, err = tx.Prepare(insertResourceSQL)
	if err != nil {
		return nil, fmt.Errorf("prepare resource sql error, %v", err)
	}

	defer resStmt.Close()

	_, err = resStmt.Exec(ins.Id, ins.Vendor, ins.Region, ins.Zone, ins.CreateAt, ins.ExpireAt, ins.Category, ins.InstanceId,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.SyncAccount, ins.PublicIP,
		ins.PrivateIP, ins.PayType, ins.ResourceHash, ins.DescribeHash)
	if err != nil {
		return nil, fmt.Errorf("execute resource sql error, %v", err)
	}

	descStmt, err = tx.Prepare(insertDescribeSQL)
	if err != nil {
		return nil, fmt.Errorf("prepare describe sql error, %v", err)
	}

	defer descStmt.Close()

	_, err = descStmt.Exec(ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec, ins.OSType, ins.OSName,
		ins.SerialNumber, ins.ImageID, ins.InternetMaxBandwidthOut, ins.InternetMaxBandwidthIn,
		ins.KeyPairName, ins.SecurityGroups)

	if err != nil {
		return nil, fmt.Errorf("execute describe sql error, %v", err)
	}

	return ins, nil
}

func (i *impl) QueryHost(ctx context.Context, q *host.QueryHostRequest) (*host.Set, error) {
	i.log.Debugf("sql: %s args: %s %s", queryHostSQL, q.Offset(), q.PageSize)

	stmt, err := i.db.Prepare(queryHostSQL)

	if err != nil {
		return nil, fmt.Errorf("prepare query host sql error, %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(q.Offset(), q.PageSize)
	if err != nil {
		return nil, err
	}

	set := host.NewDefaultSet()

	for rows.Next() {
		ins := host.NewDefaultHost()
		err := rows.Scan(&ins.Id, &ins.Vendor, &ins.Region, &ins.Zone, &ins.CreateAt, &ins.ExpireAt,
			&ins.Category, &ins.Type, &ins.InstanceId, &ins.Name,
			&ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
			&ins.PublicIP, &ins.PrivateIP, &ins.PayType, &ins.ResourceHash, &ins.DescribeHash,
			&ins.Id, &ins.CPU,
			&ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName,
			&ins.SerialNumber, &ins.ImageID, &ins.InternetMaxBandwidthOut, &ins.InternetMaxBandwidthIn,
			&ins.KeyPairName, &ins.SecurityGroups)
		if err != nil {
			return nil, err
		}
		set.Add(ins)
	}
	return set, nil
}

func (ins *impl) DescribeHost(context.Context, *host.DescribeHostRequest) (*host.Host, error) {
	return nil, nil
}

func (ins *impl) UpdateHost(context.Context, *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (ins *impl) DeleteHost(context.Context, *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
