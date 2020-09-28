import { Format } from './ts/time'

// 授权数据
interface authData {
    sid: string
    accessToken: string
    url?: string
}

export class Socket {
    // socket 连接
    webSocket: WebSocket

    // 授权信息
    auth: authData

    // 授权请求
    authRequest = 2

    // 授权响应
    authResponse = 3

    // 服务器响应数据
    serverReplyMsg = 6

    // 连接端发送消息
    clientSendMsg = 4

    // 消息回执
    clientSendHeartbeat = 0

    // 消息类型
    clientSendReceipt = 7

    // 心跳检测
    heartbeatInterval: any = null

    // 处理函数
    handler: any

    /**
     * 初始化函数
     * @param url 
     * @param data 
     * @param handler 
     */
    constructor(url: string, data: any, handler:any) {
      // ws地址
      this.webSocket = new WebSocket(url)
      this.webSocket.onmessage = this.message.bind(this)
      this.webSocket.onclose = this.close.bind(this)
      this.webSocket.onopen = this.open.bind(this)
      this.auth = data
      this.handler = handler
    }
  
    // 开启事件
    open = () => this.send(this.auth, this.authRequest)
  
    // 数据接收
    message = (e: MessageEvent) => {
  
      if (!e.data) {
        console.info('出错了')
        return
      }
  
      // 接收消息的内容
      const receiver = JSON.parse(e.data)
  
      // 授权成功发送的通知，表示可以正常通行，需要上报心跳了
      if (receiver.type == this.authResponse) {
        this.heartbeatInterval = setInterval(() => {
            this.heartbeat() 
        }, 3e4)
      } else if (receiver.type == this.serverReplyMsg) {
        // 服务器发送消息过来
        let data = JSON.parse(receiver.body)
        try {
          this.callHandler(data, receiver)
        } catch (e) {
          console.info(data.type + ' 执行函数不存在')
        }
      }
    }
  
    // 数据发送
    send = (data:any, type:number = 4) => {
      this.webSocket.send(JSON.stringify({type, data: JSON.stringify(data), time: Format(new Date(), "yyyy-MM-dd hh:mm:ss")}))
    }
  
    // 消息回执
    messageReceipt = (message: any) => this.send({...message}, this.clientSendReceipt)

    // 关闭
    close = (e:CloseEvent) =>  {
      if (this.heartbeatInterval) {
        clearInterval(this.heartbeatInterval)
      }

      console.log('connection closed (' + e.code + ')')
    }
  
    // 心跳
    heartbeat = () => this.send({}, this.clientSendHeartbeat)
  
    // 调用处理函数
    callHandler = (body:any, receiver:any) => {
      if (this.handler.hasOwnProperty(body.type)) {
        this.handler[body.type].call(this, body.data, receiver)
      }
    }
}

const webSocket = new Socket("ws://localhost:3000/ws", {
    sid: 1,
    accessToken: "123456789",
    url: "/"
}, null)


