import { SimpleSSEClient } from './sse.ts'

export class EventAwareSSEClient extends SimpleSSEClient {
  private handlers: Record<string, Function> = {}

  addEventListener(eventName: string, callback: (data: any) => void) {
    this.handlers[eventName] = callback
  }

  public handleMessage(data: any) {
    if (data.event && this.handlers[data.event]) {
      this.handlers[data.event](data.payload)
    } else {
      super.handleMessage(data)
    }
  }
}
