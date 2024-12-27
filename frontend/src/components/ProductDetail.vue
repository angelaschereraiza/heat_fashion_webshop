<template>
  <div class="container">
    <header class="header">
      <button class="back-button" @click="goBack">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="icon">
          <path d="M15.41 16.59L10.83 12l4.58-4.59L14 6l-6 6 6 6z" />
        </svg>
        <span class="tooltip">Back</span>
      </button>
      <h1>Product Details</h1>
    </header>

    <div v-if="product" class="product-detail">
      <div class="product-layout">
        <div class="product-info">
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
              <span class="value">{{ product.ali_express_id }}</span>
            </div>
          </div>
        </div>
        <div class="action-section">
          <div class="quantity-selector">
            <button @click="decrementQuantity" :disabled="quantity === 1">-</button>
            <span>{{ quantity }}</span>
            <button @click="incrementQuantity">+</button>
          </div>
          <button class="add-to-cart-button" @click="addToCart">Add to Cart</button>
        </div>
      </div>
    </div>
    <div v-else>
      <div class="loading-icon"></div>
    </div>
    <div v-for="(toast, index) in toasts" :key="index" :class="['toast', toast.type]">{{ toast.message }}</div>
  </div>
</template>

<script>
export default {
  name: 'ProductDetail',
  props: ['id'],
  data() {
    return {
      product: null,
      toasts: [],
      quantity: 1,
    };
  },
  methods: {
    async fetchProduct() {
      try {
        const response = await fetch(`http://localhost:3000/products/${this.id}`, {
          credentials: 'include',
        });
        if (!response.ok) {
          throw new Error('Failed to fetch product details');
        }
        this.product = await response.json();
      } catch (error) {
        console.error('Error fetching product details:', error);
        this.showToast('Failed to fetch product details.', 'error');
      }
    },
    goBack() {
      this.$router.go(-1);
    },
    incrementQuantity() {
      this.quantity++;
    },
    decrementQuantity() {
      if (this.quantity > 1) {
        this.quantity--;
      }
    },
    async addToCart() {
      try {
        const cartItem = {
          product_id: this.product.id,
          quantity: this.quantity,
        };

        const response = await fetch('http://localhost:3000/cart/', {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(cartItem),
        });

        if (!response.ok) {
          throw new Error('Failed to add product to cart');
        }

        this.showToast('Product added to cart successfully!', 'success');
      } catch (error) {
        console.error('Error adding product to cart:', error);
        this.showToast('Failed to add product to cart.', 'error');
      }
    },
    showToast(message, type) {
      this.toasts.push({ message, type });
      setTimeout(() => {
        this.toasts.shift();
      }, 5000);
    },
  },
  mounted() {
    this.fetchProduct();
  },
};
</script>

<style scoped>
.product-detail {
  background-color: #2a2a2a;
  padding: 20px;
  border-radius: 10px;
  color: #f1f1f1;
  text-align: left;
  margin: 20px auto;
}

.product-layout {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.product-info {
  max-width: 70%;
}

.action-section {
  display: flex;
  align-self: flex-end;
  flex-direction: column;
  align-items: flex-end;
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

.add-to-cart-button {
  margin-top: 12px;
  padding: 10px 20px;
  background-color: #ff4500;
  border: none;
  color: #fff;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  border-radius: 5px;
  transition: transform 0.2s;
}

.add-to-cart-button:hover {
  background-color: #e63900;
  transform: scale(1.1);
}

.loading-icon {
  border: 4px solid rgba(255, 255, 255, 0.2);
  border-top: 4px solid #fff;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  animation: spin 1s linear infinite;
  margin: 20px auto;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
