SELECT
    client_cd
    ,COALESCE(dept_cd,'')
    ,person_cd
    ,COALESCE(executive,'')
    ,COALESCE(person_nm,'')
    ,COALESCE(person_name_kana,'')
    ,COALESCE(zip,'')
    ,COALESCE(address,'')
    ,COALESCE(tel,'')
    ,COALESCE(fax,'')
    ,COALESCE(email,'')
    ,COALESCE(mobile,'')
    ,COALESCE(remarks,'')
    ,COALESCE(search_person_nm,'')
    ,COALESCE(type,'')
    ,COALESCE(idno,'')
    ,COALESCE((last_verified_date)::TEXT, '')
    ,COALESCE(delete_flg,'')
    ,COALESCE((upd_date)::TEXT, '')
    ,COALESCE(upd_user,'')
    ,COALESCE((inp_date)::TEXT, '')
    ,COALESCE(inp_user,'')
FROM
    mst_person
WHERE
      client_cd  = $1
      AND
      person_cd = $2
      AND
      delete_flg = '0'
