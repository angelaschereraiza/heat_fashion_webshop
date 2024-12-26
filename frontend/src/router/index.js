import { createRouter, createWebHistory } from 'vue-router';
import HeaderLayout from '@/components/Header.vue';
import Homepage from '@/components/Homepage.vue';
import ProductDetail from '@/components/ProductDetail.vue';
import Cart from '@/components/Cart.vue';
import Profile from '@/components/Profile.vue';

const routes = [
  {
    path: '/',
    component: HeaderLayout,
    children: [
      { path: '', name: 'Homepage', component: Homepage },
      { path: 'product/:id', name: 'ProductDetail', component: ProductDetail, props: true },
      { path: 'cart', name: 'Cart', component: Cart },
      { path: 'profile', name: 'Profile', component: Profile }
    ]
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
