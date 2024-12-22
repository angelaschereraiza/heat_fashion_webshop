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

    <div v-if="view === 'profile'">
      <h1>User Profile</h1>
      <p>Manage your account details here.</p>
      <div class="profile-details">
        <div class="attribute">
          <span class="label">Name:</span>
          <span class="value">{{ user.firstName }} {{ user.lastName }}</span>
        </div>
        <div class="attribute">
          <span class="label">Email:</span>
          <span class="value">{{ user.email }}</span>
        </div>
        <div class="attribute">
          <span class="label">Phone Number:</span>
          <span class="value">{{ user.phoneNumber }}</span>
        </div>
        <div class="attribute">
          <span class="label">Address:</span>
          <span class="value">{{ user.address }}</span>
        </div>
        <div class="attribute">
          <span class="label">City:</span>
          <span class="value">{{ user.city }}</span>
        </div>
        <div class="attribute">
          <span class="label">Postal Code:</span>
          <span class="value">{{ user.postalCode }}</span>
        </div>
        <div class="attribute">
          <span class="label">Country:</span>
          <span class="value">{{ user.country }}</span>
        </div>
        <button @click="logout" class="button">Logout</button>
      </div>
    </div>

    <div v-else-if="view === 'login'">
      <h1>Login</h1>
      <form @submit.prevent="login">
        <div class="form-group">
          <label for="email">Email:</label>
          <input type="email" v-model="credentials.email" required />
        </div>
        <div class="form-group">
          <label for="password">Password:</label>
          <input type="password" v-model="credentials.password" required />
        </div>
        <button type="submit" class="button">Login</button>
        <p class="switch-view" @click="switchView('register')">Don't have an account? Register here.</p>
      </form>
    </div>

    <div v-else-if="view === 'register'">
      <h1>Register</h1>
      <form @submit.prevent="register">
        <div class="form-group">
          <label for="firstName">First Name:</label>
          <input type="text" v-model="newUser.firstName" required />
        </div>
        <div class="form-group">
          <label for="lastName">Last Name:</label>
          <input type="text" v-model="newUser.lastName" required />
        </div>
        <div class="form-group">
          <label for="email">Email:</label>
          <input type="email" v-model="newUser.email" required />
        </div>
        <div class="form-group">
          <label for="password">Password:</label>
          <input type="password" v-model="newUser.password" required />
        </div>
        <div class="form-group">
          <label for="confirmPassword">Confirm Password:</label>
          <input type="password" v-model="confirmPassword" required />
          <p v-if="passwordMismatch" class="error-message">Passwords do not match</p>
        </div>
        <div class="form-group">
          <label for="phoneNumber">Phone Number:</label>
          <input type="text" v-model="newUser.phoneNumber" />
        </div>
        <div class="form-group">
          <label for="address">Address:</label>
          <input type="text" v-model="newUser.address" />
        </div>
        <div class="form-group">
          <label for="city">City:</label>
          <input type="text" v-model="newUser.city" />
        </div>
        <div class="form-group">
          <label for="postalCode">Postal Code:</label>
          <input type="text" v-model="newUser.postalCode" />
        </div>
        <div class="form-group">
          <label for="country">Country:</label>
          <input type="text" v-model="newUser.country" />
        </div>
        <button type="submit" class="button" :disabled="passwordMismatch">Register</button>
        <p class="switch-view" @click="switchView('login')">Already have an account? Login here.</p>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'User',
  data() {
    return {
      view: 'login', // 'profile', 'login', or 'register'
      user: null,
      credentials: {
        email: '',
        password: '',
      },
      newUser: {
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        phoneNumber: '',
        address: '',
        city: '',
        postalCode: '',
        country: '',
      },
      confirmPassword: '',
    };
  },
  computed: {
    passwordMismatch() {
      return this.newUser.password !== this.confirmPassword && this.confirmPassword !== '';
    },
  },
  methods: {
    async login() {
      try {
        const response = await axios.post('http://localhost:3000/login', this.credentials);
        this.user = response.data.user;
        this.view = 'profile';
      } catch (error) {
        console.error('Login failed:', error);
        alert('Login failed. Please check your credentials.');
      }
    },
    async register() {
      if (this.passwordMismatch) {
        alert('Passwords do not match');
        return;
      }
      try {
        await axios.post('http://localhost:3000/users', this.newUser);
        alert('Registration successful. You can now login.');
        this.switchView('login');
      } catch (error) {
        console.error('Registration failed:', error);
        alert('Registration failed. Please try again.');
      }
    },
    logout() {
      this.user = null;
      this.credentials = { email: '', password: '' };
      this.switchView('login');
    },
    switchView(view) {
      this.view = view;
    },
    goBack() {
      this.$router.go(-1)
    }
  },
};
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

h1 {
  color: #ff4500;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
  text-align: left;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

input {
  width: 100%;
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  background-color: #2a2a2a;
  color: #fff;
  box-sizing: border-box;
}

.button {
  width: 100%;
  padding: 10px;
  background-color: #ff4500;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
  box-sizing: border-box;
}

.button:disabled {
  background-color: #555;
  cursor: not-allowed;
}

.button:hover:enabled {
  background-color: #e63e00;
}

.profile-details {
  margin-top: 20px;
  text-align: left;
  max-width: 400px;
  margin: 0 auto;
}

.attribute {
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between;
}

.label {
  font-weight: bold;
}

.value {
  color: #ff4500;
}

.switch-view {
  color: #ff4500;
  cursor: pointer;
  margin-top: 10px;
  text-align: center;
}

.switch-view:hover {
  text-decoration: underline;
}

.error-message {
  color: #ff4500;
  font-size: 14px;
  margin-top: 5px;
}
</style>
