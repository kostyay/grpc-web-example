<template>
  <div id="app">
    <Time />
  </div>
</template>

<script>
import Time from './components/Time.vue'

export default {
  name: 'App',
  components: {
    Time
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
