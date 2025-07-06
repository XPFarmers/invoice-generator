// Use Vite's built-in environment flags to determine base URL
// import.meta.env.DEV is true in development mode (npm run dev)
// import.meta.env.PROD is true in production mode (npm run build)
const base = import.meta.env.DEV ? 'http://localhost:8080' : ''

export const Routes = {
  LineItems: base + '/lineItems',
  GenerateInvoice: base + '/invoice',
}
