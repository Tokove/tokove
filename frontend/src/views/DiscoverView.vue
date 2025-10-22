<template>
  <ContentBase>
    <div class="row">
      <div class="col-8">
      </div>
      <div class="col-4">
        <DiscoverInfo />
      </div>
      
    </div>
  </ContentBase>
</template>

<script>
import ContentBase from '@/components/ContentBase.vue';
import DiscoverInfo from '@/components/DiscoverInfo.vue';

export default {
  name: "DiscoverView",
  components: {
    ContentBase ,
    DiscoverInfo,
  }
};
</script>

<style scoped>
.hot-topics {
  margin-top: 20px;
}

.topic-item {
  background: white;
  padding: 15px;
  margin-top: 10px;
  border-radius: 6px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}
</style>