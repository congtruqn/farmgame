SELECT
    client_cd
    ,person_cd
    , COALESCE(person_field,"") as person_field
    , COALESCE(division_nm,"") as division_nm
    , COALESCE(position_nm,"") as position_nm
    , COALESCE(person_nm,"") as person_nm
    , COALESCE(person_nm_jp,"") as person_nm_jp
    , COALESCE(post_cd,"") as post_cd
    , COALESCE(address,"") as address
    , COALESCE(tel,"") as tel
    , COALESCE(fax,"") as fax
    , COALESCE(email,"") as email
    , COALESCE(moby,"") as moby
    , COALESCE(remarks,"") as remarks
    , COALESCE(invsnd_dn_flg,"") as invsnd_dn_flg
    , COALESCE(invsnd_tm_flg,"") as invsnd_tm_flg
    , COALESCE(invsnd_ot_flg,"") as invsnd_ot_flg
    , COALESCE(sign_flg,"") as sign_flg
    , COALESCE(delete_flg,"") as delete_flg
    , COALESCE(upd_date,"") as upd_date
    , COALESCE(upd_user,"") as upd_user
    , COALESCE(upd_prg_id,"") as upd_prg_id
    , COALESCE(inp_date,"") as inp_date
    , COALESCE(inp_user,"") as inp_user
    , COALESCE(inp_prg_id,"") as inp_prg_id
FROM
    t_m_person
WHERE
      client_cd  = ?
      AND
      person_cd = ?
      AND
      delete_flg = '0'
