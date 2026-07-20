import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', () => {
  const authToken = ref('')
  const data = ref({})

  // Add a computed property

  let localdata = localStorage.getItem('data')
  if (localdata !== undefined) {
    data.value = localdata
  }
  // TODO: Handle parse errors here
  function getData() {
    let localData = localStorage.getItem('data')
    if (localData) {
      return JSON.parse(localData)
    }
    return {}
  }

  function setData(d) {
    data.value = d
    localStorage.setItem('data', JSON.stringify(data.value))
  }

  function setAuthToken(username, password) {
    authToken.value = btoa(`${username}:${password}`)
  }
  // Ewwww, returning data doesn't seem fine
  return { authToken, setAuthToken, data, getData, setData }
})
