UPDATE
    t_m_person
SET
	person_cd = ?,
	person_field = ?,
	division_nm = ?,
	position_nm = ?,
	person_nm = ?,
	person_nm_jp = ?,
	post_cd = ?,
	address = ?,
	tel = ?,
	fax = ?,
	email = ?,
	moby = ?,
	remarks = ?,
	invsnd_dn_flg = ?,
	invsnd_tm_flg = ?,
	invsnd_ot_flg = ?,
	sign_flg = ?,
    upd_date = NOW(),
    upd_user = ?,
    upd_prg_id = ?
WHERE 
    client_cd = ?
AND
    person_cd = ?