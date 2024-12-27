<template>
  <div class="container">
    <header class="header">
      <button class="back-button" @click="goBack">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="icon">
          <path d="M15.41 16.59L10.83 12l4.58-4.59L14 6l-6 6 6 6z" />
        </svg>
        <span class="tooltip">Back</span>
      </button>
      <h1>Shopping Cart</h1>
    </header>

    <div v-if="cartItems.length > 0" class="cart-detail">
      <div v-for="(item, index) in cartItems" :key="item.id" class="cart-item">
        <div class="item-info">
          <h2 class="item-name">{{ item.product.name }}</h2>
          <p class="item-price">Price: CHF {{ item.product.price }}</p>
        </div>
        <div class="quantity-selector">
          <button @click="decrementQuantity(index)" :disabled="item.quantity === 1">-</button>
          <span>{{ item.quantity }}</span>
          <button @click="incrementQuantity(index)">+</button>
        </div>
        <button class="remove-button" @click="removeItem(index)">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="icon">
            <path d="M3 6h18v2H3V6zm2 3h14v13H5V9zm3 2v9h2v-9H8zm6 0v9h2v-9h-2zM10 3h4v2h5v2H5V5h5V3z" />
          </svg>
          <span class="tooltip">Remove</span>
        </button>
      </div>
    </div>
    <p v-else>Your cart is currently empty.</p>
    <div v-for="(toast, index) in toasts" :key="index" :class="['toast', toast.type]">{{ toast.message }}</div>
  </div>
</template>

<script>
export default {
  name: 'Cart',
  data() {
    return {
      cartItems: [],
      toasts: [],
    };
  },
  methods: {
    async fetchCart() {
      try {
        const response = await fetch('http://localhost:3000/cart/', {
          credentials: 'include',
        });
        if (!response.ok) {
          throw new Error('Failed to fetch cart');
        }
        const data = await response.json();
        this.cartItems = data.cart_items.map(item => ({
          id: item.id,
          cart_id: item.cart_id,
          product: {
            id: item.product.id,
            name: item.product.name,
            price: item.product.price,
          },
          quantity: item.quantity,
        }));
      } catch (error) {
        console.error('Error fetching cart:', error);
        this.showToast('Failed to fetch cart.', 'error');
      }
    },
    incrementQuantity(index) {
      this.cartItems[index].quantity++;
      this.updateCartItem(this.cartItems[index]);
    },
    decrementQuantity(index) {
      if (this.cartItems[index].quantity > 1) {
        this.cartItems[index].quantity--;
        this.updateCartItem(this.cartItems[index]);
      }
    },
    async updateCartItem(item) {
      try {
        const response = await fetch(`http://localhost:3000/cart/`, {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            product_id: item.product.id,
            quantity: item.quantity,
          }),
        });
        if (!response.ok) {
          throw new Error('Failed to update cart item');
        }
        this.showToast('Cart updated successfully!', 'success');
      } catch (error) {
        console.error('Error updating cart item:', error);
        this.showToast('Failed to update cart.', 'error');
      }
    },
    async removeItem(index) {
      const item = this.cartItems[index];
      try {
        const response = await fetch(`http://localhost:3000/cart/${item.product.id}`, {
          method: 'DELETE',
          credentials: 'include',
        });
        if (!response.ok) {
          throw new Error('Failed to remove item from cart');
        }
        this.cartItems.splice(index, 1);
        this.showToast('Item removed from cart!', 'success');
      } catch (error) {
        console.error('Error removing item from cart:', error);
        this.showToast('Failed to remove item from cart.', 'error');
      }
    },
    showToast(message, type) {
      this.toasts.push({ message, type });
      setTimeout(() => {
        this.toasts.shift();
      }, 5000);
    },
    goBack() {
      this.$router.go(-1);
    },
  },
  mounted() {
    this.fetchCart();
  },
};
</script>

<style scoped>
.cart-detail {
  background-color: #2a2a2a;
  padding: 20px;
  border-radius: 10px;
  color: #f1f1f1;
  text-align: left;
  margin: 20px auto;
}

.cart-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding: 10px;
  border-bottom: 1px solid #444;
}

.item-info {
  flex: 1;
}

.item-name {
  font-size: 18px;
  margin: 0 0 10px 0;
}

.item-price {
  font-size: 16px;
  color: #ff4500;
}

.quantity-selector {
  display: flex;
  align-items: center;
  gap: 10px;
}

.quantity-selector button {
  background-color: #444;
  color: white;
  border: none;
  padding: 5px 10px;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
}

.quantity-selector button:disabled {
  background-color: #777;
  cursor: not-allowed;
}

.quantity-selector span {
  font-size: 18px;
  font-weight: bold;
  color: white;
}

.remove-button {
  margin-left: 16px;
  position: relative;
  background-color: #ff4500;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 5px;
  border-radius: 5px;
  transition: background-color 0.2s;
}

.remove-button:hover {
  background-color: #e63900;
  transform: scale(1.1);
}

.remove-button .icon {
  width: 24px;
  height: 24px;
  fill: white;
}
</style>

