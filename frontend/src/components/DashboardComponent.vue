<template>
  <v-container fluid>
    <!-- Header Section -->
    <v-row class="mb-4">
      <v-col cols="12">
        <v-card>
          <v-card-title class="d-flex justify-space-between align-center">
            <div>
              <!-- <h2 class="text-h4 font-weight-bold">Invoice Dashboard</h2> -->
              <p class="text-subtitle-1 text-medium-emphasis mt-1">
                Manage and view all your invoices and quotations
              </p>
            </div>
            <v-btn
              color="primary"
              size="large"
              prepend-icon="mdi-plus"
              @click="createNewInvoice"
            >
              New Invoice
            </v-btn>
          </v-card-title>
        </v-card>
      </v-col>
    </v-row>

    <!-- Data Table Card -->
    <v-card>
      <v-card-title class="d-flex justify-space-between align-center">
        <div>
          <h3 class="text-h5">Invoices & Quotations</h3>
          <p class="text-body-2 text-medium-emphasis">
            {{ filteredInvoices.length }} of {{ invoices.length }} records
          </p>
        </div>

        <!-- Search and Filter Controls -->
        <div class="d-flex gap-4 align-center mr-4">
          <v-text-field
            v-model="search"
            prepend-inner-icon="mdi-magnify"
            label="Search invoices..."
            single-line
            hide-details
            density="compact"
            max-width="700px"
            min-width="550px"
            clearable
            class="mr-4"
          />

          <v-select
            v-model="statusFilter"
            :items="statusOptions"
            label="Status"
            density="compact"
            hide-details
            max-width="400px"
            min-width="350px"
            clearable
          />
        </div>
      </v-card-title>

      <v-card-text>
        <!-- Loading State -->
        <div v-if="loading" class="d-flex justify-center align-center py-8">
          <v-progress-circular indeterminate color="primary" size="64" />
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="d-flex justify-center align-center py-8">
          <v-alert type="error" class="max-width-400">
            <template #title>Error Loading Invoices</template>
            {{ error }}
            <template #append>
              <v-btn color="error" variant="text" @click="fetchInvoices">
                Retry
              </v-btn>
            </template>
          </v-alert>
        </div>

        <!-- Data Table -->
        <v-data-table
          v-else
          :headers="headers"
          :items="filteredInvoices"
          :search="search"
          :loading="loading"
          hover
          class="elevation-1"
          item-value="ID"
        >
          <!-- Client Name Column -->
          <template #item.clientName="{ item }">
            <div>
              <div class="font-weight-medium">{{ item.clientName || 'N/A' }}</div>
              <div class="text-caption text-medium-emphasis">{{ item.clientAddress || 'No address' }}</div>
            </div>
          </template>

          <!-- Issue Date Column -->
          <template #item.issueDate="{ item }">
            <v-chip
              :color="getDateColor(item.issueDate)"
              size="small"
              variant="tonal"
            >
              {{ formatDate(item.issueDate) }}
            </v-chip>
          </template>

          <!-- Total Amount Column -->
          <template #item.total="{ item }">
            <div class="text-right">
              <div class="font-weight-bold text-primary">
                R{{ parseFloat(item.total).toFixed(2) }}
              </div>
              <div class="text-caption text-medium-emphasis">
                {{ item.lineItems.length }} items
              </div>
            </div>
          </template>

          <!-- Deposit Column -->
          <template #item.deposit="{ item }">
            <div class="text-right">
              <div class="font-weight-medium">
                R{{ parseFloat(item.deposit).toFixed(2) }}
              </div>
              <div class="text-caption text-medium-emphasis">
                {{ item.depositPercentage }}%
              </div>
            </div>
          </template>

          <!-- Balance Column -->
          <template #item.balance="{ item }">
            <div class="text-right">
              <div class="font-weight-medium">
                R{{ parseFloat(item.balance).toFixed(2) }}
              </div>
            </div>
          </template>

          <!-- Status Column -->
          <template #item.quotation="{ item }">
            <div class="d-flex justify-center">
              <v-chip
                :color="item.quotation ? 'warning' : 'success'"
                size="small"
                variant="tonal"
              >
                <v-icon start size="small">
                  {{ item.quotation ? 'mdi-file-document-outline' : 'mdi-receipt' }}
                </v-icon>
                {{ item.quotation ? 'Quotation' : 'Invoice' }}
              </v-chip>
            </div>
          </template>

          <!-- Actions Column -->
          <template #item.actions="{ item }">
            <div class="d-flex justify-center gap-1">
              <v-btn
                icon="mdi-eye"
                size="small"
                variant="text"
                color="info"
                @click="viewInvoice(item)"
                title="View Details"
              />
              <v-btn
                icon="mdi-pencil"
                size="small"
                variant="text"
                color="warning"
                @click="editInvoice(item)"
                title="Edit Invoice"
              />
              <v-btn
                icon="mdi-download"
                size="small"
                variant="text"
                color="success"
                @click="downloadInvoice(item)"
                title="Download PDF"
              />
              <v-btn
                icon="mdi-archive"
                size="small"
                variant="text"
                color="secondary"
                @click="archiveInvoice(item)"
                title="Archive Invoice"
              />
            </div>
          </template>

          <!-- No Data State -->
          <template #no-data>
            <div class="d-flex flex-column align-center justify-center py-8">
              <v-icon size="64" color="grey-lighten-1" class="mb-4">
                mdi-file-document-outline
              </v-icon>
              <h3 class="text-h6 text-grey">No invoices found</h3>
              <p class="text-body-2 text-grey-lighten-1">
                Create your first invoice to get started
              </p>
              <v-btn
                color="primary"
                class="mt-4"
                @click="createNewInvoice"
              >
                Create Invoice
              </v-btn>
            </div>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <!-- Confirmation Dialog for Archive -->
    <v-dialog v-model="archiveDialog" max-width="400">
      <v-card>
        <v-card-title>Archive Invoice</v-card-title>
        <v-card-text>
          Are you sure you want to archive this invoice? This action cannot be undone.
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn @click="archiveDialog = false">Cancel</v-btn>
          <v-btn color="error" @click="confirmArchive">Archive</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Routes } from '@/utils/routes'
import type { Invoice } from '@/models/invoice'

const router = useRouter()

// Reactive data
const invoices = ref<Invoice[]>([])
const loading = ref(false)
const error = ref<string | null>(null)
const search = ref('')
const statusFilter = ref<string | null>(null)
const archiveDialog = ref(false)
const selectedInvoice = ref<Invoice | null>(null)

// Table headers
const headers = [
  { title: 'Client', key: 'clientName', sortable: true },
  { title: 'Issue Date', key: 'issueDate', sortable: true },
  { title: 'Total', key: 'total', sortable: true, align: 'end' as const },
  { title: 'Deposit', key: 'deposit', sortable: true, align: 'end' as const },
  { title: 'Balance', key: 'balance', sortable: true, align: 'end' as const },
  { title: 'Status', key: 'quotation', sortable: true, align: 'center' as const },
  { title: 'Actions', key: 'actions', sortable: false, align: 'center' as const }
]

// Status filter options
const statusOptions = [
  { title: 'All', value: null },
  { title: 'Quotations', value: 'quotation' },
  { title: 'Invoices', value: 'invoice' }
]

// Computed properties
const filteredInvoices = computed(() => {
  let filtered = invoices.value

  // Apply status filter
  if (statusFilter.value === 'quotation') {
    filtered = filtered.filter(invoice => invoice.quotation)
  } else if (statusFilter.value === 'invoice') {
    filtered = filtered.filter(invoice => !invoice.quotation)
  }

  return filtered
})

// Methods
const fetchInvoices = async () => {
  loading.value = true
  error.value = null

  try {
  const response = await fetch(Routes.Invoices)
  if (!response.ok) {
    const errorData = await response.json()
    // Now you have structured error info:
    // errorData.error, errorData.message, errorData.code
    throw new Error(errorData.message)
  }
  invoices.value = await response.json()
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Failed to fetch invoices'
    console.error('Error fetching invoices:', err)
  } finally {
    loading.value = false
  }
}

const formatDate = (dateString: string) => {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleDateString('en-GB', {
    day: '2-digit',
    month: 'short',
    year: 'numeric'
  })
}

const getDateColor = (dateString: string) => {
  if (!dateString) return 'grey'

  const date = new Date(dateString)
  const now = new Date()
  const diffTime = now.getTime() - date.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))

  if (diffDays <= 7) return 'success'
  if (diffDays <= 30) return 'warning'
  return 'grey'
}

const createNewInvoice = () => {
  router.push('/invoice')
}

const viewInvoice = (invoice: Invoice) => {
  // Navigate to invoice preview or details view
  console.log('View invoice:', invoice.ID)
  // router.push({ name: 'invoice-details', params: { id: invoice.ID } })
}

const editInvoice = (invoice: Invoice) => {
  // Navigate to edit invoice view
  console.log('Edit invoice:', invoice.ID)
  // router.push({ name: 'edit-invoice', params: { id: invoice.ID } })
}

const downloadInvoice = (invoice: Invoice) => {
  // Download PDF functionality
  console.log('Download invoice:', invoice.ID)
  // Implement PDF download logic
}

const archiveInvoice = (invoice: Invoice) => {
  selectedInvoice.value = invoice
  archiveDialog.value = true
}

const confirmArchive = async () => {
  if (!selectedInvoice.value) return

  try {
    // Implement archive API call
    console.log('Archiving invoice:', selectedInvoice.value.ID)
    // await archiveInvoiceAPI(selectedInvoice.value.ID)

    // Remove from local list
    const index = invoices.value.findIndex(inv => inv.ID === selectedInvoice.value?.ID)
    if (index > -1) {
      invoices.value.splice(index, 1)
    }

    archiveDialog.value = false
    selectedInvoice.value = null
  } catch (err) {
    console.error('Error archiving invoice:', err)
  }
}

// Lifecycle
onMounted(() => {
  fetchInvoices()
})
</script>

<style scoped>
.max-width-400 {
  max-width: 400px;
}

/* Custom table styling for better column spacing */
:deep(.v-data-table) {
  --v-data-table-header-height: 56px;
}

:deep(.v-data-table th) {
  padding: 0 8px !important;
}

:deep(.v-data-table td) {
  padding: 8px !important;
}


:deep(.v-data-table th:nth-child(1)) {
  min-width: 200px;
}

:deep(.v-data-table th:nth-child(2)) {
  min-width: 120px;
}

:deep(.v-data-table th:nth-child(3)) {
  min-width: 100px;
}

:deep(.v-data-table th:nth-child(4)) {
  min-width: 100px;
}

:deep(.v-data-table th:nth-child(5)) {
  min-width: 100px;
}

:deep(.v-data-table th:nth-child(6)) {
  min-width: 120px;
}

  :deep(.v-data-table th:nth-child(7)) { /* Actions */
  min-width: 160px;
}
</style>
