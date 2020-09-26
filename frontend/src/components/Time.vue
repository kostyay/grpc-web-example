<template>
  <div class="hello">
    <h1>gRPC-Web Example</h1>
    <div>
      Click the button below to send <pre>GetCurrentTimeRequest</pre> message.
    </div>
    <div>
      <input type="button" class="button" value="Refresh Time" @click="getTime()">
    </div>
    <div>
      {{ lastTimeResponse }}
    </div>
  </div>
</template>

<script>
import { TimeServiceClient } from '../jsclient/time/v1/time_service_grpc_web_pb';
import { GetCurrentTimeRequest } from '../jsclient/time/v1/time_service_pb';

export default {
  name: 'Time',
  data: function() {
    return { lastTimeResponse: 'n/a' }
  },
  methods: {
    getTime: function () {
      const client = new TimeServiceClient("http://localhost:8080", null, null);
      const enableDevTools = window.__GRPCWEB_DEVTOOLS__ || (() => {});
      enableDevTools([
        client,
      ]);

      client.getCurrentTime(new GetCurrentTimeRequest(), {}, (err, response) => {
        this.lastTimeResponse = response.getCurrentTime();
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
pre {
  background: #f4f4f4;
  border: 1px solid #ccc;
  color: #666;
  page-break-inside: avoid;
  font-family: monospace;
  font-size: 15px;
  line-height: 1.6;
  margin-bottom: 1.6em;
  max-width: 100%;
  overflow: auto;
  padding: 0.1em 0.1em;
  display: inline;
  word-wrap: break-word;
}
.button {
  background-color: purple;
  border-radius: 4px;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  font-weight: bold;
  cursor: pointer;
}
</style>
