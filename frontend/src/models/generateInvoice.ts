import { LineItem } from '@/models/lineItems'

export class GenerateInvoice {
  name: string = ''
  address: string = ''
  timeToOrder: string = ''
  depositPercentage: number = 0
  lineItems: LineItem[] = []

  // constructor(name: string, address: string, timeToOrder: string, depositpercentage: number) {
  //   this.name = name
  //   this.address = address
  //   this.timeToOrder = timeToOrder
  //   this.depositpercentage = depositpercentage
  //   this.lineItems = []
  // }

  setClientName(name: string) {
    this.name = name
  }

  setClientAddress(address: string) {
    this.address = address
  }

  setTimeToOrder(time: string) {
    this.timeToOrder = time
  }

  setDepositPercentage(percent: number) {
    this.depositPercentage = percent
  }

  addLineItem(item: LineItem) {
    this.lineItems.push(item)
  }

  getLineItems(): LineItem[] {
    return this.lineItems
  }

  getGrandTotal(): number {
    return this.lineItems.reduce((sum, item) => sum + item.total, 0)
  }

  toJSON() {
    return {
      clientName: this.name,
      clientAddress: this.address,
      timeToOrder: this.timeToOrder,
      depositPercentage: this.depositPercentage.toFixed(0),
      lineItems: this.lineItems.map(item => item.toJSON())
    }
  }
}
