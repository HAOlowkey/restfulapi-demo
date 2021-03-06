package impl

const (
	insertResourceSQL = `INSERT INTO resource (
id,vendor,region,zone,create_at,expire_at,category,type,instance_id,
name,description,status,update_at,sync_at,sync_accout,public_ip,
private_ip,pay_type,resource_hash,describe_hash) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	insertDescribeSQL = `INSERT INTO host (
resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
serial_number,image_id,internet_max_bandwidth_out,
internet_max_bandwidth_in,key_pair_name,security_groups) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
	queryHostSQL        = `SELECT * FROM resource LEFT JOIN host ON resource.id = host.resource_id ORDER BY create_at DESC LIMIT ?,?`
	queryHostSQLKeyWord = `SELECT * FROM resource LEFT JOIN host ON resource.id = host.resource_id WHERE name like ? ORDER BY create_at DESC LIMIT ?,?`
	describeHostSQL     = `SELECT * FROM resource LEFT JOIN host ON resource.id = host.resource_id WHERE id=?`

	updateResourceSQL = `UPDATE resource SET vendor=?,region=?,zone=?,expire_at=?,name=?,description=? WHERE id = ?`

	updateDescribeSQL = `UPDATE host SET cpu=?,memory=? WHERE resource_id = ?`

	deleteResourceSQL = `DELETE FROM resource where id = ?`
	deleteDescribeSQL = `DELETE FROM host where resource_id = ?`
)
