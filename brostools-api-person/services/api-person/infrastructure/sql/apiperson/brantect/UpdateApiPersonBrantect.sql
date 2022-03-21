UPDATE
    mst_person
SET 
	dept_cd = $1,
	person_cd = $2,
	executive = $3,
	person_nm = $4,
	person_name_kana = $5,
	zip = $6,
	address = $7,
	tel = $8,
	fax = $9,
	email = $10,
	mobile = $11,
	remarks = $12,
	search_person_nm = $13,
	type = $14,
	idno = $15,
    upd_date = NOW(),
    upd_user = $16
WHERE
    client_cd = $17
AND 
    person_cd = $18