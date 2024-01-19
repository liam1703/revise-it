import { ref } from 'vue'


const userEmail = ref('')
const firstName = ref('')
const lastName = ref('')
const userId = ref('')
const isLoggedIn = ref(false)

export {
  userEmail,
  firstName,
  lastName,
  userId,
  isLoggedIn
}
