<template>
<div class="container">
    <h2 class="section-title">Our Products</h2>
    <div class="product-grid">
    <div
        v-for="product in products"
        :key="product.id"
        class="product-card"
        @click="goToProductDetail(product.id)"
    >
        <div class="product-details">
        <h3 class="product-name">{{ product.name }}</h3>
        <p class="product-price">CHF {{ product.price }}</p>
        <p class="product-id">ID: {{ product.id }}</p>
        <p class="product-ali-id">AliExpress ID: {{ product.aliExpressID }}</p>
        </div>
    </div>
    </div>
</div>
</template>
  
<script>
  import axios from 'axios'
  
  export default {
    name: 'ProductManager',
    data() {
      return {
        products: []
      }
    },
    methods: {
      async fetchProducts() {
        try {
          const response = await axios.get('http://localhost:3000/products/')
          this.products = response.data
        } catch (error) {
          console.error('Error fetching products:', error)
        }
      },
      goToProductDetail(productId) {
        this.$router.push({ name: 'ProductDetail', params: { id: productId } })
      }
    },
    mounted() {
      this.fetchProducts()
    }
  }
</script>
  
<style scoped>
  .section-title {
    font-size: 24px;
    color: #ff4500;
    text-align: center;
    margin-bottom: 20px;
  }
  
  .product-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 20px;
  }
  
  .product-card {
    background-color: #2a2a2a;
    padding: 15px;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
    transition: transform 0.3s, box-shadow 0.3s;
  }
  
  .product-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 6px 10px rgba(0, 0, 0, 0.3);
  }
  
  .product-details {
    text-align: center;
  }
  
  .product-name {
    font-size: 20px;
    font-weight: bold;
    margin: 10px 0;
  }
  
  .product-price {
    font-size: 18px;
    color: #ff4500;
  }
  
  .product-id,
  .product-ali-id {
    font-size: 14px;
    color: #ccc;
  }
</style>
  