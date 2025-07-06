import { createRouter, createWebHistory } from 'vue-router'
import DashboardView from '@/views/DashboardView.vue'
import InvoiceView from '@/views/InvoiceView.vue'
// import InvoicePreview from '@/views/InvoicePreview.vue'
import MainLayout from '@/components/layouts/MainLayout.vue'
import ArchivedView from '@/views/ArchivedView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: MainLayout,
      children: [
        { path: '', redirect: '/dashboard' },
        { path: 'dashboard', name: 'dashboard', component: DashboardView },
        { path: 'invoice', name: 'invoice', component: InvoiceView },
        { path: 'archived', name: 'archived', component: ArchivedView },
        // { path: 'preview', name: 'invoicePreview', component: InvoicePreview },
      ],
    },
  ],
})

export default router
