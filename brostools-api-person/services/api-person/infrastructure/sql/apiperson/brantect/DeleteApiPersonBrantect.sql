UPDATE
    mst_person
SET 
	delete_flg = '1',
    upd_date = NOW(),
    upd_user = $1
WHERE
    client_cd = $2
AND 
    person_cd = $3