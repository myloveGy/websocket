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

    /** 
     * 发送消息信息
     */
    public typeClientAuth: string = 'auth'  // 发送授权
    public typeClientHeartbeat: string = 'heartbeat' // 发送心跳
    public typeClientMessage: string = 'message' // 发送消息
    public typeClientClose: string = 'close' // 发送主动关闭
    public typeClientMessageReceipt: string = 'message receipt' // 发送消息回复

    /**
     * 回复消息
     */
    public typeServerAuth: string = 'auth response' // 授权回复
    public typeServerMessage: string = 'message response' // 消息回复
    public typeServerHeartbeat: string = 'heartbeat response' // 心跳回复
  
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
    open = () => this.send(this.auth, this.typeClientAuth)
  
    // 数据接收
    message = (e: MessageEvent) => {
  
      if (!e.data) {
        console.info('出错了')
        return
      }
  
      // 接收消息的内容
      const receiver = JSON.parse(e.data)
  
      // 授权成功发送的通知，表示可以正常通行，需要上报心跳了
      if (receiver.type == this.typeServerAuth) {
        this.heartbeatInterval = setInterval(() => {
            this.heartbeat() 
        }, 3e4)
      } else if (receiver.type == this.typeServerMessage) {
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
    send = (data:any, type: string = '') => {
      this.webSocket.send(JSON.stringify({
        type: type || this.typeClientMessage, 
        data: JSON.stringify(data), 
        time: Format(new Date(), "yyyy-MM-dd hh:mm:ss")
      }))
    }
  
    // 消息回执
    messageReceipt = (message: any) => this.send({...message}, this.typeClientMessageReceipt)

    // 关闭
    close = (e:CloseEvent) =>  {
      if (this.heartbeatInterval) {
        clearInterval(this.heartbeatInterval)
      }

      console.log('connection closed (' + e.code + ')')
    }
  
    // 心跳
    heartbeat = () => this.send({}, this.typeClientHeartbeat)
  
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


