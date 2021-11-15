<script lang="ts">
  import { userStore, UserStore } from "./store";
  import { onMount } from "svelte";
  import { getUser } from "./infratructure/paraphrace";

  import Landing from "./pages/Landing.svelte";
  import ParaphraseDash from "./pages/Landing.svelte";

  import { Router, Route } from "svelte-routing";

  let user: UserStore;

  $: url = window.location.pathname;

  userStore.subscribe((v) => (user = v));

  onMount(async () => {
    if (user.id === -1 || user.token === "") {
      try {
        const u = await getUser();
        userStore.set(u);
      } catch (e) {
        console.error("Error getting user", e);
        alert("Unable to access API. Contact support");
      }
    }
  });
</script>

<Router {url}>
  <Route path="/paraphrase" component={ParaphraseDash} />
  <Route path="/" component={Landing} />
</Router>
