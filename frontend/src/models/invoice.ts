export interface LineItem {
  ID: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt: string | null
  invoiceId: number
  qty: string
  description: string
  amount: string
  total: string
}

export interface Invoice {
  ID: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt: string | null
  clientName: string
  clientAddress: string
  issueDate: string
  timeToOrder: string
  depositPercentage: string
  deposit: string
  balance: string
  total: string
  quotation: boolean
  archived?: boolean
  lineItems: LineItem[]
}
