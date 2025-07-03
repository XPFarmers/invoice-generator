<template>
  <v-form @submit.prevent="submitForm">

    <v-row justify="end" class="mt-4 mb-6 mr-8">
      <v-btn :color="Colours.Crimson" class="mr-4" @click="clearForm">Clear</v-btn>
      <v-btn :color="Colours.ForestGreen" type="submit">Submit</v-btn>
    </v-row>

    <v-card elevation="4" class="mx-8" rounded="lg">

      <!-- Client name -->
      <v-row class="ml-2 mt-2 mr-2">
        <v-col cols="12">
          <v-text-field v-model="invoice.name" label="Client Name" type="text" outlined dense required />
        </v-col>
      </v-row>

      <!-- Client Address -->
      <v-row class="ml-2 mr-2">
        <v-col cols="12">
          <v-text-field v-model="invoice.address" label="Client Address" type="text" outlined dense required />
        </v-col>
      </v-row>

      <!-- Time To Order -->
      <v-row class="ml-2 mr-2">
        <v-col cols="12">
          <v-text-field v-model="invoice.timeToOrder" label="Time to Order" type="number" outlined dense min="0"
            required />
        </v-col>
      </v-row>

      <v-row class="ml-2 mr-2">
        <v-col cols="12">
          <v-slider v-model="invoice.depositPercentage" :min="0" :max="100" step="1" label="Deposit %"
            thumb-label="always" :color="Colours.DodgerBlue" show-ticks>
            <template #append>
              <span class="font-weight-medium text-caption ml-2">
                {{ invoice.depositPercentage }}%
              </span>
            </template>
          </v-slider>
        </v-col>
      </v-row>

      <v-row class="ml-2 mr-2 mb-4">
        <v-col cols="6" class="text-center">
          <div class="text-title font-weight-medium">
            Deposit: <span :style="{ color: Colours.ForestGreen }">R {{ deposit.toFixed(2) }}</span>
          </div>
        </v-col>
        <v-col cols="6" class="text-center">

          <div class="text-title font-weight-medium">
            Balance: <span :style="{ color: Colours.Crimson }">R {{ balance.toFixed(2) }}</span>
          </div>
        </v-col>
      </v-row>



    </v-card>

    <v-card class="mx-8 mt-6 mb-6" elevation="4" rounded="lg">
      <v-card-title class="text-h6 font-weight-bold mt-4 mb-4">Add Line Items</v-card-title>

      <!-- <v-card-text>
      <v-form @submit.prevent="submitForm"> -->
      <!-- Timestamp Field -->
      <!-- <v-row class="mb-4">
          <v-col cols="12" md="6">
            <v-text-field v-model="formData.timestamp" label="Timestamp" type="date" outlined dense required />
          </v-col>
        </v-row> -->

      <!-- Line Items Table -->
      <v-table density="compact" class="mb-4">
        <thead>
          <tr>
            <th>QTY</th>
            <th>Description</th>
            <th>Amount</th>
            <th>Total</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(entry, index) in entries" :key="index">
            <td>{{ entry.qty }}</td>
            <td>{{ entry.description }}</td>
            <!-- <td>{{ entry.amount }}</td> -->
            <td>R {{ entry.amount }}</td>
            <td>R {{ entry.total }}</td>
            <td>
              <v-btn icon :color="Colours.Crimson" variant="plain" @click="
                deleteEntry(index)">
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </td>
          </tr>
        </tbody>
        <!-- <tr>
          <td colspan="5">
            <v-divider class="mt-16" />
          </td>
        </tr> -->
        <tfoot>
          <tr>
            <!-- <td colspan="3" class="text-right font-weight-bold">Grand Total:</td>
          <td colspan="2" class="font-weight-bold">R {{ Number(grandTotal).toFixed(2) }}</td> -->
            <td colspan="5" class="text-right">
              <v-card class="mt-12 d-inline-block" variant="outlined">
                <v-card-text class="text-h6 text-right">
                  Grand Total: R {{ Number(grandTotal).toFixed(2) }}
                </v-card-text>
              </v-card>
            </td>

          </tr>
        </tfoot>
      </v-table>
    </v-card>

    <v-card class="mx-8 mt-4" elevation="4" rounded="lg">
      <v-card-text>
        <!-- Entry Form Fields -->
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field v-model="formData.qty" label="QTY" type="number" outlined dense min="0" required />
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="formData.description" label="Description" type="text" outlined dense required />
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="formData.amount" label="Amount" type="number" outlined dense min="0" prefix="R"
              required />
          </v-col>
        </v-row>

        <!-- Action Buttons -->
        <v-row class="mt-6 mr-1 mb-4" justify="end">
          <v-btn :color="Colours.Crimson" class="mr-4" @click.prevent="clearEntries">Clear Entries</v-btn>
          <v-btn :color="Colours.DodgerBlue" @click.prevent="addEntry">Add Entry</v-btn>
        </v-row>

        <!-- <v-divider class="my-6" thickness="2" color="black" /> -->


      </v-card-text>
    </v-card>
  </v-form>

</template>

<script setup lang="ts">
import { Colours } from '@/utils/colours'
import { LineItem } from '@/models/lineItems'
import { ref } from 'vue'

import { computed } from 'vue'
import { Routes } from '@/utils/routes'
import { GenerateInvoice } from '@/models/generateInvoice'


const invoice = ref(new GenerateInvoice())

const formData = ref({
  qty: 0,
  description: '',
  amount: 0
})

// const entries = ref<LineItem[]>([])
// const grandTotal = computed(() =>
//   entries.value.reduce((sum, entry) => sum + entry.total, 0)
// )

const entries = computed(() => invoice.value.getLineItems())
const grandTotal = computed(() => invoice.value.getGrandTotal())

const deposit = computed(() => {
  return grandTotal.value * (invoice.value.depositPercentage / 100)
})

const balance = computed(() => {
  return grandTotal.value - deposit.value
})


function clear() {
  formData.value.qty = 0
  formData.value.description = ''
  formData.value.amount = 0
}

function clearForm() {
  clear()
  clearEntries()
  invoice.value = new GenerateInvoice()
}

function clearEntries() {
  invoice.value.lineItems = []
}

function deleteEntry(index: number) {
  entries.value.splice(index, 1)
}

function addEntry() {
  const entry = new LineItem(formData.value.qty, formData.value.description, formData.value.amount)
  invoice.value.addLineItem(entry)
  clear()
}

// function submitForm() {
//   const payload = entries.value.map(e => e.toJSON())

//   console.log('Ready to submit:', payload)
//   clearForm()
// }

async function submitForm() {
  // console.log("Invoice = ", invoice.value)
  // const payload = entries.value.map(e => e.toJSON())

  // try {
  //   const response = await fetch(Routes.LineItems, {
  //     method: 'POST',
  //     headers: {
  //       'Content-Type': 'application/json',
  //     },
  //     body: JSON.stringify(payload),
  //   })

  //   if (!response.ok) throw new Error('Failed to submit')

  //   console.log('✅ Submitted successfully')
  //   clearForm()
  // } catch (err) {
  //   console.error('❌ Submission failed:', err)
  // }

  const payload = invoice.value.toJSON()

  try {
    const response = await fetch(Routes.GenerateInvoice, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    })

    if (!response.ok) throw new Error('Failed to submit')
    console.log('✅ Submitted successfully')

    clearForm()
  } catch (err) {
    console.error('❌ Submission failed:', err)
  }
}


</script>
