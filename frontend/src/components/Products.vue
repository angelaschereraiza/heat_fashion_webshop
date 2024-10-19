<template>
  <div class="container">
    <h2>All Products</h2>
    <ul class="product-list">
      <li v-for="product in products" :key="product.id" class="product-item">
        <span>{{ product.name }}</span> - <span>CHF {{ product.price }}</span>
      </li>
    </ul>
    <div class="form-container">
      <h2>Create a New Product</h2>
      <form @submit.prevent="submitProduct" class="product-form">
        <div class="form-group">
          <label for="name">Product Name:</label>
          <input v-model="newProduct.name" id="name" required placeholder="Enter product name" />
        </div>
        <div class="form-group">
          <label for="price">Price:</label>
          <input v-model.number="newProduct.price" id="price" required type="number" step="0.01" placeholder="Enter price" />
        </div>
        <button type="submit" class="btn">Add Product</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ProductManager',
  data() {
    return {
      products: [],
      newProduct: {
        name: '',
        price: null
      }
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
    async submitProduct() {
      try {
        console.log("Submitting product:", this.newProduct)
        await axios.post('http://localhost:3000/products/', this.newProduct)
        this.newProduct.name = ''
        this.newProduct.price = null
        this.fetchProducts()
      } catch (error) {
        console.error('Error adding product:', error)
      }
    }
  },
  mounted() {
    this.fetchProducts()
  }
}
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

h2 {
  text-align: center;
  margin-bottom: 20px;
}

.product-list {
  list-style: none;
  padding: 0;
}

.product-item {
  padding: 10px;
  background-color: #f9f9f9;
  margin-bottom: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.form-container {
  margin-top: 30px;
}

.product-form {
  display: flex;
  flex-direction: column;
}

.form-group {
  margin-bottom: 15px;
}

label {
  margin-bottom: 5px;
  font-weight: bold;
}

input {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
  width: 100%;
}

.btn {
  padding: 10px;
  background-color: #5cb85c;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  width: 50%;
}

.btn:hover {
  background-color: #4cae4c;
}
</style>
