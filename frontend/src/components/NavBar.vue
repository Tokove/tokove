<template>
  <div>
    <!-- 顶部蓝色通知栏 -->
    <div class="top-blue-bar d-flex justify-content-between align-items-center px-4">
      <span class="text-white small">Welcome to tokove!</span>
      <a href="#" class="text-white text-decoration-underline small">了解更多</a>
    </div>

    <!-- 主导航栏 -->
    <nav class="navbar navbar-expand-lg navbar-light bg-light shadow-sm main-navbar">
      <div class="container d-flex align-items-center justify-content-between">
        <!-- 品牌Logo -->
        <router-link class="navbar-brand d-flex align-items-center" :to="{ name: 'home' }">
          <img src="/favicon.ico" alt="Tokove" height="52" />
          <span class="nav-logo-lang ms-2">Tokove</span>
        </router-link>

        <!-- 移动端菜单按钮 -->
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNav"
          aria-controls="navbarNav"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>

        <!-- 导航菜单 -->
        <div class="collapse navbar-collapse justify-content-center" id="navbarNav">
          <ul class="navbar-nav d-flex flex-row gap-4 align-items-center">
            <li class="nav-item">
              <router-link 
                class="nav-link" 
                :to="{ name: 'home' }"
                :class="{ 'active': $route.name === 'home' }"
              >
                首页
              </router-link>
            </li>
            <li class="nav-item">
              <router-link 
                class="nav-link" 
                :to="{ name: 'discover' }"
                :class="{ 'active': $route.name === 'discover' }"
              >
                发现
              </router-link>
            </li>
            <li class="nav-item">
              <router-link 
                class="nav-link" 
                :to="{ name: 'moment' }"
                :class="{ 'active': $route.name === 'moment' }"
              >
                动态
              </router-link>
            </li>
            <li class="nav-item">
              <router-link 
                class="nav-link" 
                :to="{ name: 'friends' }"
                :class="{ 'active': $route.name === 'friends' }"
              >
                好友
              </router-link>
            </li>
            <li class="nav-item">
              <router-link 
                class="nav-link" 
                :to="{ name: 'messages' }"
                :class="{ 'active': $route.name === 'messages' }"
              >
                消息
              </router-link>
            </li>
            <li class="nav-item">
              <router-link 
                class="nav-link" 
                :to="{ name: 'likes' }"
                :class="{ 'active': $route.name === 'likes' }"
              >
                喜欢
              </router-link>
            </li>
          </ul>

          <!-- 搜索栏 -->
          <div class="search-box d-flex align-items-center ms-4">
            <input
              type="text"
              class="form-control search-input"
              placeholder="搜索..."
            />
            <i class="bi bi-search ms-2 search-icon"></i>
          </div>
        </div>

        <!-- 登录注册按钮（仅桌面显示） -->
        <div class="d-none d-lg-flex gap-2">
          <router-link 
            class="btn btn-primary text-white px-3" 
            :to="{ name: 'login' }"
          >
            登录
          </router-link>
          <router-link 
            class="btn btn-outline-primary text-primary bg-white px-3" 
            :to="{ name: 'register' }"
          >
            注册
          </router-link>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
export default {
  name: "NavBar",
  // 监听路由变化，确保激活状态正确更新
  watch: {
    $route() {
      this.$nextTick(() => {
        // 可选：添加路由切换时的额外逻辑
      });
    }
  }
};
</script>

<style scoped>
/* 顶部蓝色通知栏 */
.top-blue-bar {
  background-color: #00084D;
  height: 40px;
  position: fixed;
  width: 100%;
  top: 0;
  z-index: 1050; /* 确保在导航栏上方 */
}

/* 主导航栏 */
.main-navbar {
  position: fixed; /* 固定定位 */
  top: 40px; /* 位于蓝色通知栏下方 */
  left: 0;
  right: 0;
  z-index: 1040; /* 低于通知栏，高于页面内容 */
  height: 70px; /* 明确高度，方便计算页面偏移 */
}

/* Logo样式 */
.navbar-brand img {
  height: 52px;
  transition: transform 0.2s ease;
}

.navbar-brand img:hover {
  transform: scale(1.05);
}

.nav-logo-lang {
  color: #00153c;
  font-weight: 700;
  font-size: 1.6rem;
  margin-left: 10px;
  letter-spacing: 0.5px;
}

/* 导航链接样式 */
.nav-link {
  color: #333 !important;
  font-weight: 600;
  font-size: 1.1rem;
  border-radius: 8px;
  padding: 8px 14px;
  transition: all 0.2s ease;
  position: relative;
}

/* 激活状态样式 */
.nav-link.active {
  color: #007acc !important;
  background-color: rgba(0, 122, 204, 0.1);
}

.nav-link.active::after {
  content: '';
  position: absolute;
  bottom: 2px;
  left: 50%;
  transform: translateX(-50%);
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background-color: #007acc;
}

.nav-link:hover:not(.active) {
  background-color: rgba(0, 122, 204, 0.08);
  color: #007acc !important;
}

/* 搜索框样式 */
.search-box {
  position: relative;
}

.search-input {
  width: 180px;
  border-radius: 20px;
  padding: 6px 12px;
  font-size: 0.95rem;
  border: 1px solid #ccc;
  transition: all 0.3s ease;
}

.search-input:focus {
  width: 240px;
  outline: none;
  box-shadow: 0 0 6px rgba(0, 122, 204, 0.3);
  border-color: #007acc;
}

.search-icon {
  font-size: 1.2rem;
  color: #555;
  cursor: pointer;
  transition: color 0.2s ease;
}

.search-icon:hover {
  color: #007acc;
}

/* 导航栏背景 */
.navbar {
  backdrop-filter: blur(10px);
  background: rgba(255, 255, 255, 0.9) !important;
  padding: 0.5rem 1rem;
}

/* 响应式调整 */
@media (max-width: 992px) {
  /* 移动端导航菜单 */
  .navbar-collapse {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background-color: white;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    padding: 1rem;
    z-index: 1040;
  }

  .navbar-nav {
    flex-direction: column !important;
    align-items: flex-start !important;
    gap: 0.5rem !important;
  }

  .search-box {
    margin-left: 0 !important;
    margin-top: 1rem;
    width: 100%;
  }

  .search-input {
    width: 100% !important;
  }

  .search-input:focus {
    width: 100% !important;
  }
}

@media (max-width: 576px) {
  .nav-logo-lang {
    font-size: 1.3rem;
  }
  
  .main-navbar {
    height: 60px;
  }
  
  .navbar-brand img {
    height: 40px;
  }
}
</style>