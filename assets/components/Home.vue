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
        <button type="submit" :disabled="disabled">发送</button>
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

  // 消息列表
  data: Array<IMesage> = []

  // 输入消息
  message: string = ''

  // 是否可以点击
  disabled: boolean = true

  created() {
    this.socket = new webSocket("ws://localhost:8080/ws", {
      sid: '123456',
      access_token: "12121212",
      url: "/"
    }, {
      messageResponse: (content: string) => this.data.push({
        source: 'system',
        content,
      }),
      authResponse: () => {
        this.disabled = false
      }
    })
  }

  submit() {
    if (!this.message) {
      console.info('请输入消息内容')
      return
    }

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
  width: 50%;
  margin: 20px auto;
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
  width: 100%;
  line-height: 35px;
  form {
    width: 100%;
    display: flex;

    input {
      flex-grow: 1;
      display: block;
      height: 35px;
      padding: 4px 11px;
      margin: 0;
      background-color: #fff;
      background-image: none;
      border: 1px solid #d9d9d9;
      border-radius: 4px;
      box-sizing: border-box;
    }

    button {
      margin-left: 10px;
      height: 35px;
      width: 100px;
      display: block;
      padding: 4px 11px;
      background-image: none;
      border: 1px solid #13c2c2;
      border-radius: 4px;
      box-sizing: border-box;
      color: #fff;
      background-color: #13c2c2;
    }
  }
}
</style>