export class SimpleSSEClient {
  private eventSource: EventSource | null = null
  private retryCount = 0
  private maxRetries = 3
  private retryInterval = 5000
  private onOpenFunction: Function | null = null
  private onErrorFunction: Function | null = null

  constructor(private url: string) {}

  connect() {
    if (this.eventSource) {
      this.close()
    }
    this.eventSource = new EventSource(this.url)

    this.eventSource.onopen = () => {
      this.retryCount = 0 // 重置重试计数器
      console.log('SSE连接成功', this.retryCount)
      if (this.onOpenFunction) {
        this.onOpenFunction()
      }
    }

    this.eventSource.onmessage = (e) => {
      try {
        const parsedData = JSON.parse(e.data)
        this.handleMessage(parsedData)
      } catch (err) {
        console.error('数据解析失败:', err)
      }
    }

    this.eventSource.onerror = (e) => {
      console.log('SSE连接错误:', e)
      if (this.onErrorFunction) {
        this.onErrorFunction()
      }
      if (this.retryCount >= this.maxRetries) {
        this.close()
        return
      }

      setTimeout(() => {
        this.retryCount++
        this.reconnect()
      }, this.retryInterval)
    }
  }

  public setOnOpenFunction(f: Function) {
    this.onOpenFunction = f
  }

  public setOnErrorFunction(f: Function) {
    this.onErrorFunction = f
  }

  public handleMessage(data: any) {
    // 需要时在此处添加数据格式校验
    console.log('收到消息:', data)
  }

  public reconnect() {
    this.close()
    this.connect()
  }

  public close() {
    this.eventSource?.close()
    this.eventSource = null
  }
}
