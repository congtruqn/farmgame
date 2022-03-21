SELECT 
  COALESCE(client_id, '') as client_id
  ,COALESCE(client_cd, '') as client_cd
  ,COALESCE(audit_flg, '') as audit_flg
  ,COALESCE(pwd, '') as pwd
  ,COALESCE(nr_flg, '') as nr_flg
  ,COALESCE(db_ip::TEXT, '') as db_ip
  ,COALESCE(file_dir_path, '') as file_dir_path
  ,COALESCE(db_ip_global::TEXT, '') as db_ip_global
  ,COALESCE(de_flg, '') as de_flg
  ,COALESCE(de_auto_no_flg, '') as de_auto_no_flg
  ,COALESCE(de_csv_dl_flg, '') as de_csv_dl_flg
  ,COALESCE(de_cond_save_flg, '') as de_cond_save_flg
  ,COALESCE(de_data_set_flg, '') as de_data_set_flg
  ,COALESCE(tm_auto_no_flg, '') as tm_auto_no_flg
  ,COALESCE(tm_csv_dl_flg, '') as tm_csv_dl_flg
  ,COALESCE(tm_cond_save_flg, '') as tm_cond_save_flg
  ,COALESCE(tm_data_set_flg, '') as tm_data_set_flg
  ,COALESCE(tm_area_search_flg, '') as tm_area_search_flg
  ,COALESCE(tm_date_cst_flg, '') as tm_date_cst_flg
  ,COALESCE(tm_img_at_flg, '') as tm_img_at_flg
  ,COALESCE(tm_task_action_flg, '') as tm_task_action_flg
  ,COALESCE(tm_disp_mst_flg, '') as tm_disp_mst_flg
  ,COALESCE(tm_madrid_prot_set_flg, '') as tm_madrid_prot_set_flg
  ,COALESCE(lm_flg, '') as lm_flg
  ,COALESCE(lm_api_status, '') as lm_api_status
  ,COALESCE(lm_api_ip, '') as lm_api_ip
  ,COALESCE(mo_url_num_flg, '') as mo_url_num_flg
  ,COALESCE(mo_his_save_period_flg, '') as mo_his_save_period_flg
  ,COALESCE(mo_auto_getting_flg, '') as mo_auto_getting_flg
  ,COALESCE(time_stamp_use_flg, '') as time_stamp_use_flg
  ,COALESCE(time_stamp_licence_flg, '') as time_stamp_licence_flg
  ,COALESCE(tm_el_app_flg, '') as tm_el_app_flg
  ,COALESCE(tm_ipdl_sync_flg, '') as tm_ipdl_sync_flg
  ,COALESCE(de_hague_set_flg, '') as de_hague_set_flg
  ,COALESCE(dn_date_cst_flg, '') as dn_date_cst_flg
  ,COALESCE(db_name, '') as db_name
  ,COALESCE(json_ip_arrow, '') as json_ip_arrow
  ,COALESCE(ip_arrow_flg, '') as ip_arrow_flg
  ,COALESCE(last_password_change_days::TEXT, '') as last_password_change_days
  ,COALESCE(last_password_change_days_flg, '') as last_password_change_days_flg
  ,COALESCE(last_password_change_unit, '') as last_password_change_unit
  ,COALESCE(sign_up_flg, '') as sign_up_flg
  ,COALESCE(sign_up_auth_key, '') as sign_up_auth_key
  ,COALESCE(json_domain_mail, '') as json_domain_mail
  ,COALESCE(nm_mng_use_flag, '') as nm_mng_use_flag
  ,COALESCE(tm_mng_use_flag, '') as tm_mng_use_flag
  ,COALESCE(wm_mng_use_flag, '') as wm_mng_use_flag
  ,COALESCE(db_instance, '') as db_instance
FROM 
  mst_client_id
WHERE
  client_cd = $1