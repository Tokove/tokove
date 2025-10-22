‘这是HomeView.vue<template>
  <div class="home-view">
    <!-- 顶部 Banner -->
    <section class="banner d-flex flex-column justify-content-center align-items-center text-center">
      <h1 class="banner-title gradient-text" data-aos="fade-down">
        欢迎来到 Tokove
      </h1>
      <p class="banner-subtitle gradient-text" data-aos="fade-up" data-aos-delay="200">
        在这里，你可以发现好友的每一个动态，分享属于你的精彩瞬间
      </p>
      <router-link class="btn btn-primary btn-lg mt-3" :to="{name:'register'}" data-aos="zoom-in" data-aos-delay="400">
        加入我们
      </router-link>
    </section>

    <!-- 功能入口卡片 -->
    <section class="features container my-5">
      <div class="row g-4">
        <div class="col-md-4" data-aos="fade-up">
          <div class="card feature-card text-center p-4">
            <i class="bi bi-people-fill fs-1 mb-2"></i>
            <h5>发现好友</h5>
            <p>快速找到你的好友，了解他们的最新动态</p>
          </div>
        </div>
        <div class="col-md-4" data-aos="fade-up" data-aos-delay="100">
          <div class="card feature-card text-center p-4">
            <i class="bi bi-chat-dots-fill fs-1 mb-2"></i>
            <h5>消息</h5>
            <p>随时收发消息，不错过任何重要信息</p>
          </div>
        </div>
        <div class="col-md-4" data-aos="fade-up" data-aos-delay="200">
          <div class="card feature-card text-center p-4">
            <i class="bi bi-heart-fill fs-1 mb-2"></i>
            <h5>喜欢</h5>
            <p>对喜欢的内容点赞，留下你的足迹</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="text-center py-4 bg-light mt-5">
      © 2025 Tokove. All Rights Reserved.
    </footer>
  </div>
</template>

<script>
export default {
  name: "HomeView",
  mounted() {
    if (window.AOS) window.AOS.refresh();
  },
};
</script>

<style scoped>

.banner {
  background: #ffffff;
  height: 320px;
  padding-top: 200px; /* 原来 60px → 改大，让文字往下 */
  padding-bottom: 60px;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 改成 flex-start 让文字从顶部 padding 开始 */
  align-items: center;
}

.banner .btn {
  font-weight: 700;
  font-family: 'Segoe UI', 'PingFang SC', 'Helvetica', 'Arial', sans-serif;
  margin-top: 20px; 
}

.banner-title {
  font-size: 3rem;       /* 标题更大 */
  font-weight: 800;      /* 加粗 */
  font-family: 'Segoe UI', 'PingFang SC', 'Helvetica', 'Arial', sans-serif; /* 圆润字体 */
}

.banner-subtitle {
  font-size: 1.4rem;     /* 副标题稍大 */
  font-weight: 600;      /* 半加粗 */
  font-family: 'Segoe UI', 'PingFang SC', 'Helvetica', 'Arial', sans-serif; /* 圆润字体 */
  margin-top: 10px;
}


.feature-card {
  border: none;
  border-radius: 12px;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  cursor: pointer;
  background-color: #fff;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
}

.gradient-text {
  /* 1. 定义渐变色背景 */
  background: linear-gradient(90deg, #1D62EB, #000A50);
  /* 2. 补充标准属性（虽然主流浏览器仍依赖前缀，但未来可能统一） */
  background-clip: text;
  /* 3. WebKit 内核前缀（核心生效属性，必须保留） */
  -webkit-background-clip: text;
  /* 4. 文字颜色透明（标准属性 + 前缀） */
  color: transparent; /* 标准属性，部分浏览器可能支持 */
  -webkit-text-fill-color: transparent; /* WebKit 内核核心属性，必须保留 */
}

</style>