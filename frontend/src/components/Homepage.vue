<template>
  <div class="homepage">
    <header class="main-header">
      <h1>Heat Fashion</h1>
      <p>Stay warm and stylish with our electrically heated clothing.</p>
    </header>

    <!-- Produktübersicht -->
    <ProductManager />
  </div>
</template>

<script>
import ProductManager from './ProductManager.vue'

export default {
  name: 'Homepage',
  components: {
    ProductManager
  }
}
</script>

<style scoped>
  .main-header {
    text-align: center;
    margin-bottom: 30px;

    h1 {
      font-size: 36px;
      color: #ff4500; 
      margin-bottom: 10px;
    }

    p {
      font-size: 18px;
      color: #f1f1f1;
      margin: 0;
    }
  }

  .homepage {
    margin: 0;
    padding: 0;
  }

  @media (max-width: 600px) {
    .product-grid {
      grid-template-columns: 1fr;
    }

    .main-header h1 {
      font-size: 28px;
    }

    .main-header p {
      font-size: 16px;
    }
  }
</style>
