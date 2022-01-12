<template>
<div style="position: relative;width: auto; height: 100%;" id="left_menu">
  <div data-height="100%" data-scrollbar="true" style="width: auto; height: 100%;" data-init="true">
    <ul class="nav">
      <li class="has-sub active">
        <a href="#">
          <i class="fa fa-gears"></i>
            <span>Quản lý khách hàng</span>
          <b class="caret pull-right"></b>
        </a>
        <ul class="sub-menu">
          <li><a href="/listproject">Quản lý dự án</a></li>
          <li><a href="/toquote">Quản lý báo giá</a></li> 
        </ul>
      </li>
      <li class="has-sub">
        <a href="#">
          <i class="fa fa-gears"></i>
            <span>Quản lý tài khoản</span>
          <b class="caret pull-right"></b>
        </a>
        <ul class="sub-menu">
          <li><a href="/users">Quản lý người dùng</a></li>
          <li><a href="/administrator/listuseraction">Quản lí tác vụ</a></li>
          <li><a href="/administrator/listgroupuser">Quản lí Nhóm Quyền</a></li>
        </ul>
      </li>
    </ul>
  </div>
</div>
</template>