import { defineStore } from 'pinia'
import api from '../services/api'

export const useJobsStore = defineStore('jobs', {
 state: () => ({
  jobs: [],  
  search: '',
  filter: 'all',
  loading: false
  }),

  getters: {
    filteredJobs(state) {
  if (!Array.isArray(state.jobs)) return []

  return state.jobs
    .filter(j => state.filter === 'all' || j.status === state.filter)
    .filter(j => j.name?.toLowerCase().includes(state.search.toLowerCase()))
    }
  },

  actions: {

   async fetchJobs() {
  try {
    const res = await api.get('/jobs')
    this.jobs = Array.isArray(res.data) ? res.data : []
  } catch (err) {
    console.log(err)
    this.jobs = []
  }
},

    async createJob(name) {
      await api.post('/jobs', { name })
      await this.fetchJobs()
    },

    async deleteJob(id) {
      await api.delete(`/jobs/${id}`)
      await this.fetchJobs()
    },

    async updateJob(id, data) {
      await api.patch(`/jobs/${id}`, data)
      await this.fetchJobs()
    }
  }
})