<script setup lang="ts">
import { ref } from 'vue'
import { getTime } from '../client/timeService'

const lastTimeResponse = ref('')

const fetchTime = async () => {
  const time = await getTime();
  lastTimeResponse.value = time;
}
</script>

<template>
  <div class="time-dialog">
    <h1 class="green">
      gRPC-Web Example
    </h1>
    <h3>
      Click the button below to send <pre>GetCurrentTimeRequest</pre> message.
    </h3>
    <button @click="fetchTime" class="button">Get Time</button>
    <div class="time-response" v-if="lastTimeResponse">
      The time is <i>{{ lastTimeResponse }}</i>
    </div>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.time-dialog h1,
.time-dialog h3 {
  text-align: center;
}

.button {
  margin: 0 auto;
  display: block;
  border-radius: 8px;
  padding: 10px 20px;
  font-size: 1.2rem;
  cursor: pointer;
  background-color: #007bff;
  color: #fff;
  border: none;
  transition: background-color 0.2s;
}

.button:hover {
  background-color: #0056b3;
}

.time-response {
  margin-top: 20px;
  font-size: 1.2rem;
  color: #fff;
  text-align: center;
  background-color: #333;
  padding: 15px;
  border-radius: 8px;
  max-width: 80%;
  margin-left: auto;
  margin-right: auto;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
</style>
