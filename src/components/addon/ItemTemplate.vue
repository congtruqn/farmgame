<template>
  <div class="select_list">
    <span class="slect_item">{{ item }}</span>
  </div>
</template>
<script>
  export default {
    props: {
      item: { required: true },
    },
  };
</script>