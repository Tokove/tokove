<template>
  <div 
    class="content-base-container" 
    :style="{ paddingTop: topPadding + 'px' }"
  >
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: "ContentBase",
  props: {
    topPadding: {
      type: Number,
      default: 125 // 这里可以根据实际情况继续增大，比如150、160
    }
  }
};
</script>

<style scoped>
.content-base-container {

  background-color: #f9fafb;
}
</style>