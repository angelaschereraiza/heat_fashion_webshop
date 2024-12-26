<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
  export default {
    name: 'App',
    mounted() {
    document.title = "Heat Fashion";
  }
  }
</script>

<style>
  body {
      background-color: #000;
      color: #fff;
      font-family: "Helvetia", Arial, sans-serif;
      margin: 0;
  }

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

.toast {
  position: fixed;
  bottom: 40px;
  right: 40px;
  background-color: #333;
  color: #fff;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  font-size: 18px;
  z-index: 1000;
  margin-top: 10px;
  transition: opacity 0.3s;
}

.toast.success {
  background-color: #4caf50;
}

.toast.error {
  background-color: #f44336;
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
  width: 32px;
  height: 32px;
  border-radius: 5px;
  cursor: pointer;
  font-size: 20px;
}

.quantity-selector button:disabled {
  background-color: #777;
  cursor: not-allowed;
}

.quantity-selector span {
  font-size: 20px;
  font-weight: bold;
  color: white;
}
</style>
