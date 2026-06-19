<script setup>
import { onMounted, ref } from 'vue'
import { useJobsStore } from './stores/jobs'

const store = useJobsStore()
const newJob = ref("")

onMounted(() => {
  store.fetchJobs()
  setInterval(() => store.fetchJobs(), 2000)
})

function addJob() {
  if (newJob.value) {
    store.createJob(newJob.value)
    newJob.value = ""
  }
}
</script>

<template>
  <div class="container">

    <h1>MLOps Control Center</h1>

    <div class="grid" style="grid-template-columns: 2fr 1fr; margin: 20px 0;">

      <input v-model="store.search" placeholder="Search jobs..." />

      <select v-model="store.filter">
        <option value="all">All</option>
        <option value="running">Running</option>
        <option value="completed">Completed</option>
        <option value="failed">Failed</option>
      </select>

    </div>

    <div style="display:flex; gap:10px; margin-bottom:20px;">
      <input v-model="newJob" placeholder="New training job..." />
      <button class="btn btn-primary" @click="addJob">
        Create Job
      </button>
    </div>

    <div class="grid" style="grid-template-columns: repeat(3, 1fr);">

      <div v-for="job in store.filteredJobs" :key="job.id" class="card">

        <h2>{{ job.name }}</h2>

        <span
          class="badge"
          :class="{
            'badge-running': job.status === 'running',
            'badge-completed': job.status === 'completed',
            'badge-failed': job.status === 'failed'
          }"
        >
          {{ job.status }}
        </span>

        <p style="margin-top:10px;">
          Accuracy: {{ job.accuracy }}%
        </p>

        <div style="display:flex; gap:8px; margin-top:12px;">

          <button class="btn btn-danger" @click="store.deleteJob(job.id)">
            Delete
          </button>

          <button class="btn btn-primary"
            @click="store.updateJob(job.id, { status: 'running' })">
            Restart
          </button>

        </div>

      </div>

    </div>

  </div>
</template>