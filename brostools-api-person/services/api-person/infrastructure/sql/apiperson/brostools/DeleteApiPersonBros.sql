UPDATE
    t_m_person
SET
	delete_flg = '1',
    upd_date = NOW(),
    upd_user = ?,
    upd_prg_id = ?
WHERE 
    client_cd = ?
AND
    person_cd = ?