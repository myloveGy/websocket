import {Format} from './time'

// 授权数据
interface IAuth {
    sid: string
    access_token: string
    url?: string
}

export default class Socket {
    // socket 连接
    webSocket: WebSocket

    // 授权信息
    auth: IAuth

    /**
     * 发送消息信息
     */
    public typeClientAuth: string = 'auth'  // 发送授权
    public typeClientHeartbeat: string = 'heartbeat' // 发送心跳
    public typeClientMessage: string = 'message' // 发送消息
    public typeClientClose: string = 'close' // 发送主动关闭
    public typeClientMessageReceipt: string = 'messageReceipt' // 发送消息回复

    /**
     * 回复消息
     */
    public typeServerAuth: string = 'authResponse' // 授权回复
    public typeServerMessage: string = 'messageResponse' // 消息回复
    public typeServerHeartbeat: string = 'heartbeatResponse' // 心跳回复

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
    constructor(url: string, data: IAuth, handler: any) {
        // ws地址
        console.info('WebSocket Info: connection: ' + url)
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
        }

        this.callHandler(receiver)
    }

    // 数据发送
    send = (data: any, type: string = '') => {
        this.webSocket.send(JSON.stringify({
            type: type || this.typeClientMessage,
            data: JSON.stringify(data),
            time: Format(new Date(), "yyyy-MM-dd hh:mm:ss")
        }))
    }

    // 消息回执
    messageReceipt = (message: any) => this.send({...message}, this.typeClientMessageReceipt)

    // 关闭
    close = (e: CloseEvent) => {
        if (this.heartbeatInterval) {
            clearInterval(this.heartbeatInterval)
        }

        console.log('WebSocket Error: connection closed (' + e.code + ')')
    }

    // 心跳
    heartbeat = () => this.send({}, this.typeClientHeartbeat)

    // 调用处理函数
    callHandler = (receiver: any) => {
        if (this.handler.hasOwnProperty(receiver.type)) {
            this.handler[receiver.type].call(this, receiver.data, receiver)
        } else {
            console.error(`WebSocket Error: 消息类型[${receiver.type}]未处理，消息内容:`, receiver.data)
        }
    }
}