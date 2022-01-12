<template>
  <div id="wapper">
    <div class="header">
      <div class="container">
        <div class="logo-header">
          <a href="/" class="navbar-brand"><img src="@/assets/images/logo.jpg"></a>
        </div>
        <div class="header_slogan">
          <h1>Ứng dụng đọc hoá đơn điện tử</h1>
          <p>Ứng dụng kiểm tra tính hợp lệ của hóa đơn điện tử. Và kiểm tra thông tin, nội dung của hóa đơn điện tử định dạng xml</p>
        </div>
        <div id="clear"></div>
      </div>
    </div>
    <div class="container-login">
      <router-view />
    </div>
  </div>
</template>

<script>
  export default {
    name: 'MyNameComponent'   // y
  }
</script>