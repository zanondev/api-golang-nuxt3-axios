<!-- frontend/pages/tasks/index.vue -->

<template>
  <div>
    <h1>Tarefas</h1>
    <ul>
      <li v-for="task in tasks" :key="task.id">{{ task.title }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  data() {
    return {
      tasks: [],
    };
  },
  async asyncData() {
    try {
      // Faça uma solicitação GET para buscar as tarefas da API
      const response = await this.$axios.get('/tasks');
      return { tasks: response.data };
    } catch (error) {
      console.error('Erro ao buscar tarefas:', error);
      return { tasks: [] };
    }
  },
});
</script>
