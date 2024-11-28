<template>
  <div class="dashboard">
    <h1>Welcome to the Dashboard</h1>

    <div v-if="isAuthenticated">
      <h2>Your Clients</h2>
      <ClientList />
    </div>
    <div v-else>
      <p>Please log in to access your dashboard.</p>
      <button @click="goToLogin">Login</button>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import ClientList from './ClientList.vue';

export default {
  name: 'DashboardPage',
  components: {
    ClientList,
  },
  setup() {
    const router = useRouter(); // Используем useRouter в Composition API
    const isAuthenticated = ref(false);

    onMounted(() => {
      isAuthenticated.value = !!localStorage.getItem('token');

      // Если не авторизован, перенаправляем на страницу логина
      if (!isAuthenticated.value) {
        router.push("/login");
      }
    });

    const goToLogin = () => {
      router.push("/login");
    };

    return {
      isAuthenticated,
      goToLogin
    };
  }
};
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

h1 {
  font-size: 2em;
}

button {
  padding: 10px 15px;
  font-size: 16px;
  cursor: pointer;
}
</style>
