<template>
  <div class="container">
    <header class="header">
      <button class="back-button" @click="goBack">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="icon">
          <path d="M15.41 16.59L10.83 12l4.58-4.59L14 6l-6 6 6 6z" />
        </svg>
        <span class="tooltip">Zurück</span>
      </button>
      <h1>Product Details</h1>
    </header>

    <div v-if="product" class="product-detail">
      <h2 class="product-name">{{ product.name }}</h2>
      <div class="product-attributes">
        <div class="attribute">
          <span class="label">Price:</span>
          <span class="value">CHF {{ product.price }}</span>
        </div>
        <div class="attribute">
          <span class="label">ID:</span>
          <span class="value">{{ product.id }}</span>
        </div>
        <div class="attribute">
          <span class="label">AliExpress ID:</span>
          <span class="value">{{ product.aliExpressID }}</span>
        </div>
      </div>
    </div>
    <p v-else>Loading product details...</p>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'ProductDetail',
    props: ['id'],
    data() {
      return {
        product: null
      }
    },
    methods: {
      async fetchProduct() {
        try {
          const response = await axios.get(`http://localhost:3000/products/${this.id}`)
          this.product = response.data
        } catch (error) {
          console.error('Error fetching product details:', error)
        }
      },
      goBack() {
        this.$router.go(-1)
      }
    },
    mounted() {
      this.fetchProduct()
    }
  }
</script>

<style scoped>
.container {
  max-width: 900px;
  margin: 0 auto;
  padding: 20px;
  background-color: #1a1a1a;
  color: #f1f1f1;
  font-family: 'Arial', sans-serif;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
}

.header {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 20px;
}

.back-button {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: transparent;
  cursor: pointer;
  padding: 0;
  transition: transform 0.2s;
  border-radius: 10px;
  position: relative;
}

.back-button:hover {
  transform: scale(1.1);
}

.back-button:hover .tooltip {
  opacity: 1;
  visibility: visible;
}

.tooltip {
  position: absolute;
  bottom: 60px;
  left: 50%;
  transform: translateX(-50%);
  background-color: #333;
  color: #fff;
  padding: 5px 10px;
  border-radius: 5px;
  font-size: 14px;
  white-space: nowrap;
  opacity: 0;
  visibility: hidden;
  transition: opacity 0.2s ease, visibility 0.2s ease;
}

.icon {
  fill: #ff4500;
  width: 46px;
  height: 46px;
}

.product-detail {
  background-color: #2a2a2a;
  padding: 20px;
  border-radius: 10px;
  color: #f1f1f1;
  text-align: left;
  margin: 20px auto;
}

.product-name {
  font-size: 24px;
  margin: 0 0 20px 0;
  text-align: left;
}

.product-attributes {
  display: grid;
  grid-template-columns: 150px 1fr;
  gap: 10px;
  text-align: left;
  max-width: 500px;
}

.attribute {
  display: contents;
}

.label {
  font-weight: bold;
  padding-right: 10px;
  text-align: left;
}

.value {
  text-align: left;
  color: #ff4500;
}
</style>



