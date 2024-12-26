<template>
<div>
    <header class="main-header">
    <div class="header-menu">
        <h1 class="logo" @click="goToHomepage">
          <img src="@/assets/logo.png" alt="Logo" class="logo-image" />Heat Fashion
        </h1>
        <nav>
            <button @click="goToProfile" class="menu-button">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="icon">
                    <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
                </svg>
            </button>
            <button @click="goToCart" class="menu-button">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" stroke="currentColor" stroke-width="0.5" stroke-linecap="round" stroke-linejoin="round" class="icon">
                    <circle cx="9" cy="21" r="1"></circle>
                    <circle cx="20" cy="21" r="1"></circle>
                    <path d="M1 1h4l2.68 13.39a1 1 0 0 0 1 .78h9.72a1 1 0 0 0 1-.78L23 6H6"></path>
                </svg>
            </button>
        </nav>
    </div>
    </header>
    <main>
    <router-view />
    </main>
</div>
</template>  
  
<script>
  export default {
    methods: {
      goToHomepage() {
        this.$router.push({ name: 'Homepage' });
      },
      goToCart() {
        this.$router.push({ name: 'Cart' });
      },
      goToProfile() {
        this.$router.push({ name: 'Profile' });
      }
    }
  }
</script>
  
<style scoped>
.main-header {
  position: fixed;
  top: 0;
  width: 100%;
  background-color: #2a2a2a;
  z-index: 1000;
  display: flex;
  justify-content: center;
  padding: 10px 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
}

.header-menu {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1400px;
  width: 100%;
}

.logo {
  font-size: 28px;
  font-weight: bold;
  color: #ff4500;
  margin: 0;
  cursor: pointer;
}

.logo-image {
  width: 23px;
  height: 23px;
  object-fit: contain;
  margin-right: 4px;
}

nav {
  display: flex;
  gap: 20px;
}

.menu-button {
  background-color: transparent;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 5px;
  border-radius: 5px;
  transition: background-color 0.2s;
}

.menu-button:hover {
  background-color: #444;
}

.icon {
  width: 32px;
  height: 32px;
  fill: #ff4500;
}

main {
  margin-top: 70px;
  padding: 20px;
}

</style>
