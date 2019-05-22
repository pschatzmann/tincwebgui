<!--
 Openid Connect Callback
-->

<script>
import { SecurityService } from '@/services/SecurityService'

export default {
    name: "callback",
    mounted() {
      SecurityService.signinRedirectCallback().then(()=>{
        this.$store.dispatch('setLoggedIn', true)
        this.$router.push("/")
      })
    },
  }
</script>

