<template>
  <div class="container">
    <header class="header">
      <button class="back-button" @click="goBack">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="icon">
          <path d="M15.41 16.59L10.83 12l4.58-4.59L14 6l-6 6 6 6z" />
        </svg>
        <span class="tooltip">Back</span>
      </button>
      <h1 v-if="view === 'register'">Register</h1>
      <h1 v-else-if="view === 'login'">Login</h1>
      <h1 v-else-if="view === 'profile'">User Profile</h1>
    </header>

    <div v-if="view === 'register'">
      <form @submit.prevent="register">
        <div class="form-row">
          <div class="form-group half-width">
            <label for="firstName">First Name</label>
            <input type="text" v-model="newUser.first_name" required />
          </div>
          <div class="form-group half-width">
            <label for="lastName">Last Name</label>
            <input type="text" v-model="newUser.last_name" required />
          </div>
        </div>

        <div class="form-group">
          <label for="mail">Email</label>
          <input type="email" v-model="newUser.mail" required />
        </div>

        <div class="form-row">
          <div class="form-group half-width">
            <label for="password">Password</label>
            <input type="password" v-model="newUser.password" required />
          </div>
          <div class="form-group half-width">
            <label for="confirmPassword">Confirm Password</label>
            <input type="password" v-model="confirmPassword" required />
            <p v-if="passwordMismatch" class="error-message">Passwords do not match</p>
          </div>
        </div>

        <div class="form-row">
          <div class="form-group quarter-width">
            <label for="countryCode">Country Code</label>
            <select v-model="newUser.country_code" required>
              <option v-for="code in countryCodes" :value="code.value" :key="code.value">
                {{ code.label }}
              </option>
            </select>
          </div>
          <div class="form-group three-quarters-width">
            <label for="phoneNumber">Phone Number</label>
            <input type="text" v-model="newUser.phone_number" />
          </div>
        </div>

        <div class="form-group">
          <label for="address">Address</label>
          <input type="text" v-model="newUser.address" />
        </div>

        <div class="form-row">
          <div class="form-group quarter-width">
            <label for="postalCode">Postal Code</label>
            <input type="text" v-model="newUser.postal_code" />
          </div>
          <div class="form-group three-quarters-width">
            <label for="city">City</label>
            <input type="text" v-model="newUser.city" />
          </div>
        </div>

        <div class="form-group">
          <label for="country">Country</label>
          <select v-model="newUser.country" required>
            <option v-for="country in countries" :value="country.value" :key="country.value">
              {{ country.label }}
            </option>
          </select>
        </div>

        <button type="submit" class="button" :disabled="passwordMismatch">Register</button>
        <p class="switch-view" @click="switchView('login')">Already have an account? Login here.</p>
      </form>
    </div>

    <div v-else-if="view === 'profile'">
      <h2>Welcome, {{ user.first_name }} {{ user.last_name }}</h2>
      <div class="profile-details">
        <div class="attribute">
          <span class="label">Email:</span>
          <span class="value">{{ user.mail }}</span>
        </div>
        <div class="attribute">
          <span class="label">Phone:</span>
          <span class="value">{{ user.country_code }} {{ user.phone_number }}</span>
        </div>
        <div class="attribute">
          <span class="label">Address:</span>
          <span class="value">{{ user.address }}</span>
        </div>
        <div class="attribute">
          <span class="label">Postal Code:</span>
          <span class="value">{{ user.postal_code }}</span>
        </div>
        <div class="attribute">
          <span class="label">City:</span>
          <span class="value">{{ user.city }}</span>
        </div>
        <div class="attribute">
          <span class="label">Country:</span>
          <span class="value">{{ user.country }}</span>
        </div>
        <button @click="logout" class="button">Logout</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'UserForm',
  data() {
    return {
      view: 'register',
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
        countryCode: '',
        phoneNumber: '',
        address: '',
        city: '',
        postalCode: '',
        country: '',
      },
      confirmPassword: '',
      countries: [
        { label: 'Australia', value: 'AU' },
        { label: 'Austria', value: 'AT' },
        { label: 'Belgium', value: 'BE' },
        { label: 'Canada', value: 'CA' },
        { label: 'Denmark', value: 'DK' },
        { label: 'Finland', value: 'FI' },
        { label: 'France', value: 'FR' },
        { label: 'Germany', value: 'DE' },
        { label: 'Ireland', value: 'IE' },
        { label: 'Italy', value: 'IT' },
        { label: 'Netherlands', value: 'NL' },
        { label: 'Norway', value: 'NO' },
        { label: 'Sweden', value: 'SE' },
        { label: 'Switzerland', value: 'CH' },
        { label: 'United Kingdom', value: 'UK' },
        { label: 'United States', value: 'US' },
      ],
      countryCodes: [
        { label: '+61 (Australia)', value: '+61' },
        { label: '+43 (Austria)', value: '+43' },
        { label: '+32 (Belgium)', value: '+32' },
        { label: '+45 (Denmark)', value: '+45' },
        { label: '+358 (Finland)', value: '+358' },
        { label: '+33 (France)', value: '+33' },
        { label: '+49 (Germany)', value: '+49' },
        { label: '+353 (Ireland)', value: '+353' },
        { label: '+39 (Italy)', value: '+39' },
        { label: '+31 (Netherlands)', value: '+31' },
        { label: '+47 (Norway)', value: '+47' },
        { label: '+46 (Sweden)', value: '+46' },
        { label: '+41 (Switzerland)', value: '+41' },
        { label: '+44 (United Kingdom)', value: '+44' },
        { label: '+1 (USA)', value: '+1' },
      ],
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
        const response = await axios.post('http://localhost:3000/login/', this.credentials, {
          withCredentials: true,
        });
        this.user = response.data;
        this.view = 'profile';
      } catch (error) {
        alert('Login failed. Please check your credentials.');
      }
    },
    async register() {
      if (this.passwordMismatch) {
        alert('Passwords do not match');
        return;
      }
      try {
        await axios.post('http://localhost:3000/register/', this.newUser, {
          withCredentials: true,
        });
        alert('Registration successful! Please login.');
        this.switchView('login');
      } catch (error) {
        alert('Registration failed. Please try again.');
      }
    },
    logout() {
      this.user = null;
      this.credentials = { email: '', password: '' };
      this.view = 'login';
    },
    switchView(view) {
      this.view = view;
    },
    goBack() {
      this.$router.go(-1);
    },
  },
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  background-color: #2a2a2a;
  color: #fff;
  border-radius: 10px;
}

.form-row {
  display: flex;
  justify-content: space-between;
  gap: 20px;
}

.form-group {
  margin-bottom: 15px;
  flex: 1;

  &.half-width {
    flex: 1;
  }

  &.quarter-width {
    flex: 0.25;
  }

  &.three-quarters-width {
    flex: 0.75;
  }
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

input,
select {
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
  font-size: 16px;

  &:disabled {
    background-color: #555;
    cursor: not-allowed;
  }

  &:hover:enabled {
    background-color: #e63e00;
  }
}

.error-message {
  color: #ff4500;
  font-size: 14px;
  margin-top: 5px;
}

.switch-view {
  color: #ff4500;
  cursor: pointer;
  margin-top: 10px;
  text-align: center;

  &:hover {
    text-decoration: underline;
  }
}

.profile-details {
  .attribute {
    margin-bottom: 10px;
    display: flex;
    justify-content: space-between;

    .label {
      font-weight: bold;
    }

    .value {
      color: #ff4500;
    }
  }
}
</style>
