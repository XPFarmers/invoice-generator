export class LineItem {
  qty: number
  description: string
  amount: number
  total: number

  constructor(qty: number, description: string, amount: number | string) {
    this.qty = Math.round(Number(qty))
    this.description = description
    this.amount = parseFloat(parseFloat(String(amount)).toFixed(2))
    this.total = parseFloat((this.qty * this.amount).toFixed(2))
  }

  toJSON() {
    return {
      qty: String(this.qty),
      description: this.description,
      amount: this.amount.toFixed(2),
      total: this.total.toFixed(2)
    }
  }
}
