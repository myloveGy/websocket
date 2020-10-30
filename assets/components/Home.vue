<template>
  <div class="content">
    <div class="message-content">
      <div
          v-for="(item, key) in data"
          class="message" :class="{'me': item.source === 'me', 'system': item.source !== 'me'}"
          :key="key"
      >
        <span>{{ item.content }}</span>
      </div>
    </div>
    <div class="form">
      <form @submit.prevent.stop="submit">
        <input type="text" v-model="message"/>
        <button type="submit">提交</button>
      </form>
    </div>
  </div>
</template>
<script lang="ts">
import {Component, Vue} from 'vue-property-decorator'
import webSocket from '../ts/websocket'

interface IMesage {
  source: string
  content: string
}

@Component
export default class Home extends Vue {

  socket: any

  data: Array<IMesage> = []

  message: string = ''

  created() {
    this.socket = new webSocket("ws://localhost:3000/ws", {
      sid: '123456',
      access_token: "12121212",
      url: "/"
    }, {
      messageResponse: (content: string) => this.data.push({
        source: 'system',
        content,
      })
    })
  }

  submit() {
    const data: IMesage = {
      source: "me",
      content: this.message,
    }

    this.message = ''
    this.socket.send(data)
    this.data.push(data)
  }
}
</script>

<style lang="less">
.content {
  background: red;
}

.me {
  color: green;
  text-align: right;
}

.system {
  color: orange;
  text-align: left;
}

.message {
  padding: 5px;
}

.form {
  flex: 1 1 auto;
  height: 35px;
}
</style>