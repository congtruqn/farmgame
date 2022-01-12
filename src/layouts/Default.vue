<template>
    <div class="wrapper" id="wrapper">
      <header class="header" id="header">
        <div class="container">
          <AdminHeader />
        </div>
      </header>

      <div class="content content1">
        <div class="container">
          <router-view />
        </div>
      </div>
    </div>
</template>
<script>
import AdminHeader from '@/components/partials/AdminHeader'
import AdminMenu from '@/components/partials/AdminMenu'
  export default {
    name: 'defaultLayout',
    components: {
        AdminHeader,
        AdminMenu
    }
  }
</script>